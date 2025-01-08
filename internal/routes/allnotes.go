package routes

import (
	"encoding/json"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/pkg/httperror"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Получения всех записей
func (h *handler) allNotes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var allNotes []AllNotes
	if err := h.db.Find(&allNotes).Error; err != nil {
		h.logger.Error(err)
		httperror.WriteJSONError(w, "Ошибка при получения данных из бд", err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(allNotes); err != nil {
		h.logger.Error(err)
		httperror.WriteJSONError(w, "Ошибка при отправке данных клиенту", err, http.StatusInternalServerError)
	}

}
