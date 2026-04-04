# 03-Servisler Arası İletişim Protokolleri

Mikroservis mimarisinde sistemin performansı ve ölçeklenebilirliği, servisler arası iletişim stratejilerine doğrudan bağlıdır. Bu modül, senkron ve asenkron haberleşme protokollerini teknik detaylarıyla ele alır.

---

## 🏗️ İletişim Protokollerinin Karşılaştırılması

| Protokol | Özellikler | Kullanım Senaryosu (Use-case) |
| :--- | :--- | :--- |
| **REST (HTTP/1.1)** | Metin tabanlı (JSON), yaygın destek, esnek. | Dış dünyaya açık API'ler (Public Gateway). |
| **gRPC (HTTP/2)** | İksili (Binary), yüksek performans, tip güvenli. | Servisler arası (Internal) yüksek hızlı iletişim. |
| **Messaging** | Asenkron, kuyruk tabanlı, yüksek hata toleransı. | Olay odaklı (Event-driven) mimari ve uzun süren işlemler. |

---

## 🔌 gRPC (Google Remote Procedure Call)

gRPC, modern mikroservis ekosisteminde iç haberleşme için standart haline gelmiştir:
- **Protocol Buffers (.proto)**: Mesaj yapılarının ve servis arayüzlerinin dilden bağımsız (Language-agnostic) tanımlanmasını sağlar.
- **HTTP/2 Desteği**: Multiplexing ve Header Compression ile düşük gecikme süresi sağlar.
- **Güçlü Tiplendirme (Strong Typing)**: Derleme zamanında hata kontrolü imkanı sunar.

---

## 📩 Event-Driven Architecture (EDA)

Servislerin birbirlerinden tamamen bağımsız (Decoupled) bir şekilde, "olaylar" (Events) üzerinden haberleştiği modeldir:
- **Publisher (Yayıncı)**: İşlemi tamamlayan ve bir olay (örneğin: `OrderCreated`) yayınlayan servis.
- **Subscriber (Abone)**: Bu olayı dinleyen ve kendi iş mantığını tetikleyen servisler (örneğin: Stok Güncelleme, Bildirim Gönderimi).

---

[Geri - 02-İş Ayrıştırma](../02-decomposition/README.md) | [İleri - 04-Veri Yönetimi](../04-data-management/README.md)
