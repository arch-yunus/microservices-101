<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices Architecture Banner" />

  # Microservices Architecture: The Ultimate Hyper-Masterclass 📚
  ### Modern Dağıtık Sistemler, Yazılım Ekonomisi ve Mühendislik Ansiklopedisi
  
  [![Architecture](https://img.shields.io/badge/Architecture-Microservices-6366f1?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)
  [![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![OpenTelemetry](https://img.shields.io/badge/Observability-OpenTelemetry-f59e0b?style=for-the-badge)](https://opentelemetry.io)
  [![SRE Tooling](https://img.shields.io/badge/Strategy-SRE_Standard-ef4444?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)

  **"Dağıtık sistemlerin devasa karmaşıklığı karşısında, mimari disiplin sadece bir tercih değil, bir hayatta kalma stratejisidir. Bu depo, monolitik sistemlerin hantallığından kurtulup, mikroservislerin otonom ve sınırsız ölçeklenebilir dünyasına mühendislik perspektifiyle atılan dev bir adımdır."**

  ---
</div>

## 🌐 Bölüm I: Yazılım Mimarilerinin Mühendislik ve Ekonomi Evrimi

Yazılım dünyası, başlangıcından bu yana karmaşıklığı (complexity) yönetmek için sürekli form değiştirmiştir. Yazılım mimarisi, sadece teknik bir tercih değil, aynı zamanda bir **"Zaman-Maliyet-Ölçek"** denklemidir. Mikroservisler, bu evrimin en olgun, ancak aynı zamanda operasyonel maliyeti en yüksek halkalarından biridir.

### 🏛️ 1. Geleneksel Monolit: Devlerin Yükselişi ve Düşüşü (The Giant)
Tüm iş mantığı, veritabanı erişim katmanları (DAL), kullanıcı arayüzü (UI) bileşenleri ve global konfigürasyonların tek bir bellek alanında (address space) ve tek bir süreç (process) içinde paketlendiği yapıdır. 
- **Ekonomik Perspektif:** Day 1 (Başlangıç) maliyeti en düşük mimaridir. Basit deployment, kolay hata ayıklama ve tek bir veritabanı üzerinden atomik işlemler (ACID) sunar. 
- **Teknik Entropi:** Uygulama büyüdükçe "Yazılım Entropisi" (Entropy) yasası işlemeye başlar. Kod tabanı devasa bir yığın (Big Ball of Mud) haline gelir. Tek bir satır kod değişikliği, tüm sitemin regresyon testlerine tabi tutulmasını ve devasa bir build sürecini zorunlu kılar. Bu durum, piyasaya çıkış hızını (Time-to-Market) öldüren en büyük etkendir.

### 📡 2. SOA (Service-Oriented Architecture): ESB Prangaları
Mikroservislerin doğrudan atasıdır. Servislerin birbirleriyle merkezi bir **Enterprise Service Bus (ESB)** üzerinden konuştuğu yapıdır. 
- **Temel Fark ve Hata:** SOA'da "Ağır, hantal ve aşırı akıllı" bir merkezi boru hattı (ESB) varken, tüm iş mantığı bu merkezi yapıda boğulurdu. Mikroservis felsefesi ise tam zıttını savunur: **"Akıllı Uç Noktalar ve Basit Boru Hatları"** (Smart Endpoints, Dumb Pipes).

### 🧩 3. Mikroservis Devrimi: Otonom Birimlerin Dansı (The Swarm)
2011 yılında Venedik'te düzenlenen yazılım mimarisi çalıştayında kristalleşen bu felsefe; bir uygulamayı, her biri kendi veritabanına sahip, bağımsız olarak dağıtılabilen ve belirli bir iş yeteneğini (**Business Capability**) temsil eden küçük, otonom servisler kümesi olarak tasarlamaktır.
- **Conway Yasası ve Reverse Conway Maneuver:** "Sistemleri tasarlayan kurumlar, bu sistemlerin yapısında kendi iletişim yapılarını kopyalarlar." Mikroservisler, bu yasayı tersine çevirerek (Reverse Conway Maneuver) ekiplerin de otonom, çevik ve dikey olarak bölünmüş (Vertical Slices) olmasını sağlar.

---

## 🏗️ Bölüm II: Dağıtık Sistemlerin 4 Temel Sütunu (Advanced Masterclass)

Bir sistemi "mikroservis" yapan şey, kodun fiziksel boyutu değil, aşağıdaki dört mühendislik sütununa olan sarsılmaz bağlılığıdır:

### 1. Bounded Context (Sınırlı Bağlam & DDD Derinliği)
Domain-Driven Design (DDD) prensibi olan Bounded Context, her servisin kendi "semantik dilini" (**Ubiquitous Language**) belirlemesidir.
- **Context Mapping:** Bir e-ticaret sisteminde "Ürün" nesnesi, Satış bağlamında fiyat ve KDV iken, Kargo bağlamında sadece ağırlık ve hacimdir. Bu sınırların net çizilmemesi, dağıtık sistemlerdeki en büyük karmaşa kaynağı olan "Paylaşımlı Nesne" (Shared Kernel) problemine yol açar. Bu projede her servis kendi `domain` katmanıyla bu sınırları korur.

### 2. Otonomi & Data Sovereignty (Veri Egemenliği)
Her servis, kendi veritabanına, kendi teknoloji yığınına (Heterogeneous Stack) ve kendi yaşam döngüsüne sahip olmalıdır. 
- **Kritik Mühendislik Kuralı:** "Paylaşımlı Veritabanı" (Shared Database) mimari bir intihardır. Eğer iki servis aynı tabloya SQL ile erişiyorsa, otonomiden bahsedilemez. Veri otonomisi, dağıtık sistemlerde **Nihai Tutarlılık** (Eventual Consistency) zorunluluğunu da beraberinde getirir.

### 3. Ölçeklenebilirlik (The AKF Scale Cube)
Mikroservisleri ölçeklendirmek, sadece instance sayısını artırmak değildir. **AKF Scale Cube** perspektifiyle:
- **X-Axis:** Load balancer arkasında instance kopyalamak (Cloning).
- **Y-Axis:** Fonksiyonel olarak servislere bölmek (Functional Decomposition - Mikroservislerin asıl gücü).
- **Z-Axis:** Veriyi coğrafi veya müşteri bazlı bölmek (Sharding).
Bu projede biz Y-Axis ölçeklendirmeyi mikroservis mimarisiyle, X-Axis ölçeklendirmeyi ise Docker ile uyguluyoruz.

### 4. Hata Toleransı: Cascading Failure ve Resilience
Dağıtık sistemlerde ağ güvenilmezdir (Fallacies of Distributed Computing). Önemli olan hatanın bir servisden diğerine bir "domino taşı" gibi yayılmasını engellemektir (**The Titanic Effect** - Titanik gemisindeki su sızdırmaz bölmelerin hatası gibi).
- **Graceful Degradation:** Bir servis çöktüğünde sistem tamamen durmak yerine, hayati fonksiyonlarını devam ettirecek "Zarif Gerileme" moduna geçmelidir. **Circuit Breaker** (Devre Kesici), açık/kapalı/yarı-açık durumlarıyla bu güvenliğin baş aktörüdür.

---

## 📡 Bölüm III: İleri Seviye İletişim Protokolleri & Ağ Mühendisliği

Sistemin "Sinir Sistemi" olan iletişim katmanı, gecikme (latency) ve güvenilirlik dengesini belirler:

### 1. Senkron İletişim: gRPC & Protobuf v3 (High Performance)
İsteği gönderen servis, cevap gelene kadar "bloklanır". Bu projede iç servis iletişimi için **gRPC** (Google Remote Procedure Call) tercih edilmiştir.
- **HTTP/2 ve Binary Framing:** JSON gibi metin tabanlı formatların getirdiği parse maliyeti ve payload yükü, ikili (binary) formatla minimuma indirilir. Ağ bandwidth kullanımı %50-70 oranında düşer.
- **Streaming & Multiplexing:** Tek bir TCP bağlantısı üzerinde binlerce isteğin eşzamanlı taşınması ve stream desteği ile büyük verilerin parçalı iletimi sağlanır.
- **Strict Contract Management:** `.proto` dosyaları, servisler arası değişmez bir anayasa oluşturarak uyumsuzluk hatalarını (Incompatibility) daha derleme (compile) aşamasında önler.

### 2. Asenkron İletişim: RabbitMQ & Event-Driven Architecture (EDA)
Servislerin zaman bazlı bağımlılığını (Temporal Coupling) ortadan kaldıran katmandır. 
- **Competing Consumers Pattern:** RabbitMQ üzerindeki bir kuyruğu birden fazla instance dinleyerek iş yükünü (load-balancing) otomatik olarak paylaşır.
- **Idempotency & Retry:** Mesajın garantili iletimi için "At-least-once" stratejisi uygulanır. Bu durum, aynı mesajın iki kez gelmesi riskine karşı servislerimizde **Idempotency** (Aynı güçlülük) kontrolünü zorunlu kılar.

---

## 💾 Bölüm IV: Dağıtık Sistem Zorlukları & Teorik Altyapı (Mühendislik Masterclass)

### 1. CAP Teoremi vs. PACELC: Gerçek Dünya Seçimleri
CAP Teoremi'nin (Consistency, Availability, Partition Tolerance) ötesine geçen **PACELC**, modern dağıtık sistemlerin gerçek rehberidir.
- **PACELC:** Bir sistem bölünmüşken (Partitioned) CAP geçerlidir. Ancak sistem normal çalışırken (Else) **Latens (Latency)** ve **Tutarlılık (Consistency)** arasında bir seçim yapmanız gerekir. Düşük gecikme mi istersiniz, yoksa anlık verinin her yerde aynı olmasını mı?

### 2. Saga Deseni: Dağıtık İşlemlerin Yönetimi
Atomik (ACID) işlemlerin mikroservis dünyasında yapılamaması nedeniyle ortaya çıkan **Saga Pattern**, süreci adımlara böler.
- **Telafi Edici İşlemler (Compensating Transactions):** Eğer 3. adımda bir hata oluşursa, 1. ve 2. adımların etkilerini "geri alacak" (UNDO) işlemler otomatik olarak tetiklenir. Bu projede **Choreography-based Saga** (Olay tabanlı koordinasyon) yaklaşımı temellendirilmiştir.

### 3. Dağıtık Sistemlerin 8 Yanılgısı (L. Peter Deutsch)
Mühendislerin düştüğü en büyük hatalar: Ağ güvenilirdir, latency sıfırdır, bandwidth sonsuzdur, ağ güvenlidir, topoloji değişmez, bir admin vardır, taşıma maliyeti sıfırdır, ağ homojendir. Biz bu sistemde tüm bu yanılgıları varsayarak **Circuit Breaker** ve **Health Checks** kurguladık.

---

## 🛡️ Bölüm V: API Gateway ve Edge Systems (The Fortress)

Mikroservisler dış dünyaya asla doğrudan açılmamalıdır. **API Gateway**, sistemin "Zırhlı Kapısı"dır:

- **BFF (Backend for Frontend):** Farklı cihaz türleri (Web, Mobil, IoT) için özelleştirilmiş veri paketleri (Payload Optimization) sunar.
- **Load Shedding & Request Collapsing:** Sistem aşırı yüklendiğinde, Gateway henüz isteği iç servislere iletmeden "yük atma" (Shedding) yaparak çekirdek servislerin çökmesini engeller.
- **Mutual TLS (mTLS) & Security Headers:** İç servisler arası iletişimde ek güvenlik katmanları ve Gateway seviyesinde JWT tabanlı kimlik doğrulama ile uçtan uca güvenlik sağlanır.

---

## 🔍 Bölüm VI: Modern Gözlemlenebilirlik ve SRE Perspektifi

Sistemin "iç durumunu" dışarıdan ne kadar iyi anlayabildiğiniz, debugging maliyetinizi ve uptime oranınızı (SLA) belirler. **Observability** üç ana sütuna (Three Pillars) ve iki metoda dayanır:

### 1. Üç Sütun (MELT)
- **Metrics (Prometheus):** Sayısal veriler. Sistemin "ateşini" (Hata oranı, CPU, Memory) ölçeriz.
- **Distributed Tracing (OpenTelemetry & Jaeger):** Bir isteğin gateway'den başlayıp DB'ye kadar olan yolculuğunu "izleriz" (X-Ray görüşü).
- **Logs:** Detaylı olay kayıtları (Ancak TraceID ile korelasyonu şarttır).

### 2. RED ve USE Metotları
- **RED Method:** (Rate, Errors, Durations) - Servislerin sağlığını ölçmek için kullanılır.
- **USE Method:** (Utilization, Saturation, Errors) - Donanım ve kaynak kullanımını ölçmek için kullanılır.
Bu projede Modül 6 ile **OpenTelemetry** üzerinden bu verileri topluyoruz.

---

## 🛣️ Bölüm VII: Proje Müfredatı ve Uygulama Aşamaları

Bu depo, yukarıdaki teorik bilgileri pratiğe döken **6 aşamalı bir mühendislik masterclass'ı** sunar:

| Seviye | Modül | Teknik Derinlik (Technical Depth) | Kazanımlar | Durum |
| :--- | :--- | :--- | :--- | :---: |
| 🏗️ | **Modül 1: Mimari** | Clean Architecture (DDD tabanlı) | Domain, Use Case, Entity ve Adapter katmanlarının Postgres ile ayrıştırılması. | ✅ |
| 🔌 | **Modül 2: İletişim** | gRPC & Protobuf v3 | Servisler arası senkron, binary ve kontrat tabanlı (IDL) iletişim. | ✅ |
| 🛡️ | **Modül 3: Gateway** | Edge Service (Fortress) | Echo Framework ile merkezi yönlendirme ve güvenlik katmanlarının inşası. | ✅ |
| 📩 | **Modül 4: Mesajlaşma** | RabbitMQ (EDA) | Temporal Decoupling ve Event-Driven mimarinin asenkron tüketicilerle kurgulanması. | ✅ |
| 🐒 | **Modül 5: Stabilite** | Circuit Breaker & Health | Cascade failure önleme, gRPC health probes ve hata izolasyonu (Resilience). | ✅ |
| 🔍 | **Modül 6: Gözlem** | OpenTelemetry & Jaeger | Distributed Tracing ve context propagation ile sistemin tamamen izlenebilir olması. | ✅ |

---

## 🚀 Hızlı Başlangıç (The SRE Quickstart)

Sistemi tek bir komutla, tüm altyapı bileşenleriyle birlikte (Postgres, RabbitMQ, Jaeger, Services) ayağa kaldırabilirsiniz. `make` komutu tüm Docker orkestrasyonunu sizin yerinize yönetir:

```bash
# 1. Projeyi yerel makinenize klonlayın
git clone https://github.com/arch-yunus/microservices-101.git

# 2. Infra ve Mikroservis ekosistemini izole ağlarda başlatın
make up
```

---

## 📘 Bölüm VIII: İleri Seviye Konseptler, Anti-Desenler ve Teknik Sözlük

### 📜 Teknik Sözlük
- **Idempotency:** Bir işlemin (örneğin "Bakiye Düş") iki kez tetiklense bile (Ağ gecikmesi nedeniyle) veritabanında sadece bir kez etkili olması garantisidir.
- **Infrastructure as Code (IaC):** Bu projedeki Docker Compose dosyalarıyla altyapının da kod olarak tanımlanıp reproduciable (tekrar üretilebilir) olmasıdır.
- **Immutable Infrastructure:** Bir servis güncellendiğinde "yama" yapmak yerine, eski sürümün tamamen silinip yeni kopyanın sıfırdan oluşturulmasıdır.

### 🚫 Mikroservis Anti-Desenleri
- **Nana-services:** Servisleri o kadar küçük parçalara bölmek ki, ağ gecikmesinin iş mantığından daha fazla zaman alması.
- **Distributed Monolith:** Servislerin birbirine o kadar bağımlı olması ki, tek bir servisi deploy etmek için hepsini aynı anda kapatıp açmak zorunda kalmak.
- **Entity Services:** Sadece veritabanındaki bir tabloyu temsil eden servisler (CRUD servisler). Bu, otonomiyi ve business felsefesini öldürür.

---

## 🧪 Bölüm IX: Dağıtım Stratejileri ve Chaos Engineering

### 1. Modern Deployment Stratejileri
- **Canary Release:** Yeni sürüm önce kullanıcıların %1'ine açılır. Hata yoksa kademeli olarak artırılır.
- **Dark Launching:** Özelliğin kod olarak sisteme koyulması ancak "Feature Flag"lar ile sadece belirli kişilere açılması.
- **Shadow Deployment (Mirroring):** Gerçek trafiğin bir kopyasının sessizce yeni sürüme yönlendirilip performansının canlı veride test edilmesi.

### 2. Chaos Engineering (Netflix İlkeleri)
Mikroservisler "bozulmaya hazır" olmalıdır. **Chaos Monkey** gibi araçlarla sisteme rastgele hatalar (gecikme, servis durdurma) enjekte edilerek sistemin bu durumlardaki tepkisi ve otomatik iyileşme (Auto-healing) yeteneği ölçülür.

---

## 📚 Kaynakça ve İleri Okuma (The Bibliography)

Bu mühendislik ansiklopedisini hazırlarken faydalanılan, dünyada kabul görmüş başucu eserleri:
- *Microservices Patterns* - **Chris Richardson** (Saga ve Veri mimarileri için kutsal kitap).
- *Building Microservices* - **Sam Newman** (Mimari evrim ve ekipler arası iletişim stratejileri).
- *Domain-Driven Design* - **Eric Evans** (Mimari sınırların belirlenmesi ve Ubiquitous Language).
- *Designing Data-Intensive Applications* - **Martin Kleppmann** (Dağıtık sistemlerin teorik temelleri).
- *Site Reliability Engineering (SRE) Book* - **Google Engineering** (Observability ve Güvenilirlik prensipleri).

<div align="center">
  <br/>
  <sub>Engineering Excellence Series | **© 2026 arch-yunus** | "Complex systems that work are invariably found to have evolved from simple systems that worked."</sub>
</div>
