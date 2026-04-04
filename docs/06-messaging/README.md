# 06-Mesajlaşma ve Olay Odaklı Mimari (Event-Driven Architecture)

Dağıtık sistemlerin ölçeklenebilirliği ve hata toleransı, servisler arasındaki bağımlılığın (Coupling) minimize edilmesine bağlıdır. **Olay Odaklı Mimari (Event-Driven Architecture - EDA)**, servisleri asenkron bir şekilde birbirine bağlayarak bu ihtiyacı karşılar.

---

## 🏗️ Senkron ve Asenkron Haberleşme Analizi

- **Senkron (gRPC/REST)**: Bir istek gönderilir ve yanıt gelene kadar beklenir. Genellikle anlık veri ihtiyacı olan durumlarda (örneğin: Ürün fiyatının anlık kontrolü) tercih edilir.
- **Asenkron (Messaging)**: Bir olay (Event) bir mesaj kuyruğuna (Message Broker) bırakılır ve gönderici servis işine devam eder. Sistemi "Non-blocking" (bloklamayan) hale getirir.

---

## 🧠 Temel Kavramlar

- **Producer (Üretici)**: Bir iş akışı sonucunda olay (Event) üreten ve bu olayı broker'a ileten bileşendir. (Örnek: `Order Service`).
- **Consumer (Tüketici)**: Belirli bir kuyruğu dinleyen ve gelen mesajları işleyerek kendi görevini yerine getiren bileşendir. (Örnek: `Notifier Service`).
- **Message Broker**: Mesajların güvenli bir şekilde depolanmasını, yönlendirilmesini ve tüketilmesini sağlayan merkezi altyapıdır. (Örnek: **RabbitMQ**).

---

## 🎯 Neden Olay Odaklı Mimari?

1.  **Hata Toleransı (Fault Tolerance)**: Alıcı servis o an çevrimdışı olsa bile mesajlar broker'da bekler. Servis tekrar aktif olduğunda veri kaybı yaşamadan işleme devam eder.
2.  **Yatay Ölçekleme (Scalability)**: Yoğun trafik durumunda, aynı kuyruğu dinleyen birden fazla tüketici örneği (Consumer Instance) çalıştırılarak iş yükü dağıtılabilir.
3.  **Performans ve UX**: Uzun süren işlemler (örneğin: E-mail gönderimi, fatura oluşturma) arka plana atılarak istemciye anında yanıt dönülebilir.

---

## 🚀 Projedeki Uygulama Akışı

Sistemimizde şu Event-Driven senaryosu kurgulanmıştır:
1.  **Sipariş Oluşturma**: `Order Service`, başarılı bir siparişin ardından `order_created` olayını RabbitMQ'ya iletir.
2.  **Mesaj Kuyruğu**: RabbitMQ, mesajı ilgili kuyrukta muhafaza eder.
3.  **Bildirim Gönderimi**: `Notifier Service`, kuyruktan gelen mesajı tüketir ve asenkron olarak bildirim işlemini (simülasyon) gerçekleştirir.

---

[Geri - 05-API Gateway](../05-api-gateway/README.md) | [Ana README](../../README.md)
