package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	server := http.ListenAndServe(":3675", router)
	log.Fatal(server)
	fmt.Println("el servidor esta corriendo en localhosto:3675")

}


