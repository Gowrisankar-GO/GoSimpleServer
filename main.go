package main

import (
	"log"
	"net/http"
)

var (
	port = "8080"
)

func WelcomeGreet(w http.ResponseWriter, r *http.Request) {

	message := "Hello world, Its me `Gowrisankar` ~ Golang Developer\n"

	byteSlice := []byte(message)

	w.Write(byteSlice)

}

func main() {

	server := http.NewServeMux()

	server.HandleFunc("/", WelcomeGreet)

	log.Printf("server started in http://localhost:%s", port)

	log.Fatal(http.ListenAndServe(":"+port, server))
}
