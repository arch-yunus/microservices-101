# Katkıda Bulunma Rehberi (Contributing Guide)

Bu projenin gelişimine katkıda bulunmak istediğiniz için teşekkür ederiz! `microservices-101`, topluluk destekli bir eğitim platformudur ve her türlü katkı (hata düzeltmeleri, dokümantasyon iyileştirmeleri, yeni özellikler) değerlidir.

## 🤝 Nasıl Katkıda Bulunabilirsiniz?

### 1. Hata Raporlama (Issue Tracking)
Bir hata (bug) tespit ettiyseniz veya mimari bir tutarsızlık fark ettiyseniz, lütfen GitHub üzerinden bir [Issue](https://github.com/arch-yunus/microservices-101/issues) başlatın. Raporunuzda şu bilgileri vermeye özen gösterin:
- Hatanın kısa bir özeti.
- Hatayı yeniden üretme adımları.
- Beklenen davranış ve gerçekleşen davranış.

### 2. Yeni Modüller ve Özellik Önerileri
Eğitim müfredatında eksik olduğunu düşündüğünüz konuları (örneğin: Service Mesh, Sidecar Pattern, Distributed Tracing vb.) önermekten çekinmeyin.

### 3. Kod ve Dokümantasyon İyileştirmeleri
Yazım hataları, kod temizliği veya içerik zenginleştirme için doğrudan bir **Pull Request (PR)** gönderebilirsiniz.

---

## 🚀 Pull Request Süreci

Proje standartlarını korumak adına, gönderilen PR'ların şu kriterlere uyması beklenmektedir:

1.  **Açıklayıcı Başlık**: Yapılan değişikliği net bir şekilde ifade eden bir başlık seçin.
2.  **Referans**: Eğer bir Issue ile ilgiliyse, açıklama kısmında `Closes #123` şeklinde referans verin.
3.  **Dil ve Üslup**: Proje genelindeki teknik ve profesyonel dile sadık kalın. Analojileri sınırlı ve mühendislik ekseninde tutun.
4.  **Test Edilebilirlik**: Eğer kod değişikliği yaptıysanız, `make build-all` komutunun başarılı olduğundan ve sistemin orkestre edilebildiğinden emin olun.

## 📝 Kod Yazım Standartları
- Go kodları için `gofmt` standartlarını takip edin.
- Mikroservisler arasında Clean Architecture prensiplerine uyun.
- Log mesajlarını profesyonel ve takip edilebilir seviyede (Info/Error) tutun.

Bu proje, açık kaynak topluluğunun gücüyle büyümektedir. Desteğiniz için teşekkürler!
