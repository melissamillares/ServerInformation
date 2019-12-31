package main

import (
	"net"
	"net/url"
	"net/http"

	"html/template"

	"fmt"
	"strings"

	"github.com/likexian/whois-go"		

	"database/sql"
	_ "github.com/lib/pq"
	//"github.com/go-chi/chi"	
)

var input string
// 
var tmpl = template.Must(template.ParseFiles("index.html"))

func main() {
	//fmt.Println(isURL("https://google.com"))	

	/* r := chi.NewRouter()

	r.Get("/", h)
	http.ListenAndServe(":3000", r) */

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

		// insert data
		insertDomainsDB(database, outputHost, "", "")
		for _, value := range outputIPs {  		
			insertServersDB(database, value, "A", outputInfoCountry, outputInfoName, outputHost)	
		}
		//query(database)
	}	

	// close the connection to the DB
	defer database.Close()
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

// get the information s from the query whois
func getInfoWhoIs(s string, ips []net.IP) string {
	for _, value := range ips {
		result, err := whois.Whois(value.String())
		
		if err != nil {
			//fmt.Println(result)
			panic(err)
		}

		splitResult := strings.Split(result, "\n")
				
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

func getTittle() string {
	
	return ""
}

func connDB() *sql.DB {      
	db, err := sql.Open("postgres",
		"postgresql://maxroach@localhost:26257/prueba_truora?ssl=true&sslmode=require&sslrootcert=certs/ca.crt&sslkey=certs/client.maxroach.key&sslcert=certs/client.maxroach.crt")

	if err != nil {
		fmt.Println(err)
	}
	// close the connection to the DB
	//defer db.Close()
	
	// Create the "servers" table.
	_, ers := db.Exec("CREATE TABLE IF NOT EXISTS servers(address STRING PRIMARY KEY, ssl_grade STRING, country STRING, owner STRING, domain STRING, FOREIGN KEY (domain) REFERENCES domains(uri))")
	// Create the "domains" table
	_, erd := db.Exec("CREATE TABLE IF NOT EXISTS domains(url STRING PRIMARY KEY, logo STRING, tittle STRING)")

	if ers != nil || erd != nil {
		fmt.Println(ers)
		fmt.Println(erd)
	}
	
	return db
} 

func insertDomainsDB(db *sql.DB, url string, logo string, tittle string) {
	_, err := db.Exec("INSERT INTO domains (url, logo, tittle) VALUES ($1, $2, $3);", url, logo, tittle) 	

	if err != nil {
		fmt.Println("Error inserting data")
	}
}

// insert the data to the servers database, db: the database
func insertServersDB(db *sql.DB, address net.IP, ssl_grade string, country string, owner string, domain string) {		
	a := address.String()
	_, err := db.Exec("INSERT INTO servers (address, ssl_grade, country, owner, domain) VALUES ($1, $2, $3, $4, $5);", a, ssl_grade, country, owner, domain)

	if err != nil {
		fmt.Println(err)
	}
}

func query(db *sql.DB) {	
	rows, err := db.Query("SELECT address, ssl_grade, country, owner FROM servers")
	
    if err != nil {
        fmt.Println(err)
	}
	
	defer rows.Close()
	
	fmt.Println(rows)
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
