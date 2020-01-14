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
		r.Get("/", listDomainServers)  // GET    BUSCAR POR ID???
		r.Post("/", addDomain) // POST 
	})
	r.Route("/getalldomains", func(r chi.Router) { // second endpoint
		r.Get("/", listAllDomains)  // GET 
	})
	http.ListenAndServe(":3000", r)

	return r
}

func addDomain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
				
	defer r.Body.Close()	
	var domain *Domain
	var url string
	var changed bool
	servers := []Server{}	
	// obtain the url value in the request body (e.g. url="https://truora.com")
	err := json.NewDecoder(r.Body).Decode(&url)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	servers = getServers(url) // get the information to be saved in database
	domain = getDomain(url, servers) // get the information to be saved in database	

	// check if it exists before inserting in database
	exists := existsDomain(hostName(url))
	if exists == true {
		//var servs []Server
		servers = getUpdatedServers(url)
		domain = getUpdatedDomain(url, servers)
		domain.updateDomain()

		u := domain.URL // get the host from the domain
		dID := domain.getDomainID(u) // get the id from the domain

		for _, server := range servers {
			server.DomainID = dID
			server.updateServer(dID)

			last := server.getServerOneHourAgo()											
			changed = equalServers(last, server)
			if changed == true {								
				domain.Servers_Changed = true
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
	/*domain.insertDomainsDB()	
	u := domain.URL // get the url from the domain
	dID := domain.getDomainID(u) // get the id from the domain 

	for _, server := range servers {
		var servs []Server	
		server.DomainID = dID		
		server.insertServersDB()
		
		last := server.getServerOneHourAgo()											
		changed = equalServers(last, server)
		
		if changed == true {								
			servs = getUpdatedServers(url)
			for _, s := range servs {
				s.updateServer(dID)
			}

			domain = getUpdatedDomain(url, servs)
			domain.updateDomain()	
		} 									
	}  */

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(domain)	
}

func listDomainServers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	dom := &Domain{}
	d := dom.getDomain() 

	enc := json.NewEncoder(w)
	enc.SetIndent("", "	")
	enc.Encode(d)
}

func listAllDomains(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	dom := &Domain{}
	d := dom.getDomains()
	
	items := Items{		
		Domains: d,
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "	")
	enc.Encode(items)
}
