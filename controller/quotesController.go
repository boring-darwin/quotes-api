package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/ashoknitin/quotes-api/models"
	"github.com/ashoknitin/quotes-api/service"
	"github.com/gorilla/mux"
)

type QuoteController struct {
	service service.QuoteService
}

func NewQuotesController(service service.QuoteService) *QuoteController {
	return &QuoteController{
		service: service,
	}
}

// Route declaration
func (s *QuoteController) Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/welcome", s.welcomeHandler).Methods("GET")
	r.HandleFunc("/api/v1/quote/{id}", s.getQuotesByIdHandler).Methods("GET")
	r.HandleFunc("/api/v1/quote/all/count", s.getTotalNumberOfQuotesHandler).Methods("GET")
	r.HandleFunc("/api/v1/quote", s.addQuoteHandler).Methods("POST")
	return r
}

func (s *QuoteController) addQuoteHandler(w http.ResponseWriter, r *http.Request) {
	var quote models.Quote
	err := json.NewDecoder(r.Body).Decode(&quote)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}
	_, err = s.service.AddQuote(r.Context(), quote)
	if err != nil {
		io.WriteString(w, "unable to insert")
	}
}

func (s *QuoteController) getQuotesByIdHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])
	res, err := s.service.GetQuotesById(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	io.WriteString(w, string(res))
	// w.WriteHeader(http.StatusOK)
}

func (s *QuoteController) getTotalNumberOfQuotesHandler(w http.ResponseWriter, r *http.Request) {
	res, err := s.service.GetTotalNumberOfQuotes(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *QuoteController) welcomeHandler(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, service.TestMethod())
	// vars := mux.Vars(r)

	response := service.Welcome()

	io.WriteString(w, string(response))
	// io.WriteString(w, "Hello, world!\n")
	w.WriteHeader(http.StatusOK)
	// fmt.Fprint(w,"id: %v\n", vars["id"])
}
