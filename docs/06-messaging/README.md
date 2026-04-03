# Mikroservislerde Mesajlaşma ve Olay Odaklı Tasarım (Event-Driven)

Servisleriniz birbirini beklemek (senkron) zorunda kaldığında, bir tanesinin yavaşlığı tüm sistemi yavaşlatır. **Olay Odaklı Mimari (Event-Driven Architecture)**, servisleri birbirinden tamamen kopararak (decoupling) sistemin performansını ve dayanıklılığını artırır.

---

## 1. Senkron vs. Asenkron Haberleşme

- **Senkron (gRPC/REST):** İsteği gönderirsiniz ve cevap gelene kadar beklersiniz. (Örn: Sipariş verirken ürünün varlığını kontrol etmek).
- **Asenkron (Messaging):** İşinizi bitirir ve sisteme bir "Olay" (Event) fırlatırsınız. Kimin ne zaman yapacağıyla ilgilenmezsiniz. (Örn: Sipariş oluşunca e-mail göndermek).

---

## 2. Temel Kavramlar

- **Producer (Üretici):** Bir olayı (event) oluşturan ve mesaj kuyruğuna (Message Broker) gönderen servis. (Bizim projemizde: `Order Service`).
- **Consumer (Tüketici):** Kuyruğu dinleyen ve gelen mesajlara göre kendi işini yapan servis. (Bizim projemizde: `Notifier Service`).
- **Message Broker:** Mesajları güvenle taşıyan ve kuyruğa alan aracı sistem. (Bizim projemizde: **RabbitMQ**).

---

## 3. Neden Event-Driven?

1. **Ölçeklenebilirlik:** Eğer çok fazla e-mail gönderilecekse, `Notifier Service`'ten 10 tane kopya çalıştırabiliriz.
2. **Hata Dayanıklılığı:** Bildirim servisi o an çökse bile mesajlar RabbitMQ'da bekler. Servis ayağa kalkınca kaldığı yerden devam eder. Hiçbir sipariş bildirimsiz kalmaz!
3. **Performans:** Sipariş servisi e-mailin gönderilmesini beklemez. Saniyeler içinde kullanıcıya "Siparişiniz alındı" der ve geçer.

---

## Bizim Uygulamamızda:

Bu aşamada şu akışı kurduk:
1. **Order Service:** Sipariş işlemini bitirince `order_created_queue` isimli kuyruğa bir mesaj fırlatır.
2. **RabbitMQ:** Bu mesajı belleğinde tutar.
3. **Notifier Service:** Kuyruğu sürekli dinler ve mesaj düştüğü anda "Bildirim Gönderildi" simülasyonunu yapar.

---
[Geri - Ana README](../../README.md)
