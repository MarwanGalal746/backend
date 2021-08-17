package handlers

import (
	"github.com/tamiat/backend/domain/content"
	"github.com/tamiat/backend/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()
	ch:= ContentHandlers{service.NewContentService(content.NewContentRepositoryDb())}
	router.HandleFunc("/contents",ch.getAllContents).Methods(http.MethodGet)
	router.HandleFunc("/contents/{content_id:[0-9]+}",ch.getContent).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
