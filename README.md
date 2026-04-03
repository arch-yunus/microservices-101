<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices 101 Banner" />

  # Mikroservis Masterclass: Mühendislik Serüveni 🚀
  ### Adım Adım, Uygulamalı Mikroservis İnşa Etme Rehberi
  
  [![Eğitim Vizyonu](https://img.shields.io/badge/Eğitim-Masterclass-blue?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)
  [![Go Versiyonu](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![Durum](https://img.shields.io/badge/Öğrenme-Odaklı-success?style=for-the-badge)](LICENSE)

  **"Bak evladım; okursan unutursun, görürsen hatırlarsın, ama yaparsan öğrenirsin."**

  ---
</div>

## 🎓 Hoş Geldiniz: Bu Reponun Amacı Nedir?

Selamun Aleyküm çırağım! Bu repo, mikroservis dünyasını "karşıdan izlemek" için değil, **bizzat inşa ederek öğrenmek** için tasarlandı. Burada devasa bilgi yığınları arasında boğulmayacaksın; her adımda yeni bir servis ekleyecek, sistemleri bozacak ve gerçek bir mimar gibi düşüneceksin.

---

## 🗺️ Adım Adım Yol Haritası (Müfredat)

Aşağıdaki adımları sırayla takip ederek, temelden en ileri seviyeye kadar mikroservis mimarisini öğreneceksin:

| Aşama | Konu | Temel Kavram | Zorluk |
| :--- | :--- | :--- | :---: |
| 🛡️ 1 | **Temeller & Go** | Proje Yapısı ve Go Modülleri | 🟢 |
| 🔌 2 | **İletişim (gRPC)** | Servisler Arası Doğrudan Hat | 🟡 |
| 🔐 3 | **Güvenlik (Gateway)** | Zırhlı Kapı ve JWT Giriş Kartı | 🔴 |
| 📩 4 | **Mesajlaşma (RabbitMQ)** | Asenkron Haberleşme ve Mektuplar | 🟣 |
| 🐒 5 | **Dayanıklılık (Kaos)** | Sigorta (Circuit Breaker) ve Hata Yönetimi | 🟠 |

---

## 🧪 Ders 1: Servislerin Doğuşu (Fundamentals)

**Baba'nın Analojisi:** Mikroservisleri bir restoran mutfağı gibi düşün. Her servisin tek bir görevi var (Çorbacı, Pilavcı). Biz önce bu tezgahları kuracağız.

### 💻 Laboratuvar:
Projeyi ayağa kaldırmak için şu komutu çalıştır:
```bash
# Tüm dükkanı (Database, Redis, RabbitMQ) açıyoruz.
make up
```

### 🎯 Meydan Okuma:
Tüm servisler ayağa kalktıktan sonra `services/product-service` loglarını izle. Servis veritabanına bağlanabiliyor mu? Eğer bağlantı yoksa, Docker container'ının ismini kontrol et!

---

## 🔌 Ders 2: Servislerin Konuşması (gRPC)

**Baba'nın Analojisi:** Artık mutfaktaki ustaların birbiriyle hızlı konuşması lazım. "Ürün var mu?" "Var." gRPC, bu ustaların arasındaki doğrudan ve çok hızlı dahili hattır.

### 💻 Laboratuvar:
```bash
# Order servisi ile Product servisini gRPC üzerinden konuştur:
make run-product
make run-order
```
`order-service` içinden ürün sorgulaması yapıldığında loglarda "gRPC Request" ifadesini görmelisin!

### 🎯 Meydan Okuma:
`proto/product/product.proto` dosyasını aç ve yeni bir alan (field) ekle: `string description`. Sonra tüm kodda bu alanı taşımaya çalış. Bakalım her şey hala çalışıyor mu?

---

## 🔐 Ders 3: Zırhlı Kapı (API Gateway)

**Baba'nın Analojisi:** Müşteri kime sipariş verecek? Doğrudan mutfağa giremez. Kapıda bir **Garson (Gateway)** durur, siparişi alır ve içeriye iletir.

### 💻 Laboratuvar:
```bash
# Gateway'i çalıştır ve dışarıdan (Port 8080) istek at:
cd services/gateway-service && go run cmd/api/main.go
```

### 🎯 Meydan Okuma:
JWT token'ı olmadan Gateway'den içeri girmeye çalış (Unauthorized - 401). Bakalım garson seni içeri alacak mı?

---

## 📩 Ders 4: Mektuplar ve Kuyruklar (Messaging)

**Baba'nın Analojisi:** Sipariş bittiğinde e-mail atmak için garsonun beklemesine gerek yok. Bir not yazar ve **Posta Kutusuna (RabbitMQ)** bırakır. Bildirim servisi gelince oradan alır.

### 🎯 Meydan Okuma:
`notifier-service`'i durdur ve bir sipariş oluştur. Sonra servisi geri aç. Mesajın kuyrukta beklemiş olduğunu ve servis açılınca hemen işlendiğini gözlemle!

---

## 🚀 Mezuniyet ve İleri Seviye
Tüm bu adımları tamamladıysan; tebrikler çırağım, artık bir kalfasın! İleri seviye konular (Kubernetes, Service Mesh, Chaos Engineering) için `docs/` klasöründeki derinleşme rehberlerini inceleyebilirsin.

<div align="center">
  <br/>
  <sub>Bu repo, mikroservislerin ruhunu bizzat yaparken öğrenmen için inşa edildi. | **arch-yunus**</sub>
</div>
