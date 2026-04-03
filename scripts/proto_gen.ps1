# ?? microservices-101: Proto Generation Script (Windows/PowerShell)
# Not: protoc, protoc-gen-go ve protoc-gen-go-grpc kurulu olmaldr.

$ProtoDir = "proto/product"
$ProtoFile = "product.proto"

Write-Host "?? gRPC Kodlar Üretiliyor..." -ForegroundColor Cyan

# protoc komutunu çaltır
protoc --go_out=. --go_opt=paths=source_relative `
       --go-grpc_out=. --go-grpc_opt=paths=source_relative `
       $ProtoDir/$ProtoFile

if ($LASTEXITCODE -eq 0) {
    Write-Host "?? Bajarılı! gRPC dosyalar güncellendi." -ForegroundColor Green
} else {
    Write-Host "?? Hata Oluştu! 'protoc' yüklü mü kontrol et." -ForegroundColor Red
}
