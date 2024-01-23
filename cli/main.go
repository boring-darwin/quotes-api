package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ashoknitin/quotes-api/models"
)

func main() {
	// useColor := flag.Bool("color", false, "display colorized output")
	// flag.Parse()

	// if *useColor {
	getQuote()
	return
	// }
}

func getQuote() {
	c := getTotalNumberOfQuotes()
	// URL of the API you want to call
	apiURL := "http://45.79.122.143:9090/api/v1/quote/" + strconv.Itoa(getRandomId(c))

	// Make a GET request to the API
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var quote models.Quote
	// Print the response body
	// fmt.Println("Response Body:", string(body))
	json.Unmarshal(body, &quote)
	formateOutput(quote)
	// return quote.Quote, quote.Author
}

func getTotalNumberOfQuotes() int {
	// URL of the API you want to call
	apiURL := "http://45.79.122.143:9090/api/v1/quote/all/count"

	// Make a GET request to the API
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return 0
	}

	var count models.Count
	json.Unmarshal(body, &count)
	return count.Count
}

func getRandomId(maxRange int) int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := 0
	for randomNumber == 0 {
		randomNumber = rand.Intn(int(maxRange + 1))
	}

	return randomNumber
}

func formateOutput(quote models.Quote) {
	fmt.Println()
	dots := strings.Repeat(".", len(quote.Quote)+len(quote.Author)+3)
	fmt.Printf(" %s\n", dots)
	fmt.Printf("| %s", quote.Quote)
	padding := strings.Repeat(" ", len(quote.Author)+1)
	fmt.Printf("%s |\n", padding)
	spaces := strings.Repeat(" ", len(quote.Quote))
	fmt.Printf("| %s-%s |\n", spaces, quote.Author)
	fmt.Printf(" %s", dots)
	fmt.Printf("\n\n")
}
