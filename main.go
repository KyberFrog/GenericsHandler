package main

import (
	"genericsTest/TestHandler"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/test", TestHandler.SharedWrapper)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal("ListenError")
	}
}
