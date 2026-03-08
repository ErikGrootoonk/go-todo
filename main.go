package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/erikdev/go-todo/config"
	"github.com/erikdev/go-todo/database"
	"github.com/erikdev/go-todo/handlers"
	"github.com/erikdev/go-todo/store"
)

func main() {
	cfg := config.Load()

	db := database.Open(cfg.DBPath)
	defer db.Close()

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	todoStore := store.NewTodoStore(db)
	todoHandler := handlers.NewTodoHandler(todoStore, tmpl)

	http.HandleFunc("/", todoHandler.HandleIndex)
	http.HandleFunc("/add", todoHandler.HandleAdd)
	http.HandleFunc("/toggle", todoHandler.HandleToggle)
	http.HandleFunc("/delete", todoHandler.HandleDelete)

	fmt.Printf("Listening on http://localhost:%s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
