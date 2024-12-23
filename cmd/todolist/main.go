package main

import (
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/config"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/routes"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func main() {
	router := httprouter.New()

	cfg := config.GetConfig() // Читаем конфигурацию приложения

	handler := routes.NewHandler(cfg)
	handler.Routes(router)

	start(router, cfg)
}

// Функция start запускает приложение
func start(router *httprouter.Router, cfg *config.Config) {

	const wri time.Duration = 15 * time.Second
	server := &http.Server{
		Addr:         cfg.Port,
		Handler:      router,
		WriteTimeout: wri,
		ReadTimeout:  wri,
		IdleTimeout:  wri,
	}

	log.Println("Server started ...")
	log.Fatal(server.ListenAndServe())

}
