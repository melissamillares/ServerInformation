package main

import (
	//"fmt"
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

func addServers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")		
	server := Server{}
	defer r.Body.Close()	

	err := json.NewDecoder(r.Body).Decode(&server)
	server.insertServersDB()
	/* servers := getServers("https://truora.com")

	for _, value := range servers {
		server = value
		server.insertServersDB()
	} */

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	//w.WriteHeader(http.StatusOK)
	//return server
}

func addDomain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")			
	defer r.Body.Close()	
	domain := Domain{}
	//server := Server{}

	domain = getDomain("https://truora.com")
	servers := getServers("https://truora.com")

	for _, server := range servers {
		server.insertServersDB()
	}
	domain.insertDomainsDB()

	err := json.NewDecoder(r.Body).Decode(&domain)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
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
		
		fmt.Fprintf(w, pageurl)		
	} */  
}

func listDomainServers(w http.ResponseWriter, r *http.Request) {
	dom := &Domain{}
	d := dom.GetDomain()
	enc := json.NewEncoder(w)
	enc.SetIndent("", "	")
	enc.Encode(d)
}
