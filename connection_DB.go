package main

import (
	"fmt"
	//"net"
	"time"
	"strings"
	"io/ioutil"
	"database/sql"
	_ "github.com/lib/pq"
)

// struct with the server information, the info has to be parsed to JSON 
type Server struct {    
	Address string    	`json:"address"`
	SSL_grade string  	`json:"ssl_grade"`
	Country string   	`json:"country"`
	Owner string      	`json:"owner"` 
	Domain string 	  	`json:"-"` 
	Created time.Time 	`json:"-"`
}

// struct with the server information
type Domain struct {
	URL string				`json:"-"`
	Servers []Server		`json:"servers"`
	Servers_Changed bool    `json:"servers_changed"`
	SSL string				`json:"ssl_grade"`
	Previous_SSL string 	`json:"previous_ssl_grade"`
	Logo string 			`json:"logo"`
	Title string 			`json:"title"`
	Is_Down bool 			`json:"is_down"`	
	Created time.Time		`json:"-"`
}

// variables for connection with database
var user string
var host string
var port string
var database string

// read the file with the database connection information
func readFile() {
	f, err := ioutil.ReadFile("info.txt")    
    
    if err != nil {
		panic(err)
	}
	splitResult := strings.Split(string(f), "\n")
	for _, val := range splitResult {
		if strings.Contains(val, "user") {
			u := strings.Trim(val, "user: ")			
			user = u														
		} else if strings.Contains(val, "h") {
			u := strings.Trim(val, "h: ")			
			host = u														
		} else if strings.Contains(val, "port") {
			u := strings.Trim(val, "port: ")			
			port = u														
		} else if strings.Contains(val, "db") {
			u := strings.Trim(val, "db: ")			
			database = u														
		} 
	}	
}

// make the connection with the database
func connDB() *sql.DB {   
	readFile()
	d := "postgresql://"+user+"@"+host+":"+port+"/"+database+"?ssl=true&sslmode=require&sslrootcert=certs/ca.crt&sslkey=certs/client."+user+".key&sslcert=certs/client."+user+".crt"   
	db, err := sql.Open("postgres", d)
	
	if err != nil {
		fmt.Println(err)
	}
	// close the connection to the DB
	//defer db.Close()
	// Create the "domains" table
	_, erd := db.Exec(`CREATE TABLE IF NOT EXISTS domains(url STRING, servers_changed BOOL, ssl_grade STRING, previous_ssl STRING, logo STRING, title STRING, is_down BOOL, created STRING, PRIMARY KEY (url, created))`)
	// Create the "servers" table.
	_, ers := db.Exec(`CREATE TABLE IF NOT EXISTS servers(address STRING, ssl_grade STRING, country STRING, owner STRING, domain STRING, created STRING, PRIMARY KEY (address,created), FOREIGN KEY (domain) REFERENCES domains(url))`)	

	if ers != nil || erd != nil {
		fmt.Println(ers)
		fmt.Println(erd)		
	}
	
	return db
} 

//
func (d Domain) insertDomainsDB() {
	db := connDB()
	
	q, err := db.Prepare(`INSERT INTO domains (url, servers_changed, ssl_grade, previous_ssl, logo, title, is_down, created) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`)
	q.Exec(d.URL, d.Servers_Changed, d.SSL, d.Previous_SSL, d.Logo, d.Title, d.Is_Down, d.Created)
	
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
}

// insert the data to the servers database
func (s Server) insertServersDB()  {		
	db := connDB()
	
	q, err := db.Prepare(`INSERT INTO servers (address, ssl_grade, country, owner, domain, created) VALUES ($1, $2, $3, $4, $5, $6)`)
	q.Exec(s.Address, s.SSL_grade, s.Country, s.Owner, s.Domain, s.Created)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
}

// 
func (s *Server) GetServers() []Server {	
	db := connDB()	
	servers := []Server{} // servers array 
	rows, _ := db.Query(`SELECT address, ssl_grade, country, owner, domain FROM servers `)

	for rows.Next() {
		rows.Scan(&s.Address, &s.SSL_grade, &s.Country, &s.Owner, &s.Domain) 
		servers = append(servers, Server{
			Address: s.Address,
			SSL_grade: s.SSL_grade,
			Country: s.Country,
			Owner: s.Owner,
			Domain: s.Domain,
		})
	}
	
	defer db.Close()
    return servers 
}

func (d *Domain) GetDomain() Domain {	
	db := connDB()
	serv := Server{}
	domain := Domain {}
	s := serv.GetServers()

	rows, err := db.Query(`SELECT servers_changed, ssl_grade, previous_ssl, logo, title, is_down FROM domains`)
	
    if err != nil {
		fmt.Println(err)		
	}
	
	defer rows.Close()
		
    for rows.Next() {    
        rows.Scan(&d.Servers_Changed, &d.SSL, &d.Previous_SSL, &d.Logo, &d.Title, &d.Is_Down)        
		domain = Domain {
				Servers: s,
				Servers_Changed: d.Servers_Changed,
				SSL: d.SSL,
				Previous_SSL: d.Previous_SSL,
				Logo: d.Logo,
				Title: d.Title,
				Is_Down: d.Is_Down,
			} 				
	}

	defer db.Close()
    return domain
}
