package main

import (		
	_ "github.com/lib/pq"
)

func (s *Server) serversSameDomain() []Server {
	db := connDB()
	//var server Server
	var servers []Server

	rows, _ := db.Query(`SELECT address, ssl_grade, country, owner, created, updated FROM servers WHERE domain = $1`, s.Domain)	
	defer rows.Close()

	for rows.Next() {    
		rows.Scan(&s.Address, &s.SSL_grade, &s.Country, &s.Owner, &s.Created, &s.Updated)
		servers = append(servers, Server{
				Address: s.Address, 
				SSL_grade: s.SSL_grade,
				Country: s.Country,
				Owner: s.Owner,
				Created: s.Created,
				Updated: s.Updated,
		})		
	}	
	defer db.Close()

	return servers
}

func (s *Server) updateServer(domainID int) {
	db := connDB()

	q, _ := db.Prepare(`UPDATE servers SET (ssl_grade, country, owner, updated) = ($1, $2, $3, $4) WHERE domainID = $5`)
		//s.Address, s.SSL_grade, s.Country, s.Owner, s.Updated, domainID)
	q.Exec(s.SSL_grade, s.Country, s.Owner, s.Updated, domainID)

	defer db.Close()
}

// 
func (s *Server) getServers(dID int) []Server {	
	db := connDB()	
	servers := []Server{} // servers array 
	rows, _ := db.Query(`SELECT address, ssl_grade, country, owner, domain FROM servers WHERE domainID = $1`, dID)
	defer rows.Close()

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
