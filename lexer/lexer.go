package lexer

import (
	"fmt"
	"json_parser/utils"
	"slices"
)

var json_tokens = []string{"{", "}", "[", "]", ":", ","}

var whitespace = " "

var quote = "\""

var numbers = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ".", "e", "-", "+"} // this includes floating nu,ber as well as negative ones

func Lex(input string) ([]string, error) {
	Tokens := []string{}
	l := len(input)
	idx := 0

	for idx < l {

		if idx >= l {
			break
		}
		c := input[idx]

		if input[idx] == 10 || input[idx] == 9 {
			idx += 1
			continue
		}
		if string(c) == whitespace {
			idx += 1
			continue
		}

		if slices.Contains(json_tokens, string(c)) {
			Tokens = append(Tokens, string(c)) // Why not c but string(c)
			idx += 1
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
			idx = curr //  Because at every beginning there is in increment

			continue
		}

		if string(c) == quote {

			str := `"` // Should be included
			idx += 1
			c = input[idx]
			for idx < l && string(c) != quote {
				c = input[idx]
				str += string(input[idx])
				idx += 1
			}
			// if string(c) == quote {
			// 	str += `"`
			// }
			Tokens = append(Tokens, str)

			continue
		}
		str := ""
		for idx < l && utils.IsAlphaNumeric(c) {

			str += string(c)
			idx += 1
			c = input[idx]

		}
		Tokens = append(Tokens, str)

	}
	return Tokens, nil

}
