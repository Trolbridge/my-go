package app

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func PrintLine() {
	println("print line test message")
}

// fd 0 is stdin
func Listen() rune {
	state, err := terminal.MakeRaw(0)
	fmt.Println("setting stdin to raw")
	if err != nil {
		log.Fatalln("setting stdin to raw:", err)
	}

	in := bufio.NewReader(os.Stdin)
	r, _, err := in.ReadRune()
	if err != nil {
		log.Println("stdin:", err)
		if err := terminal.Restore(0, state); err != nil {
			log.Println("warning, failed to restore terminal:", err)
		}
		os.Exit(1)
	}
	return r

}
func Quit() {
	state, err := terminal.MakeRaw(0)
	fmt.Println("setting stdin to raw")
	if err != nil {
		log.Fatalln("setting stdin to raw:", err)
	}
	if err := terminal.Restore(0, state); err != nil {
		log.Println("warning, failed to restore terminal:", err)
	}
	os.Exit(1)
}
