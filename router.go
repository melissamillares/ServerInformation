package main

import (
	"fmt"
	"net/http"
	"html/template"
	_ "github.com/lib/pq"
	//"github.com/go-chi/chi"	
)

var tmpl = template.Must(template.ParseFiles("index.html"))

func h(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)		
	
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }

		pageurl := r.FormValue("purl")
		
		fmt.Fprintf(w, pageurl)		
	}  
}
