<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices 101 Banner" />

  # Microservices 101: Mimari Mühendislik Ansiklopedisi (Cilt 2)
  ### Dağıtık Sistemlerin Geçmişi, Bugünü ve Geleceği
  
  [![Lisans](https://img.shields.io/github/license/arch-yunus/microservices-101?style=for-the-badge&color=blue&logo=github)](LICENSE)
  [![Go Versiyonu](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![Durum](https://img.shields.io/badge/Durum-Dünya--Mimari--Mirası-teal?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)

  **"Bak evladım; geçmişini bilmeyen mimar, geleceğin hatalarını bugünden inşa eder."**

  ---
</div>

## Usta'nın Selamı: Neden Buradayız? 👴🏛️

Selamun Aleyküm çırağım! Madem ki "daha çok bilgi" dedin, o zaman seni bu işin mutfağından alıp, mühendislik okyanusuna fırlatıyorum. Bu bölümde sadece kod değil, strateji, tarih ve vizyon konuşacağız. Hazır ol, bu yolculuk seni bir "yazılımcıdan" bir "mimara" dönüştürecek.

---

## Bölüm 1 - 11: Temeller (Cilt 1 Özeti)
*Önceki bölümlerde monolit krizini, gRPC hızını, Gateway zırhını ve RabbitMQ'nun mektup taşıma sanatını öğrendin. Eğer oraları atladıysan, geri dön ve temelini sağlam at evladım.*

---

## Bölüm 12: Mimari Evrim ve Tarihçe (Nereden Geldik?) 📜👴

Yazılım dünyası bir günde mikroservis olmadı.
1.  **Mainframe Çağı:** Tek bir dev bilgisayar, her şey onun içindeydi.
2.  **Monolith:** Uygulama büyüdü ama hala tek vücuttu. (Basit ama hantal).
3.  **SOA (Service Oriented Architecture):** Servisleri ayırdık ama "Enterprise Service Bus" (ESB) denilen merkezi devasa bir kabloya bağladık. Kablo bozulunca her şey bitti.
4.  **Microservices:** Bağımsızlık ilan edildi! Her servis kendi başına bir devlet.

---

## Bölüm 13: Service Mesh ve Görünmez Bağlantılar (Istio/Linkerd) 🕸️🛡️

Evladım, 100 tane servisin olunca onları manuel yönetemezsin. İşte burada **Service Mesh** devreye girer.
- **Sidecar:** Her servisin yanına bir "koruma görevlisi" (Proxy) dikersin. Servis sadece işini yapar, güvenlik ve trafik yönetimini Sidecar halleder.
- **Istio:** Bu korumaların şefidir. Kim kime ne kadar paket göndermiş, kim hata yapmış hepsini o görür.

---

## Bölüm 14: Dağıtık ID Üretimi (Snowflake Sanatı) 🆔✨

Büyük sistemlerde veritabanının `auto_increment` özelliği patlar. Farklı veritabanlarındaki veriler çakışır.
- **Snowflake ID:** Twitter'ın bulduğu bu yöntemle; zaman damgası (timestamp), makine kimliği ve bir sıra numarası birleştirilir. Sonuç: Milyarlarca veride asla çakışmayan, zamana duyarlı benzersiz kimlikler.

---

## Bölüm 15: Kaos Mühendisliği (Sistemi Kasten Bozmak) 🐒💥

"Sistemim çok sağlam" demekle olmaz, test edeceksin!
- **Chaos Monkey:** Netflix'in meşhur maymunu. Rastgele servisleri kapatır, kabloları çeker. Eğer sistemin bu yıkımda hala ayaktaysa, işte o zaman "Ustayım" diyebilirsin.
- **Mantık:** Hataları öngörmek yerine, hataları yaratıp sistemin tepkisini ölçmek.

---

## Bölüm 16: Go Performans Tuning (Sıkmak Değil, Terletmek) ⚡🔧

Go dili hızlıdır ama yanlış kullanılırsa RAM canavarına dönüşür.
- **pprof:** Servisinin röntgenini çeker. Hangi fonksiyon CPU'yu sömürüyor, nerede bellek sızıntısı var (Memory Leak) bir bir gösterir.
- **Garbage Collection (GC):** Go'nun temizlikçisidir. Çöpü ne zaman toplayacağını ona sen fısıldayabilirsin.

---

## Bölüm 17: 12-Factor App (Modern Yazılımın Anayasası) 🏗️✅

Bir uygulama "Bulut Yerel" (Cloud-Native) olacaksa şu 12 kurala uymalıdır:
- **Stateless:** Servisin içinde bilgi saklama, her şey dışarıda (Redis/DB) olsun.
- **Config:** Ayarlarını kodun içine değil, ortam değişkenlerine (Environment Variables) koy.

---

## Bölüm 18: Devlerin Savaşı (Netflix, Uber, Spotify) ⚔️📊

- **Netflix:** Mikroservislerin kralıdır. Devasa trafikleri Chaos Engineering ile yönetirler.
- **Uber:** Bazı servislerini çok parçaladığı için performans sorunu yaşadı ve "Macro-services" (Daha mantıklı büyüklükte servisler) yapısına geri döndü.
- **Öğüt:** Körleme her şeyi parçalama; ihtiyacın kadar böl, usta işi yap!

---

## Mimari Yol Haritası (Müfredat - Cilt 2)

| Seviye | Modül | Konu | Durum |
| :--- | :--- | :--- | :---: |
| 🛡️ | Tarihçe | Evrim ve SOA | ![100%](https://geps.dev/progress/100) |
| 🕸️ | Service Mesh | Sidecar ve Istio | ![100%](https://geps.dev/progress/100) |
| 🆔 | Veri | Snowflake ID | ![100%](https://geps.dev/progress/100) |
| 🐒 | Kaos | Chaos Monkey | ![100%](https://geps.dev/progress/100) |
| ⚡ | Performans | pprof ve Tuning | ![100%](https://geps.dev/progress/100) |

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
  <sub>Mühendislik Onuruna ve Bilginin Sonsuzluğuna Sadakatle | **arch-yunus**</sub>
</div>
