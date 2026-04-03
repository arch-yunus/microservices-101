<div align="center">
  <img src="./assets/banner.png" width="100%" alt="Microservices 101 Banner" />

  # Microservices Engineering Academy: Mimari Mühendislik Külliyatı
  ### Dağıtık Sistemler Tasarımı ve Uygulama Müfredatı
  
  [![Akademik Standart](https://img.shields.io/badge/Eğitim-Uluslararası--Standart-blue?style=for-the-badge)](https://github.com/arch-yunus/microservices-101)
  [![Go Versiyonu](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
  [![Müfredat Durumu](https://img.shields.io/badge/Müfredat-Tamamlandı-success?style=for-the-badge)](LICENSE)

  **"Sistematik bir yaklaşım olmadan inşa edilen mimari, sadece şans eseri ayakta duran bir yapıdır."**

  ---
</div>

## 🎓 Program Tanıtımı ve Öğrenme Metodolojisi

Hoş geldiniz. Bu platform, mikroservis mimarisini sadece teorik olarak değil, **pedagojik bir yaklaşımla** ve **sektör standartlarında** öğretmek amacıyla tasarlanmış kapsamlı bir müfredattır. Eğitimimiz, Bloom Taksonomisi'nin "Uygulama" ve "Analiz" seviyelerini hedefleyerek, katılımcıları bir "Yazılım Mimarı" vizyonuna taşımayı amaçlar.

### 🎯 Genel Öğrenme Kazanımları (Learning Outcomes)
Bu programı başarıyla tamamlayan bir mühendis:
1.  **Mimari Karar Verme:** Dağıtık sistemlerdeki trade-off'ları (ödünleşim) analiz edebilir.
2.  **Sistem Parçalama:** Domain-Driven Design ilkelerini kullanarak karmaşık sistemleri yönetilebilir parçalara bölebilir.
3.  **Güvenlik ve Dayanıklılık:** Zero Trust ve Circuit Breaker gibi ileri seviye kalıpları uygulayabilir.
4.  **Stratejik Bakış:** Sistem tasarımı mülakatlarında ve gerçek dünya projelerinde sürdürülebilir altyapılar kurabilir.

---

## 🧬 Modül 101: Mimari Temeller ve Dağıtık Düşünme

### Bölüm 1: Monolitik Yapıdan Mikroservislere Dönüşüm
- **Öğrenme Hedefi:** Yazılım mimarisinin tarihsel evrimini ve mikroservis geçişinin kök nedenlerini anlamak.
- **Teorik Altyapı:** Conway Yasası ve Amdahl Yasası perspektifinde ölçeklenebilirlik analizi.
- **Kontrol Sorusu:** Bir projenin mikroservise geçiş zamanının geldiğini gösteren 3 temel sinyal nedir?

### Bölüm 2: Parçalama Sanatı ve Bounded Context (DDD)
- **Öğrenme Hedefi:** İş birimlerini (Domain) mantıksal sınırlara bölerek veri izolasyonunu sağlamak.
- **Uygulama Senaryosu:** `Order` ve `Product` servisleri arasında "Veri Sahipliği" (Ownership) kuralının uygulanması.
- **Usta Sırrı:** "Don't share databases!" kuralının ihlal edildiği senaryolarda oluşacak zincirleme hatalar.

---

## 🧬 Modül 201: İletişim, Güvenlik ve Altyapı

### Bölüm 3: gRPC ve Yüksek Performanslı İletişim
- **Öğrenme Hedefi:** Protobuf protokolü ile servisler arası telepatik ve tip güvenli bağlantı kurmak.
- **Teorik Fark:** JSON tabanlı REST ile Binary tabanlı gRPC'nin CPU ve Network üzerindeki yük kıyaslaması.

### Bölüm 4: API Gateway ve Zero Trust Security
- **Öğrenme Hedefi:** Merkezi bir giriş noktası üzerinden kimlik doğrulama (JWT) ve ağ izolasyonu sağlamak.
- **Kontrol Sorusu:** Gateway'de yapılan kimlik doğrulamanın (Authentication) iç servislerde tekrar yapılması gerekir mi? (Zero Trust yaklaşımı).

### Bölüm 5: Event-Driven Design ve RabbitMQ
- **Öğrenme Hedefi:** Asenkron mesajlaşma ile servislerin zamansal bağımlılığını (Temporal Decoupling) koparmak.
- **Kriz Senaryosu:** Tüketici servis (Notifier) kapalıyken sipariş akışının nasıl kesintisiz devam edebileceğinin analizi.

---

## 🧬 Modül 301: İleri Mimari Stratejiler ve Sistem Tasarımı

### Bölüm 10-18: Service Mesh, Snowflake ID ve Kaos Mühendisliği
- **Öğrenme Hedefi:** Karmaşık topolojilerdeki trafiği (Istio) yönetmek ve sistemi planlı kaos (Chaos Monkey) ile güçlendirmek.
- **Kazanım:** Beklenmedik hatalara karşı "Resilience" (Esneklik) katsayısını artırma.

### Bölüm 19-23: Sistem Tasarımı ve Uyumluluk (Bible)
- **Öğrenme Hedefi:** Uluslararası standartlarda (GDPR, PCI-DSS) güvenli ve ölçeklenebilir (Uber/Netflix vaka analizleri) yapılar tasarlamak.
- **Kapanış:** A'dan Z'ye mimari lügatın (Idempotency, Sharding, Saga) profesyonel jargonuna hakimiyet.

---

## 📚 Profesyonel Kaynakça ve İleri Okuma (Bibliography)
Eğitim içeriği, aşağıdaki endüstri standartları ve akademik kaynaklar referans alınarak hazırlanmıştır:
- **Newman, S. (2015).** *Building Microservices.* O'Reilly Media.
- **Fowler, M. (2014).** *Microservices Guide.* martinfowler.com
- **Richardson, C. (2018).** *Microservices Patterns.* Manning Publications.
- **The Twelve-Factor App.** 12factor.net

---

## 🧪 Laboratuvar: Uygulama Rehberi

```bash
# 1. Müfredat Ortamını Hazırla (Tezgahı Kur)
make up

# 2. Modüler Testleri Başlat
make run-product
make run-order
```

<div align="center">
  <br/>
  <sub>Profesyonel Mühendislik Standartlarına ve Akademik Disipline Sadakatle | **arch-yunus**</sub>
</div>
