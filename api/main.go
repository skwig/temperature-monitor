package main

import (
	"log"
	"net/http"
	"temperaturemonitor/api/endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	e := endpoints.Endpoints{}

	r := gin.Default()

	endpoints.RegisterHandlers(r, &e)

	s := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	log.Fatal(s.ListenAndServe())
}
