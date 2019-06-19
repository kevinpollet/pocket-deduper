package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const redirectURI = "http://localhost:8000"

var consumerKey = os.Getenv("CONSUMER_KEY")

func main() {
	requestBody, _ := json.Marshal(map[string]string{
		"consumer_key": consumerKey,
		"redirect_uri": redirectURI,
	})
	client := http.Client{}
	req, _ := http.NewRequest("POST", "https://getpocket.com/v3/oauth/request", bytes.NewBuffer(requestBody))
	req.Header.Add("Content-Type", "application/json; charset=UTF8")
	req.Header.Add("X-Accept", "application/json")

	res, _ := client.Do(req)

	var responseBody map[string]interface{}
	json.NewDecoder(res.Body).Decode(&responseBody)

	fmt.Printf("https://getpocket.com/auth/authorize?request_token=%s&redirect_uri=%s\n", responseBody["code"], redirectURI)

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestBody2, _ := json.Marshal(map[string]string{
			"consumer_key": consumerKey,
			"code":         responseBody["code"].(string),
		})

		req2, _ := http.NewRequest("POST", "https://getpocket.com/v3/oauth/authorize", bytes.NewBuffer(requestBody2))
		req2.Header.Add("Content-Type", "application/json; charset=UTF8")
		req2.Header.Add("X-Accept", "application/json")
		res2, _ := client.Do(req2)
		var responseBody2 map[string]interface{}
		json.NewDecoder(res2.Body).Decode(&responseBody2)
		fmt.Println(responseBody2)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
