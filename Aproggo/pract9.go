// programa para  convertir entero a cadena y viceversa
package main

import (
    "fmt"
    "strconv"  // De este paquete utilizaremos Itoa para convertir número entero a su cadena de caracteres(string) y la función Atoi lo contrario, de cadena numérica a número.
)

func main() {
	var vvalor int
    vvalor = 120
    fmt.Println("El valor de la variable entera vvalor es:",vvalor)
    // Use Itoa to convert int to string.
    vresult := strconv.Itoa(vvalor)
    fmt.Println("\n")
    fmt.Println("vvalor convertida a cadena(utilizando la función strconv.Itoa) y asignada a la variable result:",vresult)
    fmt.Println("El valor de verdad de la proposición de que el valor de vresult:",vresult,"es igual a cadena \"120\" es:",vresult=="120")
    fmt.Println("El valor de vresult + la cadena \"naranjas\" es:",vresult+" naranjas")
    fmt.Println("\n")
    // Use Atoi to convert string to int.
    voriginal, _:= strconv.Atoi(vresult)
    fmt.Println("vresult convertida a entero(utilizando la función strconv.Atoi) y asignada a la variable voriginal:",voriginal)
    fmt.Println("El valor de verdad de la proposición de que el valor de voriginal:",voriginal,"es igual al entero 120 es:",voriginal==120)
    fmt.Println("El valor de vriginal multiplicado por ",2,"es: ",voriginal*2)

}

