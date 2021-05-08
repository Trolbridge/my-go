package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// fd 0 is stdin
func Listen() rune {
	//	// fd 0 is stdin
	//	state, err := terminal.MakeRaw(0)
	//	if err != nil {
	//		log.Fatalln("setting stdin to raw:", err)
	//	}
	//	defer func() {
	//		if err := terminal.Restore(0, state); err != nil {
	//			log.Println("warning, failed to restore terminal:", err)
	//		}
	//	}()

	in := bufio.NewReader(os.Stdin)
	//	for {
	r, _, err := in.ReadRune()
	if err != nil {
		log.Println("stdin:", err)
		//			break
		os.Exit(1)
	}
	fmt.Printf("read rune %q\r\n", r)
	if r == 'q' {
		//			break
	}
	//	}
	return r

}

func Quit() {
	//	fmt.Println("setting stdin to raw")
	//	if err != nil {
	//		log.Fatalln("setting stdin to raw:", err)
	//	}
	//	if err := terminal.Restore(0, state); err != nil {
	//		log.Println("warning, failed to restore terminal:", err)
	//	}
	os.Exit(1)
}
