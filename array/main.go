package main

import "fmt"

func main() {
	var angka [3]int = [3]int{10, 20, 30}
	fmt.Println("isi array angka:", angka)
	fmt.Println(angka[1])

	angka[1] = 200
	fmt.Println("isi array angka setelah diubah:", angka)
}
