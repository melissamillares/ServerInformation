package main

import (	
	"net/http"
	"html/template"
	"encoding/json"
	_ "github.com/lib/pq"
	"github.com/go-chi/chi"	
)

var tmpl = template.Must(template.ParseFiles("index.html"))

func routes() http.Handler {
	r := chi.NewRouter()

	r.Route("/servers", func(r chi.Router) { // first endpoint
		r.Get("/", listDomainServers)  // GET 
		r.Post("/", addDomain) // POST 
	})
	http.ListenAndServe(":3000", r)

	return r
}

func addDomain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")			
	defer r.Body.Close()	
	var domain *Domain
	var url string
	servers := []Server{}	
	// obtain the url value in the request body (e.g. url="https://truora.com")
	err := json.NewDecoder(r.Body).Decode(&url)

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	servers = getServers(url) // get the information to be saved in database
	domain = getDomain(url, servers) // get the information to be saved in database

	domain.insertDomainsDB()	
	u := domain.URL // get the url from the domain
	dID := domain.GetDomainID(u) // get the id from the domain	

	for _, server := range servers {
		server.DomainID = dID		
		server.insertServersDB()
	} 

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(domain)
	/* tmpl.Execute(w, nil)		
	
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }

		pageurl := r.FormValue("purl")	
	} */  
}

func listDomainServers(w http.ResponseWriter, r *http.Request) {
	dom := &Domain{}
	d := dom.GetDomain() 

	enc := json.NewEncoder(w)
	enc.SetIndent("", "	")
	enc.Encode(d)
}
