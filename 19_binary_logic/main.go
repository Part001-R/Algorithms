package main

import "fmt"

func main() {

	//
	// Побитовое
	//

	// Инициализация в бинарном виде
	fmt.Println(" --- Инициализация в бинарном виде")
	var binVar = 0b00000000_00000011 // 3
	fmt.Printf("Значение в десятичном представлении: <%d>\n", binVar)
	fmt.Printf("Значение в бинарном представлении: <%b>\n", binVar)
	fmt.Println()

	// OR
	fmt.Println(" --- OR")
	var varA = 0b00000000_00000010 // 2
	var varB = 0b00000000_00000001 // 1
	varC := varA | varB
	fmt.Printf("Результат OR в десятичном представлении: <%d>\n", varC)
	fmt.Printf("Результат OR в бинарном представлении: <%b>\n", varC)
	fmt.Println()

	// AND
	fmt.Println(" --- AND")
	varA = 0b00000000_00000010 // 2
	varB = 0b00000000_00000011 // 3
	varC = varA & varB
	fmt.Printf("Результат AND в десятичном представлении: <%d>\n", varC)
	fmt.Printf("Результат AND в бинарном представлении: <%b>\n", varC)
	fmt.Println()

	// XOR
	fmt.Println(" --- XOR")
	varA = 0b00000000_00000010 // 2
	varB = 0b00000000_00000011 // 3
	varC = varA ^ varB
	fmt.Printf("Результат XOR в десятичном представлении: <%d>\n", varC)
	fmt.Printf("Результат XOR в бинарном представлении: <%b>\n", varC)
	fmt.Println()

	// NOT побитовое
	fmt.Println(" --- NOT")
	varA = 0b00000000_00000010 // 2
	varC = ^varA
	fmt.Printf("Исходное значение: <%d>\n", varA)
	fmt.Printf("Результат NOT в десятичном представлении. Дополнение до одного.: <%d>\n", varC)
	fmt.Printf("Результат NOT в бинарном представлении. Дополнение до одного.: <%b>\n", varC)
	varC += 1
	fmt.Printf("Результат NOT в десятичном представлении. Дополнение до двух.: <%d>\n", varC)
	fmt.Printf("Результат NOT в бинарном представлении. Дополнение до двух.: <%b>\n", varC)
	fmt.Println()

	// SHL
	fmt.Println(" --- SHL")
	varA = 0b00000000_00000010 // 2
	varC = varA << 1
	fmt.Printf("Результат SHL в десятичном представлении: <%d>\n", varC)
	fmt.Printf("Результат SHL в бинарном представлении: <%b>\n", varC)
	fmt.Println()

	// SHR
	fmt.Println(" --- SHR")
	varA = 0b00000000_00000010 // 2
	varC = varA >> 1
	fmt.Printf("Результат SHR в десятичном представлении: <%d>\n", varC)
	fmt.Printf("Результат SHR в бинарном представлении: <%b>\n", varC)
	fmt.Println()

	fmt.Println(" ---------------------- ")
	fmt.Println("")

	//
	// Логическое
	//

	// NOT
	fmt.Println(" --- XOR")
	var varAA bool = false
	varCC := !varAA
	fmt.Printf("Результат логического NOT: <%t>\n", varCC)
	fmt.Println()

	// AND
	fmt.Println(" --- AND")
	varAA = false
	var varBB bool = false
	varCC = varAA && varBB
	fmt.Printf("Результат логического AND: <%t>\n", varCC)
	fmt.Println()

	// OR
	fmt.Println(" --- OR")
	varAA = false
	varBB = true
	varCC = varAA || varBB
	fmt.Printf("Результат логического OR: <%t>\n", varCC)
	fmt.Println()

	// AND-NOT
	fmt.Println(" --- AND-NOT")
	varAA = false
	varBB = true
	varCC = !(varAA && varBB)
	fmt.Printf("Результат логического AND-NOT: <%t>\n", varCC)
	fmt.Println()

	// OR-NOT
	fmt.Println(" --- OR-NOT")
	varAA = false
	varBB = true
	varCC = !(varAA || varBB)
	fmt.Printf("Результат логического OR-NOT: <%t>\n", varCC)
	fmt.Println()
}
