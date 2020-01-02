package main

import (
	"fmt"
	"net/http"
	"time"
	"html/template"
	_ "github.com/lib/pq"
	//"github.com/go-chi/chi"	
)

var input string
// 
var tmpl = template.Must(template.ParseFiles("index.html"))

func main() {
	//fmt.Println(isURL("https://google.com"))	

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
		fmt.Println(outputInfoName)
		// server country
		outputInfoCountry := getInfoWhoIs("Country: ", outputIPs)		
		fmt.Println(outputInfoCountry)
		// domain html tittle
		outputTittle := getTitle(input, "title")
		fmt.Println(outputTittle)
		
		// insert data
		insertDomainsDB(database, outputHost, "", outputTittle, time.Now())
		for _, value := range outputIPs {  		
			insertServersDB(database, value, "A", outputInfoCountry, outputInfoName, outputHost, time.Now())	
		}
		//query(database)
	}	

	// close the connection to the DB
	defer database.Close()
}

func h(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)		
	
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }

		pageurl := r.FormValue("purl")
		
		fmt.Fprintf(w, pageurl)		
	}  
}
