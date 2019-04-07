// calculo del área de un rombo cuyo diagonal mayor es dos veces el diagonal menor
package main
import (
        "fmt"
        )
func main() {
    var diagonal_menor = 2.5 // Declarar variable longitud de la diagonal menor de un rombo se le asigna el valor 2.5
    diagonal_mayor := diagonal_menor*2      // Declarar y asignación dinamica para variable diagonal mayor  = dos veces la menor
    fmt.Println("Area del rombo de",diagonal_menor,"mts de longitud en la diagonal menor y",diagonal_mayor,"mts de longitud en la diagonal mayor, es:",diagonal_menor*diagonal_mayor/2,"mts cuadrados")
}
