package main

import (
	"fmt"
	"llcc/handler"
)

func main() {
	// binary addition
	result := handler.AddBinary("11", "1")
	fmt.Println("result", result)
}
