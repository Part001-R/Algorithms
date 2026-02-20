package example

import "errors"

var (
	// Превышен номер строки.
	ErrOverNumbStr = errors.New("Превышен номер строки")
)
