package main

import (
	"fmt"
	pb "github.com/arch-yunus/microservices-101/proto/product"
	"github.com/arch-yunus/microservices-101/services/product-service/internal/handler"
	"github.com/arch-yunus/microservices-101/services/product-service/internal/repository"
	"github.com/arch-yunus/microservices-101/services/product-service/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("?? Ürün Servisi (Product Service) gRPC Sunucusu balatiliyor...")

	// 1. Altyapı ve Servislerin Kurulumu (Dependency Injection)
	repo := repository.NewMemoryProductRepo()
	productSvc := service.NewProductService(repo)

	// 2. gRPC Sunucusu Yaplandrmas
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("?? Port Dinleme Hatas: %v", err)
	}

	grpcServer := grpc.NewServer()
	productHandler := handler.NewGrpcProductHandler(productSvc)
	pb.RegisterProductServiceServer(grpcServer, productHandler)

	// 3. Zarif Kapan (Graceful Shutdown) Stratejisi
	// Bir kanal (channel) olusturup Isletim Sistemi sinyallerini dinliyoruz.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Println("?? gRPC Sunucusu :50051 portunda dinliyor...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("?? Sunucu Hatas: %v", err)
		}
	}()

	// Sinyal gelene kadar burada bloklanrz.
	<-stop

	fmt.Println("\n?? Sunucu Kapatiliyor... (Graceful Shutdown)")
	grpcServer.GracefulStop()
	fmt.Println("?? Sunucu Guvenli Bir Sekilde Durduruldu.")
}
