package tt2maps

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"

	"github.com/cryptix/go-tenten"
)

func init() {
	rand.Seed(time.Now().Unix())

	http.HandleFunc("/", somewhereHandler)
}

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "Hello, world!")
}

var somewhereTmpl = template.Must(template.New("somewhere").Parse(`
	<h1>Hello</h1>
	<p>({{.lat}},{{.lon}}) is: {{.tt}}</p>
	<h4>Take me there</h4>
	<ul>
		<li><a href="http://www.openstreetmap.org/#map=5/{{.lat}}/{{.lon}}">OpenStreetMap</a> </li>
		<li><a href="http://maps.google.com/?q={{.lat}},{{.lon}}">Google Maps</a></li>
	</ul>
`))

func somewhereHandler(rw http.ResponseWriter, req *http.Request) {
	lat := rand.Float64() * 180
	lon := rand.Float64() * 360
	tt := tenten.Encode(lat, lon)

	somewhereTmpl.Execute(rw, map[string]interface{}{
		"lat": fmt.Sprintf("%.2f", lat),
		"lon": fmt.Sprintf("%.2f", lon),
		"tt":  tt,
	})
}
