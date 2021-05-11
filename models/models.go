package models

import(
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"context"
	"log"
	"fmt"
	"time"
	"encoding/json"
)

type Quote struct {
	
	Quote string `json:"quote"`
	Author string `json:"author"`
	Category string	`json: "category`
}

// func (qu Quote) String() string {
// 	return fmt.Sprintf("[ %d, %d, %d]", qu.Quote, qu.Author, qu.Category)
// }

type QuoteModel struct {
	DB *mongo.Client
}

func (q QuoteModel) ListAllDB() string {
	
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	// err := q.DB.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	databases, err := q.DB.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	return databases[0]
}

func (q QuoteModel) GetQuotesById(id int) ([]byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	var quote Quote

	collection := q.DB.Database("quotes").Collection("test")

	// filter := bson.M{"id":id}
	err := collection.FindOne(ctx, bson.M{"id":id}).Decode(&quote)

	// fmt.Println()

	if err != nil {
		fmt.Println(err)		
	}

	res, _ := json.Marshal(quote)

	// fmt.Println(string(res))
	
	return res

}