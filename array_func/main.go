package main

import "fmt"

func main() {
	angka := [5]int{10, 20, 30, 40, 50}
	fmt.Println("panjang array: ", len(angka))
	fmt.Println("index ke-1: ", angka[1])

	angka[1] = 200
	fmt.Println("index ke-1 setelah diubah: ", angka[1])

	for index, value := range angka {
		fmt.Println("Index ke-", index, " = ", value)
	}

	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{4, 5, 6}

	fmt.Println("arr1 == arr2 ? ", arr1 == arr2)
	fmt.Println("arr1 != arr2 ? ", arr1 != arr2)
}
