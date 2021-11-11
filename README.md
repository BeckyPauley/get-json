# JSON

## Following this awesome book: Everyday Go by Alex Ellis
https://openfaas.gumroad.com/l/everyday-golang

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

User-Agent Header - allows a site to understand what sort of ttraffic it's receiving. 

robots.txt - tells a search engine files a crawler can/can't access from site. https://developers.google.com/search/docs/advanced/robots/intro  

### Revisit later:

I'd like to understand the User-Agent Header in more depth https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/User-Agent 

