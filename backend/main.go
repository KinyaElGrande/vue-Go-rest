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
}

type allVideos []video

var videos = allVideos{
	{
		ID:          "101",
		Name:       "Introduction to Golang",
		Description: "<p>Getting to Learn Vue js and GoLang beiby</p>",
	},
	{
		ID:          "102",
		Name:       "Introduction to Golang 102",
		Description: "<p>Getting to Learn Vue js and GoLang beiby</p>",
	},
	{
		ID:          "103",
		Name:       "Introduction to Golang 103",
		Description: "<p>Getting to Learn Vue js and GoLang beiby</p>",
	},
	{
		ID:          "104",
		Name:       "Introduction to Golang 104",
		Description: "<p>Getting to Learn Vue js and GoLang beiby</p>",
	},
	{
		ID:          "105",
		Name:       "Introduction to Golang 105",
		Description: "<p>Getting to Learn Vue js and GoLang beiby</p>",
	},
	{
		ID:          "106",
		Name:       "Introduction to Golang 106",
		Description: "<p>Getting to Learn Vue js and GoLang beiby</p>",
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