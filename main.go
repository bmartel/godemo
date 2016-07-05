package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path"
)

func main() {

	// index route (as it is, will match anything /* if no other routes match, useful for client side app renders)
	http.HandleFunc("/", serveTemplate)

	// about route (strictly matches /about)
	http.HandleFunc("/about", about)

	// api route (strictly matched /api)
	http.HandleFunc("/api", api)

	// static file serving (serve any files in the ./static directory under the /static/* route)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// about handler
func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the about page"))
}

// api handler
func api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	items := []string{"item1", "item2", "item3"}
	resp, _ := json.Marshal(items)

	w.Write(resp)
}

// compile and serve templates
func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", "index.html")

	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}
