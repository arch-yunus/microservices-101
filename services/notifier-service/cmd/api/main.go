package main

import (
	"fmt"
	"github.com/arch-yunus/microservices-101/services/notifier-service/internal/consumer"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("📧 Notifier Service (Bildirim Merkezo) baslatiliyor...")

	// 1. Olayları Ayrtı Bir Goroutine de Dinle
	go consumer.ListenOrderEvents()

	// 2. Graceful Shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	fmt.Println("📧 Notifier Service: 7/24 Aktif ve Dinliyor.")

	// Sinyal gelene kadar burada bloklanrz.
	<-stop
	fmt.Println("\n📧 Notifier Service Guvenli Bir Sekilde Kapatiliyor...")
}
