package main

import (
	"fmt"
	"strconv"
)

func main() {
	// bool to string
	truth := true
	str := strconv.FormatBool(truth)
	fmt.Println("boolean ke string:", str)

	// string to bool
	val, _ := strconv.ParseBool("true")
	fmt.Println("string ke boolean:", val)
}
