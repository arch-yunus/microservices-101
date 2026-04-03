# 02-decomposition: Servis Parcalama & DDD

Bu modül, bir sistemin mikroservis parçalarına nasıl ayrlacana dair en kritik konuyu işler. Yanlış parçalanm bir mikroservis yapısı, **Dağıtk Monolit** kbusuna yol açar.

## ?? Domain-Driven Design (DDD)

DDD, yazlm gelistirmeyi i mantna (Business Logic) odaklayan bir yaklaımıdr. En kritik bileşenleri:

1.  **Bounded Context:** Bir modelin geçerli olduğu sirdr. Örn: "Product" katalog servisinde bir tanm iken, sipariş servisinde sadece bir "ID" ve "Price" dır.
2.  **Aggregate:** Bir veri grubunun (Örn: Order ve OrderItem) tutarlı kalmasn salayan bir birimdir.

## ?? Parcalama Stratejileri

- **Business Capability:** İş fonksiyonlarnı temel alarak (Örn: Sipariş, Ödeme, Katalog).
- **Sub-domain:** DDD alt-alanlarnı kullanarak (Core, Supporting, Generic).

> [!IMPORTANT]
> Servislerin birbirine **Geç Bağlı (Loosely Coupled)** ve **Yksek Uyumlu (High Cohesion)** olması şarttır.

---

[Geri - 01-intro/README.md](../01-intro/README.md) | [İleri - 03-communication/README.md](../03-communication/README.md)
