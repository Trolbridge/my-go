package main

import (
	"fmt"
	"my-go/pkg/app"
)

func main() {
	//	fmt.Println("vim-go")
	//	app.PrintLine()
	var r rune
	r = app.Listen()

	fmt.Printf("read rune %q\r\n", r)
	if r == 'q' {
		app.Quit()
	}
}
