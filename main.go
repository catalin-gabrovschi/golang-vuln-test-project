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

	x := make([]int, 0)
	for i := 0; i < 10; i++ {
		fmt.Println(x[i])
	}
}
