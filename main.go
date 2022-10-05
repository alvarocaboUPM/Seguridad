package main

import (
	"fmt"
	"os"
	_ "strings"
	"workspace/algoritmos"
	_ "workspace/libs"
)

func main(){
	//text_raw := os.Args[1:]
	key := "PROFUNDIS"
	fmt.Println(algoritmos.Encrypt(os.Args[1], key))
}