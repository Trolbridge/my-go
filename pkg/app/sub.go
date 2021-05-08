package app

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// fd 0 is stdin
func Listen() rune {
	// fd 0 is stdin
	state, err := terminal.MakeRaw(0)
	if err != nil {
		log.Fatalln("Error: setting stdin to raw:", err)
	}
	defer func() {
		if err := terminal.Restore(0, state); err != nil {
			log.Println("Error: failed to restore terminal:", err)
		}
	}()

	in := bufio.NewReader(os.Stdin)
	//	for {
	r, _, err := in.ReadRune()
	if err != nil {
		log.Println("Error stdin:", err)
		//			break
		os.Exit(1)
	}
	fmt.Printf("You entered: %q\r\n", r)
	//	}
	return r

}
