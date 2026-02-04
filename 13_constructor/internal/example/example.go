package example

import "sync"

var initOnce sync.Once

// Геттеры значений экземпляра.
type instHndl struct {
	GetValA func() int
	GetValB func() int
}

// Состояние экземпляра.
type instData struct {
	varA int
	varB int
}

var inst instData
var instGet *instHndl

// Конструктор.
func New(a, b int) *instHndl {
	initOnce.Do(func() {
		inst.varA = a
		inst.varB = b

		instGet = &instHndl{
			GetValA: getValA,
			GetValB: getValB,
		}
	})
	return instGet
}

// Геттер.
func getValA() int {
	return inst.varA
}

// Геттер.
func getValB() int {
	return inst.varB
}
