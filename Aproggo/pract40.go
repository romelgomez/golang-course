package main
import "fmt"
func main() {
var calificacion int = 10
if calificacion < 6 { 
   fmt.Println("Reprobaste.") 
   } else if calificacion >= 6 && calificacion <= 8 {
   fmt.Println("Aprobaste.")
   } else if calificacion == 9 {
   fmt.Println("Aprobaste. Te fue muy bien.")
   } else {
   fmt.Println("Felicidades. Aprobaste con calificación perfecta")
   }
   fmt.Println("Tu calificación fue de: ", calificacion)
}
