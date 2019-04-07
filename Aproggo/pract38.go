package main
import (
        "fmt"

        )

func main() {

    var precio int //Declarar variable y tipo antes de escanear, esto es obligatorio
    fmt.Print("Dinos cual es el precio del producto: ")
	fmt.Scanln(&precio)
	var iva float64 //Declarar variable y tipo antes de escanear, esto es obligatorio
    fmt.Print("Porcentaje del iva a pagar: ")
	fmt.Scanln(&iva)
    var tasa float64 = float64(precio)*iva/100
    fmt.Println("La tasa a cobrar por iva es:",tasa)
    var total float64 = float64(precio)+tasa 
    fmt.Printf("El producto que tiene como precio %d, con un porcentaje de iva de %2.0f  tiene un costo total de %.2f Bolivares.",precio,iva,total)
   
}
