# 01-intro: Mikroservis Nedir?

Bu modl, monolitik bir sitemden mikroservis mimarisine neden ve nasl geileceine dair temel teoriyi ele alr.

## ?? Giris

Mikroservis mimarisi, byk bir yazlm sistemini, her biri belirli bir i mantna (Business Logic) sahip, bamsz olarak daltlabilen ve leklenilebilen kk servislerin birleşimi olarak tasarlama yaklaımıdr.

### Monolitik Mimari (The Monolith)
- **Avantaj:** Basit dağıtım, kolay test etme.
- **Dezavantaj:** Tek bir hata tm sistemi korur, teknoloji bağlılığı, uzun daltm sreleri.

### Mikroservis Mimarisi
- **Avantaj:** Teknolojik serbestlik, bamsz lekleme, ekal ve hzl daltm.
- **Dezavantaj:** Dağıtk sistem karmaklığı, veri tutarllğı zorluklar, operasyonel yuk.

## ?? Ne Zaman Mikroservis?

Eer projeniz:
- Byk bir ekip (100+ gelistirici) ise,
- Farklı bileşenlerin farklı ölçeklendirme gereksinimleri (CPU vs Memory) varsa,
- Hızlı srekli teslimat (Continuous Delivery) kritik ise,

... mikroservis doru tercihtir.

---

[Geri - Ana README](../../README.md) | [İleri - 02-decomposition/README.md](../02-decomposition/README.md)
