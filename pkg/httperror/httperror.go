package httperror

import (
	"encoding/json"
	"net/http"
)

// Структура для возврата ошибок
type ErrorResponse struct {
	Code     int    `json:"code"`
	CodeText string `json:"codeText"`
	Message  string `json:"message"`
	Error    string `json:"error,omitempty"` // Исходная ошибка в виде строки, исключаем, если пустое
}

// Функция для возврата ошибок в формате JSON
func WriteJSONError(w http.ResponseWriter, message string, err error, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	errorResponse := ErrorResponse{
		Code:     code,
		CodeText: http.StatusText(code),
		Message:  message,
	}

	if err != nil {
		errorResponse.Error = err.Error() // возвращаем текст ошибки
	}

	json.NewEncoder(w).Encode(errorResponse)
}
