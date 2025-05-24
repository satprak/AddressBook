package main

import (
	"addressbook/handlers"
	"addressbook/service"
	"addressbook/storage"
	"net/http"
)

func main() {
	store := storage.NewInMemoryStorage()
	svc := service.NewContactService(store)
	handler := handlers.NewContactHandler(svc)

	http.HandleFunc("/create", handler.Create)
	http.HandleFunc("/update", handler.Update)
	http.HandleFunc("/delete", handler.Delete)
	http.HandleFunc("/search", handler.Search)

	http.ListenAndServe(":5000", nil)
}
