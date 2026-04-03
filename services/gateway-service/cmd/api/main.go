package main

import (
	"fmt"
	"github.com/arch-yunus/microservices-101/services/gateway-service/internal/auth"
	"github.com/arch-yunus/microservices-101/services/gateway-service/internal/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("?? API Gateway (The Fortress) balatiliyor...")

	// 1. Gateway Handler ve Middleware Kurulumu
	// Tüm istekler önce JWT kontrolünden geçer, sonra iç servislere yönlendirilir.
	gwHandler := auth.JWTAuthMiddleware(handler.NewGatewayHandler())

	server := &http.Server{
		Addr:    ":8080",
		Handler: gwHandler,
	}

	// 2. Zarif Kapan (Graceful Shutdown)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Println("?? API Gateway :8080 portunda dis dunyaya acik.")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("?? Gateway Hatas: %v", err)
		}
	}()

	// Sinyal gelene kadar burada bloklanrz.
	<-stop
	fmt.Println("\n?? API Gateway Guvenli Bir Sekilde Kapatiliyor...")
}
