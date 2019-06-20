package main

import (
	"fmt"
	"kevinpollet/pocket-remove-duplicates/internal/pkg/pocketclient"
	"log"
	"net/http"
	"os"
)

func main() {
	pocketClient := pocketclient.PocketClient{
		ConsumerKey: os.Getenv("CONSUMER_KEY"),
	}

	code, _ := pocketClient.GetRequestToken()

	fmt.Printf(pocketClient.GetAuthorizeURL(code))

	router := http.NewServeMux()
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}

	router.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		writer.WriteHeader(200)
		writer.Write([]byte("Hello"))
		fmt.Println(pocketClient.GetAccessToken(code))
	})

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
