package main

import (
	"fmt"
	"json_parser/lexer"
	"json_parser/parser"
)

func main() {
	json := `{
  "name": "Alice Johnson",
  "age": 34,
  "isEmployed": true,
  "hasDriverLicense": false,
  "salary": 75340.75,
  "department": "Engineering",
  "email": "alice.johnson@example.com",
  "phone": null,
  "address": "123 Main St, Springfield",
  "zipCode": "90210",
  "country": "USA",
  "language": "English",
  "married": true,
  "children": null,
  "taxRate": 0.27,
  "bonus": 10000,
  "isRemote": false,
  "accessLevel": "admin",
  "startDate": "2022-01-15",
  "endDate": null,
  "bio": "Loves building backend systems.",
  "active": true,
  "lastLogin": "2025-07-09T13:45:00Z",
  "projectsCount": 15,
  "performanceScore": 4.8,
  "probation": false,
  "feedback": null,
  "gender": "female",
  "timezone": "America/New_York",
  "contract": "permanent",
  "vacationDays": 14,
  "sickDays": 2,
  "overtimeHours": 12.5,
  "bonusEligible": true,
  "equipmentProvided": false,
  "linkedIn": "alice-johnson",
  "slackID": "U1234567890",
  "notes": null
}
`
	val, _ := lexer.Lex(json)

	res := parser.ParseSimpleObject(val)

	fmt.Println(res["taxRate"].(float64) * 100)

}
