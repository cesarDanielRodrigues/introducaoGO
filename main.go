package main

import "fmt"

func main() {
	var gavetas []string
    gavetas = append(gavetas,"copos", "panos", "pratos")

    var len = len(gavetas)

    //slice[x:x-1] se quiser ir até o fim da fila usar o numero do length, ou passar vazio no segundo digito
    gavetas = gavetas[0:3]
    fmt.Println(gavetas, len)
}
