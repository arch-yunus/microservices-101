# Mikroservis Mimarisi: Teknik Derinlik (Masterclass)

Mikroservis, monolitik sistemlerin ölçeklenme ve yönetim krizine bir cevaptır. Bu rehber, dağıtık sistemlerin temel prensiplerini ve ileri seviye tasarım kalıplarını ele alır.

---

## 1. Monolith vs. Mikroservis: Mimari Karşılaştırma

### Monolitik Mimari (Geleneksel)
Tek bir kod deposu, tek bir veritabanı ve tek bir dağıtım birimi.
- **Avantaj**: Geliştirme basit, bağımlılık yönetimi kolaydır.
- **Risk**: Bir hata tüm sistemi etkiler. Teknoloji bağımlılığı yüksektir ve dikey ölçekleme (Vertical Scaling) sınırlıdır.

### Mikroservis Mimarisi (Modern)
Bağımsız servisler, bağımsız veritabanları ve bağımsız dağıtım süreçleri.
- **Avantaj**: Her servis ihtiyaca göre yatay ölçeklenebilir (Horizontal Scaling). Hata izolasyonu yüksektir.
- **Zorluk**: Dağıtık sistem karmaşıklığı, veri tutarlılığı ve ağ gecikmesi yönetimi zordur.

---

## 2. Veri Sahipliği (Data Sovereignty)

Mikroservis dünyasının altın kuralı: **Her servisin kendi veritabanı vardır.**
- **Hatalı Yaklaşım**: Tüm servislerin aynı veritabanına bağlanması (Distributed Monolith).
- **Doğru Yaklaşım**: Order Service (PostgreSQL), Search Service (Elasticsearch), Session Service (Redis).

---

## 3. Servisler Arası Haberleşme Kanalları

Servisler iki temel yolla etkileşime girer:

### A. Senkron İletişim (Request-Response)
- **gRPC / REST**: Bir istek gönderilir ve yanıt beklenir.
- **Dikkat**: Bir servis yavaşsa, onu çağıran tüm zincir yavaşlar (Cascading Failure).

### B. Asenkron İletişim (Event-Driven)
- **Message Brokers (Kafka, RabbitMQ)**: Bir olay (event) yayınlanır ve ilgili servisler bu olayı tüketir.
- **Sonuç**: Sistem çok daha dayanıklı (Resilient) hale gelir.

---

## 4. Saga Pattern: Dağıtık İşlem Yönetimi

Dağıtık sistemlerde global bir `ROLLBACK` yoktur. Bunun yerine **Saga Pattern** kullanılır:
1. Sipariş Servisi işlemi başlatır.
2. Stok Servisi stoğu düşer.
3. Ödeme Servisi ödemeyi alır.

Eğer bir adımda hata olursa, önceki adımları geri alacak "Telafi İşlemleri" (Compensating Transactions) tetiklenir.

---

## 5. Circuit Breaker: Sistemin Sigortası

Bir servis sürekli hata veriyorsa veya çok yavaşsa, ona istek göndermeye devam etmek sistemi kilitler. **Circuit Breaker**, o servise giden yolu geçici olarak kapatarak sistemin geri kalanının hayatta kalmasını sağlar.

---

## 6. Gözlemlenebilirlik (Observability)

Karanlıkta yol bulmak için şu araçlar şarttır:
- **Centralized Logging (ELK/Loki)**: Tüm logların tek bir merkezde toplanması.
- **Distributed Tracing (Jaeger)**: Bir isteğin servisler arasındaki yolculuğunun izlenmesi.
- **Metrics (Prometheus/Grafana)**: Kaynak kullanımının ve performansın takibi.

---

## Sonuç: Ne Zaman Geçilmeli?

Mikroservis operasyonel bir maliyet getirir. 
- Küçük ekipler ve projeler için: **Monolith** ile başlayın.
- Büyük ekipler ve yüksek ölçek ihtiyacı için: **Mikroservis**'e geçin.

---
[Geri - Ana README](../README.md)
