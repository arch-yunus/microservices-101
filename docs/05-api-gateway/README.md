# Mikroservis Giriş Kapısı: API Gateway ve Güvenlik

Sisteminiz büyüdüğünde, dış dünyadaki müşterilerle iç dünyadaki servisleriniz arasına bir "Zırhlı Kapı" (API Gateway) koymak zorunluluğu doğar.

---

## 1. API Gateway Nedir?

API Gateway, tüm istemci isteklerini (Mobile, Web, IoT) karşılayan ve bu istekleri ilgili iç servislere yönlendiren merkezi bir giriş noktasıdır.

### Temel Görevleri:
- **Yönlendirme (Routing):** `/products` isteğini Ürün servisine, `/orders` isteğini Sipariş servisine gönderir.
- **Kimlik Doğrulama (Authentication):** "Bu kullanıcı kim? Girmeye yetkisi var mı?" sorusunu tek bir merkezde çözer.
- **Yük Dengeleme (Load Balancing):** İstekleri servislerin farklı kopyalarına dağıtır.
- **Hız Sınırlama (Rate Limiting):** Bir kullanıcının saniyede kaç istek yapabileceğini belirleyerek sistemi korur.

---

## 2. Güvenlik Stratejisi: JWT (JSON Web Token)

Mikroservislerde her servisin içinde ayrı ayrı kullanıcı tablosu tutmak yerine, kullanıcıyı bir kez Gateway girişinde doğrularız.

- **Mekanizma:** Gateway, kullanıcı adını ve şifresini doğrularsa kullanıcıya bir "Giriş Kartı" (Token) verir.
- **İç Haberleşme:** Kullanıcı sonraki her isteğinde bu kartı gösterir. Gateway kartın sahte olmadığını kontrol eder ve isteği içeriye buyur eder.

---

## 3. Ağ İzolasyonu (Network Lockdown)

Gerçek mikroservis dünyasında, dışarıdaki birisi asla Ürün servisine veya Veritabanına doğrudan bağlanamamalıdır.
- **Çözüm:** Tüm servisleri Docker içinde özel bir ağa (bridge network) koyarız. Dış dünyaya sadece Gateway'in portunu (8080) açarız.

---

## Bizim Uygulamamızda:

Bu projede, Go diliyle yazılmış minimalist bir `gateway-service` ekledik. Bu servis:
1. **JWT Middleware** ile istekleri denetler.
2. **Reverse Proxy** ile istekleri iç ağdaki servislere yönlendirir.

---
[Geri - Ana README](../../README.md)
