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

	defer func() {
		if err := terminal.Restore(0, state); err != nil {
			log.Println("warning, failed to restore terminal:", err)
		}
	}()

	return in := bufio.NewReader(os.Stdin)
}
