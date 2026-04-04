package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	pb "github.com/arch-yunus/microservices-101/proto/product"
	"github.com/arch-yunus/microservices-101/services/product-service/internal/handler"
	"github.com/arch-yunus/microservices-101/services/product-service/internal/repository"
	"github.com/arch-yunus/microservices-101/services/product-service/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("🚀 Ürün Servisi (Product Service) gRPC Sunucusu başlatılıyor...")

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

	// 2. Altyapı ve Servislerin Kurulumu (Dependency Injection)
	repo := repository.NewPostgresProductRepository(db)
	productSvc := service.NewProductService(repo)

	// 2. gRPC Sunucusu Yaplandrmas
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("?? Port Dinleme Hatas: %v", err)
	}

	grpcServer := grpc.NewServer()
	productHandler := handler.NewGrpcProductHandler(productSvc)
	pb.RegisterProductServiceServer(grpcServer, productHandler)

	// ?? Health Check Servisi Kaydı
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

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
