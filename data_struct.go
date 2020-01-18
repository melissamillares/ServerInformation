package main

import (	
	"time"
)

// struct with the server information, the info has to be parsed to JSON
// DomainID is the id of the domain the server belongs to 
type Server struct { 
	ID int 				`json:"-"`   
	Address string    	`json:"address"`
	SSL_grade string  	`json:"ssl_grade"`
	Country string   	`json:"country"`
	Owner string      	`json:"owner"` 
	DomainID int 		`json:"-"` 
	Domain string 	  	`json:"-"`
	Created time.Time 	`json:"-"`
	Updated time.Time 	`json:"-"`
}

// struct with the server information
type Domain struct {
	ID int 					`json:"-"` 
	URL string				`json:"url"`
	Servers []Server		`json:"servers"`
	Servers_Changed bool    `json:"servers_changed"`
	SSL string				`json:"ssl_grade"`
	Previous_SSL string 	`json:"previous_ssl_grade"`
	Logo string 			`json:"logo"`
	Title string 			`json:"title"`
	Is_Down bool 			`json:"is_down"`	
	Created time.Time		`json:"-"`
	Updated time.Time 		`json:"-"`
}

// struct with a Domain array
type Item struct {	
	URL string			`json:"url"`	
	Domain Domain 		`json:"info"`
	// sacar los datos por separado en lugar de poner todo el domain	
}

type Items struct {		
	Items []Item 		`json:"items"`
	// sacar los datos por separado en lugar de poner todo el domain	
}
