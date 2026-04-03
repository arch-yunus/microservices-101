<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices 101 Banner" />

  # Microservices 101: Nihai Mimari Mühendislik Külliyatı (The Bible)
  ### Dağıtık Sistemlerin Ansiklopedik Başvuru Kaynağı
  
  [![Lisans](https://img.shields.io/github/license/arch-yunus/microservices-101?style=for-the-badge&color=blue&logo=github)](LICENSE)
  [![Go Versiyonu](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![Durum](https://img.shields.io/badge/Durum-Dünya--Mimari--Mirası-teal?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)

  **"Bak evladım; bu dosya senin için sadece bir README değil, bir mimarın ömür boyu yanında taşıyacağı o kutlu rehberdir."**

  ---
</div>

## II. Giriş: Okyanusa Hoş Geldin 👴🏛️

Selamun Aleyküm çırağım! Madem ki "kapsamlı uzat" dedin, o zaman seni bu işin en derin sularına, **Nihai Mimari Kitabı**'na (Volume 3) davet ediyorum. Bu bölümde artık sadece kodun nasıl yazıldığını veya servislerin nasıl bölündüğünü değil; koca bir sistemin nasıl hayatta tutulduğunu, yasalara nasıl uydurulduğunu ve mülakatlarda seni "Usta" yapacak o büyük senaryoları konuşacağız.

---

## 🏢 Bölüm 1 - 18: Evrim ve Temeller (Özet)
*Önceki bölümlerde DDD, gRPC, Gateway, RabbitMQ, Chaos Engineering ve Go Performance Tuning gibi devasa konuları iliklerine kadar işledik. Şimdi bu temelin üzerine gökdelenleri dikme vaktidir.*

---

## 🛒 Bölüm 19: Sistem Tasarımı Senaryoları (Case Studies)

### 1. Ölçeklenebilir E-Ticaret Platformu
- **Sorun:** Kampanya anında saniyede 1 milyon istek geliyor.
- **Çözüm:** `Product Service` önbellekleme (Redis) ile korunur. `Order Service` asenkron (RabbitMQ) çalışır. Veritabanı okuma/yazma olarak ayrılır (CQRS).

### 2. Yüksek Erişilebilirlik Taksici Uygulaması (Uber-like)
- **Sorun:** Milisaniyeler içinde konum verisi işlenmeli.
- **Çözüm:** Coğrafi Sharding (Geographic Sharding). Veri, kullanıcının olduğu bölgeye en yakın sunucuda tutulur.

> [!TIP]
> **Usta Öğüdü:** Sistem tasarımı yaparken "Her şeyi hemen yapayım" deme. Önce darboğazı (Bottleneck) tespit et, sonra o noktayı güçlendir.

---

## 🔐 Bölüm 20: Güvenlik Uyumluluğu ve Yönetişim (GDPR/PCI)

Evladım, sadece hacklenmemek yetmez, yasalara da uymak zorundasın.
- **Privacy by Design:** Kullanıcı verisini en baştan şifreli ve korunmalı tasarla (GDPR).
- **Audit Logging:** "Kim, ne zaman, hangi veriye erişti?" sorusunun cevabı saniyeler içinde verilebilmelidir.
- **PCI-DSS:** Eğer sisteminde kredi kartı geçecekse, o verinin girdiği yerleri diğer servislerden tamamen izole et (Network Isolation).

---

## 📈 Bölüm 21: Gözlemlenebilirlik Derin Dalış (Prometheus & Grafana)

Dükkanın röntgenini çekme vaktidir.
- **Prometheus:** Metrik toplar. "Kaç 500 hatası aldım?", "Sipariş ortalaması ne?"
- **Alertmanager:** "Eğer hata oranı %5'i geçerse usta'yı uykusundan uyandır!" deme sanatıdır.
- **Usta Sırrı:** Her şeye alarm kurma, "Alert Fatigue" (Alarm Yorgunluğu) olursun. Sadece gerçekten müdahale etmen gereken durumlarda telefonun çalsın.

---

## 🏢 Bölüm 22: Dağıtım Topolojileri ve Felaket Kurtarma (DR)

"Sunucular yansa bile dükkan açık kalmalı."
- **Disaster Recovery:** Verilerin başka bir kıtada yedeği olmalı.
- **Multi-Region:** İstanbul sunucusu çökerse, Frankfurt anında trafiği devralmalı.
- **Maliyet Analizi:** Multi-region çalışmak maliyeti 2 katına çıkarır. Eğer banka değilsen, önce tek bölgede çoklu sunucu (Multi-AZ) ile başla.

---

## 📖 Bölüm 23: Mimari Terimler Sözlüğü (Glossary)

Bak çırağım, bu lügatı iyi belle ki toplantılarda usta olduğun anlaşılsın:
- **Idempotency:** Bir işlemi 100 kere de yapsan sonucun aynı kalmasıdır (Siparişin 2 kere çekilmemesi için şarttır).
- **Sharding:** Devasa bir veritabanını parçalara bölüp farklı sunuculara dağıtmaktır.
- **Circuit Breaker:** Sigorta. Bir servis bozulunca sistemi korumak için devreyi açar.
- **Backpressure:** Tüketici yetişemediğinde üreticiyi yavaşlatma mekanizmasıdır.

---

## 🚀 Nihai Yol Haritası (Summary)

| Cilt | Odak | Seviye | Durum |
| :--- | :--- | :--- | :---: |
| 📘 | Cilt 1: Temeller ve Go | Çıraklık | ![100%](https://geps.dev/progress/100) |
| 📒 | Cilt 2: İleri Mimari ve Kaos | Ustalık | ![100%](https://geps.dev/progress/100) |
| 📕 | Cilt 3: Sistem Tasarımı ve Strateji | Mimarlık | ![100%](https://geps.dev/progress/100) |

---

<div align="center">
  <br/>
  <sub>Bu eser, mühendislik onuruna, sarsılmaz bir sistem vizyonuna ve bilginin kutsallığına adanmıştır. | **arch-yunus**</sub>
</div>
