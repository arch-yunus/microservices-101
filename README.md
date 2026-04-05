<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices Architecture Banner" />

  # Microservices Architecture: The Engineering Tome (Infinity Edition) 📜
  ### Mimari Egemenlik, Dijital Ekonomi ve Küresel Dağıtık Sistemler Mühendisliği
  
  [![Architecture](https://img.shields.io/badge/Architecture-Microservices-6366f1?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)
  [![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![OpenTelemetry](https://img.shields.io/badge/Observability-OpenTelemetry-f59e0b?style=for-the-badge)](https://opentelemetry.io)
  [![Chaos Engineering](https://img.shields.io/badge/Resilience-Chaos_Engineering-000000?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)

  **"Dağıtık sistemlerin sonsuz entropisi karşısında, mimari egemenlik sadece bir mühendislik tercihi değil; bir kurumun dijital dünyada hayatta kalma ve ölçeklenme iradesidir. Bu depo, monolitik sistemlerin yapısal prangalarından kurtulup, mikroservislerin otonom, sınırsız ölçeklenebilir ve kaos altında bile elastik kalabilen dünyasına atılan en derin ve radikal mühendislik mühürüdür."**

  ---
</div>

## 🌐 Bölüm I: Mimari Egemenliğin Felsefesi ve Dijital Ekonomi

Yazılım dünyası, başlangıcından bu yana karmaşıklığı (complexity) yönetmek için sürekli form değiştirmiştir. Yazılım mimarisi, sadece teknik bir tercih değil, aynı zamanda bir **"Zaman-Maliyet-Egemenlik"** ve **"Ekip Topolojisi"** denklemidir.

### 🏛️ 1. Geleneksel Monolit: Devlerin Yükselişi ve Teknik Borç (The Giant)
Tüm iş mantığının tek bir süreç (process) içinde paketlendiği yapıdır. 
- **Ekonomik Perspektif:** Day 1 maliyeti en düşük mimaridir. Ancak uygulama büyüdükçe **"Yazılım Entropisi"** (Software Entropy) devreye girer. Bu durum, piyasaya çıkış hızını (Time-to-Market) öldüren ve kurumsal çevikliği yok eden bir **Teknik Borç** sarmalına yol açar.
- **Brooks Yasası:** "Gecikmiş bir yazılım projesine daha fazla insan eklemek, onu daha da geciktirir." Mikroservisler, bu yasayı kırmak için ekipler arası bağımlılıkları minimize eden tek gerçek panzehirdir.

### 📡 2. SOA (Service-Oriented Architecture): ESB Prangaları
Mikroservislerin doğrudan atasıdır. Servislerin merkezi bir **Enterprise Service Bus (ESB)** üzerinden konuştuğu yapıdır. 
- **Temel Hata:** SOA'da "Hantal ve Akıllı" bir merkezi yapı (ESB) varken, mikroservis felsefesi **"Akıllı Uç Noktalar ve Basit Boru Hatları"** (Smart Endpoints, Dumb Pipes) yaklaşımını savunur.

### 🧩 3. Mikroservis Devrimi: Otonom Birimlerin Egemenliği (The Swarm)
Bir uygulamayı, her biri kendi veritabanına sahip, bağımsız dağıtılabilen ve belirli bir iş yeteneğini temsil eden otonom hücreler kümesi olarak tasarlamaktır.
- **Reverse Conway Maneuver:** "Sistemleri tasarlayan kurumlar, bu sistemlerin yapısında kendi iletişim yapılarını kopyalarlar." Mikroservisler, bu yasayı tersine çevirerek ekiplerin de dikey olarak bölünmüş (Vertical Slices) otonom hücreler haline gelmesini sağlar.

---

## 🏗️ Bölüm II: "Shared Nothing" Ara Temelleri (Infinity Foundations)

### 1. Bounded Context (DDD & Advanced Context Mapping)
Her servisin kendi semantik dilini (**Ubiquitous Language**) belirlemesidir.
- **Context Mapping Patterns:** **Customer-Supplier**, **Shared Kernel**, **Anti-Corruption Layer (ACL)** ve **Conformist** gibi desenlerle servisler arası sınırların nasıl yönetileceği belirlenir.
- **Event Storming:** İş akışlarını "Domain Events" üzerinden modelleyerek Bounded Context sınırlarını keşfetme tekniğidir.

### 2. Otonomi & Polyglot Architecture
Her servis, kendi veritabanına ve teknoloji yığınına sahip olmalıdır.
- **Polyglot Persistence:** Tek bir veritabanı türüyle her sorunu çözmek yerine (One-size-fits-all), işe en uygun veritabanını (Relational for ACID, Cache for speed, NoSQL for scale) seçme özgürlüğüdür.

### 3. Ölçeklenebilirlik: The AKF Scale Cube
- **X-Axis:** Instance kopyalamak (Cloning).
- **Y-Axis:** Fonksiyonel ayrıştırma (Functional Decomposition).
- **Z-Axis:** Veri bölümleme ve Kiracı İzolasyonu (**Sharding/Tenant Isolation**).
Bu projede biz Y-Axis ölçeklendirmeyi mimariyle, X-Axis'i Docker ile uyguluyoruz.

---

## 📡 Bölüm III: Protokol ve Ağ Egemenliği (Networking Supremacy)

### 1. Senkron: gRPC & Protobuf v3
- **HTTP/2 & Binary Framing:** Veriler ikili (binary) formatta taşınarak ağ bandwidth kullanımını %70'e kadar düşürür.
- **Deadlines & Cancellation Propagation:** Bir istek zamanaşımına uğradığında, bu sinyal tüm servis zinciri boyunca iletilerek gereksiz işlem yükü (Zombie Processes) önlenir.
- **L7 Load Balancing (Envoy/Istio):** İstek bazlı (Request-based) yük dengeleme ile sistemdeki yük en verimli şekilde dağıtılır.

### 2. Asenkron: RabbitMQ & EDA (Event-Driven Architecture)
- **Temporal Decoupling:** Servislerin zaman bazlı bağımlılığını ortadan kaldırır. 
- **Competing Consumers:** Kuyruğu birden fazla instance dinleyerek iş yükünü otomatik paylaşır. **DLX (Dead Letter Exchange)** ve **Message Compaction** ile mesaj güvenliği en üst düzeye çıkarılır.

---

## 💾 Bölüm IV: Dağıtık Sistem Teorisi (World-Class Research)

### 1. PACELC: Gerçek Dünya Seçimleri
Bir sistem bölünmüşken (Partitioned) CAP geçerlidir; normal çalışırken (Else) **Latens (Latency)** ve **Tutarlılık (Consistency)** arasında seçim yapılır.

### 2. Konsensüs ve Mantıksal Saatler (Consensus & Clocks)
Dağıtık sistemlerde "mutlak zaman" yoktur. **Raft** veya **Paxos** gibi konsensüs algoritmaları ve **Lamport Timestamps** (Mantıksal Saatler) kullanılarak olayların sırası (Causal Ordering) belirlenir.

### 3. Saga Deseni: Nihai Tutarlılık
Birden fazla servisi ilgilendiren işlemlerde hatayı geri almak için **Compensating Transactions** kullanılır. **Semantic Locking** ile dağıtık verinin tutarlılığı korunur.

---

## 🛡️ Bölüm V: Edge Intelligence & API Gateway

- **API Composition & Request Collapsing:** Gateway, farklı servislerden gelen verileri birleştirir ve özdeş istekleri tek bir isteğe düşürerek iç sistemleri korur.
- **Load Shedding:** Sistem aşırı yüklendiğinde, Gateway henüz isteği iç servislere iletmeden düşük öncelikli istekleri reddederek çekirdek servislerin çökmesini engeller.
- **Mutual TLS (mTLS):** İç servisler arası iletişimde her iki tarafın da sertifika doğruladığı en üst düzey güvenlik katmanıdır.

---

## 🔍 Bölüm VI: SRE & Modern Gözlemlenebilirlik (Global Scale)

### 1. MELT ve RED Metotları
Metrics, Events, Logs, Traces. Sistemin "ateşini" ölçer, "yolculuğunu" (Tracing) izleriz. **RED Method** (Rate, Errors, Durations) ile servis sağlığı anlık izlenir.

### 2. Error Budget & SLO/SLA
SLA (Service Level Agreement), SLO (Objective) ve SLI (Indicator) arasındaki ilişkiyle sistem başarısı ölçülür. **Hata Bütçesi (Error Budget)** bitene kadar yeni özellikler (features), biterse sadece kararlılık çalışmaları (stability) yapılır.

---

## 🛣️ Bölüm VII: Proje Müfredatı ve Uygulama Aşamaları

Bu depo, yukarıdaki teorik bilgileri pratiğe döken **6 aşamalı bir mühendislik masterclass'ı** sunar:

| Modül | Teknik Derinlik | Kazanımlar | Durum |
| :--- | :--- | :--- | :---: |
| 🏗️ **Modül 1** | **Mimari (DDD)** | Clean Architecture, PostgreSQL, Domain Isolation | ✅ |
| 🔌 **Modül 2** | **İletişim (gRPC)** | Protobuf v3, Binary Framing, Inter-service calls | ✅ |
| 🛡️ **Modül 3** | **Gateway (Edge)** | Fortress Pattern, Routing, Middleware | ✅ |
| 📩 **Modül 4** | **Mesajlaşma (EDA)** | RabbitMQ, Asynchronous Consumers, DLX | ✅ |
| 🐒 **Modül 5** | **Stabilite (Resilience)** | Circuit Breaker, Health Checks, Gobreaker | ✅ |
| 🔍 **Modül 6** | **Gözlem (Observability)** | OTel, Jaeger, Context Propagation | ✅ |

---

## 🛠️ Teknik Sözlük & Gelişmiş Anti-Desenler

### 🚫 Anti-Desenler (Karanlık Yüz)
- **Nana-services:** Aşırı küçük servislerin ağ gecikmesini artırması.
- **Distributed Monolith:** Servislerin birbirine o kadar bağımlı olması ki, tek bir servisi deploy etmenin diğerlerini bozması.
- **Entity Services:** Sadece veritabanı tablolarını temsil edenCRUD servisler.

### 📜 Sözlük
- **Idempotency:** Bir işlemin birden fazla kez uygulanmasının sonucu değiştirmemesi.
- **Infrastructure as Code (IaC):** Altyapının kod olarak (Docker Compose) tanımlanıp tekrar üretilebilir olması.
- **Immutable Infrastructure:** Servis güncellendiğinde yama yapmak yerine kopyanın tamamen değiştirilmesi.

---

## 🚀 Dağıtım (Deployment) & Chaos Engineering Masterclass

### 1. Stratejiler
- **Canary & Shadow Deployment:** Yeni sürümün kademeli açılması veya gerçek trafiğin sessiz kopyalanması.
- **Weighted Traffic Shifting:** Trafiğin yüzdesel olarak (99/1 → 50/50) yeni servislere aktarılması.

### 2. Kaos Mühendisliği (Chaos Engineering)
Sisteme rastgele hatalar (gecikme, servis durdurma) enjekte edilerek sistemin **Blast Radius** (Etki Alanı) analizi ve **Auto-healing** yeteneği test edilir. "Failing is not an option, but recovery is a requirement."

---

## 📚 Kaynakça (The Bibliography)

- *Microservices Patterns* - **Chris Richardson**
- *Designing Data-Intensive Applications* - **Martin Kleppmann**
- *Domain-Driven Design* - **Eric Evans**
- *Building Microservices* - **Sam Newman**
- *Site Reliability Engineering (SRE) Book* - **Google Engineering**

<div align="center">
  <br/>
  <sub>Engineering Excellence Series | **© 2026 arch-yunus** | "Complex systems that work are invariably found to have evolved from simple systems that worked."</sub>
</div>
