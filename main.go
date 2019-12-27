package main

import (
	"net"
	"net/url"
	"fmt"		
	"github.com/likexian/whois-go"		
	"strings"
)

// variable representing the url without http
var hoststring string

func main() {
	//fmt.Println(isURL("https://google.com"))		
	fmt.Print("ingrese: ")
	var input string
	fmt.Scanf("%s", &input)

	outputURL := isURL(input)
	fmt.Println(outputURL)	

	outputHost := hostName(input)
	fmt.Println(outputHost)

	outputIPs := getIP(outputHost)
	fmt.Println(outputIPs)	
	
	outputInfoName := getInfoWhoIs("OrgName: ", outputIPs)
	fmt.Println(outputInfoName)

	outputInfoCountry := getInfoWhoIs("Country: ", outputIPs)
	fmt.Println(outputInfoCountry)
}

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

// gets the information s from the query whois
func getInfoWhoIs(s string, ips []net.IP) string {
	for _, value := range ips {
		result, err := whois.Whois(value.String())
		
		if err != nil {
			//fmt.Println(result)
			panic(err)
		}

		splitResult := strings.Split(result, "\n")
		//fmt.Printf("%q\n", splitResult)
				
		// buscar en splitresult orgname
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
