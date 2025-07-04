package main

import (
	"errors"
	"fmt"
)

type CustomErr1 struct {
	Code    int
	Message string
}

func main() {

	value, err := div(10, 0)
	if err != nil {
		fmt.Printf("Сообщение ошибки:{%s}\n", err)

		var rxErr CustomErr1
		if errors.As(err, &rxErr) {
			fmt.Printf("Код ошибки:{%d}\n", rxErr.Code)
			fmt.Printf("Сообщение ошибки:{%s}\n", rxErr.Message)
		}
	} else {
		fmt.Printf("результат деления:{%d}\n", value)
	}

	fmt.Println("===") //=====

	value, err = div(10, 2)
	if err != nil {
		fmt.Printf("Сообщение ошибки:{%s}\n", err)

		var rxErr CustomErr1
		if errors.As(err, &rxErr) {
			fmt.Printf("Код ошибки:{%d}\n", rxErr.Code)
			fmt.Printf("Сообщение ошибки:{%s}\n", rxErr.Message)
		}
	} else {
		fmt.Printf("результат деления:{%d}\n", value)
	}

}

func (ce CustomErr1) Error() string {
	return fmt.Sprintf("fault: code{%d} message{%s}", ce.Code, ce.Message)
}

func div(val1, val2 int) (int, error) {

	if val2 == 0 {
		return 0, CustomErr1{
			Code:    0,
			Message: "делитель равен нулю",
		}
	}
	return val1 / val2, nil
}
