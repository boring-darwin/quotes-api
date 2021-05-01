package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"io"
	"github.com/ashoknitin/quotes-api/service"
	// "fmt"

	
)

func handler(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, service.TestMethod())
	// vars := mux.Vars(r)

	response, _ := service.GetJsonResponse()

	io.WriteString(w, string(response))
	// io.WriteString(w, "Hello, world!\n")
	w.WriteHeader(http.StatusOK)
	// fmt.Fprint(w,"id: %v\n", vars["id"])
}

// Route declaration
func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/quotes/{id}", handler).Methods("GET")
	return r
}

// Initiate web server
func main() {
	router := router()
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:9100",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}