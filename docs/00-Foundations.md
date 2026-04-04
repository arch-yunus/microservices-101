# Mikroservis Mimari Temelleri

Bu belge, dağıtık sistemlerin temel kavramlarını ve monolitik yapılardan mikroservis tabanlı mimarilere geçişin teorik altyapısını ele alır.

---

## 1. Monolitik Mimari (Monolithic Architecture)
Yazılım dünyasının geleneksel yaklaşımı olan monolitik mimaride; kullanıcı arayüzü, iş mantığı (business logic) ve veri erişim katmanları tek bir kod tabanında ve tek bir dağıtım biriminde (deployment unit) birleşir.

### Dezavantajlar:
- **Ölçeklenebilirlik Sorunu**: Uygulamanın sadece bir bölümü yoğun yük altındaysa bile tüm uygulamanın ölçeklenmesi gerekir.
- **Teknoloji Bağımlılığı**: Tüm sistem tek bir teknoloji yığınına (stack) mahkumdur.
- **Hata İzolasyonu**: Herhangi bir modüldeki kritik hata, tüm sistemin çökmesine neden olabilir.

---

## 2. Mikroservis Mimarisi (Microservices Architecture)
Mikroservisler; büyük ve karmaşık uygulamaları, sınırlı bağlamlara (Bounded Context) sahip, bağımsız olarak geliştirilebilir, dağıtılabilir ve ölçeklenebilir küçük servis birimlerine bölme yaklaşımıdır.

### Temel Prensipler:
- **Otonomi**: Her servis kendi veri tabanını yönetir ve kendi dağıtım döngüsüne sahiptir.
- **Hata İzolasyonu (Resilience)**: Bir serviste meydana gelen kesinti, diğer servislerin çalışmasını doğrudan etkilemez.
- **Poliglot Geliştirme**: İhtiyaca göre her servis farklı diller (Go, Rust, Python) veya veritabanları (SQL, NoSQL) ile inşa edilebilir.

---

## 3. Servisler Arası İletişim Stratejileri
Servislerin birbirleriyle nasıl etkileşime girdiği mimarinin başarısını belirler:

- **Senkron İletişim (Request-Response)**: İstemci bir istek gönderir ve yanıt bekler. En yaygın araçlar **REST** ve **gRPC**'dir.
- **Asenkron İletişim (Event-Driven)**: Servisler bir olay (event) yayınlar ve diğer servisler bu olayları tüketir. Bu yaklaşım servisler arasındaki sıkı bağımlılığı (tight coupling) azaltır.

---

## 4. Veri Yönetimi ve Dağıtık İşlemler
Mikroservis dünyasında "Database per Service" prensibi esastır. Veri bütünlüğünü sağlamak için **Saga Pattern** gibi dağıtık işlem yönetimi modelleri kullanılır.

---

## Sonuç
Mikroservisler sadece teknik bir seçim değil, aynı zamanda organizasyonel bir stratejidir. Bu mimari, ekiplerin daha hızlı yazılım geliştirmesine ve sistemi daha esnek bir şekilde yönetmesine imkan tanır.
