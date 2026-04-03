# 03-communication: Haberleşme Protokolleri

Mikroservis sisteminde her şey servisler arası iletisimdr. İletisimin performansı, sistemin genel performansıdır.

## ?? Protokollerin Savaşı

| Protokol | Ozellik | Use-case |
| :--- | :--- | :--- |
| **REST (HTTP)** | Esnek, dille bamsz, yaygn. | Dış API'ler, Public gatewayler. |
| **gRPC (HTTP2)** | Yuksek performans, tip-guvenli, binary. | Servisler arası (Internal) iletisim. |
| **Messaging** | Asenkron, kuyruk yapısı. | Event-driven, uzun sren islemler. |

### gRPC (Google Remote Procedure Call)
- **Protocol Buffers (.proto):** Mesajların dille bamsz tanımlanmasını salar.
- **Bi-directional Streaming:** Veri akını iki yonldur.

## ?? Event-Driven Architecture (EDA)

Servislerin birbirini tanımadan "olaylar" (Events) uzerinden haberlesmesidr.

- **Publisher:** "Sipariş Olusturuldu" olayını yayınlar.
- **Subscriber:** Bu olayı dinleyen "Stok" ve "Bildirim" servisleridir.

---

[Geri - 02-decomposition/README.md](../02-decomposition/README.md) | [İleri - 04-data-management/README.md](../04-data-management/README.md)
