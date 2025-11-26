package main

import (
	"fmt"
	"strconv"
)

func main() {
	var text string = "1000B"
	number, err := strconv.Atoi(text)

	if err != nil {
		fmt.Println("Error converting string to number:", err.Error())
	} else {
		fmt.Println("Converted number:", number)
	}
}
