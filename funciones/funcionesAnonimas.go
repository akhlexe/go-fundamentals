package funciones

import "fmt"

func Calculos() {

	var num3 int = 32
	var num4 int = 243

	suma := func(num1 int, num2 int) int {
		return num1 + num2 + num3 + num4
	}

	fmt.Println(suma(10, 25))
}
