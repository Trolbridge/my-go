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

//func Listen() []byte {
//	// disable input buffering
//	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
//	// do not display entered characters on the screen
//	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

//	// restore the echoing state when exiting
//	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

//	var b []byte = make([]byte, 1)
//	return b

//}

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
