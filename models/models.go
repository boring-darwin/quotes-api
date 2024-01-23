package models

import (
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Quote struct {
	Id       int    `json:"id"`
	Quote    string `json:"quote"`
	Author   string `json:"author"`
	Category string `json: "category`
}

type Count struct {
	Count int
}

func getDB() (*mongo.Client, error) {
	mongoURL := os.Getenv("quotes_mongo")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client, err
}

func ListAllDB() (string, error) {
	db, err := getDB()
	if err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	databases, err := db.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	return databases[0], nil
}

func GetQuotesById(ctx context.Context, id int) ([]byte, error) {
	var quote Quote

	db, err := getDB()
	if err != nil {
		return nil, err
	}
	collection := db.Database("quotes").Collection("quotes")

	err = collection.FindOne(ctx, bson.M{"id": id}).Decode(&quote)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res, _ := json.Marshal(quote)

	return res, err
}

func AddQuote(ctx context.Context, quote Quote) (bool, error) {
	db, err := getDB()
	if err != nil {
		return false, err
	}

	collection := db.Database("quotes").Collection("quotes")
	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return false, errors.New("unable to get the document count")
	}
	fmt.Printf("Count: %d\n", count)
	quote.Id = int(count) + 1
	_, err = collection.InsertOne(ctx, quote)
	if err != nil {
		return false, err
	}
	return true, err
}

func GetCountOfQuotes(ctx context.Context) ([]byte, error) {
	db, err := getDB()
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		return nil, err
	}

	collection := db.Database("quotes").Collection("quotes")
	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		return nil, errors.New("unable to get the document count")
	}
	c := &Count{
		Count: int(count),
	}
	res, _ := json.Marshal(c)
	return res, nil
}
