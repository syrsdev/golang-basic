package main

import "fmt"

func main() {
	s := make([]int, 3, 5)

	fmt.Println(s)
	fmt.Println("length:", len(s))
	fmt.Println("capacity:", cap(s))

	s = append(s, 10, 20)
	fmt.Println("setelah di append:", s)

	slice2 := make([]int, 4)
	slice3 := copy(slice2, s)

	fmt.Println("copied:", slice2)
	fmt.Println("jumlah elemen yang tersalin:", slice3)

	angka := []int{100, 200, 300, 400, 500}
	slice4 := angka[1:4]
	fmt.Println("slice4:", slice4)
}
