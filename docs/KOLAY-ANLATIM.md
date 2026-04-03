# ?? Mikroservis m? Nedir Bu? (En Ba?tan, En Basit Haliyle)

Selam kanka! "Mikroservis" dedikleri ?ey asl?nda b?y?k bir i?i, par?alara b?l?p her birini ayr? birine yapt?rmakt?r. Gel, bunu g?nl?k bir ?rnekle hi? kafa kar??t?rmadan, ad?m ad?m ??zelim.

---

## 1. Monolith Nedir? (Kocaman Bir Tencere ?orba) ??

Eskiden yaz?l?mlar b?yle yap?l?rd?. D???n ki dev bir kazan?n var. Bu kazan?n i?ine et, sebze, su, tuz, baharat, hatta bula??k makinesi deterjan? bile at?yorsun (??nk? tek bir yer her ?eyi yap?yor). Her ?eyi ayn? anda kaynat?yorsun.

- **Neden S?k?nt?l??** 
    - **Hata Da??l?m?:** ??er i?ineyanl??l?kla fazla tuz atarsan, o koca kazan? ??pe d?kmen gerekir. "Dur sadece tuzunu ay?ray?m" diyemezsin, ??nk? her ?ey i? i?e girmi?tir.
    - **Hantal Yap?:** Kazan o kadar b?y?kt?r ki, ate?in alt?n? a?t???nda kaynamas? saatler s?rer (Uygulaman?n a??l?m s?resi). 
    - **Teknoloji Esareti:** ??ker o kazanda odun ate?i kullan?yorsan, sadece bir k?sm? i?in "Hadi burada mikrodalga kullanal?m" diyemezsin. Her yer ayn? eski y?ntemle ya?amak zorundad?r.

---

## 2. Mikroservis Nedir? (A??k B?fe / Mini Kaseler Mant???) ??

Mikroservis ise o b?y?k kazan? k?r?p, i?indeki her yeme?i k???k kaselere b?lmek demektir. Bir kasede salata, birinde k?fte, birinde tatl?, birinde i?ecek var.

- **G?zelli?i Ne?** 
    - **?zg?rl?k:** Salatac? usta isterse marulu elle do?rar, k?fteci usta k?fteyi en modern f?r?nda pi?irir. Kimse birbirine kar??maz. Tatl?c? usta "Ben art?k ba?ka bir felsefeyle tatl? yapaca??m" derse sadece onun kasesi de?i?ir. (**Independent Deployment**)
    - **Patlama Korumas?:** Salatadan sinek mi ??kt?? Salatacy? mutfaktan d??ar? atar, kaseyi k?rars?n. Ama insanlar k?fteyi ve tatl?y? yemeye, i?eceklerini i?meye devam eder. D?kkanda i?ler durmaz. (**Resilience / Dayan?kl?l?k**)
    - **Ak?ll? B?y?me:** E?er m??teriler sadece salata istiyorsa, b?t?n mutfa?? b?y?tmek yerine sadece salatac? ustan?n yan?na 3 tane daha yard?mc? verirsin. K?fteci usta oturmaya devam edebilir, bo?a masraf yapmazs?n. (**Scalability / ?l?ekleme**)

---

## 3. Peki Bunlar Nas?l Konu?uyor? (Mutfaktaki Telsizler) ??

K?fteci ile Tatl?c? farkl? tezgahlarda, hatta belkide farkl? ?ehirlerdeler. ?nsanlar sipari? verince birbirlerinden haber almalar? laz?m.

1.  **gRPC / REST (Anl?k Telefon G?r??mesi):** 
    K?fteci usta telsize basar: "Usta, salatac?ya s?yle hemen bir marul g?ndersin, acil!" Yan?t gelene kadar telsizin ba??nda bekler. Marul gelirse i?ine devam eder, gelmezse telsizle kavga eder. Bu y?ntem h?zl?d?r ama iki taraf?n da o an telsiz ba??nda olmas? ?artt?r.
    
2.  **Message Broker (Mantar Pano / PTT Posta):** 
    K?fteci usta mantar panoya bir not asar: "Arkada?lar sipari? al?nd?, k?fte pi?iyor. ??i olan gelsin notu okusun." Salataci usta mutfa?a girince panoya bakar, notu g?r?r ve i?ini yap?p o da panoya yaz?r: "Salata haz?r, servis?i gelsin als?n." Kimse kimseyi beklemek zorunda de?ildir. Biri uykudayken bile i?ler panoda birikir. (**Asynchronous Communication**)

---

## 4. Birinci Alt?n Kural: Veritaban? S?rr? (Kilitli Dolaplar) ??

Mikroservis d?nyas?n?n en ?nemli kural? ?udur: **Herkesin kendi buzdolab? vard?r ve anahtar? sadece kendisindedir.**

- **Neden?** E?er salatac? usta marul bitince gidip k?fteci ustan?n dolab?n? kar??t?r?rsa, k?fteci usta dolab?n yerini de?i?tirdi?inde salatac? usta a? kal?r. 
- **??z?m:** Salataci marul mu istiyor? Dolaba kendisi dalmayacak. K?fteci ustadan (veya sistemden) "Bana marul ver" diye rica edecek. Herkes kendi malzemesinden (Veritaban?) sorumludur. Buna "Database per Service" diyoruz. Kar??t?r?rsak mikroservis de?il, "Ba??ml?l?k Karma?as?" olur.

---

## 5. ?zetle Neden Bu Kadar U?ra??yoruz? ??

1.  **H?z:** Yeni bir ?zellik eklemek i?in koca tencereyi ba?tan kaynatm?yoruz, sadece kasesini ekliyoruz.
2.  **G?ven:** Tek bir k?k hata t?m d?kkandaki insanlar? zehirlemesin diye.
3.  **Ekip ?al??mas?:** 100 ustay? tek bir tencerenin ba??na toplarsan birbirlerine ka??k f?rlat?rlar. Herkesi kendi tezgah?na (Servis) koyuyoruz ki i?ler t?k?r t?k?r y?r?s?n.

---

### ?? Unutma Kanka:
Mikroservis, "B?y?k lokma ye ama b?y?k i?e tek ba??na kalk??ma" demektir. Her i?i uzman?na veriyoruz, aradaki telsiz hatt?n? da iyi kuruyoruz, gerisi keyif ?ay?! ?☕

---
[Geri - Ana README](../README.md)
