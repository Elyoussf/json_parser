package lexer

import (
	"fmt"
	"slices"
)

var json_tokens = []string{"{", "}", "[", "]", ":", ","}

var whitespace = " "

var quote = "\""

var numbers = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ".", "e", "-"} // this includes floating nu,ber as well as negative ones

var null = "null"

var true = "true"

var false = "false"

func Lex(input string) ([]string, error) {
	Tokens := []string{}
	l := len(input)
	idx := -1

	for idx < l {

		idx += 1
		if idx >= l {
			break
		}
		c := input[idx]
		if string(c) == whitespace {
			continue
		}

		if slices.Contains(json_tokens, string(c)) {
			Tokens = append(Tokens, string(c)) // Why not c but string(c)
			continue
		}

		if slices.Contains(numbers, string(c)) {

			curr := idx
			for curr < l && slices.Contains(numbers, string(input[curr])) {
				curr += 1
			}

			if curr >= l {
				// Error here we cannot have a number that goes beyond the strung without at least closing curly brackets
				return []string{}, fmt.Errorf("not a valid json format")
			}

			Tokens = append(Tokens, input[idx:curr])
			idx = curr - 1 //  Because at every beginning there is in increment

			continue
		}

		if string(c) == quote {
			str := "\"" + string(c) // Should be included
			idx += 1
			c = input[idx]
			for idx < l && string(c) != quote {
				str += string(c)
				idx += 1
				c = input[idx]
			}
			if string(c) == quote {
				str = "\"" + string(c)
			}
			Tokens = append(Tokens, str)
			idx -= 1
			continue
		}
		str := ""
		for isAlphaNumeric(c) {
			str += string(c)
			idx += 1
			c = input[idx]
		}
		Tokens = append(Tokens, str)
		idx -= 1
	}
	return Tokens, nil

}

func isAlphaNumeric(c byte) bool {
	return (c <= 'z' && c >= 'a') || (c <= '9' && c >= '0')
}
