package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "https://www.example.com:8000/user?username=roach"

	// Parse url
	result, _ := url.Parse(s)
	fmt.Println(result.Scheme) // https
	fmt.Println(result.Host)
	fmt.Println(result.Path)     // /user
	fmt.Println(result.Port())   // 8000
	fmt.Println(result.RawQuery) // username=roach

	// Extract the query components into a Values struct
	vals := result.Query()
	fmt.Println(vals["username"]) // roach

	// Create a URL from components
	newUrl := &url.URL{
		Scheme:   "https",
		Host:     "www.example.com",
		Path:     "/args",
		RawQuery: "x=1&y=2",
	}

	// Convert this structure to string
	s = newUrl.String()
	fmt.Println(s) // https://www.example.com/args?x=1&y=2
}
