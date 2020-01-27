package main

import (	
	"net/http"		
	"encoding/json"
	_ "github.com/lib/pq"
	"github.com/go-chi/chi"	
)

func routes() http.Handler {	
	r := chi.NewRouter()

	r.Route("/domain", func(r chi.Router) { // first endpoint		  
		r.Post("/", addDomain) 
	})
	r.Route("/getalldomains", func(r chi.Router) { // second endpoint
		r.Get("/", listAllDomains) 
	})
	http.ListenAndServe(":3000", r)

	return r
}

func addDomain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.Header.Set("Connection", "close")				
	defer r.Body.Close()

	var domain *Domain
	var url string
	var equal bool
	servers := []Server{}	
	// obtain the url value in the request body (e.g. url="https://truora.com")
	err := json.NewDecoder(r.Body).Decode(&url)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	_, e := isURL(url)
	if e == nil {
		servers = getServers(url) // get the information to be saved in database
		domain = getDomain(url, servers) // get the information to be saved in database	
		u := domain.URL // get the host from the domain
		
		lastServers := servers // save the servers in a different variable
		pssl := domain.getDomainSSL(u) // obtain the previous SSL grade from the last domain saved
		// check if it exists before inserting in database
		exists := existsDomain(hostName(url))
		if exists == true {				
			servers = getUpdatedServers(url)
			domain = getUpdatedDomain(url, servers)		
			domain.updateDomain()		
			dID := domain.getDomainID(u) // get the id from the domain
			
			if (len(lastServers) == len(servers)) { // if the number of new servers is the same				
				for i, server := range servers {
					// if there is a difference of one hour from the creation or update
					if compareOneHourAgo(server, lastServers[i]) { 
						equal = equalServers(lastServers[i], server)
						if equal == false {	// if the last servers and the new servers are different						
							domain.Servers_Changed = true 					
							domain.updateServersChangedDomain()							
						}
						domain.Previous_SSL = pssl // set the previous SSL grade from the last domain saved  
						domain.updateServersPrevious()												
					}
					server.DomainID = dID
					server.updateServer(dID) 
				}
			} else if (len(lastServers) < len(servers)) { // if there is more servers
				for _, server := range servers {
					if server.existsServer() == false { // if the server is not in the DB
						server.DomainID = dID		
						server.insertServersDB() // save the new servers in the DB
						domain.updateDomain()
					} 
					
					domain.Servers_Changed = true					
					domain.updateServersChangedDomain()
					domain.Previous_SSL = pssl
					domain.updateServersPrevious()

					server.DomainID = dID
					server.updateServer(dID)
				}												
			}			
		} else {
			domain.insertDomainsDB()	
			u := domain.URL // get the url from the domain
			dID := domain.getDomainID(u) // get the id from the domain

			for _, server := range servers {			
				server.DomainID = dID		
				server.insertServersDB()
			}
		}  
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(domain)
	} else if e != nil {
		json.NewEncoder(w).Encode("error")
		return		
	}		
}

func listAllDomains(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	dom := &Domain{}
	item := []Item{}
	items := Items{}
	domains := dom.getDomains()
	
	for _, domain := range domains {
		item = append(item, Item{	
			URL: domain.URL,
			Domain: domain,
		})
		items = Items{
			Items: item,
		}
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "	")
	enc.Encode(items)
}
