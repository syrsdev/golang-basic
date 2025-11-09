package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Syrs Gantenk"
	fmt.Println("text asli:", text)

	// ubah jadi huruf kecil
	fmt.Println("LowerCase:", strings.ToLower(text))
	// ubah jadi huruf besar
	fmt.Println("UpperCase:", strings.ToUpper(text))
	// mengecek kalimat berada di awal kalimat
	fmt.Println("Apakah kalimat awal Syrs?:", strings.HasPrefix(text, "Syrs"))
	// mengecek kalimat berada di akhir kalimat
	fmt.Println("Apakah kalimat awal Syrs?:", strings.Contains(text, "Gantenk"))

	// memisahkan string berdasarkan delimiter (jadi array)
	parts := strings.Split("apel, mangga, jeruk, pisang", ",")
	fmt.Println("Split:", parts)
	// mengganti bagian string
	replace := strings.ReplaceAll(text, "Gantenk", "Ganteng")
	fmt.Println("Replace:", replace)
}
