<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices 101 Banner" />

  # ?? Microservices 101: Mimari Manifesto
  ### Daıtık Sistemlerin Elite Rehberi & Pratik Uygulama Kampı
  
  [![License](https://img.shields.io/github/license/arch-yunus/microservices-101?style=for-the-badge&color=blue&logo=github)](LICENSE)
  [![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com)
  [![Architecture](https://img.shields.io/badge/Arch-Clean--Microservices-FF4B11?style=for-the-badge&logo=architecture&logoColor=white)](#-mimari-görünüm)

  **"Geleceğin dünyasını kucuk, bamsz ve otonom paralarla ina ediyoruz."**

  [? Yol Haritas](#-yol-haritas) • [?? Teknik Derin Dalı](#-teknik-derin-dal) • [?? Mimari](#-mimari-görünüm) • [?? Basit Anlatım](docs/KOLAY-ANLATIM.md) • [?? Katkıda Bulunma](CONTRIBUTING.md)

  ---
</div>

## ?? Vizyon & Felsefe

Bu depo, monolitik yapıların hantallığından kurtulup, bulut-yerli (cloud-native) bir sistem mimarı olma yolunda ilerlemek isteyenler için tasarlanmış **Elite** bir eğitim serüvenidir. 

> [!IMPORTANT]
> Mikroservis bir teknoloji seçimi değildir; bir **organizasyonel strateji** ve **mühendislik disiplini**dir. Bu depoda sadece kod yazmıyoruz, sistem tasarlıyoruz.

---

## ?? Teknik Derin Dalı (Deep Dives)

Aşağıdaki başlıklar, modern mikroservis mimarisinin can damarıdır. Her birini detaylıca incelemek için tıkla kanka. ??

<details>
<summary><b>?? 1. Conway Kanunu & Organizasyonel Yapı</b></summary>
<br/>
Conway Kanununa göre, yazlm mimarisi organizasyonun iletisim yapısını yansıtır. Mikroservis baısı iin ekibinizi "Two-Pizza Teams" (6-8 kisilik otonom ekipler) haline getirmelisiniz. Otonom ekipler yoksa, otonom servisler de yoktur.
</details>

<details>
<summary><b>?? 2. 12-Factor App Prensipleri</b></summary>
<br/>
Modern bir mikroservisin "Cloud-Native" olması iin u kurallara uyması gerekir. siferelerin Environment Variables'da tutulması (Config), uygulamanın "Stateless" olması (Processes) ve logların birer olay akıı olarak ele alınması (Logs) bu anayasanın temelidir.
</details>

<details>
<summary><b>?? 3. Veri Yönetimi: Database per Service & CQRS</b></summary>
<br/>
Her servisin kendi veritabanı olmalıdır. Yazma (Command) ve Okuma (Query) islemlerini paralamak (CQRS) yuksek trafikli sistemlerde hayat kurtarır. Veri tutarlılıı iin Event-Driven Communication kullanıyoruz.
</details>

<details>
<summary><b>?? 4. Saga Pattern & Distributed Transactions</b></summary>
<br/>
Daltık bir sistemde atomik transaction yoktur. Bir islem patlarsa, daha onceki islemleri geri alacak "Compensating Transactions" (Telafi İslemleri) devreye girer. Saga, daltık dnyanın ROLLBACK mekanizmasıdır.
</details>

<details>
<summary><b>?? 5. Resilience: Circuit Breaker & Chaos Engineering</b></summary>
<br/>
Servisler ker. Circuit Breaker, patlayan bir servise gıden yolu kapatarak tm sistemin cokmesini kilitler. Chaos Engineering (Chaos Monkey) ile canlı ortamda servisleri kapatıp sistemin kendini iyilestirmesini test ederiz.
</details>

---

## ?? Eğitim Yol Haritas (Roadmap)

Seni bir sistem mimarına dnsturecek progress tablosu:

| Aşama | Modl | Odak Noktası | Durum |
| :--- | :--- | :--- | :---: |
| ?? **Faz 1** | [Giris](docs/01-intro/README.md) | Paradigma Deıişimi & Temeller | ![Done](https://img.shields.io/badge/-Tamamland-success?style=flat-square) |
| ?? **Faz 2** | [Decomposition](docs/02-decomposition/README.md) | DDD & Servis Parçalama | ![Done](https://img.shields.io/badge/-Tamamland-success?style=flat-square) |
| ?? **Faz 3** | [Communication](docs/03-communication/README.md) | gRPC & Event-Driven Archi. | ![Done](https://img.shields.io/badge/-Tamamland-success?style=flat-square) |
| ?? **Faz 4** | [Data Management](docs/04-data-management/README.md) | Saga Pattern & CQRS | ![Done](https://img.shields.io/badge/-Tamamland-success?style=flat-square) |
| ?? **Faz 5** | API Gateway | Security & Rate Limiting | ![Planned](https://img.shields.io/badge/-Yaknda-orange?style=flat-square) |
| ?? **Faz 6** | Observability | Distributed Tracing & Metrics | ![Planned](https://img.shields.io/badge/-Yaknda-orange?style=flat-square) |

---

## ?? Mimari Görünüm

Modern bir daltık sistemin anatomi haritası:

```mermaid
%%{init: {'theme': 'dark', 'themeVariables': { 'mainBkg': '#1a1a1a', 'nodeBorder': '#555' }}}%%
graph TD
    subgraph Users_Space
        Client["?? Client App"]
    end

    subgraph Entry_Gate
        Gateway["?? API Gateway <br/><small>Kong / Gloo</small>"]
    end

    subgraph Logic_Cloud
        Auth["?? Identity <br/><small>Go / JWT</small>"]
        Order["?? Order Svc <br/><small>Go / Saga</small>"]
        Stock["?? Inventory <br/><small>Go / CQRS</small>"]
    end

    subgraph Event_Mesh
        Bus["?? Message Bus <br/><small>Kafka / RabbitMQ</small>"]
    end

    Client -->|HTTPS| Gateway
    Gateway --> Auth
    Gateway --> Order
    Order -.->|Publish Event| Bus
    Bus -.->|Consume| Stock
    
    style Gateway fill:#7b1fa2,stroke:#fff,color:#fff
    style Bus fill:#ffb300,stroke:#333
    style Auth fill:#2e7d32,stroke:#fff,color:#fff
    style Order fill:#2e7d32,stroke:#fff,color:#fff
    style Stock fill:#2e7d32,stroke:#fff,color:#fff
```

---

## ?? Neden Go (Golang)?

Neden mikroservis dnyasının "Lider Dili" Go?
- **Ultra-Lightweight:** saniyeler içinde kalkan containerlar.
- **Concurrency:** Goroutines ile binlerce e-zamanlı islem.
- **Type Safety:** Hata payını minimize eden guclu tip sistemi.

---

## ?? Katkda Bulunma

Bu bir topluluk ve geliim projesidir. Sen de bu manifestoya katkda bulunabilirsin!
?? **[CONTRIBUTING.md](CONTRIBUTING.md)** belgesine goz at.

---

<div align="center">
  <br/>
  <img src="https://img.shields.io/badge/Designed_with-??-red?style=for-the-badge" alt="Built with love" />
  <br/>
  <sub>Mastering Microservices Architecture ?? <b>arch-yunus</b></sub>
</div>
