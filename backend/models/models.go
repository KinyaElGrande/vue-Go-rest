package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type Video struct {
	ID    	primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,omitempty"`
	Name  	string             	`json:"name" bson:"name,omitempty"`
	Description string 			`json:"description" bson:"description,omitempty"`
	Tags		*Tag 			`json:"tags" bson:"tags,omitempty"`
}
type Tag struct {
	Name       string `json:"name" bson:"name,omitempty"`
}