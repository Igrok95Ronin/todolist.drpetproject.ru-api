package routes

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Получения всех записей
func (h *handler) allNotes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var allNotes []AllNotes
	if err := h.db.Find(&allNotes).Error; err != nil {
		fmt.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(allNotes); err != nil {
		fmt.Println(err)
	}
}
