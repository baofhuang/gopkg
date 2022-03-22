package main

import (
	"fmt"
	
	"github.com/baofuhuang/gopkg/http"
)

func main() {
	err := http.StartServer(":11010")
	if err != nil {
		fmt.Println(err)
	}
}
