# 04-data-management: Veri Yönetimi & Tutarllk

Mikroservis sisteminde her servis bir veritabanı yneitir. Bu durum, "Veri Tutarlllığı" (Data Consistency) sorununa yol açar.

## ?? Database Per Service

- **Neden?** Her servis kendi ihtıyacına uygun veritabanını semelidir (Örn: Order iin SQL, Search iin NoSQL).
- **Hata:** Tüm servislerin aynı DB'yi (Shared Database) kullanması **En Byk Anti-pattern**'dir.

## ?? Saga Pattern (Daltık İslemler)

Birden fazla servis arasndaki tutarlllığı salamak iin kullanlan daltık ilem (Transaction) modelidr.

### Orchestration Model
- Bir **Orchestrator** (Yonlendirici), tum akısı yonetir. servislerden birinden hata gelirse "Compensating Transactions" tetikler.

### Choreography Model
- Her servis bir islemi bitirdiinde bir event yollar. Dieri bu eventi alıp kendi islemini yapar.

## ?? Eventual Consistency (Nihai Tutarllk)

Sistemde verilerin her an her yerde aynı anda guncel olması (Strong Consistency) mikroservis sistemlerinde performance sorunu yaratr. Bu yuzden **Eventual Consistency** kabul edilir.

---

[Geri - 03-communication/README.md](../03-communication/README.md) | [Ana README](../../README.md)
