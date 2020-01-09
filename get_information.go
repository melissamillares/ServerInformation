package main

import (
	"fmt"	
	"net"
	"net/url"
	"net/http"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"strings"
	"sort"
	"github.com/likexian/whois-go"		
)

// verify if the string is a URL
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
	var hoststring string

	if err != nil {
		panic(err)
	} else {
		hoststring = u.Hostname()
	}	
	return hoststring
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
				info := strings.Trim(val, s)				
				return info
			}
		}		
	}
	return ""	
}

func getSSLGrade(host string, length int) []string {
	u := "https://api.ssllabs.com/api/v3/analyze?host=" + host
	resp, err := http.Get(u)
	sslgrades := make([]string, length)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}	
	splitResult := strings.Split(string(body), ",")		
	for _, val := range splitResult {		
		if strings.Contains(val, "grade") {
			ssl := strings.Trim(val, "\"grade\":")
			for i := 0; i < length; i++ {
				sslgrades[i] = ssl
			}							
			return sslgrades
		}
	}

	return sslgrades
}

//
func getLowerGrade(ssl []string) string {
	sort.Strings(ssl)
	
	if len(ssl) == 1 {
	    return ssl[0]
	}
	last_index := len(ssl) - 1
	
	for _, val := range ssl {
	    if val == "" {
		return val
	    }
	}
	// compares if the last position is equal to the previous position plus the symbol "+"
	if ssl[last_index] == ssl[last_index- 1]+"+" {
	    return ssl[last_index - 1]
	} 

	return ssl[last_index]
}

//
func isServerDown(urlString string) bool {
	_, err := http.Get(urlString)

	if err != nil {
		return true
	}
	
	return false
}

//
//func serversChanged() bool {

//}

/* func getHTML(urlString string) string {

} */

// 
func getTitle(urlString string, s string) string {
	resp, err := http.Get(urlString)

	if err != nil {
		fmt.Println("No HTML file")
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
				//get the page title
				result := tokenizer.Token().Data			
				return result				
			}
		}	
	}
	return ""
}
