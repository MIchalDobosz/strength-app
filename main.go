package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"
	"strength-app/database"
	"strength-app/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No args provided")
		os.Exit(1)
	}

	allowedArgs := []string{
		"migrate",
		"seed",
		"serve",
	}
	if !slices.Contains(allowedArgs, args[0]) {
		fmt.Printf("Argument %s is not supported\n", args[0])
		os.Exit(1)
	}

	//TODO: get connection params form env
	db, err := database.GetDB("127.0.0.1", 3306, "strength-app", "root", "root")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch args[0] {
	case "migrate":
		if err = database.Migrate(db); err == nil {
			fmt.Println("Successfully migrated")
		}
	case "seed":
		if err = database.Seed(db); err == nil {
			fmt.Println("Successfully seeded")
		}
	case "serve":
		port := "8080"
		if len(args) > 1 {
			port = args[1]
		}
		err = serve(":"+port, db)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func serve(addr string, db *sqlx.DB) error {
	router := mux.NewRouter()

	for _, route := range Routes() {
		router.HandleFunc(route.Path, func(response http.ResponseWriter, request *http.Request) {
			handlerResponse, err := route.Handler(request, mux.Vars(request), db)
			if err != nil {
				log.Println(err)
				writeResponse(response, handlers.Response{}.Failure(500, "Server error."))
				return
			}

			writeResponse(response, handlerResponse)
		}).Methods(route.Method)
	}

	router.NotFoundHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		writeResponse(response, handlers.Response{}.NotFound("Page not found."))
	})

	if err := http.ListenAndServe(addr, router); err != nil {
		return fmt.Errorf("failed to listen on address %s: %v", addr, err)
	}

	return nil
}

func writeResponse(response http.ResponseWriter, handlerResponse handlers.Response) {
	response.WriteHeader(handlerResponse.Code)
	response.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(handlerResponse.Body); err != nil {
		log.Println(err)
	}
}
