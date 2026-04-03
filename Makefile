# microservices-101: Makefile
# Bu dosya sistem parcalarını yonetmek icin kullanılan ana merkezdir.

.PHONY: up down build test run-product

# Altyapı servislerini (DB, Broker, Cache) baslat
up:
	docker-compose -f infrastructure/docker-compose.yml up -d

# Altyapı servislerini durdur
down:
	docker-compose -f infrastructure/docker-compose.yml down

# Proto dosyalarından Go kodunu üret
# Not: protoc, protoc-gen-go ve protoc-gen-go-grpc kurulu olmalıdır.
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/product/product.proto

# Tum servislerin Docker imagelerini olustur
build:
	docker-compose -f infrastructure/docker-compose.yml build

# Servisleri yerel olarak calistir
run-product:
	cd services/product-service && go run cmd/api/main.go

run-order:
	cd services/order-service && go run cmd/api/main.go

# Tum testleri calistir
test:
	go test ./...
