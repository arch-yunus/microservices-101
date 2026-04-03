<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices 101 Banner" />

  # ?? Microservices 101: Sistem Mimarisinin Nihai Anayasas
  ### Daıtık Sistemlerin Karanlık Dünyasında Stratejik Liderlik
  
  [![License](https://img.shields.io/github/license/arch-yunus/microservices-101?style=for-the-badge&color=blue&logo=github)](LICENSE)
  [![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![Architecture](https://img.shields.io/badge/Arch-Clean--Microservices-FF4B11?style=for-the-badge&logo=architecture&logoColor=white)](#-mimari-görünüm)
  [![Status](https://img.shields.io/badge/Status-Legendary--Architect--Guide-gold?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)

  **"Karmaşıklığı yönetmek bir yetenek, otonomiyi inşa etmek ise bir sanattır."**

  [?? Kolay Anlatım](docs/KOLAY-ANLATIM.md) • [?? Master Class](docs/MASTERCLASS.md) • [? Yol Haritas](#-eğitim-yol-haritas) • [?? Hızlı Başlangı](#-hızlı-başlangıç)

  ---
</div>

## ?? Vizyon & Felsefe: Neden Mikroservis?

Mikroservisler, devasa monolitik yapıların hantallığından kurtulup, saniyeler içinde daltılabilen ve bamsz leklenen otonom hcrelerin inşasıdır. Bu bir teknolojik tercih değil, milyar dolarlık trafiği yonetebilecek bir **organizasyonel strateji**dir. 

> [!IMPORTANT]
> Mikroservis mimarisi "Design for Failure" (Hata iin Tasarla) prensibiyle yaşar. Sistemde her an bir yerlerin çökeceini kabul eder ve sistemi bu cokuste bile ayakta kalacak şekilde zırhlandırırız.

---

## ?? Hızlı Başlangıç (Quick Start)

Sistemi saniyeler iinde ayaa kaldrmak iin terminalinizi hazrlayn:

```bash
# 1. Altyapıyı (DB, Message Broker) başlat
make up

# 2. Product Service'i çalıştır
make run-product

# 3. Order Service'i çalıştır (farklı bir terminalde)
make run-order
```

---

## ?? Teknik Derin Dalı (Architectural Deep Dives)

<details>
<summary><b>?? 1. Conway Kanunu ve Ters Conway Manevrası</b></summary>
<br/>
"Sistemler, tasarlayan organizasyonun iletisim yapısını kopyalar." **Inverse Conway Maneuver** kullanarak, hedeflediiniz mimariyi elde etmek iin once ekiplerinizi (Teams) bu mimariye gore paralamalısınız.
</details>

<details>
<summary><b>?? 2. Migration: Strangler Fig Pattern</b></summary>
<br/>
Bir Monolith'i mikroservise cevirmek iin her şeyi bir anda yıkmak intihardır. Bunun yerine trafik yava yava Monolith'ten yeni servislere akıtılarak Monolith "boğulur" (Strangled) ve yok edilir.
</details>

<details>
<summary><b>?? 3. Data Patterns: CQRS & Event Sourcing</b></summary>
<br/>
Okuma (Query) ve Yazma (Command) islemlerini paralarız. Verinin son halini değil, veriyi o hale getiren tum tarihsel olayları (Events) saklayarak kusursuz bir denetim (Audit) saglarız.
</details>

---

## ?? İleri Seviye Tasarım Kalıpları

- **Anti-Corruption Layer (ACL):** Legacy sistemlerin kirli verisini filtrelemek iin kullanılır.
- **Ambassador Pattern:** Servis dısı iletisimi (Retry, Logging) yoneten bir vekil sunucu.
- **Outbox Pattern:** Veritabanı islemi ve mesaj gonderimini atomik hale getirir.

---

## ?? Mimari Görünüm

```mermaid
%%{init: {'theme': 'dark', 'themeVariables': { 'mainBkg': '#1a1a1a', 'nodeBorder': '#555' }}}%%
graph TD
    subgraph Gateways
        Edge["?? API Gateway <br/><small>Kong / Ocelot <br/> Auth / Rate Limit</small>"]
    end

    subgraph Service_Cloud
        Auth["?? Identity Svc <br/><small>Go / gRPC</small>"]
        Order["?? Order Svc <br/><small>Go / Clean Arch</small>"]
        Catalog["?? Product Svc <br/><small>Go / CQRS</small>"]
    end

    subgraph Messaging
        Kafka["?? Event Bus <br/><small>Kafka / RabbitMQ</small>"]
    end

    Edge --> Auth & Order & Catalog
    Order -.->|Sync Call| Auth
    Order -.->|Publish| Kafka
    Kafka -.->|Consume| Catalog
    
    style Edge fill:#7b1fa2,stroke:#fff,color:#fff
    style Kafka fill:#ffb300,stroke:#333
    style Auth fill:#2e7d32,stroke:#fff,color:#fff
    style Order fill:#2e7d32,stroke:#fff,color:#fff
    style Catalog fill:#2e7d32,stroke:#fff,color:#fff
```

---

## ?? Eğitim Yol Haritas (The Elite Roadmap)

| Aşama | Modl | Odak Noktası | Durum |
| :--- | :--- | :--- | :---: |
| ?? **Faz 1** | [Giris](docs/01-intro/README.md) | Paradigma Deıişimi | ![100%](https://geps.dev/progress/100) |
| ?? **Faz 2** | [Decomposition](docs/02-decomposition/README.md) | DDD & Bounded Context | ![100%](https://geps.dev/progress/100) |
| ?? **Faz 3** | [Communication](docs/03-communication/README.md) | gRPC & REST | ![100%](https://geps.dev/progress/100) |
| ?? **Faz 4** | [Data Management](docs/04-data-management/README.md) | Saga & CQRS | ![100%](https://geps.dev/progress/100) |
| ?? **Faz 5** | API Gateway | Security & OIDC | ![20%](https://geps.dev/progress/20) |
| ?? **Faz 6** | Observability | Tracing & Metrics | ![10%](https://geps.dev/progress/10) |
| ?? **Faz 7** | Cloud Native | Kubernetes & GitOps | ![0%](https://geps.dev/progress/0) |

---

<div align="center">
  <br/>
  <img src="https://img.shields.io/badge/Designed_for-Masters-gold?style=for-the-badge" alt="Designed for Masters" />
  <br/>
  <sub>Achieving Architectural Excellence ?? <b>arch-yunus</b></sub>
</div>
