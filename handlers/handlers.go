package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// Index handles requests to "/" and "/Home"
func Index(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/", "/Index", "/index", "/INDEX":
		http.Redirect(w, r, "/index", http.StatusMovedPermanently)
		if r.Method == http.MethodGet {
			serveTemplate(w, "index.html")
		}
	case "Make_referral", "make_referral":
		http.Redirect(w, r, "/make_referral", http.StatusMovedPermanently)
		
		if r.Method == http.MethodGet {
			serveTemplate(w, "make_referral.html")
		}

	// default:
	// 	if r.Method == http.MethodGet {
	// 		tmpl := template.Must(template.ParseFiles("templates/404.html"))
	// 		tmpl.Execute(w, "404: Page not Found!")

	// 	}
	}
}


// serveTemplate loads and executes a template file
func serveTemplate(w http.ResponseWriter, filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		http.Error(w, "404 Page not Found", http.StatusNotFound)
		return
	}
	tmpl := template.Must(template.ParseFiles(filename))
	errr := tmpl.Execute(w, nil)
	if errr != nil {
		log.Println("500 Internal Server Error")
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}

