# 05-Merkezi Giriş Kapısı: API Gateway ve Güvenlik Mimarisi

Mikroservis ekosistemleri büyüdükçe, dış dünyadaki istemciler ile iç ağdaki servisler arasında güvenli, kontrollü ve optimize edilmiş bir giriş noktası (API Gateway) kurmak zorunlu hale gelir.

---

## 🏗️ API Gateway Nedir?

API Gateway, tüm istemci isteklerini (Web, Mobil, IoT) karşılayan ve bu istekleri iş mantığına göre ilgili iç servislere yönlendiren (Routing) merkezi bir vekil sunucudur (Reverse Proxy).

### Temel Sorumluluklar:
- **Yönlendirme (Routing)**: İsteklerin URL yapısına göre doğru servislere (`/products` -> Product Service, `/orders` -> Order Service) iletilmesi.
- **Kimlik Doğrulama ve Yetkilendirme (AuthN/AuthZ)**: Kullanıcı erişim haklarının servislerin önünde tek bir merkezden kontrol edilmesi.
- **Yük Dengeleme (Load Balancing)**: Gelen trafiğin servislerin farklı replikalarına sağlıklı bir şekilde dağıtılması.
- **Hız Sınırlama (Rate Limiting)**: Servislerin aşırı yoğunluğa veya DDoS saldırılarına karşı korunması.
- **Protokol Dönüşümü**: Dışarıdan gelen REST/HTTP isteklerinin içeride gRPC veya başka protokollere dönüştürülmesi.

---

## 🔐 Güvenlik Stratejisi: JWT (JSON Web Token)

Dağıtık sistemlerde stateless (durumsuz) bir güvenlik modeli benimsenmelidir. Kullanıcının her serviste yeniden doğrulanması yerine, Gateway katmanında bir kez doğrulanması esastır:

1.  **Giriş İşlemi**: Kullanıcı bilgileri doğrulandıktan sonra Gateway/Auth servisi tarafından dijital imzalı bir **JWT Token** üretilir.
2.  **Erişim Kontrolü**: İstemci, sonraki her istekte bu token'ı `Authorization: Bearer <token>` başlığıyla gönderir.
3.  **Haberleşme Güvenliği**: Gateway, token'ın geçerliliğini ve imzasını kontrol ederek isteği içeriye buyur eder.

---

## 🔒 Ağ İzolasyonu (Network Segmentation)

Güvenli bir mimaride, iç servisler ve veritabanları asla doğrudan dış dünyaya açılmamalıdır:
- **Docker Network**: Tüm servisler Docker üzerinde izole bir "Bridge Network" içinde çalışır.
- **Sınırlı Erişim**: Dış dünya ile sadece API Gateway'in belirlenmiş portu (örneğin: 8080) üzerinden iletişim kurulabilir. Bu, "Saldırı Yüzeyi"ni (Attack Surface) minimuma indirir.

---

## 🚀 Projedeki Uygulama

Bu eğitim kapsamında, Go diliyle geliştirilmiş minimalist bir `gateway-service` yer almaktadır. Bu servis;
- Custom **JWT Middleware** ile yetki kontrolü yapar.
- Dinamik **Reverse Proxy** mekanizması ile trafiği orkestre eder.

---

[Geri - 04-Veri Yönetimi](../04-data-management/README.md) | [Ana README](../../README.md)

---

## 1. API Gateway Nedir?

API Gateway, tüm istemci isteklerini (Mobil, Web, IoT) karşılayan ve bu istekleri ilgili iç servislere yönlendiren merkezi bir giriş noktasıdır.

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
2. **Reverse Proxy** ile istekleri iç ağdaki servislere yönlenditir.

---
[Geri - Ana README](../../README.md)
