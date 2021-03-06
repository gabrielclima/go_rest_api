package main

import (
	"github.com/gabrielclima/go_rest_api/controllers"
	_ "github.com/gabrielclima/go_rest_api/database"
	"log"
	"net/http"
	"os"
	"runtime"
)

func main() {
	defer os.Exit(0)
	log.Fatal(http.ListenAndServe(":8080", controllers.Handlers()))
	runtime.Goexit()
}
