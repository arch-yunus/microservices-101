package consumer

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

// ListenOrderEvents RabbitMQ dan siparis olaylarını dinler
func ListenOrderEvents() {
	var conn *amqp.Connection
	var err error

	// Haberleşme katmanı baglantı denemesi (Retry Logic)
	for i := 0; i < 10; i++ {
		conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err == nil {
			break
		}
		log.Printf("?? RabbitMQ'ya bağlanılamıyor, tekrar deneniyor (%d/10): %v", i+1, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("?? RabbitMQ Bağlantı Hatası: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("?? Kanal Acma Hatası: %v", err)
		return
	}
	defer ch.Close()

	// Kuyruk Tanımlama
	q, err := ch.QueueDeclare(
		"order_created_queue", // isim
		false,                 // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)
	if err != nil {
		log.Printf("?? Kuyruk Tanımlama Hatası: %v", err)
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Printf("?? Mesaj Dinleme Hatası: %v", err)
		return
	}

	fmt.Println("?? Notifier Service: Sipariş Olayları Dinleniyor...")

	for d := range msgs {
		fmt.Printf("?? BİLDİRİM: Yeni bir sipariş alındı! Sipariş Detayı: %s\n", d.Body)
		fmt.Println("📧 Email ve SMS gönderildi (Simülasyon).")
	}
}
