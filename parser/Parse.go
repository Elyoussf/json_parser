package parser

import (
	"fmt"
	"json_parser/utils"
	"strconv"
)

func ParseSimpleObject(Tokens []string) map[string]any {
	var WaitingKey = ""
	var idx = 1
	res := make(map[string]any)
	var err error
	current := Tokens[0]

	for idx < len(Tokens)-1 && current != "}" {
		current = Tokens[idx]

		if len(WaitingKey) == 0 && current != "," && current != ":" {

			WaitingKey = current[1 : len(current)-1]

			res[WaitingKey] = nil // reserving

			idx += 1
		} else {
			if current == ":" || current == "," {
				idx += 1
				continue
			} else {
				var val any
				if utils.IsNull(current) {
					val = nil
				}
				if utils.IsBool(current) {
					val = (current == "true")
				}
				if utils.IsNumber(current) && utils.IsFloat(current) {

					val, err = strconv.ParseFloat(current, 64)
					if err != nil {
						fmt.Println("Error while converting to float")
						return nil
					}
				} else if utils.IsNumber(current) {
					val, _ = strconv.ParseInt(current, 10, 16) // Not sure 16
				}
				if val == nil && current != "null" {

					val = current
				}
				res[WaitingKey] = val

				idx += 1
				WaitingKey = ""
			}
		}
	}
	return res
}

// func ParseSimpleArray(Tokens []string) []any{
// 	res := make([]any)
// }
