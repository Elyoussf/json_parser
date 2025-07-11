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
				if current == "[" {
					open := 1
					i := idx + 1
					for Tokens[i] != "]" && open != 0 {
						if Tokens[i] == "[" {
							open += 1
						}
						if Tokens[i] == "]" {
							open -= 1
						}
						if open == 0 {
							break
						}
						i += 1
					}
					subTokens := Tokens[idx : i+1]
					val = ParseSimpleArray(subTokens)
					res[WaitingKey] = val

					idx = i + 1
					WaitingKey = ""
				}
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

func ParseSimpleArray(Tokens []string) []any {

	var arr []any

	idx := 1
	current := Tokens[idx]

	for idx < len(Tokens) && current != "]" {
		if current == "{" {
			open := 1
			i := idx + 1
			for Tokens[i] != "}" && open != 0 {
				if Tokens[i] == "{" {
					open += 1
				}
				if Tokens[i] == "}" {
					open -= 1
				}
				if open == 0 {
					break
				}
				i += 1
			}
			subTokens := Tokens[idx : i+1]
			res := ParseSimpleObject(subTokens)
			arr = append(arr, res)
			idx = i + 1
			current = Tokens[idx]
			continue
		}
		if current == "," {
			idx += 1
			current = Tokens[idx]
			continue
		}
		if utils.IsNull(current) {
			arr = append(arr, nil)
			idx += 1
			current = Tokens[idx]
			continue
		}
		if utils.IsBool(current) {
			v, _ := strconv.ParseBool(current)
			arr = append(arr, v)
			idx += 1
			current = Tokens[idx]
			continue
		}
		if utils.IsNumber(current) {
			var v any
			if utils.IsFloat(current) {
				v, _ = strconv.ParseFloat(current, 64)
			} else {
				v, _ = strconv.ParseInt(current, 10, 16) // for larger numbers
			}
			arr = append(arr, v)
			idx += 1
			current = Tokens[idx]
			continue

		}

		if utils.IsAllAlphaNumeric(current) || len(current) == 0 || (current[0] == current[len(current)-1] && string(current[0]) == `"`) {
			arr = append(arr, current)
			idx += 1
			current = Tokens[idx]
			continue
		}
	}
	return arr
}
