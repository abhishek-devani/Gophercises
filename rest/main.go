package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

var client *mongo.Client

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	// setting response type
	response.Header().Set("content-type", "application/json")

	var person Person
	json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("demo").Collection("people")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_ = cancel
	result, _ := collection.InsertOne(ctx, person)
	// return result of inserting
	json.NewEncoder(response).Encode(result)
}

func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request) // we will have route parameters in the request
	// that we are going to extract
	id, _ := primitive.ObjectIDFromHex(params["id"]) // when we pass id we have to converted into usable mongodb object id
	var person Person
	collection := client.Database("demo").Collection("people")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_ = cancel
	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + ` "}`))
		return
	}
	json.NewEncoder(response).Encode(person)
}

func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var people []Person
	collection := client.Database("demo").Collection(("people"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_ = cancel

	// bson.m{} is an empty object which means everything
	cursor, err := collection.Find(ctx, bson.M{}) // it means return everything from the collection
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + ` "}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + ` "}`))
		return
	}
	json.NewEncoder(response).Encode(people)
}

func UpdatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)

	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR", err)
	}

	var person Person
	json.NewDecoder(request.Body).Decode(&person)

	collection := client.Database("demo").Collection("people")

	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": person}

	result, _ := collection.UpdateOne(context.Background(), filter, update)
	json.NewEncoder(response).Encode(result)
}

func DeletePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := client.Database("demo").Collection("people")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_ = cancel
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + ` "}`))
		return
	}
	json.NewEncoder(response).Encode(result)

}

func main() {
	fmt.Println("starting the application...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_ = cancel

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)

	router := mux.NewRouter()
	router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", UpdatePersonEndpoint).Methods("PUT")
	router.HandleFunc("/person/{id}", DeletePersonEndpoint).Methods("DELETE")
	http.ListenAndServe(":3000", router)
}
