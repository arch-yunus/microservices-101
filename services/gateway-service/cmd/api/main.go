package main

import (
	"context"
	"fmt"
	"github.com/arch-yunus/microservices-101/services/gateway-service/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("🚀 API Gateway (The Fortress) başlatılıyor...")

	e := echo.New()

	// 1. Middleware Kurulumu
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 2. Handler Kurulumu
	gwHandler := handler.NewGatewayHandler()

	// 3. API Rotaları
	api := e.Group("/api/v1")
	{
		api.GET("/products/:id", gwHandler.GetProduct)
		api.POST("/orders", gwHandler.CreateOrder)
	}

	// 4. Zarif Kapanış (Graceful Shutdown)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Println("🚀 API Gateway :8080 portunda dış dünyaya açık.")
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("?? Gateway Hatası: %v", err)
		}
	}()

	// Sinyal gelene kadar burada bloklanırız.
	<-stop
	fmt.Println("\n🚀 API Gateway Güvenli Bir Şekilde Kapatılıyor...")
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
