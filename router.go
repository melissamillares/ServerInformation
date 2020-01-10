package main

import (
	//"fmt"
	//"bytes"
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
	domain := Domain{}	
	servers := []Server{}

	//buf := new(bytes.Buffer)
    //buf.ReadFrom(r.Body)
    //newStr := buf.String()

	servers = getServers("https://gnula.nu/")
	domain = getDomain("https://gnula.nu/", servers)

	domain.insertDomainsDB()	
	u := domain.URL // get the url from the domain
	dID := domain.GetDomainID(u)
	for _, server := range servers {
		server.DomainID = dID		
		server.insertServersDB()
	} 

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
	} */  
}

func listDomainServers(w http.ResponseWriter, r *http.Request) {
	dom := &Domain{}
	d := dom.GetDomain() 

	enc := json.NewEncoder(w)
	enc.SetIndent("", "	")
	enc.Encode(d)
}
