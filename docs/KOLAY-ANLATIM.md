# Mikroservis Nedir? (Teknik Olmayan Bir Bakış)

Mikroservis, büyük bir işi bağımsız parçalara bölüp her birini birer uzman gibi çalıştırma sanatıdır. Bu kavramı zihninizde canlandırmak için bir restoran mutfağını örnek alalım.

---

## 1. Monolith: Tek Bir Dev Tencere

Eskiden yazılımlar dev bir kazanın içindeki çorba gibi yapılırdı. Et, sebze, su, tuz ve baharat aynı anda aynı kazanda kaynardı.

- **Kritik Sorun**: Eğer birisi yanlışlıkla fazla tuz atarsa, tüm tencereyi çöpe dökmeniz gerekir. "Sadece tuzu ayırayım" diyemezsiniz çünkü her şey birbirine karışmıştır.
- **Hız Sorunu**: Tencere o kadar büyüktür ki, altını açtığınızda ısınması saatler sürer (Uygulamanın açılma süresi).

---

## 2. Mikroservis: Bağımsız Servis Kaseleri

Mikroservis, o büyük kazanı kırıp içindeki her yemeği ayrı kaselere bölmek demektir. Bir kasede salata, birinde köfte, birinde tatlı ve birinde içecek var.

- **Özgürlük**: Salatacı usta marulu istediği gibi doğrar, köfteci usta eti en modern fırında pişirir. Kimse birbirinin işine karışmaz.
- **Dayanıklılık (Resilience)**: Salatadan sinek mi çıktı? Sadece salata kasesini çöpe atarsınız. Diğer müşteriler köfte ve tatlılarını yemeğe devam edebilir. Sistem tamamen çökmez.
- **Akıllı Ölçekleme (Scalability)**: Eğer herkes salata istiyorsa, tüm mutfağı büyütmek yerine sadece salatacı ustanın yanına birkaç yardımcı verirsiniz.

---

## 3. Haberleşme: Mutfaktaki Telsizler

Kaseler ayrı tezgahlarda olduğu için ustaların birbirleriyle konuşması gerekir.

1.  **Senkron (Anlık)**: Köfteci usta telsize basar: "Bana salata lazım!" der ve cevap gelene kadar telsizin başında bekler. (gRPC / REST)
2.  **Asenkron (Mesajlaşma)**: Köfteci usta bir mantar panoya not asar: "Ben köfteyi pişirdim, lazım olan alsın." Diğer ustalar gelip panodaki notları istedikleri zaman okur. (Kafka / RabbitMQ)

---

## 4. Altın Kural: Veritabanı Sırrı

Her servisin kendi buzdolabı vardır ve anahtarı sadece kendisindedir. Salatacı, köftecinin dolabını izinsiz açamaz. Eğer marul lazımsa kimseden gizlice almaz, sistem üzerinden talep eder.

---

## Sonuç

Mikroservis, "Büyük lokma ye ama işleri parçalara bölerek yönet" demektir. Her işi uzmanına verir, aradaki iletişimi (network) iyi kurarsanız sisteminiz asla yıkılmaz.

---
[Geri - Ana README](../README.md)
