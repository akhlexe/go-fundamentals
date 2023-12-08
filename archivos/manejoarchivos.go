package archivos

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var filename string = "./archivos/txt/example.txt"

func GrabaTexto() {
	texto := "Texto a guardar"
	archivo, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error al crear el archivo: ", err)
		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTexto() {
	texto := "Texto nuevo a añadir"

	if !Append(filename, texto) {
		fmt.Println("Error al concatenar el contenido")
		return
	}
}

func Append(filename, texto string) bool {
	archivo, err := os.OpenFile(filename, os.O_RDONLY|os.O_APPEND, 0644)

	if err != nil {
		return false
	}

	_, err = archivo.WriteString(texto)

	archivo.Close()

	return err == nil
}

func LeoArchivo() {
	archivo, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error al leer el archivo ", filename)
		return
	}

	fmt.Println(string(archivo))
}

func LeoArchivoConBufio() {
	archivo, err := os.Open(filename)

	if err != nil {
		return
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("---> " + registro)
	}

	archivo.Close()
}
