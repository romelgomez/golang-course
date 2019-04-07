// calculo del área de un círculo
package main
import (
        "fmt"
        "math"
        )
func main() {

    var radio float64 = 5 // Declarar variable radio del círculo, se le asigna el valor 5
    const Pi float64 = 3.14159265358979323846  // constante Pi
    fmt.Println("Area del círculo de radio",radio,"mts es:",Pi*radio*radio,"mts cuadrados")
    fmt.Println("Area del círculo de radio",radio,"mts es:",math.Pi*radio*radio,"mts cuadrados")
    fmt.Printf("Area del círculo de radio %2f mts es %6.2f mts cuadrados \n",radio,math.Pi*radio*radio,)
    fmt.Printf("Area del círculo de radio %2.2f mts es %6.2f mts cuadrados",radio,math.Pi*radio*radio,)
}
