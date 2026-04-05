<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices Architecture Banner" />

  # Microservices Architecture: The Engineering Bible (Absolute Masterclass) 📚
  ### Üst Düzey Dağıtık Sistemler, Dijital Ekonomi ve Yazılım Egemenliği Ansiklopedisi
  
  [![Architecture](https://img.shields.io/badge/Architecture-Microservices-6366f1?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)
  [![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![OpenTelemetry](https://img.shields.io/badge/Observability-OpenTelemetry-f59e0b?style=for-the-badge)](https://opentelemetry.io)
  [![Chaos Engineering](https://img.shields.io/badge/Resilience-Chaos_Engineering-000000?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)

  **"Dağıtık sistemlerin kaotik doğası karşısında, mimari disiplin sadece bir tercih değil; bir hayatta kalma ve egemenlik stratejisidir. Bu depo, monolitik sistemlerin hantallığından kurtulup, mikroservislerin otonom, sınırsız ölçeklenebilir ve hata toleranslı dünyasına mühendislik perspektifiyle atılan en radikal ve kapsamlı adımdır."**

  ---
</div>

## 🌐 Bölüm I: Yazılım Mimarilerinin Ekonomik ve Teknik Evrimi

Yazılım dünyası, başlangıcından bu yana karmaşıklığı (complexity) yönetmek için sürekli form değiştirmiştir. Yazılım mimarisi, sadece teknik bir tercih değil, aynı zamanda bir **"Zaman-Maliyet-Egemenlik"** denklemidir. Mikroservisler, bu evrimin en olgun, ancak operasyonel maliyeti (Overhead) en yüksek halkasıdır.

### 🏛️ 1. Geleneksel Monolit: Devlerin Yükselişi (The Giant)
Tüm iş mantığının tek bir süreç (process) içinde paketlendiği yapıdır. 
- **Ekonomik Perspektif:** Day 1 maliyeti en düşük mimaridir. Ancak uygulama büyüdükçe **"Yazılım Entropisi"** devreye girer. Kod tabanı devasa bir yığın (Big Ball of Mud) haline gelir. Bu durum, piyasaya çıkış hızını (Time-to-Market) öldüren ve kurumsal çevikliği (Agility) yok eden bir **Teknik Borç** (Technical Debt) sarmalına yol açar.

### 📡 2. SOA (Service-Oriented Architecture): ESB Hantallığı
Mikroservislerin atasıdır. Servislerin merkezi bir **Enterprise Service Bus (ESB)** üzerinden konuştuğu yapıdır. 
- **Hata:** SOA'da "Akıllı Boru Hattı" (ESB) varken, tüm iş mantığı bu merkezi yapıda boğulurdu. Mikroservis felsefesi ise bunun tam zıttıdır: **"Akıllı Uç Noktalar ve Basit Boru Hatları"** (Smart Endpoints, Dumb Pipes).

### 🧩 3. Mikroservis Devrimi: Otonom Birimlerin Egemenliği (The Swarm)
Bir uygulamayı, her biri kendi veritabanına sahip, bağımsız dağıtılabilen ve belirli bir iş yeteneğini temsil eden otonom servisler kümesi olarak tasarlamaktır.
- **Conway Yasası:** Mikroservisler, ekiplerin dikey olarak bölünmüş (Vertical Slices) otonom hücreler haline gelmesini sağlayarak kurumsal egemenliği (Sovereignty) artırır.

---

## 🏗️ Bölüm II: Dağıtık Sistemlerin 4 Temel Sütunu (Ultra Masterclass)

### 1. Bounded Context (DDD & Event Storming)
Her servisin kendi semantik dilini (**Ubiquitous Language**) belirlemesidir.
- **Event Storming:** İş akışlarını "Domain Events" üzerinden modelleyerek Bounded Context sınırlarını keşfetme tekniğidir. Bu projede her servis, kendi bağlamında (Context) bir "Egemenlik Alanı" oluşturur.

### 2. Otonomi & Polyglot Architecture
Her servis, kendi veritabanına ve teknoloji yığınına sahip olmalıdır.
- **Polyglot Persistence:** "One size fits all" yaklaşımını reddederek, işe en uygun veritabanını (Postgres, Redis, Mongo vb.) seçme özgürlüğüdür. Paylaşımlı veritabanı otonomiyi öldüren en büyük anti-desendir.

### 3. Ölçeklenebilirlik: AKF Scale Cube
- **X-Axis:** Instance kopyalamak (Cloning).
- **Y-Axis:** Fonksiyonel ayrıştırma (Functional Decomposition).
- **Z-Axis:** Veri bölümleme (Sharding/Tenant Isolation).
Bu projede biz Y-Axis ölçeklendirmeyi mimariyle, X-Axis'i Docker ile uyguluyoruz.

### 4. Hata Toleransı: Cascade Failures & Chaos Engineering
Dağıtık sistemlerde ağ güvenilmezdir. Önemli olan hatanın bir domino taşı gibi yayılmasını önlemektir (**The Titanic Effect**). **Circuit Breaker** (Devre Kesici), açık/kapalı/yarı-açık durumlarıyla sistemin geri kalanını korur.

---

## 📡 Bölüm III: Protokol Mühendisliği & Ağ Katmanı

### 1. Senkron: gRPC & Protobuf v3
- **HTTP/2 & Binary Framing:** Veriler JSON gibi metin değil, ikili formatta taşınır. Bandwidth kullanımını %50-70 düşürür.
- **Deadlines & Cancellation Propagation:** Bir istek iptal edildiğinde veya zamanaşımına (timeout) uğradığında, bu sinyal tüm servis zinciri boyunca iletilerek gereksiz işlem yükü önlenir.

### 2. Asenkron: RabbitMQ & EDA (Event-Driven Architecture)
- **Temporal Decoupling:** Servislerin zaman bazlı bağımlılığını ortadan kaldırır. 
- **Competing Consumers:** Kuyruğu birden fazla instance dinleyerek iş yükünü otomatik paylaşır. **Manual ACKs** ve **DLX (Dead Letter Exchange)** ile mesaj kaybı imkansız hale getirilir.

---

## 💾 Bölüm IV: Dağıtık Sistem Teorisi (The PhD Masterclass)

### 1. CAP Teoremi vs. PACELC
PACELC, modern sistemlerin asıl rehberidir. Bir sistem bölünmüşken (Partitioned) CAP geçerlidir; normal çalışırken (Else) **Latens (Latency)** ve **Tutarlılık (Consistency)** arasında seçim yapılır.

### 2. Saat Senkronizasyonu (Clock Synchronization)
Dağıtık sistemlerde "mutlak zaman" yoktur. **Mantıksal Saatler** (Logical Clocks) ve **Lamport Timestamps** kullanılarak olayların sırası (Causal Ordering) belirlenir.

### 3. Saga Deseni: Nihai Tutarlılık
Birden fazla servisi ilgilendiren işlemlerde hatayı geri almak için **Telafi Edici İşlemler** (Compensating Transactions) kullanılır. **Semantic Locking** ile dağıtık verinin tutarlılığı korunur.

---

## 🛡️ Bölüm V: Edge Intelligence & API Gateway (The Fortress)

- **API Composition:** Farklı servislerden gelen verilerin Gateway seviyesinde birleştirilip istemciye tek bir cevap olarak dönülmesidir.
- **Request Collapsing:** Aynı anda gelen özdeş isteklerin tek bir isteğe düşürülerek iç sistemlerin korunmasıdır.
- **Mutual TLS (mTLS):** İç servisler arası iletişimde her iki tarafın da birbirinin sertifikasını doğruladığı en üst düzey güvenlik katmanıdır.

---

## 🔍 Bölüm VI: SRE & Modern Gözlemlenebilirlik

### 1. Üç Sütun (MELT)
Metrics, Events, Logs, Traces. Sistemin "ateşini" (Metrics) ölçer, "yolculuğunu" (Tracing) izleriz.

### 2. RED ve USE Metotları
- **RED:** Rate, Errors, Durations (Servis sağlığı).
- **USE:** Utilization, Saturation, Errors (Kaynak kullanımı).
**OpenTelemetry** ile bu veriler merkezi bir Jaeger UI üzerinden izlenebilir hale getirilmiştir.

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

## 🛠️ Teknik Sözlük & Anti-Desenler

### 🚫 Anti-Desenler (Karanlık Yüz)
- **Nana-services:** Aşırı küçük servisler, ağ gecikmesini artırır.
- **Distributed Monolith:** Servislerin birbirine aşırı sıkı bağımlı olması.
- **Entity Services:** Sadece veritabanı tablolarını temsil eden CRUD servisler.

### 📜 Sözlük
- **Idempotency:** Bir işlemin birden fazla kez tetiklense bile veritabanında sadece bir kez etkili olması.
- **Infrastructure as Code (IaC):** Altyapının kod olarak tanımlanıp (Docker Compose) tekrar üretilebilir olması.
- **Immutable Infrastructure:** Servis güncellendiğinde yama yapmak yerine kopyanın tamamen değiştirilmesi.

---

## 🚀 Dağıtım (Deployment) & Chaos Engineering

### 1. Stratejiler
- **Canary Release:** Yeni sürümün kademeli olarak açılması.
- **Shadow Deployment:** Gerçek trafiğin sessizce yeni sürüme kopyalanması.
- **Weighted Traffic Shifting:** Trafiğin yüzdesel olarak yeni servislere aktarılması.

### 2. Kaos Mühendisliği (Chaos Engineering)
Sisteme rastgele hatalar (gecikme, servis durdurma) enjekte edilerek sistemin dayanıklılığı ve **Auto-healing** yeteneği test edilir. "Failing is not an option, but recovery is a requirement."

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
