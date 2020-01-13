package main

import (		
	_ "github.com/lib/pq"
)

func existsDomain(url string) bool {
	db := connDB()
	var resp bool

	rows, _ := db.Query(`SELECT * FROM domains WHERE url = $1`, url) 
	defer rows.Close()

	if rows.Next() {
		resp = true
	} else {
		resp = false
	}

	defer db.Close()

	return resp
}

func (d *Domain) updateDomain() {
	db := connDB()
	
	q, _ := db.Prepare(`UPDATE domains SET (servers_changed, ssl_grade, previous_ssl, logo, title, is_down, updated) = ($1, $2, $3, $4, $5, $6, $7)`)		
	q.Exec(d.Servers_Changed, d.SSL, d.Previous_SSL, d.Logo, d.Title, d.Is_Down, d.Updated)
	
	defer db.Close()
}

func (d *Domain) getDomainID(host string) int {
	db := connDB()
	var id int

	rows, _ := db.Query(`SELECT id FROM (SELECT * FROM domains ORDER BY created) WHERE url = $1`, host)	
	defer rows.Close()

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

	rows, _ := db.Query(`SELECT id, servers_changed, ssl_grade, previous_ssl, logo, title, is_down FROM domains`)	    	
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

	rows, err := db.Query(`SELECT id, servers_changed, ssl_grade, previous_ssl, logo, title, is_down 
		FROM domains ORDER BY created DESC`)
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
