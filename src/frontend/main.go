package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"

	. "github.com/cnnrznn/wordcounter/util"
)

var (
	tpl *template.Template
)

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

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}

func doBackendReq(URL *url.URL, wcr *WCResponse) error {
	resp, err := http.Get(fmt.Sprintf("http://%s/wordcount?url=%s",
		os.Getenv("BACKEND_API_ADDR"),
		url.QueryEscape(URL.String())))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, wcr); err != nil {
		return err
	}

	return nil
}

func handleIndex(w http.ResponseWriter, req *http.Request) {
	// check method, path
	if req.Method != http.MethodGet {
		http.Error(w, "Method not implemented", http.StatusNotImplemented)
		return
	}

	log.Println(req.Method, req.URL.Path, req.URL.Query())

	// check url is valid
	URL, err := url.Parse(req.URL.Query().Get("url"))
	if err != nil {
		http.Error(w, "Not a valid url", http.StatusBadRequest)
		return
	}

	wcr := WCResponse{}

	if len(URL.String()) > 0 {
		if err := doBackendReq(URL, &wcr); err != nil {
			http.Error(w, "Backend request failed", http.StatusInternalServerError)
			return
		}
	}

	// order wordcounts descending
	sort.Slice(wcr.Wordcount, func(i, j int) bool {
		if wcr.Wordcount[i].Val > wcr.Wordcount[j].Val {
			return true
		}
		return false
	})

	err = tpl.ExecuteTemplate(w, "index", map[string]interface{}{
		"Wordcount": wcr.Wordcount,
		"Time":      wcr.Time,
		"URL":       wcr.URL,
	})
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
