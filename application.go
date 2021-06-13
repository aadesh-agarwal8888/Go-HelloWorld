package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", SayHello)

	log.Println("Starting Server on 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))

}

func SayHello(res http.ResponseWriter, req *http.Request) {
	log.Println("Saying Hello")

	res.Write([]byte("Heelo World"))
}
