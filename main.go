package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// Cette fonction doit respecter cette signature pour être une HandlerFunc
func hello(w http.ResponseWriter, r *http.Request) {
	// Ecriture d'un message en réponse HTTP
	fmt.Fprintf(w, "Hello Gophers")
}
func search(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	p := r.URL.Query().Get("p")
	fmt.Fprintf(w, "Search for term=%v. Current Page=%v", q, p)
}
func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "index.html")
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() failed: %v", err)
		}
		fmt.Fprintf(w, "Go login POST. value=%v\n", r.PostForm)
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "go" && password == "gopher" {
			fmt.Fprintf(w, "Cool, vous êtes connecté")
		} else {
			fmt.Fprintf(w, "Désolé, les informations sont érronés")
		}
	}
}

func temp(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"upperCase": func(str string) string {
			return strings.ToUpper(str)
		},
	}

	tmpl, err := template.New("template.html").Funcs(funcMap).ParseFiles("template.html", "navbar.html", "footer.html")
	if err != nil {
		http.NotFound(w, r)
	}

	page := struct {
		Titre string
		Name  string
		Nav   map[string]string
	}{
		Titre: "Templating",
		Name:  "samuel michaux",
		Nav:   map[string]string{"home": "/", "search": "/search", "login": "/login"},
	}

	tmpl.Execute(w, page)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// Associe un chemin à une fonction
	http.HandleFunc("/", hello)
	http.HandleFunc("/search", search)
	http.HandleFunc("/login", login)
	http.HandleFunc("/template", temp)

	// Création du serveur HTTP écoutant sur le port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
