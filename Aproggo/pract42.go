// programa pra determinar si un año es bisiesto o no
package main
import (
       "fmt"
       "math"
    )

func main() {
        
var anobis int //Declarar variable que guardará el posible año bisiesto 
  
  fmt.Printf("Ingresa un año: ")
  fmt.Scanln(&anobis)
  
  if math.Floor(float64(anobis)/4)*4==float64(anobis)  {       // para chequear que  el año es bisiesto (si es divisible por 4 o múltiplo de 4 
		fmt.Println("El año ",anobis," es bisiesto")
    }  else {
	    fmt.Println("El año ",anobis," no es bisiesto")
     }
   nota := `El proceso implementado es que al año ingresado se le divide por 4. Al resultado se extrae la parte entera
   utiizando la función Ceil. Lo obtenido se multiplica por 4 y este último valor resultante se le compara con el
   año ingresado. De ser verdad la igualdad, el año es bisiesto sino no es bisiesto. Ejemplo: 2018/4 es 504,5. si la parte entera de este último 504*4 no da 2018.
   por tanto 2018 no es bisiesto.   Veamos 2020 es bisiento ya que 2020/4 es 505 exacto y 505*4 es 2020`
        fmt.Println("\n",nota)   
    }
