package main

import (
	"fmt"
	"net/http"
	"strconv"
	"stresstest/lib"
	"time"
)

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {

	streams := r.FormValue("streams")
	if len(streams) <= 0 {
		http.Error(w, "'streams' param is required", 400)
	}
	streamsInt, errInt := strconv.Atoi(streams)
	if errInt != nil {
		http.Error(w, "'streams' param must be integer", 400)
	}

	itter := r.FormValue("iteration")
	if len(itter) <= 0 {
		http.Error(w, "'iteration' param is required", 400)
	}
	itterInt, errItter := strconv.Atoi(itter)
	if errItter != nil {
		http.Error(w, "'iteration' param must be integer", 400)
	}

	method := r.FormValue("method")
	if len(method) <= 0 {
		method = "GET"
	}

	url := r.FormValue("url")
	if len(url) <= 0 {
		http.Error(w, "'url' param is required", 400)
	}
	fmt.Println(itter)
	fmt.Println(streams)
	fmt.Println(url)

	strain := lib.Strain{Url: url, Streams: streamsInt, Method: method, Itt: itterInt, Finishing: 0}
	strain.Requests(w)

	for true {
		if strain.Finishing >= streamsInt {
			fmt.Fprintf(w, "Finished, streams: "+strconv.Itoa(len(strain.Listing)))
			fmt.Fprintf(w, "<br> \n")
			for index, element := range strain.Listing {
				fmt.Fprintf(w, strconv.Itoa(index)+" -> "+element)
				fmt.Fprintf(w, "<br> \n")
			}

			break
		}
		time.Sleep(1)
	}
}

func main() {
	http.HandleFunc("/", HomeRouterHandler)
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		panic(err)
	}
}
