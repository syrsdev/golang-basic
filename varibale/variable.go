package main

import "fmt"

func main() {
	// deklarasi variable klasik (sebutkan tipe datanya)
	var nama string = "syrsdev"
	var umur int = 100

	// deklarasi variable otomatis (type interface)
	city := "jakarta"
	year := 2025

	// konstanta
	const pi = 3.14

	fmt.Println("nama saya " + nama)
	fmt.Println("umur saya ", umur)
	fmt.Println("saya tinggal di " + city)
	fmt.Println("tahun sekarang ", year)
	fmt.Println("nilai konstanta pi ", pi)
}
