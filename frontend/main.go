package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

var (
	tpl *template.Template
)

func handleIndex(w http.ResponseWriter, req *http.Request) {
	// check method, path
	if req.Method != http.MethodGet {
		http.Error(w, "Method not implemented", http.StatusNotImplemented)
		return
	}

	log.Println(req.Method, req.URL.Path)
	log.Println(req.URL.RawQuery)

	// valid request
	w.WriteHeader(http.StatusOK)
	err := tpl.ExecuteTemplate(w, "index", map[string]interface{}{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func handlePost(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	log.Println(req.Method, req.URL.Path)

	// parse url
	URL := req.FormValue("url")
	if URL == "" {
		http.Error(w, "Bad form data", http.StatusBadRequest)
	}
	u, _ := url.Parse("/")
	q := u.Query()
	q.Set("url", URL)
	u.RawQuery = q.Encode()

	http.Redirect(w, req, u.String(), http.StatusFound)
}

func main() {
	t, err := template.New("index").ParseGlob("templates/index.tpl")
	if err != nil {
		log.Fatal("Failed to parse templates: ", err)
	}
	tpl = t

	// dynamic content
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/post", handlePost)

	// static file server
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
