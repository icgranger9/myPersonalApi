package internal

import (
	"io"
	"fmt"
	"log"
	"net/http"
)


func GetUserHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Received GetUser request")
    io.WriteString(w, "testing user api\n")
}


func FuncHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received generic request")
    response := fmt.Sprintf("The is a generic response\n")

    w.Write([]byte(response))
}