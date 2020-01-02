package main

import (
	"net"
	"net/url"
	"net/http"
	"golang.org/x/net/html"
	"io"
	"strings"
	"github.com/likexian/whois-go"		
)

// 
func isURL(urlString string) bool {
	_, err := url.ParseRequestURI(urlString)

	if err != nil {			
		panic(err)					
	} else {	
		return true
	}		
}

//
func hostName(urlString string) string {
	u, err := url.Parse(urlString) 

	if err != nil {
		panic(err)
	} else {
		hoststring := u.Hostname()
		return hoststring
	}	
}

//
func getIP(urlString string) []net.IP {
	ips, err := net.LookupIP(urlString)

	if err != nil {			
		panic(err)					
	} else {					
		return ips
	}	
}

// get the information s from the query whois
func getInfoWhoIs(s string, ips []net.IP) string {
	for _, value := range ips {
		result, err := whois.Whois(value.String())
		
		if err != nil {
			//fmt.Println(result)
			panic(err)
		}

		splitResult := strings.Split(result, "\n")
				
		// search in splitresult orgname
		for _, val := range splitResult {
			if strings.Contains(val, s) {
				//fmt.Println(val)
				info := strings.Trim(val, s)
				//fmt.Println(info)
				return info
			}
		}		
	}
	return ""	
}

/* func getHTML(urlString string) string {

} */

// 
func getTitle(urlString string, s string) string {
	resp, err := http.Get(urlString)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//create a new tokenizer over the response body
	tokenizer := html.NewTokenizer(resp.Body)
	for {
		tokenType := tokenizer.Next()
		
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			//fmt.Println(err)
			if err == io.EOF {
				//end of the file, break out of the loop
				break
			}			
		}		
		
		if tokenType == html.StartTagToken {
			//get the token
			token := tokenizer.Token()			
			//if the name of the element is s
			if s == token.Data {
				//the next token should be the page title
				tokenType = tokenizer.Next()				
				//report the page title and break out of the loop
				result := tokenizer.Token().Data			
				return result				
			}
		}	
	}
	return ""
}
