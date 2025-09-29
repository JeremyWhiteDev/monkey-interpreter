package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMP = ">> "

// REPL stand for REad Eval Print Loop... 
// A javascript console in a web browser is a REPL:
// It allows input (Read), 
// that input is evaluated into code (Eval), 
// the output of the evaluated code is printed (Print),
// and then the program switches back to Read mode (Loop)
func Start (in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, PROMP)
		scanned := scanner.Scan()
		if (!scanned) {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}