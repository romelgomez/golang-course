// calculo del área de un trapecio
package main
import (
        "fmt"
        
        )

func main() {

    var lado_mayor int = 7 // Declarar variable longitud del lado paralelo mayor se le asigna el valor 7
    var lado_menor int = 8 // Declarar variable longitud del lado paralelo mmenor se le asigna el valor 8
    var altura int = 5   // Declarar variable altura del trapecio se le asigna valor de 5
    fmt.Println("Area del trapecio de longitudes, lado paralelo mayor:",lado_mayor,"mts, lado paralelo menor:",lado_menor,"mts, y altura:",altura,"mts, es:",(lado_mayor+lado_menor)*altura/2,"mts cuadrados")
    fmt.Println("Area real: ",(float32(lado_mayor)+float32(lado_menor))*float32(altura)/2,"mts cuadrados")

    fmt.Println("\n")
    fmt.Println("\n")
    
    // calculo del área del paralelogramo o romboide cuya altura es 5 unidades mas que su base
    
    
    var base float32 // Declarar variable base como float64 y se le asigna el valor 6.5
    base=6.5
    var altura2 float32      // Declarar variable altura como float32 y se le asigna valor base + 5
    altura2=base+5.0
    fmt.Println("Area del romboide cuya longitud de la base",base,"mts y altura",altura2,"mts es:",base*altura2,"mts cuadrados")



}
