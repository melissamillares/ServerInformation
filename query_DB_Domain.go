package main

import (		
	_ "github.com/lib/pq"
	"database/sql"
)

func existsDomain(url string) bool {
	db := connDB()
	var resp bool
	var id, title string

	rows := db.QueryRow(`SELECT id, title FROM domains WHERE url = $1`, url).Scan(&id, &title)
	//defer rows.Close()

	/* if rows != nil {
		resp = true
	} else {
		resp = false
	} */
	if rows == sql.ErrNoRows {
		resp = false
	} else {
		resp = true
	}

	defer db.Close()

	return resp
}

func (d *Domain) updateDomain() {
	db := connDB()
	
	q, _ := db.Prepare(`UPDATE domains SET (servers_changed, ssl_grade, previous_ssl, logo, title, is_down, updated) = ($1, $2, $3, $4, $5, $6, $7) WHERE url = $8`)		
	q.Exec(d.Servers_Changed, d.SSL, d.Previous_SSL, d.Logo, d.Title, d.Is_Down, d.Updated, d.URL)
	
	defer db.Close()
}

func (d *Domain) updateServersChangedDomain() {
	db := connDB()
	
	q, _ := db.Prepare(`UPDATE domains SET (servers_changed) = ($1) WHERE url = $2`)		
	q.Exec(d.Servers_Changed, d.URL)
	
	defer db.Close()
}

func (d *Domain) updateServersPrevious() {
	db := connDB()
	
	q, _ := db.Prepare(`UPDATE domains SET (previous_ssl) = ($1) WHERE url = $2`)		
	q.Exec(d.Previous_SSL, d.URL)
	
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

func (d *Domain) getDomainSSL(host string) string {
	db := connDB()
	var ssl string

	rows, _ := db.Query(`SELECT ssl_grade FROM (SELECT * FROM domains ORDER BY created) WHERE url = $1`, host)	
	defer rows.Close()

	for rows.Next() {    
        rows.Scan(&d.ID)        
		ssl = d.SSL
	}
	defer db.Close()
	return ssl
}

func (d *Domain) getDomain() Domain {	
	db := connDB()
	serv := Server{}
	domain := Domain{}
	var s []Server

	rows, _ := db.Query(`SELECT id, url, servers_changed, ssl_grade, previous_ssl, logo, title, is_down FROM domains`)	    	
	defer rows.Close()
		
    for rows.Next() {    
        rows.Scan(&d.ID, &d.URL, &d.Servers_Changed, &d.SSL, &d.Previous_SSL, &d.Logo, &d.Title, &d.Is_Down)        
		domain = Domain {
				ID: d.ID,
				URL: d.URL,
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

	rows, err := db.Query(`SELECT id, url, servers_changed, ssl_grade, previous_ssl, logo, title, is_down 
		FROM domains ORDER BY created DESC`)
    if err != nil {
		return domains
	}
	defer rows.Close()
		
    for rows.Next() {    
		rows.Scan(&d.ID, &d.URL, &d.Servers_Changed, &d.SSL, &d.Previous_SSL, &d.Logo, &d.Title, &d.Is_Down) 
		domain = Domain {
			ID: d.ID,
			URL: d.URL,
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
