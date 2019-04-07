package main
import (
        "fmt"
        "bufio"
        "os"
        "strings"
        "strconv"
        "log"
        )

func main() {
     fmt.Print("Dinos cual es el precio: ")
      leer := bufio.NewReader(os.Stdin)
      entrada, error1 := leer.ReadString('\n')  // leer hasta el separador de salto de linea
      if error1 != nil {
          log.Fatal(error1)
	  }	  
      eleccion := strings.TrimRight(entrada, "\r\n") // Remover el salto de línea de la entrada del usuario
      fmt.Println(eleccion)
      salida,error2:= strconv.Atoi(eleccion)
      if error2 != nil {
          log.Fatal(error2)
	  }	  
      var iva float64 = 3
      fmt.Println("El iva es",iva,"%.")
      var tasa float64 = float64(salida)*iva/100
      fmt.Println("Tasa a cobrar",tasa)
      var total float64 = float64(salida)+tasa 
      fmt.Println("El costo total es ",total,"Bolivares.")
    
}





