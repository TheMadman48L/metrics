package main

import (
	"net/http"

	"github.com/TheMadman48L/metrics/internal/handler/update"
	"github.com/TheMadman48L/metrics/internal/repository/memory"
)

func main() {
	memStorage := memory.New()

	updateHandler := update.New(memStorage)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /update/{type}/{name}/{value}", updateHandler.Handle)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
