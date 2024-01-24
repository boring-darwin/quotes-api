package service

import (
	"context"
	"encoding/json"

	"github.com/boring-darwin/quotes-api/models"
)

type Data struct {
	Quote  string
	Author string
}

type QuoteService struct {
}

func NewQuotesService() *QuoteService {
	return &QuoteService{}
}

func (s *QuoteService) AddQuote(ctx context.Context, quote models.Quote) (bool, error) {
	return models.AddQuote(ctx, quote)
}

func (s *QuoteService) GetQuotesById(ctx context.Context, id int) ([]byte, error) {
	return models.GetQuotesById(ctx, id)
}

func (s *QuoteService) GetTotalNumberOfQuotes(ctx context.Context) ([]byte, error) {
	return models.GetCountOfQuotes(ctx)
}

func Welcome() string {
	return "Welcome to quotes API"
}

func GetJsonResponse() ([]byte, error) {
	quote := "You miss 100% of shot you don't take"
	author := "Unkonwn"

	d := Data{quote, author}

	return json.MarshalIndent(d, "", "")
}
