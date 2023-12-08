package arreglosslice

import "fmt"

var tablaSlice []int = []int{2, 5, 6}

func MuestroSlice() {
	fmt.Println(tablaSlice)

	tablaSlice = append(tablaSlice, 22)

	fmt.Println(tablaSlice)

	
}
