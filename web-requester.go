package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}

	response, err := client.Get("http://www.google.com/")
	fmt.Println(response, err)
}
