package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	dash "github.com/medenzon/gobadge/dashboard"
	"github.com/medenzon/gobadge/dashboard/svg"
)

type data struct {
	Badges []dash.Badge `json:"badges"`
}

func new(canvas *svg.Canvas, columns int) dash.View {

	file, err := ioutil.ReadFile("data.json")

	if err != nil {
		log.Printf("readfile error %v")
	}

	var data data

	log.Printf("%s", file)

	err = json.Unmarshal(file, &data)

	if err != nil {
		log.Printf("unmarshal error %v")
	}

	return dash.View{
		Canvas: canvas,
		Header: dash.Header{
			Title: "App.FrontEnd | April 1, 2020 10:18.06",
		},
		Columns: columns,
		Badges:  data.Badges,
	}
}

// Build returns the dashboard badges as SVG.
var dashboard = func(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "image/svg+xml")
	writer.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")

	param := request.FormValue("columns")
	cols, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		cols = 1
	}

	canvas := svg.New(writer)
	view := new(canvas, int(cols))
	view.Draw()
}

func main() {

	// create new http router
	router := mux.NewRouter()

	// draw endpoint
	router.HandleFunc("/dashboard/example", dashboard).Methods("GET")

	// use goroutine to serve endpoint
	http.ListenAndServe(":8080", router)
}
