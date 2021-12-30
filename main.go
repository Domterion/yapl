package main

import (
	"fmt"
	"os"

	"github.com/domterion/yapl/repl"
)

func main() {
	fmt.Printf("Welcome to YAPL REPL\n")
	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
