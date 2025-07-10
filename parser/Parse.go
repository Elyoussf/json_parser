package parser

func Syntacticanalysis(Tokens []string) error {
	/*
		Find { Expecting a string without either whitespaces the next string should be : then this could be a json object
		OR Array , float , integer , string , null and boolean

		Then expecting closing curly bracket if no an error should be thrown with wrong expectation

	*/
	idx := 0
	l := len(Tokens)
	expected := "{"
	for idx < l {
		token := Tokens[idx]
		if token != expected {

		}
	}
}

//ParseString  --> string
//ParseBoolean --> bool
//ParseObject --> map[string]interface{}
//ParseNull --> nil
//ParseArray --> []interface{}
