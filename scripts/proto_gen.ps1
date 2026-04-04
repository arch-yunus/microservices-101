# ?? microservices-101: Proto Kod Üretim Betii (Windows/PowerShell)
# Not: protoc, protoc-gen-go ve protoc-gen-go-grpc kurulu olmalıdır.

# List of proto files to generate
$Protos = @(
    @{ Dir = "proto/product"; File = "product.proto" },
    @{ Dir = "proto/order"; File = "order.proto" }
)

Write-Host "?? gRPC Kodları Üretiliyor..." -ForegroundColor Cyan

foreach ($Proto in $Protos) {
    Write-Host "--- $($Proto.File) İşleniyor ---" -ForegroundColor Yellow
    protoc --go_out=. --go_opt=paths=source_relative `
           --go-grpc_out=. --go-grpc_opt=paths=source_relative `
           "$($Proto.Dir)/$($Proto.File)"
}

if ($LASTEXITCODE -eq 0) {
    Write-Host "?? Başarılı! Tüm gRPC dosyaları güncellendi." -ForegroundColor Green
} else {
    Write-Host "?? Hata Oluştu! 'protoc' yüklü mü kontrol et." -ForegroundColor Red
}
