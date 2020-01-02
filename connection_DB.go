package main

import (
	"fmt"
	"net"
	"time"
	"database/sql"
	_ "github.com/lib/pq"
)

func connDB() *sql.DB {      
	db, err := sql.Open("postgres",
		"postgresql://maxroach@localhost:26257/prueba_truora?ssl=true&sslmode=require&sslrootcert=certs/ca.crt&sslkey=certs/client.maxroach.key&sslcert=certs/client.maxroach.crt")

	if err != nil {
		fmt.Println(err)
	}
	// close the connection to the DB
	//defer db.Close()
	
	// Create the "domains" table
	_, erd := db.Exec("CREATE TABLE IF NOT EXISTS domains(url STRING, logo STRING, title STRING, time STRING, PRIMARY KEY (url, time))")
	// Create the "servers" table.
	_, ers := db.Exec("CREATE TABLE IF NOT EXISTS servers(address STRING, ssl_grade STRING, country STRING, owner STRING, domain STRING, time STRING, PRIMARY KEY (address,time), FOREIGN KEY (domain) REFERENCES domains(url))")	

	if ers != nil || erd != nil {
		fmt.Println(ers)
		fmt.Println(erd)
	}
	
	return db
} 

func insertDomainsDB(db *sql.DB, urlString string, logo string, title string, t time.Time) {
	_, err := db.Exec("INSERT INTO domains (url, logo, title, time) VALUES ($1, $2, $3, $4);", urlString, logo, title, t) 	

	if err != nil {
		fmt.Println(err)
	}
}

// insert the data to the servers database, db: the database
func insertServersDB(db *sql.DB, address net.IP, ssl_grade string, country string, owner string, domain string, t time.Time) {		
	a := address.String()
	_, err := db.Exec("INSERT INTO servers (address, ssl_grade, country, owner, domain, time) VALUES ($1, $2, $3, $4, $5, $6);", a, ssl_grade, country, owner, domain, t)

	if err != nil {
		fmt.Println(err)
	}
}

func query(db *sql.DB) {	
	rows, err := db.Query("SELECT * FROM servers")
	
    if err != nil {
        fmt.Println(err)
	}
	
	defer rows.Close()
	
	fmt.Println(rows)
}
