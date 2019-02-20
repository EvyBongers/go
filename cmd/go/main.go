package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	shortUrls = map[string]string{}
)

func init() {
	redirectsFile := flag.String("redirects-file", "", "CSV file that contains the redirect mapping")
	flag.Parse()

	file, err := os.Open(*redirectsFile)
	redirects, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(redirects, &shortUrls); err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", landingPageHandler)
	r.HandleFunc("/{shortUrl:[a-z]+}", redirectHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func landingPageHandler(w http.ResponseWriter, r *http.Request) {
	writeBuffer(w, "<!DOCTYPE html>")
	writeBuffer(w, "<html lang=\"en\">")
	writeBuffer(w, "  <head>")
	writeBuffer(w, "    <meta charset=\"utf-8\" />")
	writeBuffer(w, "    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">")
	writeBuffer(w, "    <title>Go</title>")
	writeBuffer(w, "    <link rel=\"stylesheet\" href=\"https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css\" integrity=\"sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T\" crossorigin=\"anonymous\">")
	writeBuffer(w, "  </head>")
	writeBuffer(w, "  <body>")
	writeBuffer(w, "    <table>")
	writeBuffer(w, "      <thead>")
	writeBuffer(w, "        <tr>")
	writeBuffer(w, "          <th>Short URL</th>")
	writeBuffer(w, "          <th>Target URL</th>")
	writeBuffer(w, "        </tr>")
	writeBuffer(w, "      </thead>")
	writeBuffer(w, "      <tbody>")
	for shortUrl, targetUrl := range shortUrls {
		writeBuffer(w, "        <tr>")
		writeBuffer(w, "          <td>/"+shortUrl+"</td>")
		writeBuffer(w, "          <td>"+targetUrl+"</td>")
		writeBuffer(w, "        </tr>")
	}
	writeBuffer(w, "      </tbody>")
	writeBuffer(w, "    </table>")
	writeBuffer(w, "  </body>")
	writeBuffer(w, "</html>")
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortUrl := mux.Vars(r)["shortUrl"]
	url, ok := shortUrls[shortUrl]
	if !ok {
		http.Error(w, fmt.Sprintf("ShortURL %s does not exist", shortUrl), http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func writeBuffer(w http.ResponseWriter, text string) {
	w.Write([]byte(text + "\n"))
}
