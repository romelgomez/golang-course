// calculo del área de un triángulo
package main
import (
        "fmt"
        )

func main() {
    var base float64 = 3.5 // Declarar variable base como float64 y se le asigna el valor 3.5
    var altura = 3.8      // Declarar variable altura como float64 y se le asigna valor 3.8
    fmt.Println("\n\nArea del triángulo de base",base,"mts y altura",altura,"mts es:",base*altura/2,"mts cuadrados")
}
