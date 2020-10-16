package main

import "fmt"

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
	fmt.Println("Hello world")
}
