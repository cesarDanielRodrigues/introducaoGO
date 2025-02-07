package main

import (
	"fmt"
	"strings"
)

func main() {

	var teste1 string = "teste 1 "
	var teste2 string = "teste 2"

	var juntar string = teste1+teste2
	fmt.Println(teste1)
	fmt.Println(strings.ToUpper(juntar))
	fmt.Println(strings.Contains(juntar,"maico"))

}
