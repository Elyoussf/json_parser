package main

import (
	"fmt"
	"json_parser/lexer"
	"json_parser/parser"
)

func main() {
	json := `{
  "name" : "Hamza",
  "age" : 21 ,
  "infos" : [123,null,"hamza",false,"   ",{"age":45}]
  }`
	val, _ := lexer.Lex(json)

	res := parser.ParseSimpleObject(val)

	fmt.Println(res)

}
