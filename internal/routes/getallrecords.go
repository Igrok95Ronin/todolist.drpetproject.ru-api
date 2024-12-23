package routes

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Получения всех записей
func (h *handler) getAllRecords(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "getAllRecords")
	fmt.Fprintf(w, "%s", h.cfg.Port)

}
