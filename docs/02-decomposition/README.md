# 02-Servis Ayrıştırma ve Domain-Driven Design (DDD)

Bu modül, bir sistemin mikroservis bileşenlerine nasıl ayrıştırılacağı konusundaki stratejik yaklaşımları ele alır. Yanlış tasarlanmış bir ayrıştırma süreci, sistemin "Dağıtık Monolit" (Distributed Monolith) yapısına dönüşmesine neden olabilir.

---

## 🏗️ Domain-Driven Design (DDD) Prensipleri

DDD, yazılım geliştirmede karmaşıklığı yönetmek için iş mantığına (Business Logic) odaklanan bir metodolojidir. Mikroservis ekosisteminde en kritik bileşenleri şunlardır:

1.  **Bounded Context (Sınırlı Bağlam)**: Bir modelin veya terimin geçerli olduğu sınırdır. Örneğin; "Ürün" (Product) kavramı Katalog Servisi'nde detaylı bir içerik iken, Sipariş Servisi'nde sadece bir "Kimlik" (ID) ve "Fiyat" (Price) verisinden ibarettir.
2.  **Aggregate (Küme)**: Veri tutarlılığını sağlamak amacıyla bir arada yönetilmesi gereken nesneler grubudur. Örneğin; Sipariş (Order) ve Sipariş Kalemleri (OrderItems) tek bir Aggregate kökü üzerinden yönetilir.

---

## 🎯 Ayrıştırma Stratejileri (Decomposition)

Sistemi servisleşirken şu iki ana strateji izlenebilir:

- **İş Yeteneklerine Göre (Decomposition by Business Capability)**: Organizasyonun iş fonksiyonlarını temel alır. Örnek: Sipariş Yönetimi, Stok Kontrolü, Müşteri Hizmetleri.
- **Alt Alanlara Göre (Decomposition by Sub-domain)**: DDD alt alanlarını (Core, Supporting, Generic) temel alır. Örnek: Ödeme sistemleri (Core), Bildirim sistemleri (Generic).

---

> [!IMPORTANT]
> Başarılı bir mikroservis mimarisi için servislerin **Düşük Bağımlılıklı (Loosely Coupled)** ve **Yüksek Uyumluluklu (High Cohesion)** olması temel şarttır.

---

[Geri - 01-Giriş](../01-intro/README.md) | [İleri - 03-İletişim Stratejileri](../03-communication/README.md)
