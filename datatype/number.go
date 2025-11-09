package main

import "fmt"

func main() {
	// signed integer
	var a int8 = -100
	var b int16 = 20000
	var c int32 = -3000000
	var d int64 = 4000000000
	var in int = 80

	// unsigned integer
	var e uint8 = 100
	var f uint16 = 50000
	var g uint32 = 6000000
	var h uint64 = 7000000000
	var uin uint = 800

	fmt.Println("Signed Integers:")
	fmt.Printf("int8: %v\n", a)
	fmt.Printf("int16: %v\n", b)
	fmt.Printf("int32: %v\n", c)
	fmt.Printf("int64: %v\n", d)
	fmt.Printf("int: %v\n", in)

	fmt.Println("\nUnsigned Integers:")
	fmt.Printf("uint8: %v\n", e)
	fmt.Printf("uint8: %v\n", f)
	fmt.Printf("uint8: %v\n", g)
	fmt.Printf("uint8: %v\n", h)
	fmt.Printf("uint: %v\n", uin)
}
