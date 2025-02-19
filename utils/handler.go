package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
)

// handleRequest обрабатывает входящий запрос
func HandleRequest(w http.ResponseWriter, r *http.Request) (map[string]string, int, error) {
	// Проверяем, есть ли тело запроса
	if r.Body == nil {
		return nil, http.StatusBadRequest, http.ErrBodyNotAllowed
	}

	// Читаем тело запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	defer r.Body.Close()

	// Проверяем, пустое ли тело
	if len(body) == 0 {
		return nil, http.StatusBadRequest, http.ErrBodyNotAllowed
	}

	// Пытаемся декодировать JSON
	var requestData map[string]string
	if err := json.Unmarshal(body, &requestData); err != nil {
		return nil, http.StatusBadRequest, err
	}

	// Проверяем наличие ключа "number"
	number, exists := requestData["number"]
	if !exists {
		return nil, http.StatusBadRequest, http.ErrMissingFile
	}

	// Проверяем, что значение состоит из 16 цифр
	validNumber := regexp.MustCompile(`^\d{16}$`).MatchString
	if !validNumber(number) {
		return nil, http.StatusBadRequest, http.ErrNotSupported
	}

	// Проверяем число по алгоритму Луна
	isValid := CardCheck(number)

	validResponse := map[string]string{
		"message": "Number is valid",
		"number":  number,
	}
	invalidResponse := map[string]string{
		"message": "Number is invalid",
		"number":  number,
	}

	if !isValid {
		// Возвращаем успешный ответ с указанием неверного числа
		return invalidResponse, http.StatusOK, nil
	} else {
		// Возвращаем успешный ответ с извлеченным числом
		return validResponse, http.StatusOK, nil
	}

}
