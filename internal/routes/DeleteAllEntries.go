package routes

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Удалить все записи
func (h *handler) DeleteAllEntries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := h.db.Unscoped().Where("1 = 1").Delete(&AllNotes{}).Error; err != nil {
		h.logger.Errorf("Ошибка при удалении всех записей: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
