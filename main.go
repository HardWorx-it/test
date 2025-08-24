package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// ContactForm представляет структуру данных контактной формы
type ContactForm struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

// Response представляет структуру ответа API
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func main() {
	// Настройка обработчиков маршрутов
	http.HandleFunc("/", serveStaticFiles)
	http.HandleFunc("/api/contact", handleContactForm)

	// Настройка порта
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Запуск сервера
	fmt.Printf("Сервер запущен на порту %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// serveStaticFiles обслуживает статические файлы
func serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	// Установка заголовков CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Обслуживание index.html для всех маршрутов
	http.ServeFile(w, r, "index.html")
}

// handleContactForm обрабатывает отправку контактной формы
func handleContactForm(w http.ResponseWriter, r *http.Request) {
	// Установка заголовков CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	// Обработка preflight запросов
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Разрешить только POST запросы
	if r.Method != "POST" {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Декодирование JSON из тела запроса
	var formData ContactForm
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&formData)
	if err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	// Валидация данных
	if formData.Name == "" || formData.Email == "" || formData.Message == "" {
		response := Response{
			Success: false,
			Message: "Все поля обязательны для заполнения",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Здесь можно добавить логику для сохранения данных в базу данных
	// или отправки email уведомления

	// Логирование полученных данных
	log.Printf("Получена новая заявка: %s, %s, %s",
		formData.Name, formData.Email, formData.Message)

	// Формирование ответа
	response := Response{
		Success: true,
		Message: "Сообщение успешно отправлено!",
	}

	// Имитация обработки запроса
	time.Sleep(1 * time.Second)

	// Отправка ответа
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
