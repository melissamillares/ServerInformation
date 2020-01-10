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
		//panic(err)					
		return false
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
// (e.g. s="Country: " returns the country associated with the IP)
func getInfoWhoIs(s string, ips []net.IP) string {
	for _, ip := range ips {
		result, err := whois.Whois(ip.String())		
		if err != nil {			
			panic(err)
		}
		// split the result from whois by \n
		splitResult := strings.Split(result, "\n")				
		// search in splitresult the string s
		for _, val := range splitResult {
			if strings.Contains(val, s) {				
				info := strings.Trim(val, s)				
				return info
			}
		}		
	}
	return ""	
}

// returns an array with the SSL grade of the host servers
// length: the length from the IPs array (associated with the host servers)
func getSSLGrade(host string, length int) []string {
	u := fmt.Sprintf("https://api.ssllabs.com/api/v3/analyze?host=%s", host)
	resp, err := http.Get(u)
	sslgrades := make([]string, length) // array with the length from the IPs array

	if err != nil {
		//panic(err)
		fmt.Println(err)
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

// ssl: array with the SSL grade from the servers
// returns 
func getLowerGrade(ssl []string) string {
	sort.Strings(ssl) // sorts the array in increasing order
	
	if len(ssl) == 1 { // if the array only has one element, returns that element
	    return ssl[0]
	}
	last_index := len(ssl) - 1
	
	for _, val := range ssl {
	    if val == "" { // if the array has one grade equal to "", returns ""
		return val
	    }
	}
	// compares if the last position is equal to the previous position concatenated 
	// with the symbol "+", returns the previous position
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

// Reads HTML file and return the data that matches s 
func getTitle(urlString string, s string) string {
	resp, err := http.Get(urlString)

	if err != nil {
		fmt.Println("No HTML file")
	}
	defer resp.Body.Close()

	//create a new tokenizer over the response body
	tokenizer := html.NewTokenizer(resp.Body)
	for {
		tokenType := tokenizer.Next() // get the token type
		
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()			
			if err == io.EOF {				
				break //end of the file
			}			
		}		
		
		if tokenType == html.StartTagToken {
			// get the token
			token := tokenizer.Token()			
			// if the name of the element is the string s (e.g. s="title")
			if s == token.Data {				
				tokenType = tokenizer.Next() 	// get the type of the next token			
				//get the page title
				result := tokenizer.Token().Data			
				return result				
			}
		}	
	}
	return ""
}

func isLast(d Domain) bool {
	
	return false
}
