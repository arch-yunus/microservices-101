# 01-Giriş: Mikroservis Mimarisine Giriş

Bu modül, monolitik sistemlerden mikroservis mimarisine geçişin nedenlerini, avantajlarını ve karşılaşılabilecek temel zorlukları teknik bir perspektifle ele alır.

---

## 🏗️ Mimari Dönüşüm

Mikroservis mimarisi, karmaşık bir yazılım sistemini, her biri belirli bir iş mantığına (Business Logic) sahip, bağımsız olarak geliştirilebilir, dağıtılabilir ve ölçeklenebilir küçük servislerin birleşimi olarak tasarlama yaklaşımıdır.

### Monolitik Mimari (The Monolith)
Geleneksel "her şey tek bir yerde" yaklaşımıdır.
- **Avantajlar**: Basit dağıtım süreçleri, kolay uçtan uca (E2E) test imkanı.
- **Dezavantajlar**: Tek bir hata tüm sistemi etkileyebilir (Single Point of Failure), teknoloji bağımlılığı yüksektir ve dağıtım süreleri sistem büyüdükçe uzar.

### Mikroservis Mimarisi
Modern, dağıtık sistem yaklaşımıdır.
- **Avantajlar**: Teknolojik serbestlik (Poliglot), bağımsız ölçeklenebilirlik, hata izolasyonu ve hızlı sürekli teslimat (CD).
- **Dezavantajlar**: Dağıtık sistem karmaşıklığı, veri tutarlılığı (Data Consistency) yönetimi ve operasyonel yükün artması.

---

## 🎯 Ne Zaman Mikroservis Tercih Edilmeli?

Mikroservis mimarisine geçiş kararı, projenin ölçeği ve ekibin yetkinliklerine bağlı olarak verilmelidir:
1.  **Ekip Ölçeği**: Geniş geliştirici ekiplerinin (100+) bağımsız çalışması gerektiğinde.
2.  **Ölçeklendirme Gereksinimleri**: Farklı sistem bileşenlerinin farklı kaynak (CPU/Memory) ihtiyaçları olduğunda.
3.  **Hızlı Teslimat Deneyimi**: Continuous Delivery süreçlerinin kritik olduğu, yüksek rekabetçi pazarlarda.

---

[Geri - Ana README](../../README.md) | [İleri - 02-Ayrıştırma Stratejileri](../02-decomposition/README.md)
