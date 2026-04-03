<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices 101 Banner" />

  # Microservices 101: Mimari Mühendislik Rehberi
  ### Daıtık Sistemleri Hassasiyetle Tasarlamak ve Yönetmek
  
  [![Lisans](https://img.shields.io/github/license/arch-yunus/microservices-101?style=for-the-badge&color=blue&logo=github)](LICENSE)
  [![Go Versiyonu](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![Durum](https://img.shields.io/badge/Durum-Profesyonel--Eğitim--Seti-teal?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)

  **"Mimari, karmaşıklığı yönetmek ve ayrıştırılmış mühendislik yoluyla otonomiyi sağlamakla ilgilidir."**

  ---
</div>

## Giriş: Mikroservislerin Felsefesi

Mikroservis mimarisi, yazılım mühendisliğinin uç ölçek ve karmaşıklığı yönetmek için geçirdiği evrimdir. Monolitik "her şey bir arada" modelinden, her servisin belirli bir iş alanında uzmanlaşmış, bağımsız ve otonom bir birim olduğu hücresel bir yapıya geçişi temsil eder.

Bu depo, güvenilirlik, ölçeklenebilirlik ve sürdürülebilirliğe odaklanarak dağıtık sistemlerde ustalaşmak için hedeflenmiş bir yol haritası sunar.

---

## Bölüm 1: Monolit Krizi ve Evrim

Yazılım basit başlar. Tek bir depo, tek bir veritabanı. Ancak organizasyonlar büyüdükçe, monolit bir darboğaz haline gelir:
- **Sıkı Bağlılık (Tight Coupling)**: Bir modüldeki değişiklik diğerini bozar.
- **Ölçekleme Sorunları**: Sadece bir özellik yük altında olsa bile tüm uygulamayı ölçeklendirmek zorundasınızdır.
- **Operasyonel Yük**: Dev muazzam derleme süreleri ve katastrofik arıza riski.

Mikroservisler, ekipleri ve teknolojiyi birbirinden ayırarak, bağımsız dağıtım ve uzmanlaşmış ölçeklendirme sağlayarak bu sorunu çözer.

---

## Bölüm 2: Parçalama Sanatı (DDD)

Bir sistemi bölmek stratejik bir karardır. **Domain-Driven Design (DDD)** kullanarak **Bounded Context**'leri (Sınırlandırılmış Bağlamlar) belirliyoruz.
- **Catalog Service**: Ürün kimliğini ve fiyatlandırmasını yönetir.
- **Order Service**: Sipariş yaşam döngüsünü ve yerine getirmeyi yönetir.

Bu projede, bu ayrımı göstermek için `services/product-service` ve `services/order-service` yapılarını farklı varlıklar olarak uyguladık.

---

## Bölüm 3: İletişim ve Protokoller

Bağımsız servisler nasıl konuşur? Performans ve esnekliği dengeliyoruz.

1.  **gRPC (Senkron)**:
    - **Kullanım**: Dahili, servisten servise çağrılar.
    - **Neden**: Protocol Buffers, yüksek performanslı ve güçlü tipli ikili iletişim sağlar. ✨
2.  **Mesajlaşma (Asenkron)**:
    - **Kullanım**: Olay tabanlı (Event-driven) mimari.
    - **Neden**: Servisleri birbirinden ayırır, böylece sistemi çökertmeden bağımsız olarak arızalanabilir ve kurtulabilirler.

---

## Bölüm 4: Veri Yönetimi Stratejileri

Altın kural: **Her servis kendi veritabanına sahiptir.**
- **Saga Deseni**: Global kilitler olmadan servisler arası işlemleri yönetmek.
- **CQRS**: Maksimum sorgu performansı için Okuma ve Yazma modellerini ayırmak.

---

## Bölüm 5: Operasyonel Mükemmellik (API Gateway)

Servisi inşa etmek mücadelenin sadece yarısıdır. Ölçekli bir şekilde çalıştırmak şunları gerektirir:
- **API Gateway**: Tüm trafiği karşılayan zırhlı bir giriş kapısı (Örn: `services/gateway-service`).
- **Zarif Kapanış (Graceful Shutdown)**: Bağlantıları temiz bir şekilde kapatmak için işletim sistemi sinyallerini yönetmek.
- **Sağlık Kontrolleri (Health Checks)**: Trafik akmaya başlamadan önce bağımlılıkların hazır olduğundan emin olmak.

---

## Laboratuvar: Uygulama Rehberi

Depo, tüm geliştirme ortamını yönetmek için bir `Makefile` ile donatılmıştır:

```bash
# Altyapıyı Başlat (PostgreSQL, RabbitMQ, Redis)
make up

# Ürün Servisini Çalıştır (Go)
make run-product

# Sipariş Servisini Çalıştır (Go)
make run-order
```

---

## Mimari Yol Haritası

| Aşama | Modül | Odak Noktası | Durum |
| :--- | :--- | :--- | :---: |
| 1 | Temeller | Paradigma Değişimi | ![100%](https://geps.dev/progress/100) |
| 2 | Clean Architecture | Go Proje Düzeni | ![100%](https://geps.dev/progress/100) |
| 3 | İletişim | gRPC Protokolleri | ![100%](https://geps.dev/progress/100) |
| 4 | Konteynerizasyon | Docker ve Ağlar | ![100%](https://geps.dev/progress/100) |
| 5 | API Gateway | Güvenlik ve Yönlendirme | ![100%](https://geps.dev/progress/100) |
| 6 | Mesajlaşma | Kafka/RabbitMQ | ![Soon](https://img.shields.io/badge/-Yakında-orange) |

<div align="center">
  <br/>
  <sub>Mühendislik Mükemmelliği | **arch-yunus**</sub>
</div>
