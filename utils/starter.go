package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// startServer запускает HTTP-сервер
func StartServer() {
	// Создаем новый роутер Chi
	r := chi.NewRouter()

	// Обработчик для GET-запросов
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// Обрабатываем запрос
		response, statusCode, err := HandleRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), statusCode)
			return
		}

		// Кодируем ответ в JSON
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
			return
		}

		// Отправляем ответ
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write(jsonResponse)
	})

	// Запускаем сервер на порту 8080
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
