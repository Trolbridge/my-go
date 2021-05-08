package main

import (
	"my-go/pkg/app"
	"os"
)

func main() {
	//	fmt.Println("vim-go")
	//	app.PrintLine()
	var r rune
	for {
		r = app.Listen()
		if r == 'q' {
			//			break
			os.Exit(1)
		}

	}
}
