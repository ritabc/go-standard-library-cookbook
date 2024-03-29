package main

import (
	"fmt"
	"net/http"
)

func main() {
	header := http.Header{}

	// Using the header as slice
	header.Set("Auth-X", "abcdef1234")
	// Add adds to the slice
	header.Add("Auth-X", "defghijkl")
	fmt.Println(header)

	// retrieving slice of values in header
	resSlice := header["Auth-X"]
	fmt.Println(resSlice)

	// get the first value
	resFirst := header.Get("Auth-X")
	fmt.Println(resFirst)

	// replace all existing values with this one
	// contrast this with Add()
	header.Set("Auth-X", "newvalue")
	fmt.Println(header)

	// remove header
	header.Del("Auth-X")
	fmt.Println(header)
}
