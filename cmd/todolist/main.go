package main

import (
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/routes"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func main() {
	router := httprouter.New()

	handler := routes.NewHandler()
	handler.Routes(router)

	start(router)
}

// Функция start запускает приложение
func start(router *httprouter.Router) {

	const wri time.Duration = 15 * time.Second
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		WriteTimeout: wri,
		ReadTimeout:  wri,
		IdleTimeout:  wri,
	}

	log.Println("Server started ...")
	log.Fatal(server.ListenAndServe())

}
