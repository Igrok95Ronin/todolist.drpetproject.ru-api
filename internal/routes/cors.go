package routes

import (
	"github.com/rs/cors"
	"net/http"
)

// Обработка cors
func CorsSettings() *cors.Cors {
	return cors.New(cors.Options{
		// Указываем разрешенные HTTP методы, только POST
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPatch,
			http.MethodDelete,
		},
		// Указываем разрешенные источники домен
		AllowedOrigins: []string{"*"}, //"http://localhost:3000",
		//"https://sites.jquery.link",
		//"https://cvetigrozny.ru",

		AllowedHeaders: []string{
			"X-Api-Password",
			"Content-Type",
			"Authorization",
		},
	})
}
