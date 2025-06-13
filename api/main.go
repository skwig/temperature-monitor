package main

import (
	"fmt"
	"log"
	"net/http"
	"temperaturemonitor/api/endpoints"
	"time"

	"github.com/gin-gonic/gin"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	e := endpoints.Endpoints{}

	r := gin.Default()

	endpoints.RegisterHandlers(r, e)

	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}
