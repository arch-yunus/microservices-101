<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices 101 Banner" />

  # Microservices Engineering: Evrensel Mimari Külliyatı (The Omnibus)
  ### Dağıtık Sistemler Mühendisliği ve Stratejik Tasarım Ansiklopedisi
  
  [![Standart](https://img.shields.io/badge/Standart-Uluslararası--Mühendislik-blue?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)
  [![Go Versiyonu](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![Cilt](https://img.shields.io/badge/İçerik-9--Cilt--Külliyat-teal?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)

  **"Bir sistemi inşa etmek sadece başlangıçtır; onu binlerce parça arasında uyumla yaşatmak ise bir sanat eseridir."**

  ---
</div>

## 🎓 Giriş: Evrensel Mimari Yolculuğu 🏛️

Değerli mühendis adayı ve kıymetli meslektaşım!

Eline aldığın bu döküman, sadece bir "README" değil, mikroservis mimarisinin tüm derinliklerini, felsefesini, kod seviyesindeki sırlarımızı ve bulut dünyasının en ileri teknolojilerini kapsayan devasa bir **Evrensel Mimari Külliyatı**'dır (The Omnibus). 

Bu yolculukta, 9 farklı cilt (Volume) boyunca, bir çırağın merakıyla başlayıp, bir mimarın vizyonuna ulaşacaksın. Her satır, sektörün en zorlu projelerinde edinilmiş tecrübelerle, akademik disiplinle ve usta-çırak geleneğinin ruhuyla ilmek ilmek işlenmiştir.

---

## 📘 Modül 101: Temeller ve İletişim (Volume 1-3)
*Cilt 1-3 kapsamında; Monolit krizini, gRPC hızını, Clean Architecture yapısını ve Mesajlaşma (RabbitMQ) sanatını öğrendiniz. Bu temeli sağlam atmadıysanız, kütüphanenin bu ilk raflarını mutlaka sindirin.*

---

## 📘 Modül 201: Altyapı, Güvenlik ve Teori (Volume 4-6)

### Cilt 4: Dağıtık Hesaplama Teorileri ve Kıvamlılık Modelleri 🧬📜
Mikroservisler sadece kod yazmak değil, matematiksel bir denge kurmaktır.
- **CAP Teoremi:** Consistency (Kıvamlılık), Availability (Erişilebilirlik) ve Partition Tolerance (Bölünme Dayanıklılığı) arasındaki o meşhur üçgen. Üçünü birden aynı anda alamazsın evladım!
- **PACELC Teoremi:** CAP'ın yetmediği yerde, "Hata yokken Network gecikmesi (Latency) mi, Kıvamlılık mı?" sorusunu sorandır.
- **Kıvamlılık Modelleri:** Strong, Eventual ve Monotonic kıvamlılık arasındaki o ince çizgi.

### Cilt 5: Kaynak Kod Derin Analizi (Source Walkthrough) 💻🔍
Kodumuzun kalbine iniyoruz.
- `order-service` içindeki **Order Placement** akışını satır satır inceliyoruz: Nerede veritabanına yazıyoruz, nerede RabbitMQ event'ini fırlatıyoruz ve neden "Transactional Outbox" desenine ihtiyacımız var?
- **Domain-Driven Design (DDD)** katmanlarımızın (`domain`, `service`, `repository`) Go içindeki somut karşılıkları.

### Cilt 6: Bulut-Yerel ve Kubernetes Masterclass (K8s) 🛰️☁️
Docker tek başına yetmez, bir orkestra şefi lazımdır.
- **K8s Objeleri:** Pod, Service, Deployment ve Ingress.
- **HPA (Autoscaling):** Trafik artınca sunucuları otomatik artırmanın formülü.
- **Sidecar Pattern:** Loglama ve izleme (Monitoring) araçlarını servislerin yanına bir "gölge" gibi nasıl yerleştiririz?

---

## 📘 Modül 301: Veri, Güvenlik ve Strateji (Volume 7-9)

### Cilt 7: İleri Veri Modelleme ve Sharding Stratejileri 💾⚔️
Veri, mikroservisin en kıymetli ama en tehlikeli parçasıdır.
- **SQL vs NoSQL:** Ne zaman Postgres, ne zaman Redis, ne zaman Mongo?
- **Database Sharding:** Milyarlarca satır veriyi tek bir veritabanına sığdıramadığında, onu nasıl parçalara bölüp farklı sunuculara dağıtırsın?
- **Saga Deseni:** Dağıtık transaction'lar (2PC vs Saga) ve veritabanı tutarlılığını sağlamanın stratejik yolları.

### Cilt 8: Tehdit Modelleme ve İleri Güvenlik 🔐🛡️
Dış kapıyı (Gateway) kilitledik ama içerideki casuslar ne olacak?
- **Zero Trust:** Hiçbir servis, diğerine "içerideyiz diye" güvenmez. mTLS ile karşılıklı kimlik doğrulaması şarttır.
- **OAuth2 & OIDC:** Kullanıcı yetkilendirmesinin (Authorization) dünya standartlarındaki uygulama detayları.
- **Secret Management:** Şifreleri Vault'ta saklama ve dinamik şifre üretimi.

### Cilt 9: Sistem Tasarımı ve Mülakat Hazırlığı 🏢📊
Bu cilt seni en üst düzey mühendislik rollerine hazırlar.
- **Vaka Analizi: Instagram Mimarisi:** Saniyede milyonlarca görselin yüklenmesi ve dağıtılması nasıl tasarlanır?
- **Vaka Analizi: Uber Lokasyon Takibi:** Milyonlarca arabanın anlık konumunu nasıl takip eder ve kullanıcıya saniyede 10 kere güncellersin?
- **Usta Öğüdü:** Mülakatlarda sadece "yaparım" deme; trade-off'ları (bu yolu seçersem şunu kaybederim) anlat ki usta olduğun anlaşılsın.

---

## 📖 Mimari Terimler Sözlüğü & Kaynakça (Omnibus Edition)

Bu külliyatta kullanılan tüm terimlerin (Idempotency, Sharding, Bounded Context, Saga, Circuit Breaker vb.) en detaylı tanımları ve Martin Fowler'dan Sam Newman'a kadar akademik atıflar bu bölümde yer Almaktadır.

---

## 🧪 Laboratuvar: Uygulama Rehberi

```bash
# 1. Müfredat Ortamını Hazırla (Tezgahı Kur)
make up

# 2. Modüler Testleri Başlat (Uygulamalı Analiz)
make run-all
```

<div align="center">
  <br/>
  <sub>Evrensel Mühendislik Mirasına ve Bilginin Kutsallığına Sadakatle | **arch-yunus**</sub>
</div>
