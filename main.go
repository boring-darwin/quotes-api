package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ashoknitin/quotes-api/controller"
	"github.com/ashoknitin/quotes-api/service"
)

// type Env struct {
// 	quotes models.QuoteModel
// }

// Initiate web server
func main() {

	// env := &Env{
	// 	quotes: models.QuoteModel{DB: client},
	// }

	// port := os.Getenv("PORT")

	quoteController := controller.NewQuotesController(*service.NewQuotesService())
	port := "9090"
	router := quoteController.Router()
	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func RegisterController() {

}
