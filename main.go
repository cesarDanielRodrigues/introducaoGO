package main

import "fmt"

func main() {
	var pessoa = map[string]int{}
    pessoa["cesar"] = 26
    pessoa["daniel"] = 13
    
    if idade, ok:=pessoa["marcelo"];ok{
        fmt.Println("A pessoa tem",idade,"anos")
    } else{
        fmt.Println("Não existe essa pessoa")
    }

}
