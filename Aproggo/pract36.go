// Programa con ingreso de datos utilizando Scanln de paquete "ftm" para el cálculo del área de rectangulos dados largo y el ancho
package main
import (
        "fmt"
        )
func main() {
    fmt.Println("Cálculo del área de un rectángulo\n")
    var largo float64 //Declarar variable y tipo antes de escanear, esto es obligatorio
    fmt.Print("Ingresa valor del largo del rectángulo: ")
	fmt.Scanln(&largo)
	var ancho float64 //Declarar variable y tipo antes de escanear, esto es obligatorio
    fmt.Print("Ingresa valor del ancho del rectángulo:  ")
	fmt.Scanln(&ancho)
    fmt.Printf("El área del rectángulo de largo %.2f y ancho %.2f es: %.2f\n",largo,ancho,largo*ancho)
   }
