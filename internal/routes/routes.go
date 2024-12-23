package routes

import (
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/handlers"
	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{}

type handler struct{}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Routes(router *httprouter.Router) {
	router.GET("/getallrecords", h.getAllRecords) // Получения всех записей
}
