# ?? microservices-101: Proto Kod Üretim Betii (Windows/PowerShell)
# Not: protoc, protoc-gen-go ve protoc-gen-go-grpc kurulu olmalıdır.

$ProtoDir = "proto/product"
$ProtoFile = "product.proto"

Write-Host "?? gRPC Kodları Üretiliyor..." -ForegroundColor Cyan

# protoc komutunu çalıstır
protoc --go_out=. --go_opt=paths=source_relative `
       --go-grpc_out=. --go-grpc_opt=paths=source_relative `
       $ProtoDir/$ProtoFile

if ($LASTEXITCODE -eq 0) {
    Write-Host "?? Başarılı! gRPC dosyaları güncellendi." -ForegroundColor Green
} else {
    Write-Host "?? Hata Oluştu! 'protoc' yüklü mü kontrol et." -ForegroundColor Red
}
