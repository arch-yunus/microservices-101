<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices 101 Banner" />

  # ?? Microservices 101: Mimari Manifesto & Teknoloji Rehberi
  ### Stratejik Karar Verme ve Sistem Tasarm Uzmanl
  
  [![License](https://img.shields.io/github/license/arch-yunus/microservices-101?style=for-the-badge&color=blue&logo=github)](LICENSE)
  [![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![Architecture](https://img.shields.io/badge/Arch-Clean--Microservices-FF4B11?style=for-the-badge&logo=architecture&logoColor=white)](#-mimari-görünüm)
  [![Status](https://img.shields.io/badge/Status-Ultimate--Master--Guide-red?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)

  **"Sadece kodlamayı değil, daltk dnyanın karmaııklığını nasıl yoneteceini oren."**

  ---
</div>

## ?? Mikroservis Mimarisi: Kökenler ve Evrim

Mikroservis, bir gecede ortaya cıkmıs bir fantezi deildir; monolitik sistemlerin "leklenemezlik" krizine bir cevaptır.

### ?? Tarihçe (Chronology)
- **2011 (Venice Workshop):** "Micro-services" terimi ilk kez bir yazlm mimarları toplantısında telaffuz edildi.
- **2012 (James Lewis):** Bu mimariyi "Microservices" olarak adlandırıp Case Study olarak sundu.
- **2014 (Martin Fowler & James Lewis):** Mikroservislerin altın çaını baslatan meşhur makale yayınlandı. Monolitik devir resmen sarsıldı.
- **2015+ (Cloud Era):** Docker ve Kubernetes'in yukselisiyle mikroservisler endstri standardı haline geldi.

---

## ?? Devlerin Savaşı: Teknoloji Karşılaştırmaları

Mikroservis dünyasında "En İyi" yoktur, "İhtiyaca En Uygun" vardır. İşte kritik savaslar:

### 1. Haberleşme: gRPC vs REST
| Ozellik | gRPC (Google RPC) | REST (HTTP/1.1) | Winner? |
| :--- | :--- | :--- | :--- |
| **Hız** | ?? Ultra Hızlı (Binary) | ?? Orta (Text-based) | **gRPC** |
| **Format** | Protocol Buffers | JSON / XML | **gRPC** |
| **Kullanım** | Servisler Arası (İç) | Client-Server (Dış) | **DURUMA GORE** |
| **Browser** | Zayıf Destek | Tam Destek | **REST** |

### 2. Mesajlama: Kafka vs RabbitMQ
| Ozellik | Apache Kafka | RabbitMQ | Winner? |
| :--- | :--- | :--- | :--- |
| **Trafik** | ?? Milyarlarca Mesaj | ?? Yuksek Trafik | **Kafka** |
| **Mantık** | Log-based (Stream) | Queue-based (Kuyruk) | **Berabere** |
| **Yapı** | Pub/Sub (Olay Odaklı) | Akıllı Yonlendirme | **RabbitMQ** |
| **Kullanım** | Big Data / Event Store | Task Queue / RPC | **DURUMA GORE** |

### 3. Orkestrasyon: Kubernetes vs Docker Swarm
- **Kubernetes (K8s):** Karmaıktır ama her şeyi yonetir. (Google'ın mimarisi). **Büyük sistemler iin tek tercih.**
- **Docker Swarm:** Basittir, ogrenmesi kolaydır. **Kucuk/Orta olcekli projeler iin ideal.**
- **HashiCorp Nomad:** K8s'e gore cok daha hafiftir, sadece "Job" odaklıdır.

### 4. Veritabanı: SQL vs NoSQL
- **PostgreSQL (SQL):** İlişkisel veri, katı tutarlılık (ACID). (Order, User, Payment servisleri).
- **MongoDB / Cassandra (NoSQL):** Yatay lekleme, esnek şema. (Product Catalog, Log, Tracking servisleri).

---

## ?? Reponun Amacı & Misyonu

Bu depo (microservices-101), senin için sadece bir "Tutorial" deildir. Amacımız:
1.  **Doru Silahı Seç:** Hangi teknolojiyi ne zaman kullanman gerektiğini ogretmek.
2.  **Maliyet Optimize Et:** Gereksiz yere Kafka kullanıp server faturasını artırmanı engellemek.
3.  **Elite Architect Yetistir:** Sektorde "Her iie mikroservis yapalım" diyen deil, "Burada monolit daha iyi olur" diyebilecek bilinçte mimarlar yetistirmek.

---

## ?? Mimari Görünüm (Strategy Map)

```mermaid
graph LR
    subgraph Clients
        App["?? Mobile/Web"]
    end

    subgraph API_Edge
        Gateway["?? API Gateway <br/>(REST)"]
    end

    subgraph Internal_HighSpeed
        Auth["?? Security <br/>(gRPC)"]
        Order["?? Order Svc <br/>(gRPC)"]
    end

    subgraph Event_Driven
        Broker["?? Event Bus <br/>(Kafka/RabbitMQ)"]
    end

    App -->|JSON/REST| Gateway
    Gateway -->|Proto/gRPC| Auth
    Gateway -->|Proto/gRPC| Order
    Order -.->|Fire Event| Broker
    
    style Gateway fill:#f9f,stroke:#333
    style Broker fill:#ff9,stroke:#333
    style Internal_HighSpeed fill:#e1f5fe,stroke:#01579b
```

---

<div align="center">
  <sub>Elite Microservices Architect Journey ?? <b>arch-yunus</b></sub>
</div>
