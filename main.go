package main

import (
	"net/url"
 	"fmt"
)

//url.ParseRequestURI 
// https://stackoverflow.com/questions/31480710/validate-url-with-standard-package-in-go
func main() {
	fmt.Println(isURL("https://google.com"))
	fmt.Println(isURL("google"))
}

func isURL(urlString string) bool {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		//panic(err)
		fmt.Println("Error: page not found")
		return false
	} else {
		return true
	}		
}
