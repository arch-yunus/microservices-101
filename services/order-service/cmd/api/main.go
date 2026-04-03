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
	fmt.Println("?? Order Service balatiliyor...")

	// 1. gRPC Bağlantıs Kurulumu (Product Service'e baılan)
	// Not: Docker ortamda localhost yerine "product-service" kullanlabilir.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("?? Baılantı Hatas: %v", err)
	}
	defer conn.Close()

	productClient := pb.NewProductServiceClient(conn)
	orderSvc := service.NewOrderService(productClient)

	// 2. Graceful Shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Println("?? Order Service islemler iin hazır (7/24 Aktif).")
		
		// Eitim Amacl: Periyodik olarak bir sipari? denemesi yapalım (Simülasyon)
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
				_, err := orderSvc.PlaceOrder(ctx, "top-tier-key-123", 1)
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
	fmt.Println("\n?? Order Service Guvenli Bir Sekilde Kapatiliyor...")
}
