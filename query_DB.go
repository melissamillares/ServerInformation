package main

import (
	"fmt"
	_ "github.com/lib/pq"
)

// 
func (s *Server) GetServers(dID int) []Server {	
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

func (d *Domain) GetDomainID(host string) int {
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

func (d *Domain) GetDomain() Domain {	
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

	s = serv.GetServers(domain.ID)
	domain.Servers = s

	defer db.Close()
    return domain
}

/* func (d *Domain) GetDomains() []Domain {	
	db := connDB()
	serv := Server{}
	domain := []Domain{}
	s := serv.GetServers(domain.ID)

	rows, err := db.Query(`SELECT servers_changed, ssl_grade, previous_ssl, logo, title, is_down FROM domains`)
	
    if err != nil {
		fmt.Println(err)		
	}
	
	defer rows.Close()
		
    for rows.Next() {    
        rows.Scan(&d.Servers_Changed, &d.SSL, &d.Previous_SSL, &d.Logo, &d.Title, &d.Is_Down)        
		domain = append(domain, Domain {
				Servers: s,
				Servers_Changed: d.Servers_Changed,
				SSL: d.SSL,
				Previous_SSL: d.Previous_SSL,
				Logo: d.Logo,
				Title: d.Title,
				Is_Down: d.Is_Down,
			})
	}

	defer db.Close()
    return domain
} */
