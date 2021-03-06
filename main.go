package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type person struct {
	Name string `json:"name"`
}

type people struct {
	Number int      `json:"number"`
	Person []person `json:"people"`
}

func main() {
	apiURL := "http://api.open-notify.org/astros.json"
	people, err := getAstros(apiURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d people found in space.\n", people.Number)

	for _, p := range people.Person {
		fmt.Printf("Let's wave to: %s\n", p.Name)
	}
}

func getAstros(apiUrl string) (people, error) {
	p := people{}

	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("UserAgent", "spacecount-tutorial")

	res, err := http.DefaultClient.Do(req)
	fmt.Printf("HTTP status: %s\n", res.Status)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return p, err
	}

	err = json.Unmarshal(body, &p) // &p passes the value of the empty struct to the method, without the & it will use a different copy of the empty struct
	if err != nil {
		log.Fatal(err)
		log.Fatalf("unable to parse value: %q, error: %s",
			string(body), err.Error())
	}
	return p, nil
}
