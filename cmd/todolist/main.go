package main

import (
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/config"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/routes"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func main() {

	db := routes.InitDB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database handle: %v", err)
	}
	defer sqlDB.Close()

	router := httprouter.New()

	logger := logging.GetLogger()

	cfg := config.GetConfig() // Читаем конфигурацию приложения

	handler := routes.NewHandler(cfg, logger, db)
	handler.Routes(router)

	// Обработка cors
	corsHandler := routes.CorsSettings().Handler(router)

	start(corsHandler, cfg, logger)
}

// Функция start запускает приложение
func start(router http.Handler, cfg *config.Config, logger *logging.Logger) {

	const wri time.Duration = 15 * time.Second
	server := &http.Server{
		Addr:         cfg.Port,
		Handler:      router,
		WriteTimeout: wri,
		ReadTimeout:  wri,
		IdleTimeout:  wri,
	}

	logger.Infof("Server started %v", cfg.Port)
	logger.Fatal(server.ListenAndServe())

}
