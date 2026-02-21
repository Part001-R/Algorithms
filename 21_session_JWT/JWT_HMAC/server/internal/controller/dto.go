package controller

// Данные регистрации пользователя.
type rxRegistration struct {
	UserName      string `json:"user_name"`       // Имя пользователя
	UserPwd       string `json:"user_pwd"`        // Пароль пользователя
	UserPwdRepeat string `json:"user_pwd_repeat"` // Подтверждение пароля
}

// Данные аутентификации пользователя.
type rxAuthentication struct {
	UserName string `json:"user_name"` // Имя пользователя
	UserPwd  string `json:"user_pwd"`  // Пароль пользователя
}

// Передаваемые данные.
type txData struct {
	LocalTime string `json:"local_time"` // Локальное время
}
