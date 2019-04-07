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
     fmt.Print("Escribe una puntuación: ")
      leer := bufio.NewReader(os.Stdin)
      entrada, error := leer.ReadString('\n')  // leer hasta el separador de salto de linea
      if error != nil {
          log.Fatal(error)
	  }	  
      entrada = strings.TrimSpace(entrada) // Remover espacios y el salto de línea de la entrada del usuario
      puntuacion, error := strconv.ParseFloat(entrada,64)
      if error != nil {
          log.Fatal(error)
	  }	  
      var status string
      if puntuacion >= 50 {
         status = "aprobado"
      }  else {
         status = "suspenso"
      }
      fmt.Println("una puntuación de ",puntuacion,"es",status)   
    
}



