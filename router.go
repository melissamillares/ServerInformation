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
		r.Get("/", listServers)  // GET /
		r.Post("/", addServer) // POST /
	})
	http.ListenAndServe(":3000", r)

	return r
}

func addServer(w http.ResponseWriter, r *http.Request) {
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

func listServers(w http.ResponseWriter, r *http.Request) {
	dom := &Domain{}
	//serv := &Server{}
	//items := serv.GetServers()
	d := dom.GetDomain()
	json.NewEncoder(w).Encode(d)
}
