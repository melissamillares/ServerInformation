package main

import (
	"fmt"
	"time"
	_ "github.com/lib/pq"
)

func (s *Server) getServerOneHourAgo() Server {
	db := connDB()
	var server Server
	var servers []Server

	rows, err := db.Query(`SELECT address, ssl_grade, country, owner, created, updated FROM servers WHERE domain = $1`, s.Domain)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

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

	q, err := db.Prepare(`UPDATE servers SET (address, ssl_grade, country, owner, updated) = ($1, $2, $3, $4, $5) WHERE domainID = $6`)
		//s.Address, s.SSL_grade, s.Country, s.Owner, s.Updated, domainID)
	q.Exec(s.Address, s.SSL_grade, s.Country, s.Owner, s.Updated, domainID)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
}

// 
func (s *Server) getServers(dID int) []Server {	
	db := connDB()	
	servers := []Server{} // servers array 
	rows, _ := db.Query(`SELECT address, ssl_grade, country, owner, domain FROM servers WHERE domainID = $1`, dID)

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

func (d *Domain) updateDomain() {
	db := connDB()
	
	q, err := db.Prepare(`UPDATE domains SET (servers_changed, ssl_grade, previous_ssl, logo, title, is_down, updated) = ($1, $2, $3, $4, $5, $6, $7)`)
		//d.Servers_Changed, d.SSL, d.Previous_SSL, d.Updated)
	q.Exec(d.Servers_Changed, d.SSL, d.Previous_SSL, d.Logo, d.Title, d.Is_Down, d.Updated)
	
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}

func (d *Domain) getDomainID(host string) int {
	db := connDB()
	var id int

	rows, err := db.Query(`SELECT id FROM (SELECT * FROM domains ORDER BY created) WHERE url = $1`, host)
	if err != nil {
		fmt.Println(err)		
	}
	defer db.Close()

	for rows.Next() {    
        rows.Scan(&d.ID)        
		id = d.ID
	}

	defer db.Close()
	return id
}

func (d *Domain) getDomain() Domain {	
	db := connDB()
	serv := Server{}
	domain := Domain{}
	var s []Server

	rows, err := db.Query(`SELECT id, servers_changed, ssl_grade, previous_ssl, logo, title, is_down FROM domains`)	
    if err != nil {
		fmt.Println(err)		
	}	
	defer rows.Close()
		
    for rows.Next() {    
        rows.Scan(&d.ID, &d.Servers_Changed, &d.SSL, &d.Previous_SSL, &d.Logo, &d.Title, &d.Is_Down)        
		domain = Domain {
				ID: d.ID,
				Servers: s,
				Servers_Changed: d.Servers_Changed,
				SSL: d.SSL,
				Previous_SSL: d.Previous_SSL,
				Logo: d.Logo,
				Title: d.Title,
				Is_Down: d.Is_Down,
			} 				
	}

	s = serv.getServers(domain.ID)
	domain.Servers = s

	defer db.Close()
    return domain
}

func (d *Domain) getDomains() []Domain {	
	db := connDB()
	serv := Server{}
	servers := []Server{}
	domain := Domain{}
	domains := []Domain{}
	//s := serv.getServers(d.ID)

	rows, err := db.Query(`SELECT id, servers_changed, ssl_grade, previous_ssl, logo, title, is_down FROM domains`)
    if err != nil {
		return domains
	}
	defer rows.Close()
		
    for rows.Next() {    
		rows.Scan(&d.ID, &d.Servers_Changed, &d.SSL, &d.Previous_SSL, &d.Logo, &d.Title, &d.Is_Down) 
		domain = Domain {
			ID: d.ID,
			Servers: servers,
			Servers_Changed: d.Servers_Changed,
			SSL: d.SSL,
			Previous_SSL: d.Previous_SSL,
			Logo: d.Logo,
			Title: d.Title,
			Is_Down: d.Is_Down,
		}    
		servers = serv.getServers(domain.ID)
		domain.Servers = servers   
		domains = append(domains, domain)		
	}	

	defer db.Close()
    return domains
} 
