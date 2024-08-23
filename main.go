package main

import (
	"fmt"

	"github.com/hashicorp/go-retryablehttp"
)

var password = "mypassword"

func main() {
	resp, err := retryablehttp.Get("/path")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	fmt.Println(password)
}
