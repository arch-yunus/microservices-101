package main

import (
	"fmt"
	pb "github.com/arch-yunus/microservices-101/proto/product"
	"github.com/arch-yunus/microservices-101/services/product-service/internal/handler"
	"github.com/arch-yunus/microservices-101/services/product-service/internal/repository"
	"github.com/arch-yunus/microservices-101/services/product-service/internal/service"
	"google.golang.org/grpc"
	"net"
)

func main() {
	fmt.Println("?? Product Service gRPC ve API balatiliyor...")

	// 1. Altyapı ve Servislerin Kurulumu
	repo := repository.NewMemoryProductRepo()
	productSvc := service.NewProductService(repo)

	// 2. gRPC Sunucusu Kurulumu
	lis, err := net.Listen("tcp", ":50051") // Standart gRPC portu
	if err != nil {
		fmt.Printf("?? Port Dinleme Hatas: %v\n", err)
		return
	}

	grpcServer := grpc.NewServer()
	productHandler := handler.NewGrpcProductHandler(productSvc)
	pb.RegisterProductServiceServer(grpcServer, productHandler)

	fmt.Println("?? gRPC Sunucusu :50051 portunda dinliyor...")
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("?? Sunucu Hatas: %v\n", err)
	}
}
