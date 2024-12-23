package routes

import (
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/config"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/handlers"
	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{}

type handler struct {
	cfg *config.Config
}

func NewHandler(cfg *config.Config) handlers.Handler {
	return &handler{
		cfg: cfg,
	}
}

func (h *handler) Routes(router *httprouter.Router) {
	router.GET("/getallrecords", h.getAllRecords) // Получения всех записей
}
