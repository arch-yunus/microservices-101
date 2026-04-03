# ?? microservices-101: Makefile
# Bu dosya sistem paralarini yonetmek iin "Elite" araclar sunar.

.PHONY: up down build test run-product

# ?? Altyap servislerini baslat
up:
	docker-compose -f infrastructure/docker-compose.yml up -d

# ?? Altyap servislerini durdur
down:
	docker-compose -f infrastructure/docker-compose.yml down

# ?? Tum servislerin Docker imagelerini olustur
build:
	docker-compose -f infrastructure/docker-compose.yml build

# ?? Product Service'i yerel olarak calistir
run-product:
	cd services/product-service && go run cmd/api/main.go

# ?? Tum testleri calistir
test:
	go test ./...
