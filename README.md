<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices 101 Banner" />

  # Microservices 101: Mimari Mühendislik ve Tasarım Ansiklopedisi
  ### Dağıtık Sistemlerde Uzmanlaşmak İçin Temel Kaynak
  
  [![Lisans](https://img.shields.io/github/license/arch-yunus/microservices-101?style=for-the-badge&color=blue&logo=github)](LICENSE)
  [![Go Versiyonu](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![Durum](https://img.shields.io/badge/Durum-Mimari--Başvuru--Kaynağı-teal?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)

  **"Mükemmel bir mimari, sadece çalışan kod değil; öngörülebilen, ölçeklenebilen ve evrimleşebilen bir yapıdır."**

  ---
</div>

## Giriş: Mikroservis Ekosistemine Bakış

Mikroservis mimarisi, devasa ve yönetilmesi imkansız hale gelen monolitik sistemlerin yarattığı hantallığa bir çözümdür. Ancak bu mimari "bedava" değildir; beraberinde dağıtık sistemlerin karmaşıklığını getirir. Bu rehber, sizi bu karmaşıklığın içinde bir usta (Architect) gibi yönetmeye hazırlayacaktır.

---

## Bölüm 1: Monolit Krizi ve Mikroservise Geçiş Stratejileri

Yazılım dünyasında hiçbir dev sistem bir gecede mikroservis olarak doğmaz. Genellikle bir Monolit'in sınırlarını zorlamasıyla başlar.

### Strangler Fig (Boğucu İncir) Deseni
Eski bir sistemi (Monolith) tamamen yıkıp yeniden yazmak yerine, servisleri birer birer dışarı çıkarıp trafiği yavaşça yeni servislere yönlendirme stratejisidir. 
- **Zamanla:** Yeni servisler monolitik yapıyı "boğar" ve sonunda eski yapı tamamen devre dışı kalır.

---

## Bölüm 2: İleri Seviye Tasarım Kalıpları (Design Patterns)

Mikroservisler arası bağımlılığı ve hata yayılımını (Cascading Failure) önlemek için şu kalıplar hayati önem taşır:

### 1. Circuit Breaker (Sigorta Deseni)
Bir servis sürekli hata veriyorsa, diğer servislerin onu aramaya devam edip kendilerini kitlemesini önler.
- **Kapalı (Closed):** Normal akış.
- **Açık (Open):** Hata eşiği aşılırsa, isteklere anında "Hata var" döner, hedef servisi yormaz.
- **Yarı Açık (Half-Open):** Servisin düzelip düzelmediğini anlamak için arada birkaç deneme yapar.

### 2. Bulkhead (Bölme Deseni)
Gemi ambarlarındaki bölmeler gibi, bir servisteki bir modülün çökmesinin tüm servisi veya sistemi etkilemesini önlemek için kaynakları (Thread pools, CPU) izole etmektir.

---

## Bölüm 3: Karar Matrisleri ve Teknoloji Kıyaslama

"En iyi teknoloji" yoktur, "ihtiyaca en uygun çözüm" vardır.

### İletişim Protokolleri
| Protokol | Hız / Performans | Kullanım Alanı | Zorluk |
| :--- | :--- | :--- | :--- |
| **REST/JSON** | Orta | Dış dünya, Mobil uygulamalar | Düşük |
| **gRPC/Proto** | **Çok Yüksek** | **Dahili (Svc to Svc)** | Orta |
| **GraphQL** | Esnek | Dinamik veri ihtiyacı olan UIlar | Yüksek |

---

## Bölüm 4: Veri Yönetimi ve Saga Topolojileri

Mikroservislerde en büyük zorluk veridir. Atomik bir `UPDATE` her zaman mümkün değildir.

### Saga Deseni (Dağıtık Transaction)
1. **Koreografi (Choreography):** Her servis bir iş yapınca "Event" fırlatır. Diğerleri bu event'i duyup kendi işini yapar. Merkezi bir yönetici yoktur. (Küçük sistemler için ideal).
2. **Orkestrasyon (Orchestration):** Bir "Saga Manager" vardır. Tüm süreci o yönetir. Hangi adımın bittiğini ve kimin hata yaptığını o bilir. (Karmaşık iş süreçleri için şarttır).

---

## Bölüm 5: API Gateway & Güvenlik (The Fortress)

Sisteminiz büyüdüğünde, dış dünyadaki müşterilerle iç dünyadaki servisleriniz arasına bir "Zırhlı Kapı" (API Gateway) koymak zorunluluğu doğar.

- **JWT (JSON Web Token):** Kullanıcıyı bir kez Gateway girişinde doğrularız. Kullanıcıya bir "Giriş Kartı" verilir ve bu kartla iç ağda dolaşabilir.
- **Rate Limiting:** Bir kullanıcının saniyede kaç istek yapabileceğini belirleyerek sistemi (DDoS gibi) saldırılardan korur.

---

## Bölüm 6: Mesajlaşma ve Olay Odaklı Tasarım (Event-Driven)

Servisleriniz birbirini beklemek zorunda kaldığında, bir tanesinin yavaşlığı tüm sistemi yavaşlatır.
- **Producer / Consumer:** Bir olay (Event) fırlatılır ve ilgili servisler bunu tüketir.
- **RabbitMQ / Kafka:** Mesajları güvenle taşıyan ve kuyruğa alan aracı sistemlerdir.

---

## Bölüm 7: Dağıtık Dünyada Test Stratejileri 🧪

Mikroservislerde "Her şeyi test ettim" demek zordur. Test piramidi burada evrilir:
- **Unit Testing**: Fonksiyonel seviyede hız ve doğruluk.
- **Integration Testing**: Veritabanı ve dış servislerle (containerized) olan bağ.
- **Contract Testing (Pact)**: "Servis A, Servis B'nin gönderdiği JSON yapısını hala anlayabiliyor mu?" sorusunu tüm sistemi çalıştırmadan test etmek.
- **E2E (End-to-End)**: Kullanıcı gözünden tüm akışın testi.

---

## Bölüm 8: Zero Trust ve İç Güvenlik ✨🔐

Dış kapıyı (Gateway) kilitlemek yetmez! İçerideki servislerin de birbirine güvenmemesi gerekir.
- **Service-to-Service Auth**: Her servisin kendi kimliği (Identity) olmalıdır.
- **mTLS (Mutual TLS)**: Servisler birbirleriyle konuşurken sadece şifreli değil, aynı zamanda birbirlerinin kimliğini doğrulayarak konuşurlar.
- **Secret Management**: Şifreler ve API anahtarları asla kodda değil, Vault veya AWS Secrets Manager gibi güvenli yerlerde tutulmalıdır.

---

## Bölüm 9: DevOps ve Otomasyon Hattı (CI/CD) 🚀🏗️

Mikroservislerin "hızlı dağıtım" gücü otomasyondan gelir:
- **Continuous Integration (CI)**: Her kod değişikliğinde testlerin otomatik çalışması.
- **Continuous Delivery (CD)**: Testten geçen kodun otomatik olarak staging/production ortamına çıkması.
- **Canary Release**: Yeni sürümü önce kullanıcıların %5'ine açıp, hata yoksa herkese yaymak.
- **Blue-Green Deployment**: Eski ve yeni sürümü paralel çalıştırıp trafiği anında kaydırmak.

---

## Bölüm 10: Kubernetes ve Bulut Yerel Gelecek 🛰️☁️

Docker tek bir konteyner için iyidir ama 100 tane servisi yönetmek için bir "Orkestra Şefi" lazımdır: **Kubernetes (K8s)**.
- **Auto-Scaling**: Trafik artınca servis kopyalarını otomatik artırır.
- **Self-Healing**: Çöken bir servisi otomatik olarak yeniden başlatır.
- **Service Discovery**: Servislerin birbirini dinamik olarak bulmasını sağlar.

---

## Bölüm 11: Veri Tutarlılığı ve Nihai Sürdürülebilirlik 💾🔄

"Nihai Tutarlılık" (Eventual Consistency) kavramına alışmalısınız.
- **ACID vs BASE**: Mikroservislerde ACID (Katı Tutarlılık) yerine BASE (Esnek ama Erişilebilir) prensipleri öne çıkar.
- **Data Sharding**: Veritabanını parçalara bölerek devasa yükleri yönetmek.

---

## Laboratuvar: Uygulama ve Deney Rehberi

Deponuzda yer alan altyapıyı kullanarak bu kavramları bizzat test edin:

```bash
# 1. Altyapıyı Başlat (Database, RabbitMQ, Redis)
make up

# 2. API Gateway (Zırhlı Kapı) Çalıştır
cd services/gateway-service && go run cmd/api/main.go

# 3. İç Servisleri Uyandır
make run-product
make run-order
```

---

## Mimari Yol Haritası (Mükemmeliyete Giden Yol)

| Adım | Konu | Odak Noktası | Durum |
| :--- | :--- | :--- | :---: |
| 1 | Temeller | Paradigma ve Evrim | ![100%](https://geps.dev/progress/100) |
| 2 | Clean Architecture | Go Proje Standartları | ![100%](https://geps.dev/progress/100) |
| 3 | İletişim | gRPC ve Protobuf | ![100%](https://geps.dev/progress/100) |
| 4 | Güvenlik | API Gateway ve JWT | ![100%](https://geps.dev/progress/100) |
| 5 | Dayanıklılık | Sigorta ve Bölme Kalıpları | ![30%](https://geps.dev/progress/30) |
| 6 | Mesajlaşma | Event-Driven Tasarım | ![100%](https://geps.dev/progress/100) |
| 10 | Kubernetes | Orchestration | ![Yakında](https://img.shields.io/badge/-Yakında-orange) |

<div align="center">
  <br/>
  <sub>Sürekli Gelişim ve Mühendislik Onuruyla | **arch-yunus**</sub>
</div>
