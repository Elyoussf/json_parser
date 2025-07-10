package main

import (
	"fmt"
	"json_parser/lexer"
)

func main() {
	json := `{
  "name": "John Doe",
  "age": 30,
  "city": "New York",
  "isStudent": false,
  "address": {
    "street": "123 Main St",
    "zipCode": "10001"
  },
  "courses": [
    "Math",
    "Physics",
    "History"
  ]
}`
	val, _ := lexer.Lex(json)

	fmt.Println(val)
}
