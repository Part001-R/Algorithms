package controller

// Представление контроллера.
type Cntr struct {
	regHash map[string]string // Хэши регистрации
	Addr    string            // Адрес сервера
	secrKey []byte            // Секретный ключ
}

// Конструктор.
func New(addr string) *Cntr {

	return &Cntr{
		regHash: make(map[string]string),
		Addr:    addr,
		secrKey: []byte("ABCD123!"),
	}
}
