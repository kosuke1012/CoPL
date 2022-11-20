package main

import (
	"fmt"
	"os"

	copl "github.com/kosuke1012/CoPL/CoPL"
)

func main() {
	f, err := os.Open("test.html")
	if err != nil {
		panic(err)
	}
	res, err := copl.ParseHTMLResult(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}
