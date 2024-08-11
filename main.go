package main

import (
	"log"
	"net/http"
	"carbon/routers"
)

func main() {
	router := routers.InitRoutes()
	log.Fatal(http.ListenAndServe(":8001", router))
}

