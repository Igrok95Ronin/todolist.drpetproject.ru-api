package main

import (
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/config"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/routes"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func main() {
	router := httprouter.New()

	logger := logging.GetLogger()

	cfg := config.GetConfig() // Читаем конфигурацию приложения

	handler := routes.NewHandler(cfg, logger)
	handler.Routes(router)

	start(router, cfg, logger)
}

// Функция start запускает приложение
func start(router *httprouter.Router, cfg *config.Config, logger *logging.Logger) {

	const wri time.Duration = 15 * time.Second
	server := &http.Server{
		Addr:         cfg.Port,
		Handler:      router,
		WriteTimeout: wri,
		ReadTimeout:  wri,
		IdleTimeout:  wri,
	}

	logger.Info("Server started ...")
	logger.Fatal(server.ListenAndServe())

}
