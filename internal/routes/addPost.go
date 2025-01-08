package routes

import (
	"encoding/json"
	"fmt"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/pkg/httperror"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Note struct {
	Note string `json:"note"`
}

// Добавить пост
func (h *handler) addPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var note Note

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		// Если произошла ошибка декодирования, возвращаем клиенту ошибку с кодом 400
		httperror.WriteJSONError(w, "Ошибка декодирования в JSON", err, http.StatusBadRequest)
		// Логируем ошибку
		h.logger.Errorf("Ошибка декодирования в JSON: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)

	if err := addPostDB(h, w, note.Note); err != nil {
		h.logger.Errorf("Ошибка при добавления записи в БД: %s", err)
	}

}

func addPostDB(h *handler, w http.ResponseWriter, note string) error {

	allNotes := AllNotes{
		Note: note,
	}

	if err := h.db.Create(&allNotes).Error; err != nil {
		httperror.WriteJSONError(w, "Ошибка при добавления записи в БД", fmt.Errorf(""), http.StatusInternalServerError)
		return err
	}

	return nil
}
