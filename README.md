<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices 101 Banner" />

  # Microservices 101: Mimari Mühendislik Külliyatı (The Sacred Text)
  ### Dağıtık Sistemlerin Felsefesi, Mühendisliği ve Savaş Yaraları
  
  [![Lisans](https://img.shields.io/github/license/arch-yunus/microservices-101?style=for-the-badge&color=blue&logo=github)](LICENSE)
  [![Go Versiyonu](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![Durum](https://img.shields.io/badge/Durum-Nihai--Mimari--Eser-teal?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)

  **"Bak evladım; kod yazarsın unutulur, ama bir sistem kurarsın ya nesiller boyu yaşar ya da bir gecede seni uykusuz bırakır. Burada sadece teknoloji değil, sistemin ruhunu öğreneceksin."**

  ---
</div>

## II. Giriş: Zihinsel Dönüşüm (Paradigm Shift) 👴🏛️

Selamun Aleyküm çırağım! Mikroservislere sadece "servisleri ayırmak" diye bakıyorsan, daha yolun başındasın demektir. Bu iş, bir zihinsel dönüşümdür. Eskiden "Her şey benim elimde, her şeye hakimim" dediğin o güvenli monolitik limandan ayrılıp, fırtınalı ama özgür dağıtık okyanuslara açılıyoruz.

Burada kural şudur: **Hata kaçınılmazdır.** Önemli olan hatanın olması değil, hatanın tüm sitemi çökertmeden nasıl yönetileceğidir. Şimdi sandalyeni daha yakın çek, çünkü her paragrafta sektörün tozunu yutmuş bir ustanın tecrübelerini iliklerine kadar hissedeceksin.

---

## 🧩 Bölüm 1: Parçalama Sanatı ve Domain-Driven Design (DDD)

Mikroservisleri rastgele bölemezsin evladım. Yanlış yerden bölersen, servisler birbirine yapışık ikiz gibi dolanır. İşte burada **Domain-Driven Design (DDD)** ve **Bounded Context** (Sınırlandırılmış Bağlam) devreye girer. Her servisin kendi lügatı, kendi sınırları ve kendi yasaları olmalıdır.

> [!CAUTION]
> **Kriz Senaryosu:** `Product` ve `Order` servislerinin aynı veritabanı tablosuna (örn: `inventory`) eriştiğini düşün. Kampanya döneminde bir servis tabloyu kilitlerse (DB Lock), diğer servis de cevap veremez hale gelir. Tüm site kilitlenir! İşte bu yüzden "Database-per-service" altın kuraldır.

- **Usta Sırrı:** Bir servisi bölmeden önce, o servisin "Veri Sahipliği"ni (Data Ownership) belirle. Eğer iki servis aynı veriye sürekli yazma ihtiyacı duyuyorsa, onları henüz ayırma vaktin gelmemiş demektir.
- **Maliyet Analizi:** Servisleri bölmek; ağ gecikmesi (latency) ve operasyonel karmaşıklığı artırır. Eğer ekibin küçükse, devasa bölünmelere girmek seni bitirir.

---

## 📡 Bölüm 2: İletişim Protokolleri (gRPC ve Telepatik Bağlar)

Mikroservisler arası iletişim, sistemin sinir sistemidir. Biz burada **gRPC** kullanıyoruz. Neden mi? Çünkü REST gibi JSON metinleriyle vakit kaybetmeyiz. İkili (Binary) formatta, Protocol Buffers kullanarak servisleri birbirine telepatik bir bağla bağlıyoruz.

- **Neden gRPC?** Tip güvenliği (Strongly typed) sağlar. Kodun bir tarafında hata yaparsan, diğer taraf derleme anında sana "Durdur!" der.
- **Usta Sırrı:** Dahili (internal) servislerde gRPC, harici (external) dünyada REST/GraphQL kullan. Protokollerin karakterine göre iş yap, usta işi budur.
- **Kriz Senaryosu:** Yüksek trafikli bir sistemde JSON parse etmek CPU'yu sömürür. gRPC kullanarak CPU kullanımını %30-40 düşürebilir ve sunucu maliyetlerini kurtarabilirsin.

---

## 🐋 Bölüm 3: Konteynırlar ve Ağ İzolasyonu (Docker)

Docker sadece "Benim makinemde çalışıyordu" sorununu çözmez evladım. Docker, her servise kendi izole dünyasını verir. Bizim projemizde her servis kendi konteynırında, dış dünyadan kopuk, sadece Gateway üzerinden haberleşecek şekilde tasarlandı.

- **Health Checks:** Bir konteynırın "çıkmış" olması, onun "sağlıklı" olduğu anlamına gelmez. Veritabanına bağlanabiliyor mu? İş yapabiliyor mu? Docker Health Check'leri bu yüzden hayati önemdedir.
- **Kriz Senaryosu:** Bir servisin bellek sızıntısı (Memory Leak) yaptığını ve sunucunun tüm RAM'ini bitirdiğini düşün. Docker limitleri (resource constraints) olmazsa, o tek servis tüm fiziksel makineyi (ve üzerindeki diğer servisleri) çökertir.

---

## 🔐 Bölüm 4: API Gateway ve Zero Trust (The Fortress)

Sistemin ön kapısı (Gateway) zırhlı olmalı. Ama daha da önemlisi, içerideki servislerin de birbirine "şüpheyle" bakmasıdır. Biz burada **JWT (JSON Web Token)** ile güvenliği sağlıyoruz.

- **Stateless Security:** Gateway kullanıcıyı doğrular, bir token verir. Diğer servisler veritabanına gidip "Bu kim?" diye sormaz. Token'ı matematiksel olarak doğrular ve işine bakar.
- **Usta Sırrı:** Token'ların içine sadece gerekli bilgileri (örn: `user_id`) koy. Devasa nesneleri token içinde taşırsan ağ trafiğini yorarsın.
- **Maliyet Analizi:** Merkezi bir yetkilendirme servisi (IAM) kurmak zordur ama güvenlik açıklarının maliyeti dükkanı kapattırır evladım.

---

## 📩 Bölüm 5: Mesajlaşma ve Event-Driven Design (RabbitMQ)

Bak çırağım, bazen "Hemen cevap ver!" demek yerine "Ben işimi yaptım, sen müsait olunca şuna bak" demek gerekir. İşte asenkron haberleşmenin gücü budur. RabbitMQ üzerinden olayları (Events) havaya fırlatırız.

- **Eventually Consistent:** "Şimdi değil ama birazdan her şey düzelecek" demektir. Sipariş servisi işini yapar, Notifier servisi 2 saniye sonra e-mail atar. Kullanıcı mutlu, sistem rahat.
- **Usta Sırrı:** Kuyruğa attığın mesajlar "Idempotent" olmalı. Yani aynı mesaj iki kere gelirse (network hatasından dolayı), sistem hata yapmamalı.
- **Kriz Senaryosu:** E-mail servisinin (Notifier) 1 saatliğine çöktüğünü düşün. Eğer gRPC ile bağlı olsaydı, Sipariş servisi de tıkanırdı. RabbitMQ sayesinde mesajlar kuyrukta bekler, servis gelince 1 saatlik işi 5 dakikada eritir.

---

## 🐒 Bölüm 6: Kaos Mühendisliği ve Esneklik (Circuit Breaker)

Sistemini o kadar sağlam kurmalısın ki, kablolar kopunca bile ayakta kalsın. **Circuit Breaker (Sigorta)** deseni tam olarak budur. Eğer `Product Service` hata verirse, `Order Service` sürekli onu arayıp kendini kilitlemez. Anında "Şu an ürün bilgisi alamıyorum, sonra dene" der ve yoluna devam eder.

- **Felsefe:** Hata beklenmedik bir misafir değildir; o ailenin bir parçasıdır. Onunla yaşamayı öğrenmelisin çırağım.
- **Kaos Deneyi:** Haftada bir kez, en kritik servislerini kasten kapat ve sistemin "Zarif Kapanış" (Graceful Shutdown) yapıp yapmadığını izle.

---

## 📜 Bölüm 7: Mimari Evrim (Tarih Bilmeyen Mimar Olamaz)

Eskiden her şey devasa demir yığınları (Mainframes) içindeydi. Sonra Monolitler geldi. Biz bugün Mikroservisleri konuşuyoruz ama yarın belki Serverless veya başka bir şey konuşacağız.

- **Strangler Fig:** Eski sistemi öldürme, onu yavaşça sar. Yeni servisler eskileri boğdukça sistem modernleşir.
- **Maliyet Analizi:** Bir monoliti bir günde mikroservise çevirmeye çalışmak, dükkanı iflasa sürükler. Bu bir maratondur, sprint değil evladım.

---

## 🚀 Bölüm 8: Öğrenme Yolu ve Ustalık Seviyeleri

1.  **Çıraklık:** `make up` komutunu çalıştır, logları izle, akışı anla.
2.  **Kalfalık:** Bir servisin kodunu değiştir, gRPC ile yeni bir alan (field) ekle ve tüm sistemin nasıl değiştiğini gör.
3.  **Ustalık:** Yeni bir servis ekle (örn: `inventory-service`) ve RabbitMQ üzerinden diğer servislerle konuştur.
4.  **Mimarlık:** Sistemin bir parçasını kasten boz ve Circuit Breaker mekanizmasının sistemi nasıl koruduğunu `pprof` ile analiz et.

---

<div align="center">
  <br/>
  <sub>Bu eser, mühendislik onuruna, usta-çırak geleneğine ve tertemiz bir kod felsefesine adanmıştır. | **arch-yunus**</sub>
</div>
