<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices Architecture Banner" />

  # Microservices Architecture: The Definitive Encyclopedia 📚
  ### Modern Dağıtık Sistemler Tasarımı ve Mühendislik El Kitabı
  
  [![Architecture](https://img.shields.io/badge/Architecture-Microservices-6366f1?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)
  [![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![Status](https://img.shields.io/badge/Status-Active_Development-success?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)
  [![License](https://img.shields.io/badge/License-MIT-success?style=for-the-badge)](LICENSE)

  **"Dağıtık sistemlerin karmaşıklığını yönetmek, sadece kod yazmak değil; doğru mimari prensipleri uygulamakla başlar. Bu depo, bir yazılımın monolitik prangalarından kurtulup, mikroservislerin özgür ve ölçeklenebilir dünyasına geçiş rehberidir."**

  ---
</div>

## 🌐 Bölüm I: Yazılım Mimarilerinin Evrimi

Yazılım dünyası, başlangıcından bu yana karmaşıklığı yönetmek için sürekli form değiştirmiştir. Mikroservisler, bu evrimin en olgun ve güncel halkalarından biridir.

### 🏛️ 1. Monolitik Mimari (The Giant)
Tüm iş mantığının, veritabanı erişiminin ve kullanıcı arayüzü kodunun tek bir uygulama paketinde toplandığı yapıdır. 
- **Avantajı:** Geliştirmesi, test etmesi ve deploy etmesi (başlangıçta) kolaydır.
- **Dezavantajı:** Uygulama büyüdükçe "Yazılım Entropisi" artar; kod birbirine girer, ölçeklenebilirlik hantallaşır ve tek bir hata tüm sistemi çökertebilir.

### 🧩 2. Mikroservis Mimarisi (The Swarm)
Bir uygulamayı, her biri belirli bir iş yeteneğini temsil eden, bağımsız olarak geliştirilebilen ve dağıtılabilen servisler kümesi olarak tasarlamaktır. 
- **Temel Felsefe:** "Do one thing and do it well" (Bir işi yap ve en iyisini yap).

---

## 🏗️ Bölüm II: Mikroservislerin 4 Temel Sütunu

Mikroservis mimarisini ayakta tutan dört kritik kavram bulunmaktadır:

### 1. Bounded Context (Sınırlı Bağlam)
Domain-Driven Design (DDD) prensibi olan Bounded Context, her servisin kendi "dilini" ve "sınırlarını" belirler. Örneğin, "Ürün" (Product) nesnesi, Satış servisinde farklı, Stok servisinde farklı özelliklere sahip olabilir. Bu sınırlar, servisler arası kafa karışıklığını önler.

### 2. Otonomi (Autonomy)
Her servis, kendi veritabanına, kendi teknoloji yığınına ve kendi yaşam döngüsüne sahip olmalıdır. Otonom bir servis, diğer servisler kapalı olsa bile kendi işini (mümkün mertebe) yapmaya devam edebilmelidir.

### 3. Ölçeklenebilirlik (Scalability)
Monolitik yapıda tüm sistemi 10 katına çıkarmanız gerekirken, mikroservislerde sadece darboğaz oluşturan (örneğin sadece Sipariş servisi) servisi ölçeklendirerek kaynak tasarrufu sağlarsınız.

### 4. Hata Toleransı (Resilience)
Bir servis çöktüğünde tüm sistemin durmaması (Cascade Failure önleme) mikroservislerin en büyük gücüdür. **Circuit Breaker** gibi desenlerle sistemin geri kalanı korunur.

---

## 📡 Bölüm III: Servisler Arası İletişim Stratejileri

Mikroservisler birbirleriyle nasıl konuşur? İki ana yöntem vardır:

### 1. Senkron İletişim (REST & gRPC)
İsteği gönderen servis, cevap gelene kadar bekler. 
- **gRPC:** Protobuf kullanarak ikili (binary) formatta çok hızlı ve hafif bir iletişim sağlar. Bu projede **Product <-> Order** iletişimi gRPC ile kurgulanmıştır.

### 2. Asenkron İletişim (Event-Driven)
Servisler birbirine doğrudan bağımlı değildir. Bir olay (event) bir mesaj kuyruğuna (Message Broker) bırakılır ve ilgilenen diğer servisler bu olayı tüketir.
- **RabbitMQ:** Bu projede sipariş oluşturulduğunda **Order Service** bir mesaj bırakır ve **Notifier Service** bunu asenkron olarak işler.

---

## 💾 Bölüm IV: Veri Yönetimi ve Dağıtık Sistem Zorlukları

Mikroservis dünyası "bedava öğle yemeği" sunmaz; beraberinde ciddi zorluklar getirir:

- **Eventual Consistency (Nihai Tutarlılık):** Dağıtık sistemlerde verinin tüm servislerde aynı anda güncel olması zordur. Veri zamanla tutarlı hale gelir.
- **CAP Teoremi:** Dağıtık bir sistemde *Tutarlılık (Consistency)*, *Erişilebilirlik (Availability)* ve *Bölünme Toleransı (Partition Tolerance)* özelliklerinden sadece ikisi aynı anda tam olarak sağlanabilir.
- **Saga Deseni:** Birden fazla servisi ilgilendiren işlemleri (Transaction) yönetmek için kullanılan bir akış kontrol yöntemidir.

---

## 🛡️ Bölüm V: API Gateway ve Güvenlik

Mikroservisler dış dünyaya doğrudan açılmamalıdır. **API Gateway**, tüm isteklerin karşılandığı "Zırhlı Kapı"dır:
- **Authentication/Authorization:** Kimlik doğrulama burada yapılır.
- **Rate Limiting:** Sistem aşırı istek yükünden korunur.
- **Routing:** Hangi isteğin hangi iç servise gideceğine karar verilir.

---

## 🛣️ Bölüm VI: Proje Müfredatı ve Uygulama

Bu depo, yukarıdaki teorik bilgileri pratiğe döken 5 aşamalı bir masterclass sunar:

| Seviye | Modül | Teknik Derinlik | Durum |
| :--- | :--- | :--- | :---: |
| 🏗️ | **Modül 1: Mimari** | Clean Architecture, PostgreSQL, Domain Logic | ✅ |
| 🔌 | **Modül 2: İletişim** | gRPC, Protobuf, Synchronous Calls | ✅ |
| 🛡️ | **Modül 3: Gateway** | Edge Service, Echo Framework, Routing | ✅ |
| 📩 | **Modül 4: Mesajlaşma** | RabbitMQ, Event-Driven, Async Consumers | ✅ |
| 🐒 | **Modül 5: Stabilite** | Circuit Breaker, Health Checks, Distributed Tracing | 🏗️ |

---

## 🛠️ Teknik Yığın (The Stack)

- **Backend:** Go (Golang) - Yüksek performanslı dağıtık sistem dili.
- **İletişim:** gRPC (Internal) & REST (External).
- **Mesaj Kuyruğu:** RabbitMQ (Asynchronous processing).
- **Veritabanı:** PostgreSQL (Relational persistence).
- **Orkestrasyon:** Docker & Docker Compose.
- **Gateway:** Echo Web Framework.

---

## 🚀 Hızlı Başlangıç

Sistemi tek bir komutla, tüm altyapı bileşenleriyle birlikte ayağa kaldırabilirsiniz:

```bash
# 1. Projeyi klonlayın
git clone https://github.com/arch-yunus/microservices-101.git

# 2. Docker Compose ile tüm ekosistemi başlatın
make up
```

---

## 📚 Kaynakça ve İleri Okuma

Bu ansiklopediyi hazırlarken faydalanılan başucu kaynakları:
- *Microservices Patterns* - Chris Richardson
- *Building Microservices* - Sam Newman
- *Domain-Driven Design* - Eric Evans
- *The Twelve-Factor App* - 12factor.net

<div align="center">
  <br/>
  <sub>Engineering Excellence Series | **© 2026 arch-yunus**</sub>
</div>
