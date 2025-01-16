package main

import (
	"fmt"

	"github.com/hashicorp/go-retryablehttp"
)

func main() {
	resp, err := retryablehttp.Get("/path")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
