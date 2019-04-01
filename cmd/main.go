package main

import (
	"log"
	"net/http"

	"github.com/fredeastside/luasgo/pkg/handler"
)

func main() {
	handler := handler.NewHandler()
	log.Fatal(http.ListenAndServe(":8080", handler))
}
