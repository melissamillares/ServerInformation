package main

import (		
	_ "github.com/lib/pq"
	"database/sql"
)

func (s *Server) existsServer() bool {
	db := connDB()
	var resp bool	

	row := db.QueryRow(`SELECT id, address FROM servers WHERE address = $1`, s.Address).Scan(&s.ID, &s.Address)	

	if row == sql.ErrNoRows {
		resp = false
	} else {
		resp = true
	}	

	defer db.Close()
	return resp
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
