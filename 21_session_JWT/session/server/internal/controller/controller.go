package controller

// Представление контроллера.
type Cntr struct {
	session map[string]string // Токены сессий
	Addr    string
}

// Конструктор.
func New(addr string) *Cntr {

	return &Cntr{
		session: make(map[string]string),
		Addr:    addr,
	}
}
