package arreglosslice

import "fmt"

var matriz [4][10]int

func MuestroArreglos() {
	matriz[2][6] = 24
	matriz[3][8] = 32

	fmt.Println(matriz)
}
