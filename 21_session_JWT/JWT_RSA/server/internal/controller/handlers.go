package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
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
	if _, ok := c.regHash[hash]; ok {
		http.Error(w, `"Пользователь уже зарегистрирован"`, http.StatusBadRequest)
		return
	}
	c.regHash[hash] = ""

	// Ответ.
	w.WriteHeader(http.StatusCreated)
}

// Аутентификация.
func (c *Cntr) Authentication(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// Чтение тела запроса.
	rxBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Ошибка чтения тела запроса: <%v>", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer func() {
		_ = r.Body.Close()
	}()

	// Получение данных запроса.
	var rxData rxAuthentication
	err = json.Unmarshal(rxBody, &rxData)
	if err != nil {
		log.Printf("Ошибка сериализации: <%v>", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Проверка хэша регистрации.
	hash, err := createHash(rxData.UserName, rxData.UserPwd)
	if err != nil {
		log.Printf("Ошибка вычисления хэша: <%v>", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	_, ok := c.regHash[hash]
	if !ok {
		log.Printf("Нет такого пользователя: <%v>", err)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	// Создание JWT.
	srcAddr := r.RemoteAddr
	rawJWT := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"forAddr": srcAddr,                                // Для кого выставляется
		"exp":     time.Now().Add(2 * time.Minute).Unix(), // Время валидности
	})

	signJWT, err := rawJWT.SignedString(c.privKey)
	if err != nil {
		log.Printf("Ошибка подписи JWT: <%v>", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authentication", signJWT)
	w.WriteHeader(http.StatusOK)
}

// Получение информации.
func (c *Cntr) Info(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/jsom")

	rxJWT := r.Header.Get("Authentication")

	// Проверка JWT.
	jwtToken, err := jwt.Parse(rxJWT, func(t *jwt.Token) (interface{}, error) {
		// Проверка алгоритма подписи (только RS256)
		if t.Method.Alg() != "RS256" {
			log.Printf("Запрещённый алгоритм подписи: %s", t.Method.Alg())
			return nil, jwt.ErrInvalidKey
		}
		// Проверка типа метода
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			log.Printf("Некорректный тип метода подписи: %T", t.Method)
			return nil, jwt.ErrInvalidKey
		}
		return c.pubKey, nil
	})
	if err != nil {
		log.Printf("Ошибка парсинга JWT: <%v>\n", err)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if !jwtToken.Valid {
		log.Printf("Токен не прошел валидацию: <%v>\n", err)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	// Проверка содержимого JWT
	clMap, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		log.Printf("Ошибка получения данных JWT: <%v>\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	remAddr := r.RemoteAddr
	jwtAddr := clMap["forAddr"]
	if jwtAddr != remAddr {
		log.Printf("Принят запрос с другого адреса. Ожидался <%s>, а принят <%s>\n", jwtAddr, remAddr)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	// Формирование данных ответа.
	var data txData
	data.LocalTime = time.Now().Format("02-01-2006 15-04-05.000")

	txData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Ошибка сериализации: <%v>", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Ответ.
	w.WriteHeader(http.StatusOK)
	w.Write(txData)
}
