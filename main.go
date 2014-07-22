package tt2maps

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/cryptix/go-tenten"
)

const (
	devLat = 1
	devLon = 1

	meanLat = 1
	meanLon = 1
)

func init() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/rand", randHandler)
	http.HandleFunc("/redir", redirHandler)
}

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "Hello, world!")
}

func randHandler(rw http.ResponseWriter, req *http.Request) {
	lat := rand.NormFloat64()*devLat + meanLat
	lon := rand.NormFloat64()*devLon + meanLon

	fmt.Fprintf(rw, "Hello! (%.3f, %.3f) is %s\n", lat, lon, tenten.Encode(lat, lon))
}

func redirHandler(rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, "http://maps.google.com/?q=10,10", http.StatusFound)
}
