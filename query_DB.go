package main

import (	
	"time"
	_ "github.com/lib/pq"
)

func (s *Server) getServerOneHourAgo() Server {
	db := connDB()
	var server Server
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
	for _, serv := range servers {
		current := time.Now()

		y1, m1, d1 := serv.Created.Date()
		serverDate := string(y1) + string(m1) + string(d1)
		y2, m2, d2 := current.Date()
		currentDate := string(y2) + string(m2) + string(d2)

		if current.Hour() - serv.Created.Hour() >= 1 || serverDate != currentDate{
			// if the difference in update is more than 1 hour
			//|| current.Hour() - serv.Updated.Hour() >= 1 {
			server = serv					
		} /* else if serverDate != currentDate {
			server = serv			
		}	 */
	}

	defer db.Close()

	return server
}

func (s *Server) updateServer(domainID int) {
	db := connDB()

	q, _ := db.Prepare(`UPDATE servers SET (address, ssl_grade, country, owner, updated) = ($1, $2, $3, $4, $5) WHERE domainID = $6`)
		//s.Address, s.SSL_grade, s.Country, s.Owner, s.Updated, domainID)
	q.Exec(s.Address, s.SSL_grade, s.Country, s.Owner, s.Updated, domainID)

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
