<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices Architecture Banner" />

  # Microservices Architecture: The Definitive Engineering Masterclass 📚
  ### Modern Dağıtık Sistemler Tasarımı ve Mühendislik Ansiklopedisi
  
  [![Architecture](https://img.shields.io/badge/Architecture-Microservices-6366f1?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)
  [![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![OpenTelemetry](https://img.shields.io/badge/Observability-OpenTelemetry-f59e0b?style=for-the-badge)](https://opentelemetry.io)
  [![Status](https://img.shields.io/badge/Status-Project_Completed-success?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)

  **"Dağıtık sistemlerin devasa karmaşıklığı karşısında, mimari disiplin sadece bir tercih değil, bir hayatta kalma stratejisidir. Bu depo, monolitik sistemlerin hantallığından kurtulup, mikroservislerin otonom ve sınırsız ölçeklenebilir dünyasına mühendislik perspektifiyle atılan dev bir adımdır."**

  ---
</div>

## 🌐 Bölüm I: Yazılım Mimarilerinin Mühendislik Evrimi

Yazılım dünyası, başlangıcından bu yana karmaşıklığı (complexity) yönetmek için sürekli form değiştirmiştir. Mikroservisler, bu evrimin en olgun ve güncel halkalarından biridir. Bu evrimi anlamadan, mikroservislerin neden var olduğunu anlamak imkansızdır.

### 🏛️ 1. Geleneksel Monolit (The Giant)
Tüm iş mantığı, veritabanı erişim katmanları (DAL), UI bileşenleri ve konfigürasyonların tek bir süreç (process) içinde paketlendiği yapıdır. 
- **Mühendislik Perspektifi:** Geliştirme sürecinin başında (Day 1) her şey mükemmel görünür; basit deployment, kolay hata ayıklama. Ancak uygulama büyüdükçe **"Yazılım Entropisi"** devreye girer. Kod tabanı devasa bir yığın (Big Ball of Mud) haline gelir.
- **Problem:** Tek bir satır kod değişikliği için tüm uygulamanın yeniden build edilip deploy edilmesi gerekir. Bu, modern **CI/CD** süreçleri için en büyük darboğazdır.

### 📡 2. SOA (Service-Oriented Architecture)
Mikroservislerin atasıdır. Servislerin birbirleriyle bir **Enterprise Service Bus (ESB)** üzerinden konuştuğu yapıdır. 
- **Temel Fark:** SOA'da "Ağır ve Akıllı" bir boru hattı (ESB) varken, mikroservislerde "Akıllı Uç Noktalar ve Basit Boru Hatları" (Smart Endpoints, Dumb Pipes) felsefesi hakimdir.

### 🧩 3. Mikroservis Devrimi (The Swarm)
2011 yılında Venedik'te düzenlenen bir yazılım mimarisi çalıştayında (Software Architecture Workshop) ilk kez bu terim modern anlamıyla kullanılmıştır. Bu, bir uygulamayı, her biri kendi veritabanına sahip, bağımsız olarak dağıtılabilen ve belirli bir iş yeteneğini (Business Capability) temsil eden küçük servisler kümesi olarak tasarlamaktır.
- **Conway Yasası:** "Sistemleri tasarlayan kurumlar, bu sistemlerin yapısında kendi iletişim yapılarını kopyalarlar." Mikroservisler, bu yasayı tersine çevirerek (Inverse Conway Maneuver) ekiplerin de otonom olmasını sağlar.

---

## 🏗️ Bölüm II: Mikroservislerin 4 Temel Sütunu (The Deep Dive)

Bir sistemi "mikroservis" yapan şey, kodun küçük olması değil, aşağıdaki dört mühendislik prensibine sıkı sıkıya bağlı olmasıdır:

### 1. Bounded Context (Sınırlı Bağlam & DDD)
Domain-Driven Design (DDD) prensibi olan Bounded Context, her servisin kendi "semantik dilini" belirlemesidir.
- **Örnek:** Bir e-ticaret sisteminde "Ürün" (Product) nesnesi:
    - **Satış Servisi:** Fiyat, KDV ve İndirim bilgilerine odaklanır.
    - **Kargo Servisi:** Ağırlık, Boyut ve Adres bilgilerine odaklanır.
- **Sonuç:** Servisler arası nesne kirliliği (leakage) önlenir ve her servis kendi alanında uzmanlaşır.

### 2. Otonomi & Data Sovereignty (Veri Egemenliği)
Her servis, kendi veritabanına, kendi teknoloji yığınına ve kendi yaşam döngüsüne sahip olmalıdır. 
- **Kritik Kural:** "Shared Database" (Paylaşımlı Veritabanı) en büyük anti-desendir. Eğer iki servis aynı tabloyu okuyup yazıyorsa, mikroservislerin otonomisini öldürmüşsünüz demektir. Dağıtık veri yönetimi zorunludur.

### 3. Ölçeklenebilirlik (Horizontal Scalability)
Monolitik yapıda tüm sistemi (gereksiz kısımlar dahil) 10 katına çıkarmanız (Vertical Scaling) gerekirken; mikroservislerde sadece darboğaz oluşturan (örneğin indirim döneminde sadece `Sipariş Servisi`) servisi binlerce kopya (instance) oluşturarak ölçeklendirirsiniz.

### 4. Hata Toleransı & Cascading Failure Önleme
Dağıtık sistemlerde hata kaçınılmazdır. Önemli olan bu hatanın tüm sistemi bir "domino taşı" gibi yıkmasını engellemekir (**The Titanic Effect**).
- **Resilience:** Bir servis çöktüğünde sistemin geri kalanı, "Graceful Degradation" (zarifçe gerileme) prensibiyle çalışmaya devam etmelidir. **Circuit Breaker** (Devre Kesici) bu mimarinin can damarıdır.

---

## 📡 Bölüm III: İletişim Protokolleri ve Ağ Mühendisliği

Mikroservisler arası iletişim, sistemin performansını ve güvenilirliğini belirleyen en kritik katmandır:

### 1. Senkron İletişim: gRPC & Protobuf v3
İsteği gönderen servis, cevap gelene kadar bekler. Bu projede **gRPC** (Google Remote Procedure Call) kullanılmıştır.
- **HTTP/2 Üzerinden Binary İletişim:** Veriler JSON gibi metin tabanlı değil, ikili (binary) formatta taşınır. Bu, ağ bandwidth kullanımını %50'ye kadar düşürür.
- **HPACK Sıkıştırma:** Headerlar bile sıkıştırılarak taşınır.
- **Strict Contract:** Protobuf dosyaları hem sunucu hem istemci için birer "Anayasa" niteliğindedir.

### 2. Asenkron İletişim: RabbitMQ & Event-Driven Architecture (EDA)
Servisler birbirine doğrudan bağımlı değildir (Temporal Decoupling). 
- **Olay Odaklı Tasarım:** Sipariş oluştuğunda "OrderCreated" olayı RabbitMQ'ya fırlatılır. `Notifier Service` bu mesajı ne zaman müsaitse o zaman işler.
- **Smart Endpoints, Dumb Pipes:** RabbitMQ sadece mesajı taşımaktan sorumludur; iş mantığı uç servislerdedir.

---

## 💾 Bölüm IV: Dağıtık Sistem Zorlukları & Teorik Altyapı

Mikroservis dünyasında "bedava öğle yemeği" yoktur. Ciddi teorik zorluklarla baş etmek gerekir:

### 1. CAP Teoremi vs. PACELC
Dağıtık bir sistemde **Tutarlılık (Consistency)**, **Erişilebilirlik (Availability)** ve **Bölünme Toleransı (Partition Tolerance)** özelliklerinden sadece ikisi aynı anda sağlanabilir.
- **PACELC:** Bir sistem bölünmüşken (Partitioned) CAP geçerlidir. Ancak sistem normal çalışırken (Else) **Latens (Latency)** ve **Tutarlılık (Consistency)** arasında bir seçim yapmanız gerekir.

### 2. Saga Deseni & Nihai Tutarlılık (Eventual Consistency)
Birden fazla servisi ilgilendiren bir işlemde bir adım hata verirse, önceki adımları geri almak için **Compensating Transactions** (Telafi Edici İşlemler) kullanılır.
- **Orkestrasyon:** Merkezi bir "Saga Manager" süreci yönetir.
- **Koreografi:** Servisler olaylar (events) üzerinden kendi aralarında haberleşip süreci ilerletir.

### 3. Dağıtık İşlem (Distributed Transaction) Fallacies
1. Ağ güvenilirdir (Değildir!)
2. Latency sıfırdır (Değildir!)
3. Bandwidth sonsuzdur (Değildir!)
4. Ağ güvenlidir (Değildir!)

---

## 🛡️ Bölüm V: API Gateway Masterclass

Mikroservisler dış dünyaya doğrudan açılmamalıdır. **API Gateway**, tüm isteklerin karşılandığı "Zırhlı Kapı" ve sistemin "Orkestratörü"dür:

- **BFF (Backend for Frontend):** Mobil ve Web için farklı Gateway katmanları sunularak gereksiz veri taşıma yükü (Overfetching) önlenir.
- **Circuit Breaking & Load Shedding:** Gateway, iç servislerin aşırı yüklendiğini hissederse istekleri henüz içeri almadan reddeder.
- **JWT & Authorization:** Kullanıcı kimliği burada doğrulanır ve iç servislere "Güvenli Context" olarak iletilir.

---

## 🔍 Bölüm VI: Modern Gözlemlenebilirlik (Observability)

Sistemin iç durumunu (internal state) dışarıdan ne kadar iyi anlayabildiğiniz, debugging maliyetinizi belirler. **Observability** üç ana sütuna (Three Pillars) dayanır:

1.  **Metrics (Prometheus):** Sistemin "ateşi" kaç? (Hata oranı, CPU yükü, İstek sayısı).
2.  **Distributed Tracing (OpenTelemetry & Jaeger):** Bir isteğin gateway'den başlayıp DB'ye gidip geri dönüşüne kadar olan tüm yolculuğun (Span'lar) görselleşmesi.
3.  **Structured Logging:** "Hata oluştu" yerine, TraceID ve SpanID içeren JSON formatında loglar.

---

## 🛣️ Bölüm VII: Proje Müfredatı ve Mühendislik Uygulaması

Bu depo, yukarıdaki teorik bilgileri pratiğe döken 6 aşamalı bir mühendislik masterclass'ı sunar:

| Modül | Teknik Derinlik | Kazanımlar | Durum |
| :--- | :--- | :--- | :---: |
| 🏗️ **Modül 1** | **Mimarinin İnşası** | Clean Architecture (Domain, Use Case, Interface), PostgreSQL Integration | ✅ |
| 🔌 **Modül 2** | **Kontratlı İletişim** | gRPC, Protobuf v3, Unary calls, Internal Service Communication | ✅ |
| 🛡️ **Modül 3** | **Stratejik Geçiş** | API Gateway (Fortress Pattern), Echo Framework, Middleware Routing | ✅ |
| 📩 **Modül 4** | **Olay Odaklı Tasarım** | RabbitMQ, Asynchronous Messaging, Consumer Robustness, EDA | ✅ |
| 🐒 **Modül 5** | **Hata Dayanıklılığı** | Circuit Breaker (Half-Open states), gRPC Health Checks, Fail-fast strategy | ✅ |
| 🔍 **Modül 6** | **X-Ray Görüşü** | Distributed Tracing, OpenTelemetry SDK, Context Propagation, Jaeger UI | ✅ |

---

## 🚀 Hızlı Başlangıç (The Quickstart)

Sistemi tek bir komutla, tüm altyapı bileşenleriyle birlikte ayağa kaldırabilirsiniz. `make` komutu tüm karmaşıklığı (`docker-compose` parametrelerini) sizin yerinize yönetir:

```bash
# 1. Projeyi yerel makinenize klonlayın
git clone https://github.com/arch-yunus/microservices-101.git

# 2. Docker Compose ile tüm ekosistemi (Gateway, Services, RabbitMQ, Postgres, Jaeger) başlatın
make up
```

---

## 📘 Bölüm VIII: İleri Seviye Konseptler ve Sözlük

Uzmanlık seviyesini belirleyen kavramlar:

### 📜 Mikroservis Sözlüğü
- **Idempotency (Aynı Güçlülük):** Bir işlemin (örneğin ödeme mesajı) birden fazla kez uygulanmasının (Network retry nedeniyle), sonucu değiştirmemesi özelliğidir.
- **Service Mesh (Istio/Linkerd):** Ağ karmaşıklığını uygulama kodundan çıkarıp altyapı katmanına (Sidecar) taşıyan yapıdır.
- **Load Shedding:** Sistem aşırı yüklendiğinde, kritik olmayan istekleri önceden belirlenmiş bir stratejiyle reddederek tamamen çökmesini engellemek.

### 🧪 Test Stratejileri (New Pyramid)
Mikroservislerde klasik test piramidi yetmez:
- **Contract Testing (PACT):** Servisler arası API değişimlerinin sistemleri bozmayacağının matematiksel garantisidir.
- **Chaos Engineering (Netflix Monkey):** Canlı sistemde rastgele servisleri kapatarak sistemin dayanıklılığını test etmek.

---

## 🚀 Bölüm IX: Dağıtım Stratejileri Masterclass

Sistemde sıfır kesinti (Zero-Downtime) ile sürüm güncellemek için kullanılan yöntemler:

1.  **Blue-Green Deployment:** Eski (Blue) ve Yeni (Green) sürüm aynı anda çalışır. Trafik aniden yeniye yönlendirilir.
2.  **Canary Release:** Yeni sürüm önce kullanıcıların küçük bir kısmına (%5) açılır. Hata yoksa genele yayılır.
3.  **Rolling Update:** Servis kopyaları tek birer birer güncellenerek trafiğin kesilmemesi sağlanır.

---

## 📚 Kaynakça ve İleri Okuma

Bu mühendislik ansiklopedisini hazırlarken faydalanılan başucu eserleri:
- *Microservices Patterns* - Chris Richardson (Saga ve Veri desenleri için İncil niteliğindedir).
- *Building Microservices* - Sam Newman (Mimari evrim ve ekipler arası iletişim).
- *Domain-Driven Design* - Eric Evans (Mimari sınırların belirlenmesi).
- *Designing Data-Intensive Applications* - Martin Kleppmann (Dağıtık sistemlerin teorik temelleri).

<div align="center">
  <br/>
  <sub>Engineering Excellence Series | **© 2026 arch-yunus** | "Complex systems that work are invariably found to have evolved from simple systems that worked."</sub>
</div>
