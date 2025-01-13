package routes

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Удалить все отмеченные записи
func (h *handler) DeleteAllMarkedEntries(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	if err := h.db.Unscoped().Where("completed = ?", true).Delete(&AllNotes{}).Error; err != nil {
		h.logger.Errorf("Ошибка при удалении всех отмеченных записей: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
