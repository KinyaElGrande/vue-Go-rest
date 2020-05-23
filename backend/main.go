package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"./models"
	"./utils"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

//Getting all videos
func getVideos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Video array
	var videos []models.Video

	//Connection mongoDB with utils class
	collection := utils.ConnectDB()

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		utils.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var video models.Video
		// & character returns the memory address of the following variable.
		err := cur.Decode(&video) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		videos = append(videos, video)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(videos) // encode similar to serialize process.
}

func getVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var video models.Video
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := utils.ConnectDB()

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&video)

	if err != nil {
		utils.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(video)
}

func createVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var video models.Video

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&video)

	// connect db
	collection := utils.ConnectDB()

	// insert our video model.
	result, err := collection.InsertOne(context.TODO(), video)

	if err != nil {
		utils.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
	
}

func updateVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var video models.Video

	collection := utils.ConnectDB()

	// Create filter
	filter := bson.M{"_id": id}

	// Read update model from body request
	_ = json.NewDecoder(r.Body).Decode(&video)

	// prepare update model.
	update := bson.D{ primitive.E{Key: "$set",Value: bson.D{
			{Key: "name",Value: video.Name},
			{Key: "description",Value: video.Description},
			{Key: "Tags",Value: bson.D{
				{Key: "name",Value: video.Tags.Name},
			}},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&video)

	if err != nil {
		utils.GetError(err, w)
		return
	}

	video.ID = id

	json.NewEncoder(w).Encode(video)
	
}

func deleteVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get params
	var params = mux.Vars(r)

	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])

	collection := utils.ConnectDB()

	// prepare filter.
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		utils.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
	
}


func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", home)
	r.HandleFunc("/videos", getVideos).Methods("GET")
	r.HandleFunc("/videos/{id}", getVideo).Methods("GET")
	r.HandleFunc("/videos", createVideo).Methods("POST")
	r.HandleFunc("/videos/{id}", updateVideo).Methods("PUT")
	r.HandleFunc("/videos/{id}", deleteVideo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3030", r))
}