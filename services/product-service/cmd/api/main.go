package main

import (
	"context"
	"database/sql"
	"fmt"
	pb "github.com/arch-yunus/microservices-101/proto/product"
	"github.com/arch-yunus/microservices-101/services/product-service/internal/handler"
	"github.com/arch-yunus/microservices-101/services/product-service/internal/pkg/otel"
	"github.com/arch-yunus/microservices-101/services/product-service/internal/repository"
	"github.com/arch-yunus/microservices-101/services/product-service/internal/service"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("🚀 Product Service başlatılıyor...")

	// 🚀 OpenTelemetry Tracing İlklendirme
	shutdown, err := otel.InitTracer("product-service")
	if err != nil {
		log.Fatal(err)
	}
	defer shutdown(context.Background())

	// 1. Veritabanı Bağlantısı
	dbAddr := os.Getenv("DATABASE_URL")
	if dbAddr == "" {
		dbAddr = "postgres://user:password@postgres:5432/micro_db?sslmode=disable"
	}

	var db *sql.DB
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", dbAddr)
		if err == nil {
			err = db.Ping()
		}
		if err == nil {
			break
		}
		log.Printf("?? Veritabanına (Postgres) bağlanılamadı, 2s içinde tekrar denenecek (%d/10): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("?? Veritabanı Bağlantı Hatası: %v", err)
	}
	defer db.Close()

	// 2. Servis Kablolaması (Dependency Injection)
	repo := repository.NewPostgresProductRepository(db)
	productSvc := service.NewProductService(repo)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("?? Port Dinleme Hatası: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
	)
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
