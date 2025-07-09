package main

import (
	"fmt"
	"json_parser/lexer"
)

func main() {
	Tkns, _ := lexer.Lex("{hamza : 2,laila : 890}")

	fmt.Println(Tkns)

}
