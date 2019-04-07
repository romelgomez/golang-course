package main
import "fmt"
func main() {
/*variable interface sin tipo asignado*/
var x interface{}
switch x.(type){ /*Retorna el tipo de x*/
/*Casos*/
 case nil: fmt.Println("Es una variable tipo nil")
 case int: fmt.Println("Es una variable tipo int")
 case float64: fmt.Println("Es una variable tipo float64")
 case int64: fmt.Println("Es una variable tipo int64")
 default: fmt.Println("No es ninguno de los tipos anteriores")
}
}
