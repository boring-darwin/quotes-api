package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"io"	
	"context"
	"github.com/ashoknitin/quotes-api/service"
	"go.mongodb.org/mongo-driver/mongo"	
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/ashoknitin/quotes-api/models"	
	"strconv"
	"os"

	
)

type Env struct {
	quotes models.QuoteModel
}



// Route declaration
func (env *Env) router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/welcome", welcomeHandler).Methods("GET")	
	r.HandleFunc("/api/v1/quote/{id}", env.getQuotesByIdHandler).Methods("GET")
	return r
}



// Initiate web server
func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://nitinashok:nitin28295@spacecrusader.hi2jd.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	

	env := &Env{
		quotes: models.QuoteModel{DB: client},
	}

	port := os.Getenv("PORT")
	router := env.router()
	srv := &http.Server{
		Handler: router,
		Addr:    ":"+port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}


func (env *Env) getQuotesByIdHandler(w http.ResponseWriter, r *http.Request) {
	
	params := mux.Vars(r)	

	id1, _ := strconv.Atoi(params["id"])
	res := env.quotes.GetQuotesById(id1)
	io.WriteString(w, string(res))
	w.WriteHeader(http.StatusOK)


}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, service.TestMethod())
	// vars := mux.Vars(r)

	response := service.Welcome()

	io.WriteString(w, string(response))
	// io.WriteString(w, "Hello, world!\n")
	w.WriteHeader(http.StatusOK)
	// fmt.Fprint(w,"id: %v\n", vars["id"])
}