// Adivinar un numero entre 1 y 100
package main
import (
        "fmt"
        "bufio"
        "os"
        "strings"
        "strconv"
        "log"
        "math/rand"
        "time"
        )

func main() {

	 segundos := time.Now().Unix()
	 fmt.Println(segundos)
	 rand.Seed(segundos)  // inicializar la raiz de los números rand a partir del valor de la variable segundos
	 target := rand.Intn(100)+1		// se genera un número aleatorio entre 1 y 100
     acertar := false
     fmt.Println("He elegido un número entre 1 y 100.") 
     fmt.Println("¿Prueba adivinarlo?")
     leer := bufio.NewReader(os.Stdin)
     
   for respuestas := 0; respuestas<10; respuestas++ {
	fmt.Println("Te quedan",10-respuestas,"respuestas")   
     entrada, error := leer.ReadString('\n')  // leer hasta el separador de salto de linea
     if error != nil {
          log.Fatal(error)
	  }	  
      entrada = strings.TrimSpace(entrada) // Remover espacios y el salto de línea de la entrada del usuario
      adivina, error := strconv.Atoi(entrada)
      if error != nil {
          log.Fatal(error)
	  }	
	  
	  //fmt.Println(target)
	  //fmt.Println(adivina)
	  if adivina == target {
		   acertar = true
	       fmt.Println("Felicitaciones. Acertaste.")
	       break       // instrucción para parar el loop
	  } else if adivina < target {
	         fmt.Println("No, el número es mayor.")
             } else {
	                fmt.Println("No, el número es menor.")
                    }
                    
   }
   if !acertar {    
           fmt.Println("Lo sentimos, no acertaste. El número era:",target)
		}
}




