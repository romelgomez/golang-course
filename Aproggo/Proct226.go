package main
  import "fmt"
  func main(){
     var res int    //Variable local
     fmt.Printf("El resultado de 122*120: %d\n", mult(122, 120))  //Llamado a función
     fmt.Printf("El resultado de 130*105: %d\n", mult(130, 105))  //Llamado con parámetros distintos
     res = mult(20, 178)                                          //Se almacena el resultado en una variable res
     fmt.Printf("Resultado almacenado en una variable: %d \n", res)
     var a,b int
     a = 5
     b = 7
     res = mult(a,b)
     fmt.Printf("El producto de %d y  %d es:  %d \n",a,b,res)
     
     
     var largo int //Declarar variable y tipo antes de escanear, esto es obligatorio
     fmt.Print("Ingresa valor del largo de un rectángulo: ")
	 fmt.Scanln(&largo)
	 var ancho int //Declarar variable y tipo antes de escanear, esto es obligatorio
     fmt.Print("Ingresa valor del ancho del mismo:  ")
	 fmt.Scanln(&ancho)
     fmt.Printf("El área del rectángulo de largo %.2d y ancho %.2d es: %.3d\n",largo,ancho,mult(largo,ancho))
  
}

/*Función que multiplica dos números*/

func mult(numero1, numero2 int) (int){            // numero1 y numero2 declaradas entera (int) indica el tipo de Variable de retorno
  var resultado int
  resultado = numero1 * numero2
  return resultado
}
