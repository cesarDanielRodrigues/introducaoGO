package main

import "fmt"

func main() {

	players := map[string]int{
		"cesar": 26,
	}
	// if value, ok := players["cesar"]; ok {
	// 	fmt.Println("verdadeiro", value, ok)
	// } else {
	// 	fmt.Println("falso", value, !ok)
	// }
    value, ok:=players["cesar"]
    fmt.Print(value, ok)
}
