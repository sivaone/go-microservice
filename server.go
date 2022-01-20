package main

import (
	"go-microservice/config"
	"go-microservice/ent"
	"go-microservice/router"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const defaultPort = ":8080"

func main() {

	client, err := config.NewEntClient()
	if err != nil {
		log.Printf("err initializing db client: %v", err)
	}

	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Printf("err closing db client: %v", err)
		}
	}(client)

	config.SetClient(client)

	r := gin.Default()
	router.RegisterRouter(r)
	_ = http.ListenAndServe(defaultPort, r)
}
