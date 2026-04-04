package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	pb "github.com/arch-yunus/microservices-101/proto/product"
	"github.com/arch-yunus/microservices-101/services/order-service/internal/repository"
	"github.com/arch-yunus/microservices-101/services/order-service/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("🚀 Sipariş Servisi (Order Service) başlatılıyor...")

	// 1. Veritabanı Bağlantısı
	dbAddr := os.Getenv("DATABASE_URL")
	if dbAddr == "" {
		dbAddr = "postgres://user:password@postgres:5432/micro_db?sslmode=disable"
	}

	var db *sql.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", dbAddr)
		if err == nil {
			if err = db.Ping(); err == nil {
				break
			}
		}
		log.Printf("?? Veritabanına bağlanılamıyor, tekrar deneniyor (%d/10): %v", i+1, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("?? Veritabanı Bağlantı Hatası: %v", err)
	}
	defer db.Close()

	// 2. gRPC Bağlantısı Kurulumu (Product Service'e bağlan)
	productSvcAddr := os.Getenv("PRODUCT_SERVICE_ADDR")
	if productSvcAddr == "" {
		productSvcAddr = "product-service:50051"
	}

	conn, err := grpc.Dial(productSvcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("?? gRPC Bağlantı Hatası: %v", err)
	}
	defer conn.Close()

	productClient := pb.NewProductServiceClient(conn)
	repo := repository.NewPostgresOrderRepository(db)
	orderSvc := service.NewOrderService(productClient, repo)

	// 2. Zarif Kapan (Graceful Shutdown)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Println("?? Sipariş Servisi islemler iin hazır (7/24 Aktif).")
		
		// Eğitim Amacl: Periyodik olarak bir sipari? denemesi yapalım (Simülasyon)
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
				_, err := orderSvc.PlaceOrder(ctx, "üst-seviye-anahtar-123", 1)
				if err != nil {
					fmt.Printf("?? Otomatik Sipari Hatas: %v\n", err)
				}
				cancel()
			case <-stop:
				return
			}
		}
	}()

	// Sinyal gelene kadar burada bloklanrz.
	<-stop
	fmt.Println("\n?? Sipariş Servisi Guvenli Bir Sekilde Kapatiliyor...")
}
