package main

import (
	"context"
	"fmt"
	pb "github.com/arch-yunus/microservices-101/proto/product"
	"github.com/arch-yunus/microservices-101/services/order-service/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("?? Order Service balatiliyor...")

	// 1. gRPC Bağlantıs Kurulumu (Product Service'e baılan)
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("?? Bağlantı Hatas: %v\n", err)
		return
	}
	defer conn.Close()

	productClient := pb.NewProductServiceClient(conn)

	// 2. Order Servis Kurulumu
	orderSvc := service.NewOrderService(productClient)

	// 3. Örnek Sipariş Ver (Eitim Amaçlı)
	ctx := context.Background()
	order, err := orderSvc.PlaceOrder(ctx, "top-tier-key-123", 2)
	if err != nil {
		fmt.Printf("?? Sipari Hatas: %v\n", err)
		return
	}

	fmt.Printf("?? Sipari Başarıyla Oluşturuldu: %+v\n", order)
}
