package routes

import (
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/config"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/handlers"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

var _ handlers.Handler = &handler{}

type handler struct {
	cfg    *config.Config
	logger *logging.Logger
	db     *gorm.DB
}

func NewHandler(cfg *config.Config, logger *logging.Logger, db *gorm.DB) handlers.Handler {
	return &handler{
		cfg:    cfg,
		logger: logger,
		db:     db,
	}
}

func (h *handler) Routes(router *httprouter.Router) {
	router.GET("/", h.allNotes)                      // Получения всех записей
	router.POST("/addpost", h.addPost)               // Добавить пост
	router.PUT("/editentry/:id", h.editEntry)        // Редактировать запись
	router.DELETE("/deleteentry/:id", h.deleteEntry) // Удалить запись
}
