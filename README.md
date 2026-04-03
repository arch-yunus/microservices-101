<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices 101 Banner" />

  # ?? Microservices 101: Mimari Manifesto & Master Class
  ### Daıtık Sistemlerin Kusursuz Yonetimi ve "Elite" Tasarım Kalıpları
  
  [![License](https://img.shields.io/github/license/arch-yunus/microservices-101?style=for-the-badge&color=blue&logo=github)](LICENSE)
  [![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![Architecture](https://img.shields.io/badge/Arch-Clean--Microservices-FF4B11?style=for-the-badge&logo=architecture&logoColor=white)](#-mimari-görünüm)
  [![Status](https://img.shields.io/badge/Status-Ultimate--Master--Guide-red?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)

  **"Kompleksiteyi parala, otonomiyi fethet. Gelecei elite mimarlar ina eder."**

  [? Yol Haritas](#-yol-haritas) • [?? Teknik Derin Dalı](#-teknik-derin-dal) • [?? Mimari](#-mimari-görünüm) • [?? Basit Anlatım](docs/KOLAY-ANLATIM.md) • [?? İleri Seviye Patternlar](#-ileri-seviye-tasarm-kalplar)

  ---
</div>

## ?? Vizyon & Felsefe

Bu repo, modern daltık sistemlerin sadece kodlanmasını değil, bir **mhendislik harikas** olarak nasıl tasarlanacaını "Elite" seviyede ele alır. Mikroservislerin "hızlı daltım", "yuksek leklebilirlik" ve "hata dayankllıı" (Fault Tolerance) vaatlerini gerçeğe dnştrmek iin gereken tum cephanelii burada bulacaksın.

---

## ?? Teknik Derin Dalı (Elite Deep Dives)

Modern bir mikroservis sisteminin "Anayasas" olan kritik kavramları aşağıda detaylandırdım.

<details>
<summary><b>?? 1. Conway Kanunu ve Ters Conway Manevrası</b></summary>
<br/>
Yazlm mimarisi, organizasyonunuzun iletisim yapısının yansmasıdır. **Inverse Conway Maneuver** kullanarak, istediiniz mimariyi elde etmek iin once ekiplerinizi (Teams) o mimariye uygun hale getirmelisiniz. Otonom ekipler = Otonom servisler.
</details>

<details>
<summary><b>?? 2. Migration: Strangler Fig Pattern</b></summary>
<br/>
Bir Monolith'i mikroservise cevirmek iin her şeyi bir anda yıkmak intihardır. Bunun yerine "Strangler Fig" (Boan İncir) pattern'ı kullanılır. Yeni ozellikler mikroservis olarak yazılır, eski Monolith'in paraları yava yava mikroservislere aktarılarak Monolith sonunda "bolar" ve yok edilir.
</details>

<details>
<summary><b>?? 3. Data Patterns: CQRS & Event Sourcing</b></summary>
<br/>
Servislerin veri tabanlarını ayırmak yetmez. Okuma ve Yazma islemlerini (CQRS) paralamak, sistemin performansını 10 kat artırabilir. Event Sourcing ile verinin "en son halini" değil, "oluşum hikayesini" tutarız. Bu sayede geriye donuk tum hataları replay ederek bulabiliriz.
</details>

<details>
<summary><b>?? 4. Güvenlik: mTLS & Zero-Trust Mesh</b></summary>
<br/>
Sadece kapıda (Gateway) gvenlik yapmak yeterli deildir. Her servisin birbiriyle sertifika (mTLS) uzerinden konusması ve "İçerideki kimseye guvenme" mantııyla hareket edilmesi (Zero Trust) zorunludur.
</details>

<details>
<summary><b>?? 5. Resilience: Rate Limiting & Circuit Breaker</b></summary>
<br/>
Sistemi asırı ykten (DDoS veya hatalı clientlar) korumak iin **Token Bucket** veya **Leaky Bucket** algoritmalarıyla Rate Limiting uygulanmalıdır. Patlayan bir servise gıden yolu kapatan Circuit Breaker ise sistemin "kendini onarma" (Self-Healing) yeteneğini sağlar.
</details>

---

## ?? İleri Seviye Tasarım Kalıpları (The Architect's Toolkit)

Mikroservis sistemlerinde hayati onem tasıyan pattern'lar katalou:

- **Sidecar Pattern:** Loglama, metrik ve gvenlik gibi "ortak" iileri bir container yanına (Sidecar) devretmek. (Kodunuz kirlenmez!).
- **Ambassador Pattern:** Servis dısı iletisimleri (Retry, Logging) yoneten bir vekil sunucu kullanımı.
- **Anti-Corruption Layer (ACL):** Eski (Legacy) bir sistemle yeni mikroservisi konusurken, eski sistemin "kirli" verisinin yeni sistemimizi bozmaması iin araya koyduumuz tercme katmanı.
- **Bulkhead Pattern:** Bir geminin paraları gibi, bir servis cokyorsa digerlerine yayılmaması iin kaynakları (Thread, CPU) izole etmek.

---

## ?? Altyapı & DevEx (Developer Experience)

Mikroservis yazmak, sadece kod yazmak deildir; altyapıyı da kodlamaktır (Infrastructure as Code).
- **IaC (Terraform):** Tum sunucuların kodla yonetilmesi.
- **Service Mesh (Istio):** Servisler arası iletisimi kod yazmadan yonetme katmanı.
- **Local Dev (Telepresence):** lokalinizde calısırken kendinizi Kubernetes cluster'ının bir parçası gibi hissetmenizi salayan elite araclar.

---

## ?? Mimari Görünüm (High-Level Design)

```mermaid
%%{init: {'theme': 'dark', 'themeVariables': { 'mainBkg': '#1a1a1a', 'nodeBorder': '#555' }}}%%
graph TD
    subgraph UI_Layer
        Web["?? Web App"]
        Mobile["?? Mobile App"]
    end

    subgraph Entry_Layer
        Edge["?? Edge Gateway <br/><small>OAuth2 / Rate Limit</small>"]
    end

    subgraph Control_Plane
        Mesh["?? Service Mesh Control <br/><small>Istio / Linkerd</small>"]
    end

    subgraph Business_Execution
        Auth["?? Identity <br/><small>Go / Vault</small>"]
        Catalog["?? Catalog Svc <br/><small>Go / MongoDB</small>"]
        Order["?? Order Svc <br/><small>Go / PostgreSQL</small>"]
        Notify["?? Push Svc <br/><small>Node.js / Redis</small>"]
    end

    subgraph Messaging_Backbone
        Kafka["?? Message Broker <br/><small>Kafka Clusters</small>"]
    end

    Web & Mobile -->|HTTPS| Edge
    Edge --> Mesh
    Mesh --> Auth & Catalog & Order
    Order -.->|Sync Event| Kafka
    Kafka -.->|Async Trigger| Notify
    
    style Edge fill:#7b1fa2,stroke:#fff,color:#fff
    style Kafka fill:#ffb300,stroke:#333
    style Auth fill:#2e7d32,stroke:#fff,color:#fff
    style Order fill:#2e7d32,stroke:#fff,color:#fff
    style Catalog fill:#2e7d32,stroke:#fff,color:#fff
    style Mesh fill:#0277bd,stroke:#fff,color:#fff
```

---

## ?? Eğitim Yol Haritas (Elite Roadmap)

Seni bir sistem mimarına dnsturecek progress tablosu:

| Aşama | Modl | Odak Noktası | Durum |
| :--- | :--- | :--- | :---: |
| ?? **Faz 1** | [Giris](docs/01-intro/README.md) | Paradigma Deıişimi & Temeller | ![Done](https://img.shields.io/badge/-Tamamland-success?style=flat-square) |
| ?? **Faz 2** | [Decomposition](docs/02-decomposition/README.md) | DDD & Servis Parçalama | ![Done](https://img.shields.io/badge/-Tamamland-success?style=flat-square) |
| ?? **Faz 3** | [Communication](docs/03-communication/README.md) | gRPC & Event-Driven Archi. | ![Done](https://img.shields.io/badge/-Tamamland-success?style=flat-square) |
| ?? **Faz 4** | [Data Management](docs/04-data-management/README.md) | Saga Pattern & CQRS | ![Done](https://img.shields.io/badge/-Tamamland-success?style=flat-square) |
| ?? **Faz 5** | API Gateway | Security, Rate Limiting & Auth | ![Planned](https://img.shields.io/badge/-Yaknda-orange?style=flat-square) |
| ?? **Faz 6** | Observability | Distributed Tracing & Metrics | ![Planned](https://img.shields.io/badge/-Yaknda-orange?style=flat-square) |
| ?? **Faz 7** | Cloud Native | Docker, K8s & GitOps | ![Planned](https://img.shields.io/badge/-Yaknda-orange?style=flat-square) |

---

## ?? Neden Go (Golang)?

Neden mikroservis dnyasının lider dili Go?
- **Sıfır Baımlılık:** Statik binary dosyaları ile Docker container'larını sniler içinde ayağa kaldırın.
- **Eszamanllk (Concurrency):** Go kanalları ve routine'leri ile binlerce request'i en duuk CPU ile yonetin.

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
