<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices Architecture Banner" />

  # Microservices Architecture: Engineering Handbook 🏗️
  ### Modern Dağıtık Sistemler Tasarımı ve Uygulama Rehberi
  
  [![Architecture](https://img.shields.io/badge/Architecture-Microservices-6366f1?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)
  [![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![License](https://img.shields.io/badge/License-MIT-success?style=for-the-badge)](LICENSE)

  **"Dağıtık sistemlerin karmaşıklığını yönetmek, sadece kod yazmak değil; doğru mimari prensipleri uygulamakla başlar."**

  ---
</div>

## 📌 Proje Hakkında

Bu depo, modern mikroservis mimarisi prensiplerini uçtan uca uygulamalı bir şekilde öğretmek amacıyla tasarlanmış profesyonel bir eğitim platformudur. Proje; **Go dili**, **gRPC**, **Event-Driven Architecture (EDA)** ve **Cloud-Native** araçlar kullanarak, gerçek dünya senaryolarına dayanan bir e-ticaret altyapısını simüle eder.

### 🛠️ Temel Teknolojiler
- **Dil**: Go (Golang)
- **Haberleşme**: gRPC (Synchronous), RabbitMQ (Asynchronous)
- **Veri Yönetimi**: PostgreSQL, Redis
- **Altyapı**: Docker, Docker Compose
- **Güvenlik**: API Gateway & JWT Authentication

---

## 🗺️ Öğrenim Müfredatı (Learning Roadmap)

Müfredat, bir mühendisin monolitik yapıdan dağıtık sistemlere geçiş sürecini kapsayan 5 ana modülden oluşmaktadır:

| Seviye | Modül | Teknik Kapsam | Durum |
| :--- | :--- | :--- | :---: |
| 🏗️ | **Modül 1: Temeller** | Clean Architecture, PostgreSQL, UUID | ✅ |
| 🔌 | **Modül 2: Haberleşme** | gRPC (Product <-> Order) | ✅ |
| 🛡️ | **Modül 3: Güvenlik** | API Gateway (Echo), Path Routing | ✅ |
| 📩 | **Modül 4: Event-Driven** | RabbitMQ (Order Created -> Notifier) | ✅ |
| 🐒 | **Modül 5: Dayanıklılık** | Circuit Breaker & Retry Logic | 🏗️ |

---

## 🚀 Başlangıç Rehberi

### Ön Gereksinimler
Sistemi çalıştırmadan önce aşağıdaki araçların kurulu olduğundan emin olun:
- [Docker & Docker Compose](https://docs.docker.com/get-docker/)
- [Go 1.21+](https://go.dev/dl/)
- Make (Opsiyonel, ancak önerilir)

### Hızlı Kurulum
Tüm altyapıyı (Veritabanları, Message Broker, Redis) tek bir komutla ayağa kaldırın:
```bash
# Projeyi klonlayın
git clone https://github.com/arch-yunus/microservices-101.git
cd microservices-101

# Altyapıyı başlatın
make up
```

---

## 📚 Dokümantasyon ve Derinlemesine Bakış

Teorik altyapıyı güçlendirmek için `docs/` dizinindeki teknik makaleleri inceleyebilirsiniz:

- [**00 - Dağıtık Sistem Temelleri**](./docs/01-intro/README.md): Monolith vs. Microservice karşılaştırması.
- [**01 - Mimari Tasarım**](./docs/02-decomposition/README.md): Bounded Context ve Servis Ayrıştırma.
- [**02 - Servisler Arası İletişim**](./docs/03-communication/README.md): gRPC ve Protobuf derinlemesine inceleme.
- [**03 - API Gateway Tasarımı**](./docs/05-api-gateway/README.md): Neden Gateway kullanmalıyız?
- [**04 - Event-Driven Architecture**](./docs/06-messaging/README.md): RabbitMQ ve Mesaj Kuyrukları.

---

## 📈 Katkıda Bulunma

Bu proje açık kaynaklı bir topluluk girişimidir. Geliştirme sürecine katkıda bulunmak, hata raporlamak veya yeni özellikler önermek için lütfen [CONTRIBUTING.md](./CONTRIBUTING.md) dosyasını inceleyin.

<div align="center">
  <br/>
  <sub>Engineering Excellence Series | **© 2026 arch-yunus**</sub>
</div>
