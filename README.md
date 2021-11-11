# JSON

Can parse json into a struct (deal with in a structured way).

Marshalling (object --> text or binary)/unmarshalling (text --> object) more info here: https://go.dev/blog/json

Declare struct e.g:
```
type people struct {
	Number int `json: "number"` 
} 
```
Number is an annotation - capital letter marks it as exportable or public. 

Using pointers - https://www.geeksforgeeks.org/pointer-to-a-struct-in-golang/

