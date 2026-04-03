# ?? Mikroservis m? Nedir Bu? (En Ba?tan, En Basit Haliyle)

Selam kanka! "Mikroservis" dedikleri ?ey asl?nda b?y?k bir i?i, par?alara b?l?p her birini ayr? birine yapt?rmakt?r. Gel, bunu g?nl?k bir ?rnekle hi? kafa kar??t?rmadan ??zelim.

---

## 1. Monolith Nedir? (Kocaman Bir Tencere ?orba) ??

Eskiden yaz?l?mlar b?yle yap?l?rd?. Bir tane devasa tenceren var, her ?eyi (et, sebze, su, tuz, baharat) i?ine at?p kaynat?yorsun. 

- **S?k?nt? Ne?** ??er tuzu fazla ka??rd?ysan, tencerenin tamam? ??p olur. ?orbay? d?zeltmek i?in her ?eyi ba?tan yapman gerekir. Ayrica o koca tencereyi masaya ta??mak i?in 10 ki?i laz?md?r (Bilgisayar?n t?m g?c?n? harcar).

## 2. Mikroservis Nedir? (A??k B?fe Mant???) ??

Mikroservis ise o yeme?i k???k kaselere b?lmektir. Bir kasede salata var, bir kasede k?fte, bir kasede tatl?.

- **G?zelli?i Ne?** 
    - Salatada sinek mi ??kt?? Sadece salatay? atars?n, insanlar k?fteyi ve tatl?y? yemeye devam eder. (**Resilience / Dayankllk**)
    - Herkes salataya sald?r?yorsa, sadece salata kasesini b?y?t?rs?n. Tatl? kasesi k???k kalabilir. (**Scalability / lekleme**)
    - Salatay? ba?ka bir usta, k?fteyi ba?ka bir usta yapabilir. Birbirlerine hi? bakmalar?na bile gerek yok. (**Independent Teams / Bamsz Ekipler**)

---

## 3. Peki Bunlar Nas?l Konu?uyor? (Telsiz Mant???) ??

K?fteci ile Tatl?c? farkl? yerlerde. ?nsanlar sipari? verince birbirlerine haber vermeleri laz?m.

1.  **gRPC / REST:** Telsizle do?rudan s?ylemek gibi. "K?fte bitti usta, salata yolla!" Yan?t beklersin.
2.  **Message Broker (RabbitMQ):** Bir mantar panoya not asmak gibi. "Sipari? al?nd?, panoya yazd?m, isteyen gidip okusun."

---

## 4. Birinci Alt?n Kural: Veritaban? S?rr? ??

En ?nemli kural ?u: **Salataci, K?ftecinin dolab?n? a?amaz!** 
Herkesin kendi dolab? (Veritaban?) var. E?er salatac? k?ftecinin dolab?ndan marul almaya kalkarsa kavga ??kar, her ?ey kar???r. Buna "Database per Service" diyoruz.

---

## 5. ?zetle Neden Yap?yoruz? ??

1.  **Bir yer patlarsa her yer patlamas?n diye.**
2.  **?ok y?k binince sadece laz?m olan yeri b?y?telim diye.**
3.  **Yeni bir teknoloji gelince (rn: Go dili) sadece bir k?sm?nda kullanabilelim diye.**

---

### ?? Unutma Kanka:
Mikroservis yapmak, 1 tencere yemek yapmak yerine 10 tane k???k kasede yemek yapmakt?r. Bula??klar? (Zorlu?u) biraz daha fazlad?r ama misafir (Kullan?c?) ?oksa en iyisi budur!

---
[Geri - Ana README](../README.md)
