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
func Listen() {
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

	in := bufio.NewReader(os.Stdin)
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
