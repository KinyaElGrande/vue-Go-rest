package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type video struct {
	ID          string `json:"ID"`
	Name       string `json:"Name"`
	Description string `json:"Description"`
	Tags		[]tag `json:"Tags"`
}
type tag struct {
	ID          string `json:"ID"`
	Name       string `json:"Name"`
}

type allVideos []video

var videos = allVideos{
	{
		ID:          "101",
		Name:       "Introduction to Go X Vue",
		Description: "<p>Getting to Learn Vue js and GoLang beiby</p>",
		Tags: []tag {
			{ ID: "1", Name: "Javascript" },
			{ ID: "3", Name: "Vue" },
			{ ID: "4", Name: "Mongo" },
		},
	},
	{
		ID:          "102",
		Name:       "Introduction to Golang 102",
		Description: "<p>Getting to Learn Vue js and GoLang beiby</p>",
		Tags: []tag {
			{ ID: "1", Name: "Javascript" },
			{ ID: "3", Name: "Vue" },
		},
	},
	{
		ID:          "103",
		Name:       "Introduction to Golang 103",
		Description: "<p>Getting to Learn Vue js and GoLang beiby</p>",
		Tags: []tag {
			{ ID: "1", Name: "Javascript" },
			{ ID: "2", Name: "Golang" },
			{ ID: "3", Name: "Vue" },
		},
	},
	{
		ID:          "104",
		Name:       "Introduction to Golang 104",
		Description: "<p>Getting to Learn Vue js and GoLang beiby</p>",
		Tags: []tag {
			{ ID: "1", Name: "Javascript" },
			{ ID: "2", Name: "Golang" },
		},
	},
	{
		ID:          "105",
		Name:       "Introduction to Golang 105",
		Description: "<p>Getting to Learn Vue js and GoLang beiby</p>",
		Tags: []tag {
			{ ID: "1", Name: "Javascript" },
			{ ID: "2", Name: "Golang" },
		},
	},
	{
		ID:          "106",
		Name:       "Introduction to Golang 106",
		Description: "<p>Getting to Learn Vue js and GoLang beiby</p>",
		Tags: []tag {
			{ ID: "1", Name: "Javascript" },
			{ ID: "2", Name: "Golang" },
			{ ID: "3", Name: "Vue" },
		},
	},
	{
		ID:          "107",
		Name:       "Introduction to Golang, Vue,VueX and Vue-Router",
		Description: "<p>Getting to Learn Vue js and GoLang beiby</p>",
		Tags: []tag {
			{ ID: "3", Name: "Vue" },
			{ ID: "2", Name: "Golang" },
		},
	},
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func getOneVideo(w http.ResponseWriter, r *http.Request) {
	videoID := mux.Vars(r)["id"]

	for _, singleVideo := range videos {
		if singleVideo.ID == videoID {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json.NewEncoder(w).Encode(singleVideo)
		}
	}
}

func getAllVideos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(videos)
}


func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", home)
	r.HandleFunc("/videos", getAllVideos).Methods("GET")
	r.HandleFunc("/videos/{id}", getOneVideo).Methods("GET")
	log.Fatal(http.ListenAndServe(":3030", r))
}