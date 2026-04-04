# ==============================================================================
# Microservices Architecture: Engineering Handbook - Makefile
# ==============================================================================
# Bu dosya, sistem bileşenlerinin yönetimi, geliştirilmesi ve orkestrasyonu 
# için merkezi bir kontrol noktası olarak tasarlanmıştır.

# Proje Değişkenleri
DOCKER_COMPOSE = docker-compose -f infrastructure/docker-compose.yml
PROTO_DIR = proto

.PHONY: help up down build test proto run-product run-order run-gateway

help:
	@echo "Kullanılabilir komutlar:"
	@echo "  make up           - Altyapı servislerini (PostgreSQL, RabbitMQ, Redis) başlatır."
	@echo "  make down         - Altyapı servislerini durdurur."
	@echo "  make build        - Tüm mikroservis imajlarını oluşturur."
	@echo "  make proto        - Proto dosyalarından Go kodlarını üretir."
	@echo "  make test         - Tüm birim testleri (Unit Tests) çalıştırır."
	@echo "  make run-product  - Product Service'i yerel ortamda başlatır."
	@echo "  make run-order    - Order Service'i yerel ortamda başlatır."
	@echo "  make run-gateway  - Gateway Service'i yerel ortamda başlatır."

# --- Altyapı Yönetimi ---
up:
	$(DOCKER_COMPOSE) up -d

down:
	$(DOCKER_COMPOSE) down

# --- Geliştirme Araçları ---
proto:
	@echo "Protobuf kodları üretiliyor..."
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	$(PROTO_DIR)/product/product.proto
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	$(PROTO_DIR)/order/order.proto

build:
	$(DOCKER_COMPOSE) build

test:
	go test -v ./...

# --- Servis Çalıştırma ---
run-product:
	go run services/product-service/cmd/api/main.go

run-order:
	go run services/order-service/cmd/api/main.go

run-gateway:
	go run services/gateway-service/cmd/api/main.go

# ==============================================================================
