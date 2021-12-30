package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/domterion/yapl/lexer"
	"github.com/domterion/yapl/token"
)

// REPL is Read Eval Print Loop

const PROMPT = ":: "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	// Always looping to check for new input to lex
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
