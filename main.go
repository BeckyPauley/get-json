package main

import (
	"encoding/json"
	"fmt"
)

type people struct {
	Number int `json: "number"` // Number is an annotation - capal leter marks it as exportable or public https://www.geeksforgeeks.org/pointer-to-a-struct-in-golang/
}

func main() {
	text := `{"people": [{"craft": "ISS", "name": "Sergey Rizhikov"}, {"craft": "ISS",
	"name": "Andrey Borisenko"}, {"craft": "ISS", "name": "Shane Kimbrough"}, {"craft":
	"ISS", "name": "Oleg Novitskiy"}, {"craft": "ISS", "name": "Thomas Pesquet"}, {"craft":
	"ISS", "name": "Peggy Whitson"}], "message": "success", "number": 6}`
	textBytes := []byte(text) // text variable --> type byte (stored in textBytes variable) as json.Unmarshal() takes type byte.

	p := people{}
	err := json.Unmarshal(textBytes, &p) // &p passes the value of the empty struct to the method, without the & it will use a different copy of the empty struct
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p.Number)
}
