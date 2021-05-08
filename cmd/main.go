package main

import (
	"fmt"
	"my-go/pkg/app"
	"os"
)

func main() {
	var r rune
	fmt.Println("Welcome to My-Go")
	for {
		r = app.Listen()
		if r == 'q' {
			//			break
			os.Exit(0)
		}
		switch r {
		case 1:
			app.vari()
		}

	}
}
