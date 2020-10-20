package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	// POST_INDEX constant value
	POST_INDEX = "post"
	// POST_TYPE constant value
	POST_TYPE = "post"
	// DISTANCE constant value for the distance
	DISTANCE = "200km"
	// Needs to update this URL if you deploy it to cloud.
	ES_URL = "http://35.184.206.35:9200"
)

// Location location of the request
type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// Post user post
type Post struct {
	// `json:"user"` is for the json parsing of this User field
	// otherwise, by default it is 'User'.
	User     string   `json:"user"`
	Message  string   `json:"message"`
	Location Location `json:"location"`
}

func main() {
	fmt.Println("started-service")
	// All the request with endpoint "/post" will be handled by handlerPost
	http.HandleFunc("/post", handlerPost)
	http.HandleFunc("/search", handlerSearch)
	// Start http server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	// Parse from body of request to get a json object.
	fmt.Println("Received one post request")
	decoder := json.NewDecoder(r.Body)
	var p Post
	if err := decoder.Decode(&p); err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Post received: %s\n", p.Message)
}

func handlerSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one request for search")
	lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lon, _ := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
	// range is optional
	ran := DISTANCE
	if val := r.URL.Query().Get("range"); val != "" {
		ran = val + "km"
	}

	fmt.Println("range is", ran)

	// Return a fake post
	p := &Post{
		User:    "1111",
		Message: "100 Most Beautiful Place",
		Location: Location{
			Lat: lat,
			Lon: lon,
		},
	}
	// json.Marshal is used to serialize Go type
	js, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
