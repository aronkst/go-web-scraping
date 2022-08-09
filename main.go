package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aronkst/go-web-scraping/web"
	_ "github.com/joho/godotenv/autoload"
	"github.com/julienschmidt/httprouter"
)

func main() {
	port := getEnv("PORT")

	router := httprouter.New()

	router.POST("/html", web.HandlerHTML)
	router.POST("/find", web.HandlerFind)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}

func getEnv(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		err := fmt.Errorf("invalid %s environment variable", name)
		log.Fatal(err)
	}

	return value
}
