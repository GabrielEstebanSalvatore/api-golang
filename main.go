package main

import (
	"log"
	"net/http"

	"github.com/GabrielEstebanSalvatore/api-golang/routes"
	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json.ID`
	Name    string `json.Name`
	Content string `json.Content`
}

type allTask []task

var tasks = allTask{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Content One",
	},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler)
	log.Fatal(http.ListenAndServe(":3000", router))
}
