package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	a := initApp()

	fmt.Println("Server started")

	log.Fatal(http.ListenAndServe(":8000", a.router))
}
