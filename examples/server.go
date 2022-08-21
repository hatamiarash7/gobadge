package main

import (
	"fmt"
	"net/http"

	dash "github.com/hatamiarash7/gobadge/dashboard"
	"github.com/hatamiarash7/gobadge/svg"
)

// Build returns the dashboard badges as SVG.
var dashboard = func(writer http.ResponseWriter, r *http.Request) {
	label := r.URL.Query().Get("label")
	tag := r.URL.Query().Get("tag")
	color := r.URL.Query().Get("color")

	writer.Header().Set("Content-Type", "image/svg+xml")
	writer.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")

	canvas := svg.New(writer)
	view := dash.View{
		Canvas: canvas,
		Badge: dash.Badge{
			Label: label,
			Tag:   tag,
			Color: color,
		},
	}

	view.Draw()
}

const (
	port  = ":8080"
	route = "/example"
)

func main() {
	http.HandleFunc(route, dashboard)
	fmt.Printf("listening on http://localhost%s%s", port, route)
	http.ListenAndServe(port, nil)
}
