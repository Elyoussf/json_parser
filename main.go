package main

import (
	"fmt"
	"json_parser/lexer"
	"json_parser/parser"
)

func main() {
	data := `{
  "name": "Hamza",
  "contacts": [
    {
      "type": "email",
      "value": "hamza@example.com"
    },
    {
      "type": "phone",
      "value": "+212600000000"
    }
  ],
  "settings": {
    "notifications": {
      "email": null,
      "sms": false
    },
    "theme": "dark"
  },
  "active": true
}

`
	val, _ := lexer.Lex(data)

	res := parser.ParseSimpleObject(val)

	fmt.Println(res)

}
