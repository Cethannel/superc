package parser

import (
	"superc/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
	func main() {
		let x = 5;
		let y = 10;
		let foobar = 838383;
	}
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
	}
}
