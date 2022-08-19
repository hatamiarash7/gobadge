package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	dash "github.com/hatamiarash7/gobadge/dashboard"
	"github.com/hatamiarash7/gobadge/svg"
)

func new(canvas *svg.Canvas) dash.View {

	file, err := ioutil.ReadFile("data.json")

	if err != nil {
		log.Printf("readfile error %v", err)
	}

	var data dash.Badge

	err = json.Unmarshal(file, &data)

	if err != nil {
		log.Printf("unmarshal error %v", err)
	}

	return dash.View{
		Canvas: canvas,
		Badge:  data,
	}
}

// Build returns the dashboard badges as SVG.
var dashboard = func(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "image/svg+xml")
	writer.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")

	canvas := svg.New(writer)
	view := new(canvas)
	view.Draw()
}

const (
	port  = ":8080"
	route = "/dashboard/example"
)

func main() {

	// draw endpoint
	http.HandleFunc(route, dashboard)

	log.Printf("listening on http://localhost%s%s", port, route)

	// use goroutine to serve endpoint
	http.ListenAndServe(port, nil)
}
