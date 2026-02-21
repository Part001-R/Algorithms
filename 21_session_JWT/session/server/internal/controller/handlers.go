package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

var wantContentType = "application/json"

// Миддлваре.
func (c *Cntr) Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Вызов обработчика
		timeStart := time.Now()
		h.ServeHTTP(w, r)
		duration := time.Since(timeStart)

		// Вывод лога запроса
		log.Printf("Принят HTTP запрос. URI:<%s>. Метод:<%s>. Продолжительность:<%v>. Адрес:<%s>", r.RequestURI, r.Method, duration, r.RemoteAddr)
	})
}

// Регистрация пользователя.
func (c *Cntr) Registration(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// Проверки.
	if http.MethodPost != r.Method {
		http.Error(w, `"Нет соответствия Content-Type"`, http.StatusForbidden)
		return
	}

	contType := r.Header.Get("Content-Type")
	if wantContentType != contType {
		http.Error(w, `"Нет соответствия Content-Type"`, http.StatusForbidden)
		return
	}

	// Чтение тела запроса.
	rxBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, `"Ошибка чтения тела запроса"`, http.StatusBadRequest)
		return
	}
	defer func() {
		_ = r.Body.Close()
	}()

	var rxData rxRegistration
	err = json.Unmarshal(rxBody, &rxData)
	if err != nil {
		http.Error(w, `"Ошибка декодирования тела запроса"`, http.StatusBadRequest)
		return
	}

	// Проверка принятых данных.
	if rxData.UserPwd != rxData.UserPwdRepeat {
		http.Error(w, `"Ошибка в пароле"`, http.StatusBadRequest)
		return
	}
	if rxData.UserPwd == "" {
		http.Error(w, `"Пароль не указан"`, http.StatusBadRequest)
		return
	}
	if rxData.UserName == "" {
		http.Error(w, `"Пользователь не указан"`, http.StatusBadRequest)
		return
	}

	// Создание токена сессии.
	hash, err := createHash(rxData.UserName, rxData.UserPwd)
	if err != nil {
		http.Error(w, `"Ошибка вычисления хэша"`, http.StatusInternalServerError)
		return
	}

	// Добавление токена сесии в InMemory.
	if _, ok := c.session[hash]; ok {
		http.Error(w, `"Пользователь уже зарегистрирован"`, http.StatusBadRequest)
		return
	}
	c.session[hash] = ""

	// Ответ.
	w.WriteHeader(http.StatusCreated)
}

// Аутентификация.
func (c *Cntr) Authentication(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// Чтение тела запроса.
	rxBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, `"Ошибка при чтении тела запроса"`, http.StatusInternalServerError)
		return
	}
	defer func() {
		_ = r.Body.Close()
	}()

	// Получение данных запроса.
	var rxData rxAuthentication
	err = json.Unmarshal(rxBody, &rxData)
	if err != nil {
		http.Error(w, `"Ошибка сериализации данных тела запроса"`, http.StatusBadRequest)
		return
	}

	// Вычисление хэша.
	hash, err := createHash(rxData.UserName, rxData.UserPwd)
	if err != nil {
		http.Error(w, `"Ошибка вычисления хэша"`, http.StatusBadRequest)
		return
	}

	_, ok := c.session[hash]
	if !ok {
		http.Error(w, `"Нет такого пользователя"`, http.StatusForbidden)
		return
	}

	w.Header().Set("authentication", hash)
	w.WriteHeader(http.StatusOK)
}

// Получение информации.
func (c *Cntr) Info(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/jsom")

	// Обработка токена сессии.
	rxSession := r.Header.Get("Authentication")

	_, ok := c.session[rxSession]
	if !ok {
		http.Error(w, `"Нет авторизации пользователя"`, http.StatusForbidden)
		return
	}

	// Формирование данных ответа.
	var data txData
	data.LocalTime = time.Now().Format("02-01-2006 15-04-05.000")

	txData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, `"Ошибка сериализации"`, http.StatusInternalServerError)
		return
	}

	// Ответ.
	w.WriteHeader(http.StatusOK)
	w.Write(txData)
}
