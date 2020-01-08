package main

/* import (
	"fmt"
	"time"		
) */

var input string

func main() {
	//t()	
	database := connDB()
	routes()
	/* fmt.Print("ingrese: ")	
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
		outputSSL := getSSLGrade(outputHost, len(outputIPs))
		// lower SSL grade
		outputLowerSSL := getLowerGrade(outputSSL)
		fmt.Println(outputLowerSSL)
		//
		outputDown := isServerDown(input)
		fmt.Println(outputDown)
		
		// insert data
		//insertDomainsDB(database, outputHost, "", outputTittle, time.Now())
		for i, value := range outputIPs {  				
			//insertServersDB(database, value, outputSSL[i], outputInfoCountry, outputInfoName, outputHost, time.Now())	
		} 
		//query(database)
	}	*/

	// close the connection to the DB
	defer database.Close() 	
}
