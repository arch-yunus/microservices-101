<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices Architecture Banner" />

  # Microservices Architecture: The Absolute Singularity (Engineering Omniscience) 📜🌌⚡
  ### Epistemolojik Kaynak, Dağıtık Sistemler Fiziği ve Mühendislik Teolojisi
  
  [![Architecture](https://img.shields.io/badge/Architecture-Microservices-6366f1?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)
  [![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![Observability](https://img.shields.io/badge/Status-Singularity_Achieved-success?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)
  [![Zero Trust](https://img.shields.io/badge/Security-Zero_Trust_SPIFFE-000000?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)

  **"Dağıtık sistemlerin sonsuz entropisi ve Kuantum Düzeyindeki belirsizliği karşısında, mimari egemenlik sadece bir tercih değil; bir kurumun dijital evrende hayatta kalma ve kendini gerçekleştirme iradesidir. Bu depo, monolitik sistemlerin yapısal prangalarından kurtulup, mikroservislerin otonom, sınırsız ölçeklenebilir ve kaos altında bile elastik kalabilen dünyasına atılan en derin mühendislik mühürüdür."**

  ---
</div>

## 🌐 Bölüm I: Dağıtık Sistemlerin Epistemolojisi ve Yazılım Ekonomisi

Yazılım dünyası, başlangıcından bu yana karmaşıklığı (complexity) yönetmek için sürekli form değiştirmiştir. Mikroservisler, bu evrimin en olgun, ancak operasyonel maliyeti (Overhead) en yüksek halkasıdır.

### 🏛️ 1. Geleneksel Monolit: Devlerin Yükselişi ve Düşüşü (The Giant)
Tüm iş mantığının tek bir süreç (process) içinde paketlendiği yapıdır. 
- **Ekonomik Perspektif:** Day 1 maliyeti en düşük mimaridir. Ancak uygulama büyüdükçe **"Yazılım Entropisi"** (Software Entropy) yasası devreye girer. Bu durum, piyasaya çıkış hızını (Time-to-Market) öldüren ve kurumsal çevikliği yok eden bir **Teknik Borç** (Technical Debt) sarmalına yol açar.
- **CAP Belirsizlik İlkesi:** Dağıtık bir sistemde aynı anda hem Tutarlılık (Consistency) hem de Erişilebilirlik (Availability) bilgisini mutlak bir doğrulukla "bilmek" imkansızdır; bu, sistemin fiziksel sınırlamasıdır.

### 🧩 2. Mikroservis Devrimi: Otonom Birimlerin Egemenliği (The Swarm)
Bir uygulamayı, her biri kendi veritabanına sahip, bağımsız dağıtılabilen ve belirli bir iş yeteneğini temsil eden otonom hücreler kümesi olarak tasarlamaktır.
- **Brooks Yasası ve Ekip Topolojileri:** "Gecikmiş bir yazılım projesine daha fazla insan eklemek, onu daha da geciktirir." Mikroservisler, bu yasayı kırmak için ekipler arası bağımlılıkları minimize eden tek gerçek panzehirdir.

---

## 🏗️ Bölüm II: Dağıtık Sistemlerin Fiziksel Temelleri (The Singularity Foundations)

### 1. Bounded Context (DDD & Davranışsal Modelleme)
Her servisin kendi semantik dilini (**Ubiquitous Language**) belirlemesidir.
- **Event Storming:** İş akışlarını "Domain Events" üzerinden modelleyerek Bounded Context sınırlarını keşfetme tekniğidir. Bu projede her servis, kendi bağlamında (Context) bir "Egemenlik Alanı" oluşturur.

### 2. Otonomi & Biyolojik "Shared Nothing" Mimari
Her servis, kendi veritabanına ve teknoloji yığınına sahip olmalıdır. Veri egemenliği (Data Sovereignty), biyolojik bir hücrenin zarı gibi servisi izole eder.

### 3. Ölçeklenebilirlik: AKF Scale Cube (Z-Axis Sharding)
- **X-Axis:** Instance kopyalamak.
- **Y-Axis:** Fonksiyonel ayrıştırma.
- **Z-Axis:** Veri bölümleme ve Kiracı İzolasyonu (**Tenant Isolation**). Bu, küresel ölçekteki sistemlerin (Facebook, Google) sunduğu mutlak ölçeklenme kapasitesidir.

---

## 📡 Bölüm III: Protokol Fiziği ve Nanosaniye Mühendisliği

Ağ katmanı, sistemin sinir sistemidir ve burada her nanosaniye kritiktir:

### 1. Senkron: gRPC, Protobuf v3 & CPU Cache Line
- **False Sharing & Cache Interaction:** gRPC'nin ikili (binary) formatı, CPU L1/L2 cache hatlarıyla (Cache Lines) en verimli şekilde etkileşime girecek şekilde optimize edilmiştir. Bu, bağlam değiştirme (context switching) maliyetlerini minimize eder.
- **Deadlines & Cancellation Propagation:** Bir istek iptal edildiğinde bu sinyal tüm servis zinciri boyunca iletilerek gereksiz işlem yükü (Zombie Processes) önlenir.

### 2. Asenkron: RabbitMQ & Oyun Teorisi (Nash Equilibrium)
Dağıtık sistemlerde retrialar (yeniden denemeler) bir "Yeniden Deneme Fırtınasına" (Retry Storm) yol açabilir.
- **Nash Dengesi:** Servislerin retrialarını ve backpressure (geri basınç) mekanizmalarını oyun teorisi perspektifiyle kurgulayarak, sistemin kaosa sürüklenmesini önlüyoruz. **DLX (Dead Letter Exchange)** ve **Manual ACKs** ile mesaj güvenliği sağlanır.

---

## 💾 Bölüm IV: Kuyruklama Teorisi ve Dağıtık Konsensüs (The PhD Tome)

### 1. Little Yasası (Queuing Theory)
Bir sistemdeki ortalama iş sayısı, geliş hızı ile ortalama harcanan zamanın çarpımına eşittir ($L = \lambda W$). Bu matematiksel model ile Gateway doygunluğunu ve P99.9 gecikme yüzdeliklerini (P99.9 tail latency) hesaplıyoruz.

### 2. Konsensüs ve Mantıksal Saatler (Clocks & Raft)
Dağıtık sistemlerde "mutlak zaman" yoktur (Time is an illusion). **Raft/Paxos** algoritmaları ve **Lamport Timestamps** (Mantıksal Saatler) kullanılarak olayların sırası (Causal Ordering) belirlenir.

---

## 🛡️ Bölüm V: Sıfır Güven (Zero-Trust) ve Kimlik Egemenliği

- **SPIFFE/SPIRE tabanlı Kimlik-merkezli Mikro-segmentasyon:** Artık sadece IP/Port bazlı değil, her bir fonksiyon çağrısının kendine ait kriptografik bir kimliği olduğu bir güvenlik modelidir.
- **Mutual TLS (mTLS):** İç servisler arası iletişimde her iki tarafın da birbirini doğruladığı mutlak güvenlik katmanıdır.

---

## 🔍 Bölüm VI: SRE & Modern Gözlemlenebilirlik (Global Standards)

### 1. Üç Sütun (MELT)
Metrics, Events, Logs, Traces. Sistemin "ateşini" ölçer, "yolculuğunu" izleriz.

### 2. Yüksek Kardinalite (High Cardinality) ve Sampling
Milyarlarca span içinden hangilerini saklayacağımıza **Tail-based Sampling** math (Geriye dönük örnekleme) ile karar veriyoruz. Bu, yüksek trafikli sistemlerde maliyet ve görünürlük dengesini sağlar.

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

## 🚀 Dağıtım (Deployment) Masterclass & Işık Hızı Kısıtları

### 1. Küresel Ağ Fiziği
**BGP Hijacking** ve **Anycast Routing**'in global servis keşfi üzerindeki etkileri göz önüne alınarak, ışık hızı kısıtları dahilinde (Speed of light) çok bölgeli (multi-regional) mimariler kurgulanır.

### 2. Kaos Mühendisliğinin Ontolojisi (Chaos Engineering)
Sisteme rastgele hatalar enjekte edilerek **Blast Radius** (Etki Alanı) analizi yapılır. "Failing is not an option, but recovery is a requirement."

---

## 📚 Kaynakça (The Omniscience Bibliography)

- *Designing Data-Intensive Applications* - **Martin Kleppmann** (Teorik Temel)
- *Microservices Patterns* - **Chris Richardson** (Pratik Uygulama)
- *Site Reliability Engineering (SRE) Book* - **Google Engineering** (Operasyonel Mükemmellik)

<div align="center">
  <br/>
  <sub>Engineering Excellence Series | **© 2026 arch-yunus** | "Complex systems that work are invariably found to have evolved from simple systems that worked."</sub>
</div>
