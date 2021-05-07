package main

import (
	"fmt"
	"log"
	"my-go/pkg/app"
)

func main() {
	//	fmt.Println("vim-go")
	//	app.PrintLine()
	var r rune
	r = app.Listen()

	for {
		r, _, err := in.ReadRune()
		if err != nil {
			log.Println("stdin:", err)
			break
		}
		fmt.Printf("read rune %q\r\n", r)
		if r == 'q' {
			break
		}
	}

}
