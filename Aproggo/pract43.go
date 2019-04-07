package main
import (
       "fmt"
       "math"
    )

func main() {
        fmt.Println("Numero: ",1.3)    
        fmt.Println("Resultado de aplicarle la función math.Floor:",math.Floor(1.3))
        fmt.Println("Resultado de aplicarle la función math.Ceil:",math.Ceil(1.3))

var numimp int //Declarar variable que guardará posible número para determinar si es par o impar 
  
  fmt.Printf("Ingresa un número: ")       // Aquí no se ha validado la entrada
  fmt.Scanln(&numimp)
  
  if math.Floor(float64(numimp)/2)*2==float64(numimp)  {       // para chequear que  el número sea impar 
		fmt.Printf("El número es par")
    }  else {
	    fmt.Printf("El número es impar")
     }
   nota := `El proceso implementado es que al número ingresado se le divide por 2. El resultado se extrae la parte entera
   utulizando la función floor. Lo obtenido se multiplica por 2 y este último valor resultante se le compara con el
   número ingresado. De ser verdad la igualdad, el número ingreado es par sino impar.`
        fmt.Println("\n",nota)   
    }
