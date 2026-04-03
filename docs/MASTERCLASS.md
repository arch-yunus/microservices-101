# ?? Mikroservis Mimarisi: Elite Masterclass

Hoşgeldin kanka. Eğer buradaysan, monolitik sistemlerin yarattıı "Bağımlılık Cehennemi"nden (Dependency Hell) sıkılmış ve gerçek sistem mimarisi nedir onu öğrenmeye gelmişsin demektir. 

Bu rehber, mikroservislerin sadece "kk kod paraları" olmadığını, bir **organizasyonel strateji** olduğunu anlatır.

---

## 1. Monolith vs. Mikroservis: Gerçek Savaş

### Monolitik Mimari (Eski Okul)
Tek bir kod deposu, tek bir veritabanı, tek bir daltm.
- ?? **Avantaj:** Bağımlılık az, daltm basit.
- ?? **Tehlike:** Bir hata her şeyi bitirir. Tek bir dilden çıkamazsın. Sistemin bir kısmı çok ykse, tm sistemi büyütmek zorundasın (Vertical Scaling).

### Mikroservis Mimarisi (Elite Okul)
Bamsz servisler, bamsz veritabanları, bamsz daltm.
- ?? **Avantaj:** Her servisi ihtiyacına gore (CPU/RAM) büyütürsün (Horizontal Scaling). Bir servis kerse, sistemin krisi devreye girer (Graceful Degradation).
- ?? **Zorluk:** "Dağıtk sistemin karmaşıklığı". Veri tutarlılığı ve iletisim yonetimi zordur.

---

## 2. Altın Kural: Veri Sahipliği (Data Sovereignty)

Mikroservis dünyasının birinci kuralı: **Her servisin kendi veritabanı vardır.**

- **Yanlış:** Tüm servislerin aynı PostgreSQL'e banlaması. (Buna "Distributed Monolith" denir ve her iki mimarinin de en kotu ozelliklerini barındırır).
- **Doru:** Order Service -> PostgreSQL, Search Service -> Elasticsearch, Session Service -> Redis.

### Neden?
Çünkü veritabanı şeması deistiinde sadece o servisi gncellemen yeterlidir. Dier servisler bundan etkilenmez.

---

## 3. Servisler Arası Haberleşme

Servisler dilsiz deildir, iki yolla haberleşirler:

### A. Senkron İletişim (Telefon Araması gibi)
- **REST / gRPC:** "Bana bu veriyi ver" dersin, yanıtı beklersin. 
- **Risk:** Eğer aradın servis yavaşsa sen de yavaşlarsın. (Cascading Failure).

### B. Asenkron İletişim (Not Bırakmak gibi)
- **Message Brokers (RabbitMQ, Kafka):** "Sipariş olustu!" diye bir mesajı havaya fırlatırsın. Dinlemek isteyen servisler (Email, Stok, Lojistik) o mesajı alıp kendi iisini yapar.
- **Teşhis:** Sistem cok daha "Resilient" (Dayanıklı) olur.

---

## 4. Saga Pattern: Dağıtk Transaction Yönetimi

Eskiden (Monolith) bir sipariş iptal olunca `ROLLBACK` yapardık, her şey geri alınırdı. Mikroservislerde `ROLLBACK` yoktur.

**Saga Pattern** devreye girer:
1. Sipariş İptal Edildi (Order Service).
2. "Sipariş İptal" mesajı yollanır.
3. Ödeme Servisi mesajı alır, parayı iade eder.
4. Stok Servisi mesajı alır, stoğu artırır.

Eer 3. adımda hata olursa, 2. adımı geri alacak bir "Telafi İslemi" (Compensating Transaction) tetiklenir.

---

## 5. Circuit Breaker: Sistemin Sigortası

Eer bir servis çok yavaşsa veya kerse, onu aramaya devam etmek tüm sistemi kitler. 
**Circuit Breaker** (Sigorta), o servise gıden yolu bir sre kapatır ve "Hata var, deneme!" der. Bu sayede sistemin geri kalanı hayatta kalır.

---

## 6. Observability: Karanlıkta Yol Bulmak

Mikroservislerde bir hata olduunda hangi serviste olduunu bilmek zordur.
- **Centralized Logging (ELK Stack):** Tüm loglar bir merkezde toplanır.
- **Distributed Tracing (Jaeger/Zipkin):** Bir isteğin hangi servislerden getiini izler.
- **Metrics (Prometheus/Grafana):** Hangi servis ne kadar RAM yiyor?

---

## Sonuç: Ne Zaman Kullanmalı?

Mikroservis **bedava değildir**. Operasyonel yk (Ops Overhead) getirir. 
- Eer tek başınaysan veya 3 kişilik bir ekipsen: **Monolith ile başlayın.**
- Eer ekibin büyüdüyse ve servisler birbirinin ayana takılmaya basladıysa: **Mikroservise geçin.**

---
[Geri - Ana README](../README.md)
