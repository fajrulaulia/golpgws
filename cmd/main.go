package main

import (
	"os"

	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	c "github.com/fajrulaulia/golpgws/services"
)

func main() {

	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("JWT_SECRET value is missing")
		os.Exit(1)
	}

	router := mux.NewRouter()
	database := c.Driver()

	c.Register(router, database)

	server(router)
}

func server(router *mux.Router) {
	if os.Getenv("SERVER_PORT") == "" {
		log.Fatal("SERVER_PORT value is missing")
		os.Exit(1)
	}

	log.Println("application running on port", os.Getenv("SERVER_PORT"))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Cache-Control"}),
		handlers.ExposedHeaders([]string{"Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}))(router),
	))
}
