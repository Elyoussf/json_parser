package parser

import (
	"bytes"
	"fmt"
	"json_parser/utils"
	"strconv"
	"strings"
)

func Syntacticanalysis(Tokens []string)( interface{}, error) {

	/*
		Find { Expecting a string without either whitespaces the next string should be : then this could be a json object
		OR Array , float , integer , string , null and boolean
		Then expecting closing curly bracket if no an error should be thrown with wrong expectation
	*/
	l := len(Tokens)
	
	if Tokens[0] != "{" || Tokens[l-1] != "}"{
		return  nil , fmt.Errorf("Invalid syntax curly brackets not found")
	}
	res := make(map[string]interface{})
	i := 1
	key := true // Means am expcting key now 
	thekey := ""
	for i < l-1 {
		if !key {
			// this is a value
			val := Tokens[i]
			if string(val[0]) == `"`{
				if string(val[len(val)-1]) != `"`{
					return nil , fmt.Errorf("Invalid format Syntax found double quote without corresponding one")
				}
				res[thekey] = val
			}else{
				if utils.IsNull(val){
					res[thekey] = nil
				}else if utils.IsBool(val){
					res[thekey] = (val == "true")
				}else if utils.IsNumber(val){
					if utils.IsFloat(val){
						f,err := strconv.ParseFloat(val,64)
						if err != nil {
							return nil , fmt.Errorf("Error occured while parsing a float value in json")
						}
						res[thekey] = f
					}
				}else if {
					//ToDO for objects and arrays 
				}
			}
		}


		res[Tokens[i]] = nil //reservation LOL
		thekey = Tokens[i]
		if i >=l-3 {
			// {key:} or {key}
			return  nil , fmt.Errorf("Invalid Syntax Found a Key without value!")
		} 
		if Tokens[i+1] != ":"{
			return  nil , fmt.Errorf("Invalid Syntax Found a Key without value!")
		}
		key := false
		i+=2

	}
}

//ParseString  --> string
//ParseBoolean --> bool
//ParseObject --> map[string]interface{}
//ParseNull --> nil
//ParseArray --> []interface{}


