package main

import (
	"fmt"
	"my-go/pkg/app"
	"os"
)

func main() {
	var r rune
	//	_ = r
	fmt.Println("Welcome to My-Go")
	fmt.Println("1. Hello World")
	fmt.Println("2. Variables")
	fmt.Println("3. Println")
	fmt.Println("#. Variables")
	fmt.Println("1. Variables")
	fmt.Println("1. Variables")
	fmt.Println("1. Variables")
	fmt.Println("1. Variables")
	fmt.Println("1. Variables")
	fmt.Println("1. Variables")
	fmt.Println("1. Variables")
	for {
		r = app.Listen()
		/*		if r == 'q' {
				os.Exit(0)
			} */
		switch r {
		case '1':
			app.HelloWorld()
		case '2':
			app.Vari()
		case 'q':
			os.Exit(0)
		}

	}
}
