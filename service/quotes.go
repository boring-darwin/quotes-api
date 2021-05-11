package service

import(
	"encoding/json"
)


type Data struct {
	Quote string
	Author string
}


func Welcome() string{
	return "Welcome to quotes API"
}

func GetJsonResponse()([]byte, error){

	quote := "You miss 100% of shot you don't take"
	author := "Unkonwn"

	d := Data{quote, author}

	return json.MarshalIndent(d, "", "")

}