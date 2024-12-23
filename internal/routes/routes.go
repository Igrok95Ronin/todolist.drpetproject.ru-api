package routes

import (
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/config"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/handlers"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{}

type handler struct {
	cfg    *config.Config
	logger *logging.Logger
}

func NewHandler(cfg *config.Config, logger *logging.Logger) handlers.Handler {
	return &handler{
		cfg:    cfg,
		logger: logger,
	}
}

func (h *handler) Routes(router *httprouter.Router) {
	router.GET("/getallrecords", h.getAllRecords) // Получения всех записей
}
