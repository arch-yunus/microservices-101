package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	pb "github.com/arch-yunus/microservices-101/proto/product"
	"github.com/arch-yunus/microservices-101/services/order-service/internal/repository"
	"github.com/arch-yunus/microservices-101/services/order-service/internal/service"
	"github.com/arch-yunus/microservices-101/services/order-service/internal/pkg/otel"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	fmt.Println("🚀 Order Service başlatılıyor...")

	// 🛡️ OpenTelemetry Tracing İlklendirme
	shutdown, err := otel.InitTracer("order-service")
	if err != nil {
		log.Fatal(err)
	}
	defer shutdown(context.Background())

	// 1. Veritabanı ve Mesaj Kuyruğu Bağlantıları (Retry Logic ile)
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
		log.Printf("⚠️ Veritabanına (Postgres) bağlanılamadı, 2s içinde tekrar denenecek (%d/10): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("❌ Veritabanı Bağlantı Hatası: %v", err)
	}
	defer db.Close()

	productSvcAddr := os.Getenv("PRODUCT_SERVICE_ADDR")
	if productSvcAddr == "" {
		productSvcAddr = "product-service:50051"
	}

	// Product Service gRPC Bağlantısı (Context Propagation Interceptor ile)
	conn, err := grpc.Dial(productSvcAddr, 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	if err != nil {
		log.Fatalf("❌ Product Service'e bağlanılamadı: %v", err)
	}
	defer conn.Close()

	productClient := pb.NewProductServiceClient(conn)
	repo := repository.NewPostgresOrderRepository(db)
	orderSvc := service.NewOrderService(productClient, repo)

	// ⚙️ gRPC Sunucusu Yapılandırması (Health Check & Tracing için)
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("❌ Port Dinleme Hatası: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
	)
	
	// Health Check Servisi Kaydı
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	go func() {
		fmt.Println("🚀 Health Check & Tracing Sunucusu :50052 portunda dinliyor...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("❌ Sunucu Hatası: %v", err)
		}
	}()

	// 2. Zarif Kapanış (Graceful Shutdown)
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
