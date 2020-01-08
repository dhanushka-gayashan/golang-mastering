package main

import "fmt"

func main() {

	const text = `Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede, gözlerden uzak, küçük ve sarı bir güneş vardır.
Bu güneşin yörüngesinde, kabaca yüz kırksekiz milyon kilometre uzağında, tamamıyla önemsiz ve mavi-yeşil renkli, küçük bir gezegen döner.
Gezegenin maymun soyundan gelen canlıları öyle ilkeldir ki dijital kol saatinin hâlâ çok etkileyici bir buluş olduğunu düşünürler.`

	// Some words are not correct, because some characters need more than 1 bytes. 1-4 bytes to represent the character as rune
	for i := 0; i < len(text); i++ {
		fmt.Printf("%c", text[i])
	}

	fmt.Println()
}
