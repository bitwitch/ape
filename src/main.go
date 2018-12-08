package main

import (
	"fmt"
	"os"
	"os/user"
	"repl"
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
