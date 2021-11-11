package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type people struct {
	Number int `json: "number"` // Number is an annotation - capal leter marks it as exportable or public https://www.geeksforgeeks.org/pointer-to-a-struct-in-golang/
}

func main() {

	url := "http://api.open-notify.org/astros.json"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // timeout fter 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("UserAgent", "spacecount-tutorial")

	res, err := spaceClient.Do(req)
	fmt.Printf("HTTP status: %s\n", res.Status)

	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	p := people{}
	err = json.Unmarshal(body, &p) // &p passes the value of the empty struct to the method, without the & it will use a different copy of the empty struct
	if err != nil {
		log.Fatal(err)
		log.Fatalf("unable to parse value: %q, error: %s",
			string(body), err.Error())
	}
	fmt.Println(p.Number)
}
