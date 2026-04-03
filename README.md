<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices 101 Banner" />

  # Microservices 101: Usta-Çırak Mimari Mektebi
  ### Dağıtık Sistemleri Temelden, Ruhunu Anlayarak İnşa Etmek
  
  [![Lisans](https://img.shields.io/github/license/arch-yunus/microservices-101?style=for-the-badge&color=blue&logo=github)](LICENSE)
  [![Go Versiyonu](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![Durum](https://img.shields.io/badge/Durum-Mimari--Okulu-teal?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)

  **"Bak evladım; her kod yazılır, ama her sistem yaşamaz. Mimari, sistemi ayakta tutan o görünmez omurgadır."**

  ---
</div>

## Usta'nın Selamı: Neden Buradayız? 👴🏛️

Selamun Aleyküm çırağım, hoş geldin!

Eskiden her şey basitti. Bir web sitemiz vardı, bir de veritabanımız. Her şey aynı sepetteydi (Monolith). Sonra dünya büyüdü, kullanıcılar arttı. O devasa sistemler artık yerinden kalkamaz, değişemez hale geldi. Bir yerini değiştirsen başka yeri patlıyordu.

İşte bu yüzden **Mikroservisleri** bulduk. Yani "Büyük lokma yeme, küçük lokmaları birleştir" dedik. Ama dikkat et; bir şeyi parçalamak kolaydır, parçaları bir arada uyum içinde tutmak asıl ustalıktır. Bu okulda sana sadece kod yazmayı değil, sistemi tasarlamayı öğreteceğim.

---

## Bölüm 1: Büyük Resim (Mutfak Benzetmesi) 🗺️✨

Mikroservisleri anlamak için profesyonel bir restoran mutfağını düşün:
- **Waiter (API Gateway):** Müşteriyi karşılar, siparişi alır. Müşteri içerideki kargaşayı görmez.
- **Chef (Service):** Sadece kendi yemeğini yapar. Çorbacı pilava karışmaz, pilavcı tatlıya.
- **Kitchen Intercom (gRPC):** Ustalar arası hızlı ve net iletişim.
- **Post Office (RabbitMQ):** Bir malzeme biterse veya bir yemek hazırsa, bağırmak yerine not bırakırsın. Diğerleri uygun olunca o notu alır.

---

## Bölüm 2: Temeller ve Parçalama Sanatı (DDD)

Evladım, her şeyi ayırmaya kalkma! Mikroservis yapayım derken "Nano-servis" çukuruna düşersin.
- **Catalog Service**: Dükkanın vitrinidir.
- **Order Service**: Kasanın kendisidir.

> [!TIP]
> **Usta Öğüdü:** Eğer iki servis sürekli birbirine soru soruyorsa (aşırı bağımlıysa), onları ayırmış değil, sadece birbirine kabloyla bağlamışsındır. Buna "Dağıtık Monolit" denir, sakın ha!

---

## Bölüm 3: İletişim ve Protokoller (gRPC ve Mesajlaşma)

Ustalar nasıl konuşur?
1.  **gRPC (Doğrudan Hat):** "Bana şu ürünün fiyatını hemen söyle." Beklersin, cevabı alırsın. Hızlıdır, disiplinlidir. ✨
2.  **Asenkron (Mektup):** "Ben bu siparişi aldım, haberiniz olsun." Mesajı bırakır gidersin. Notifier servis o notu alır, e-mailini atar. Kimse kimseyi beklemez.

---

## Bölüm 4: Mimari Zırh ve Koruma (Gateway & Security)

Dükkanın ön kapısı (Gateway) olmazsa, her önüne gelen mutfağa girer.
- **Zırhlı Kapı (Gateway):** Gireni çıkanı denetler.
- **Giriş Kartı (JWT):** "Ben bu dükkanın yetkili müşterisiyim" demenin yolu.
- **Sigorta (Circuit Breaker):** Eğer bir servis hata veriyorsa, diğer sistemleri de yakmasın diye sigortayı attırırsın.

---

## Bölüm 5: "Sakın Ha!" (En Büyük Hatalar) 🚫🛡️

Bak çırağım, piyasada çok can yakan şu hataları kulağına küpe yap:
1.  **Veritabanını Paylaşma:** "Aman ne olacak, ikisi de aynı tabloyu okusun" dediğin an mikroservis biter. Her usta kendi malzemesini kendi dolabında tutar.
2.  **Erken Parçalama:** Daha trafik yokken, 100 tane servis açma. Önce ihtiyacın olsun, sonra ayır.
3.  **Körleme Dağıtım:** Test etmeden, izlemeyi (Monitoring) kurmadan sistem dağıtma. Karanlıkta araba sürmeye benzer.

---

## Bölüm 6: Öğrenme Yolculuğu (Çıraklıktan Ustalığa) 🚀

Bu repoyu nasıl çalışmalısın?
1.  **Kalfalık:** Önce `make up` ile dükkanı aç, servisleri tek tek çalıştır.
2.  **Ustalık:** `infrastructure/docker-compose.yml` dosyasını kurcala, servislerin birbirini nasıl bulduğunu anla.
3.  **Mimarlık:** Kendi servisini ekle (örn: `payment-service`) ve RabbitMQ üzerinden diğerleriyle konuştur.

---

## Mimari Yol Haritası (Müfredat)

| Seviye | Modül | Konu | Durum |
| :--- | :--- | :--- | :---: |
| 🟢 | Temeller | Paradigma ve Mutfak Analojisi | ![100%](https://geps.dev/progress/100) |
| 🟡 | Clean Architecture | Go ile Temiz Kod Yazmak | ![100%](https://geps.dev/progress/100) |
| 🟠 | İletişim | gRPC ve Protobuf Ustası Olmak | ![100%](https://geps.dev/progress/100) |
| 🔴 | Güvenlik | Gateway ve JWT Zırhı | ![100%](https://geps.dev/progress/100) |
| 🟣 | Event-Driven | RabbitMQ ve Asenkron Güç | ![100%](https://geps.dev/progress/100) |

---

## Laboratuvar: Uygulama Rehberi

```bash
# 1. Altyapıyı Başlat (Tezgahı Kur)
make up

# 2. Servisleri Uyandır (Ustaları İşe Başlat)
make run-product
make run-order
```

<div align="center">
  <br/>
  <sub>Mühendislik Onuruna ve Usta-Çırak Geleneğine Sadakatle | **arch-yunus**</sub>
</div>
