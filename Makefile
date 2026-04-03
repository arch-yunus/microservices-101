# ?? microservices-101: Makefile
# Bu dosya sistem paralarini yonetmek iin "Elite" araclar sunar.

.PHONY: up down build test run-product

# ?? Altyap servislerini baslat
up:
	docker-compose -f infrastructure/docker-compose.yml up -d

# ?? Altyap servislerini durdur
down:
	docker-compose -f infrastructure/docker-compose.yml down

# ?? Proto dosyalarndan Go kodunu üret
# Not: protoc, protoc-gen-go ve protoc-gen-go-grpc kurulu olmaldr.
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/product/product.proto

# ?? Tum servislerin Docker imagelerini olustur
build:
	docker-compose -f infrastructure/docker-compose.yml build

# ?? Servisleri yerel olarak calistir
run-product:
	cd services/product-service && go run cmd/api/main.go

run-order:
	cd services/order-service && go run cmd/api/main.go

# ?? Tum testleri calistir
test:
	go test ./...
