package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	. "github.com/cnnrznn/wordcounter/util"
)

func handle(w http.ResponseWriter, req *http.Request) {
	// get url from query string
	URL, err := url.Parse(req.URL.Query().Get("url"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid url string", http.StatusBadRequest)
		return
	}

	// download url contents
	resp, err := http.Get(URL.String())
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to fetch URL", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// execute wordcount
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to read URL response", http.StatusInternalServerError)
		return
	}

	start := time.Now()
	counts := wordcount(string(body))
	elapsed := time.Since(start)

	// write json responsea
	w.Header().Set("Content-Type", "application/json")
	wcr := WCResponse{
		Wordcount: counts,
		Time:      elapsed.Seconds(),
		URL:       URL.String(),
	}
	json.NewEncoder(w).Encode(wcr)
}

func main() {
	http.HandleFunc("/wordcount", handle)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
