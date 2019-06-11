package main

import (
	"fmt"
	"github.com/bitwitch/ape/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s, this is the Ape programming language.\n", user.Username)
	fmt.Printf("Feel free to go apeshit!\n")
	repl.Start(os.Stdin, os.Stdout)
}
