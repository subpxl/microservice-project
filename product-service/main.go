package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// product service crud product sample api

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("product-service").Collection("products")
}
func main() {

	r := chi.NewRouter()

	r.Get("/", GetAll)
	r.Get("/{id}", GetById)
	r.Post("/", CreateProduct)
	r.Put("/{id}", UpdateProduct)
	r.Delete("/{id}", DeleteProduct)

	log.Fatal(http.ListenAndServe(":8000", nil))

}

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

var products = []Product{
	{ID: "1", Name: "banana", Price: 100},
	{ID: "2", Name: "apple", Price: 200},
	{ID: "3", Name: "kiwi", Price: 300},
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	var products []Product

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var product Product
		err := cursor.Decode(&product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		products = append(products, product)
	}
	json.NewEncoder(w).Encode(products)
}

func GetById(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	fmt.Println(id)
	// var product Product
	// err := collection.FindOne(context.TODO(), bson.D).Decode(&product)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	// json.NewEncoder(w).Encode(product)
	json.NewEncoder(w).Encode("this is product")

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	json.NewEncoder(w).Encode(product)

	json.NewEncoder(w).Encode("product created successfully")
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	fmt.Println(id)
	// var product Product
	// json.NewEncoder(w).Encode(product)
	json.NewEncoder(w).Encode("product updated successfully")

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	fmt.Println(id)
	// var product Product
	json.NewEncoder(w).Encode("product deteled successfully")
}
