package main

import (
	"fmt"
	"time"		
)

var input string

func main() {
	//r := chi.NewRouter()

	//r.Get("/", h)
	//http.ListenAndServe(":3000", r)

	database := connDB()

	fmt.Print("ingrese: ")	
	fmt.Scanf("%s", &input)

	outputURL := isURL(input)

	if outputURL {
		outputHost := hostName(input)
		//fmt.Println(outputHost)

		// array with the server IPs 
		outputIPs := getIP(outputHost)			
		// server owner
		outputInfoName := getInfoWhoIs("OrgName: ", outputIPs)
		// server country
		outputInfoCountry := getInfoWhoIs("Country: ", outputIPs)				
		// domain html tittle
		outputTittle := getTitle(input, "title")
		// SSL grade
		outputSSL := getSSLGrade(outputHost)
		
		// insert data
		insertDomainsDB(database, outputHost, "", outputTittle, time.Now())
		for _, value := range outputIPs {  		
			insertServersDB(database, value, outputSSL, outputInfoCountry, outputInfoName, outputHost, time.Now())	
		}
		//query(database)
	}	

	// close the connection to the DB
	defer database.Close()
}
