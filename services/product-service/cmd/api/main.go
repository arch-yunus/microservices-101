package main

import (
	"fmt"
	"github.com/arch-yunus/microservices-101/services/product-service/internal/repository"
	"github.com/arch-yunus/microservices-101/services/product-service/internal/service"
)

func main() {
	fmt.Println("?? Product Service balatiliyor...")

	// 1. Repository Balat (Memory - Geici)
	repo := repository.NewMemoryProductRepo()

	// 2. Servis Balat
	productSvc := service.NewProductService(repo)

	// 3. Ornek Bir urun Olustur (Eitim Amacl)
	p, err := productSvc.CreateNewProduct("Elite Architect Keyboard", 199.99)
	if err != nil {
		fmt.Printf("?? Hata: %v\n", err)
		return
	}

	fmt.Printf("?? urun Olusturuldu: %+v\n", p)

	// 4. Listele
	list, _ := repo.List()
	fmt.Printf("?? Mevcut urunler (%d): %+v\n", len(list), list)

	fmt.Println("?? Servis Ayakta!")
}
