# 04-Veri Yönetimi ve Dağıtık Veri Tutarlılığı

Mikroservis mimarisinde her servisin kendi veri kaynağını (Data Store) yönetmesi esastır. Bu yaklaşım, sistemin otonomisini artırsa da "Dağıtık Veri Tutarlılığı" (Distributed Data Consistency) sorununu beraberinde getirir.

---

## 💾 Database Per Service Prensiibi

Mikroservislerin bağımsız ölçeklenebilmesi ve teknolojik serbestlik (Polyglot Persistence) için her servis kendi veritabanına sahip olmalıdır:
- **Neden?**: Her servis kendi ihtiyacına en uygun veritabanı türünü seçebilir (Örn: Sipariş için PostgreSQL, Arama için Elasticsearch, Oturum yönetimi için Redis).
- **Anti-Pattern**: Birden fazla servisin aynı veritabanı şemasını paylaşması (Shared Database), servisler arasında sıkı bağımlılığa (Tight Coupling) yol açar ve mikroservislerin temel faydalarını ortadan kaldırır.

---

## 🔄 Dağıtık İşlem Yönetimi: Saga Pattern

Geleneksel monolitik sistemlerdeki ACID (Atomicity, Consistency, Isolation, Durability) işlemleri, dağıtık sistemlerde yerini **Saga Pattern** ile yönetilen yerel işlemler dizisine bırakır:

### 1. Orkestrasyon Modeli (Orchestration-based Saga)
Merkezi bir "Saga Orchestrator", tüm iş akışını yönetir ve servisleri koordine eder. Bir hata durumunda telafi edici işlemleri (Compensating Transactions) tetikler.

### 2. Koreografi Modeli (Choreography-based Saga)
Servisler arasında merkezi bir kontrol noktası yoktur. Her servis kendi işlemini tamamladığında bir olay (Event) yayınlar ve diğer servisler bu olayları dinleyerek kendi süreçlerini ilerletir.

---

## ⚖️ Nihai Tutarlılık (Eventual Consistency)

Dağıtık sistemlerde verilerin her an her yerde eşzamanlı olarak güncel olması (Strong Consistency), ciddi performans ve erişilebilirlik sorunlarına yol açar. Bu nedenle çoğu mikroservis yapısında, verilerin belirli bir süre sonunda (eventually) tutarlı hale geleceği **Eventual Consistency** modeli kabul edilir.

---

[Geri - 03-İletişim Stratejileri](../03-communication/README.md) | [Ana README](../../README.md)
