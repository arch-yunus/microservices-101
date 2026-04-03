package main

import (
	"context"
	"fmt"
	pb "github.com/arch-yunus/microservices-101/proto/product"
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
	fmt.Println("?? Sipariş Servisi (Order Service) balatiliyor...")

	// 1. gRPC Bağlantıs Kurulumu (Product Service'e baılan)
	// Not: Docker ortamda localhost yerine "product-service" kullanlabilir.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("?? Baılantı Hatas: %v", err)
	}
	defer conn.Close()

	productClient := pb.NewProductServiceClient(conn)
	orderSvc := service.NewOrderService(productClient)

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
