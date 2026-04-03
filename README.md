<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices 101 Banner" />

  # ?? Microservices 101
  ### Sfrdan Datk Sistem Mimarisi Rehberi
  
  ![License](https://img.shields.io/github/license/arch-yunus/microservices-101?style=for-the-badge&color=blue)
  ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
  ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
  ![Kubernetes](https://img.shields.io/badge/kubernetes-%23326ce5.svg?style=for-the-badge&logo=kubernetes&logoColor=white)
  
  *Modern, leklenip ynetilebilir ve dayankl mikroservis sistemleri tasarlamann elite yol haritas.*

  [Roadmap](#-yol-haritas) | [Teknolojiler](#-teknolojiler) | [Mimari](#-mimari) | [Katkda Bulunma](CONTRIBUTING.md)

  ---
</div>

## ?? Vizyon

Bu depo, monolitik yaplarn kstlamalarndan kurtulup moderne gemek isteyen gelitiriciler ve sistem mimarlar iin hazrlanmtr. Sadece kod yazmay deil, bir sistemin **yaam dngsn**, **gvenliğini** ve **izlenebilirliğini** u seviyede ele almay hedefler.

> [!IMPORTANT]
> Mikroservis bir teknoloji deil, bir **stratejidir**. Bu rehberde amacmz sadece kodlamak deil, daltk bir sistemin karmaııklığını nasıl yoneteceğimizi oğrenmektir.

---

## ?? Yol Haritas (Roadmap)

Eğitim sreci, bir sistem mimarnn zihnindeki u ak takip eder:

| Modl | Konu | Durum |
| :--- | :--- | :--- |
| **01** | [Giris: Paradigma Değisimi](docs/01-intro/README.md) | ?? Tamamland |
| **02** | [Servis Parcalama & DDD](docs/02-decomposition/README.md) | ?? Tamamland |
| **03** | [Haberlesme Protokolleri](docs/03-communication/README.md) | ?? Tamamland |
| **04** | [Veri Yönetimi & Tutarllk](docs/04-data-management/README.md) | ?? Tamamland |
| **05** | API Gateway & Security | ?? Yaknda |
| **06** | Observability (Tracing & Metrics) | ?? Yaknda |
| **07** | CI/CD & Deployment (K8s) | ?? Yaknda |

---

## ?? Mimari Görünüm

Sistemin temel akını aşağıda görebilirsiniz:

```mermaid
graph TD
    %% Ana Nodes
    Client["?? Web/Mobile Client"] 
    Gateway["?? API Gateway <br/>(Kong / Ocelot)"]
    Broker["?? Message Broker <br/>(RabbitMQ)"]
    
    %% Services
    Auth["?? Auth Service <br/>(Go)"]
    Order["?? Order Service <br/>(Go)"]
    Product["?? Product Service <br/>(Go)"]
    Notification["?? Notify Service <br/>(Go)"]
    
    %% Connections
    Client -->|HTTPS| Gateway
    Gateway -->|gRPC| Auth
    Gateway -->|REST/gRPC| Order
    Gateway -->|REST/gRPC| Product
    
    Order -.->|Event| Broker
    Broker -.->|Consume| Notification
    
    %% Styling
    style Gateway fill:#f9f,stroke:#333,stroke-width:2px
    style Broker fill:#ff9,stroke:#333,stroke-width:2px
    style Auth fill:#9f9,stroke:#333,stroke-width:1px
    style Order fill:#9f9,stroke:#333,stroke-width:1px
    style Product fill:#9f9,stroke:#333,stroke-width:1px
    style Notification fill:#9f9,stroke:#333,stroke-width:1px
```

---

## ?? Teknolojiler & Araclar

| Kategori | Arac | Badge |
| :--- | :--- | :--- |
| **Dil** | Go (Golang) | ![Go](https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white) |
| **İletişim** | gRPC / REST | ![gRPC](https://img.shields.io/badge/gRPC-4285F4?style=flat-square&logo=google&logoColor=white) |
| **Gateway** | Kong / Ocelot | ![Kong](https://img.shields.io/badge/Kong-003459?style=flat-square&logo=kong&logoColor=white) |
| **Veritabanı** | PostgreSQL / Redis | ![Postgres](https://img.shields.io/badge/PostgreSQL-316192?style=flat-square&logo=postgresql&logoColor=white) |
| **Broker** | RabbitMQ | ![RabbitMQ](https://img.shields.io/badge/RabbitMQ-FF6600?style=flat-square&logo=rabbitmq&logoColor=white) |
| **Gözlem** | Prometheus | ![Prometheus](https://img.shields.io/badge/Prometheus-E6522C?style=flat-square&logo=prometheus&logoColor=white) |

---

## ?? Baslangıc

Projeyi yerel ortamnzda ayağa kaldrmak iin:

```bash
# Repoyu klonlayn
git clone https://github.com/arch-yunus/microservices-101.git

# Altyap servislerini başlatın
docker-compose up -d
```

---

## ?? Lisans
Bu proje [MIT](LICENSE) lisansı ile korunmaktadr.

<div align="center">
  <sub>Built with ?? by <b>arch-yunus</b></sub>
</div>
