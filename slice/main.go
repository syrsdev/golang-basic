package main

import "fmt"

func main() {
	arr := [6]int{10, 20, 30, 40, 50, 60}

	// mengambil seluruh elemen array
	s1 := arr[:]
	fmt.Println("slice 1:", s1)

	// mengambil elemen dari index ke-2 sampai akhir
	s2 := arr[2:]
	fmt.Println("slice 2:", s2)

	// mengambil elemen dari awal sampai index ke-4
	s3 := arr[:3]
	fmt.Println("slice 3:", s3)

	// mengambil elemen dari index ke-1 sampai index ke-4
	s4 := arr[1:5]
	fmt.Println("slice 4:", s4)
}
