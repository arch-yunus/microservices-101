package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/arch-yunus/microservices-101/proto/product"
	"github.com/arch-yunus/microservices-101/services/order-service/internal/domain"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sony/gobreaker"
	"log"
	"time"
)

// OrderService sipariş mantığını yöneten yapı
type OrderService struct {
	productClient product.ProductServiceClient
	repo          domain.OrderRepository
	cb            *gobreaker.CircuitBreaker
}

func NewOrderService(pc product.ProductServiceClient, repo domain.OrderRepository) *OrderService {
	// Circuit Breaker Yapılandırması
	cbSettings := gobreaker.Settings{
		Name:        "ProductServiceBreaker",
		MaxRequests: 5,
		Interval:    10 * time.Second,
		Timeout:     30 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio >= 0.6
		},
		OnStateChange: func(name string, from, to gobreaker.State) {
			log.Printf("⚠️  Circuit Breaker [%s] Durum Değişti: %s -> %s", name, from, to)
		},
	}

	return &OrderService{
		productClient: pc,
		repo:          repo,
		cb:            gobreaker.NewCircuitBreaker(cbSettings),
	}
}

// PlaceOrder yeni bir sipariş olusturur ve olayı RabbitMQ ya fırlatır (Asenkron)
func (s *OrderService) PlaceOrder(ctx context.Context, productID string, quantity int) (*domain.Order, error) {
	// 1. gRPC ile urun bilgilerini sorgula (Senkron) - Circuit Breaker Korumalı
	var p *product.ProductResponse
	_, err := s.cb.Execute(func() (interface{}, error) {
		var err error
		p, err = s.productClient.GetProduct(ctx, &product.GetProductRequest{Id: productID})
		return p, err
	})

	if err != nil {
		if err == gobreaker.ErrOpenState {
			return nil, fmt.Errorf("kritik hata: ürün servisi geçici olarak devre dışı (Circuit Breaker OPEN)")
		}
		return nil, fmt.Errorf("urun bilgisi alinamadi: %v", err)
	}

	// 2. Sipariş objesini olustur
	order := &domain.Order{
		ID:        uuid.New().String(),
		ProductID: p.Id,
		Quantity:  quantity,
		Total:     p.Price * float64(quantity),
		CreatedAt: time.Now(),
	}

	// 3. Veritabanına kaydet
	if err := s.repo.Create(order); err != nil {
		return nil, fmt.Errorf("siparis kaydedilemedi: %v", err)
	}

	fmt.Printf("📦 Sipariş Olusturuldu ve Kaydedildi: %s için %d adet (Toplam: %.2f TL)\n", p.Name, quantity, order.Total)

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
