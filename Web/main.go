package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://lco.dev"
	fmt.Println("Website")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response after read", string(databytes))
}
