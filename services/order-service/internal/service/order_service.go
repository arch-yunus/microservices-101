package service

import (
	"context"
	"fmt"
	"github.com/arch-yunus/microservices-101/proto/product"
	"github.com/arch-yunus/microservices-101/services/order-service/internal/domain"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

// OrderService sipariş mantığını yöneten yapı
type OrderService struct {
	productClient product.ProductServiceClient
}

func NewOrderService(pc product.ProductServiceClient) *OrderService {
	return &OrderService{productClient: pc}
}

// PlaceOrder yeni bir sipariş olusturur ve olayı RabbitMQ ya fırlatır (Asenkron)
func (s *OrderService) PlaceOrder(ctx context.Context, productID string, quantity int) (*domain.Order, error) {
	// 1. gRPC ile urun bilgilerini sorgula (Senkron)
	p, err := s.productClient.GetProduct(ctx, &product.GetProductRequest{Id: productID})
	if err != nil {
		return nil, fmt.Errorf("urun bilgisi alinamadi: %v", err)
	}

	// 2. Sipariş objesini olustur
	order := &domain.Order{
		ID:        fmt.Sprintf("order-%d", time.Now().Unix()),
		ProductID: p.Id,
		Quantity:  quantity,
		Total:     p.Price * float64(quantity),
		CreatedAt: time.Now(),
	}

	fmt.Printf("?? Sipariş Olusturuldu: %s için %d adet (Toplam: %.2f TL)\n", p.Name, quantity, order.Total)

	// 3. RABBITMQ'YA OLAY FIRLAT (Asenkron Haberleşme)
	// Not: Gerçek projelerde baglantı havuzu (connection pool) kullanılmalıdır.
	go s.publishOrderCreatedEvent(order)

	return order, nil
}

func (s *OrderService) publishOrderCreatedEvent(order *domain.Order) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Printf("?? RabbitMQ Baglantı Hatası: %v", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("?? Kanal Acma Hatası: %v", err)
		return
	}
	defer ch.Close()

	q, _ := ch.QueueDeclare("order_created_queue", false, false, false, false, nil)

	body := fmt.Sprintf("Siparis ID: %s, Urun: %s, Tutar: %.2f TL", order.ID, order.ProductID, order.Total)
	err = ch.PublishWithContext(context.Background(),
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	if err != nil {
		log.Printf("?? Mesaj Gonderim Hatası: %v", err)
	} else {
		fmt.Println("🚀 Olay Havaya Fırlatıldı: order_created_queue")
	}
}
