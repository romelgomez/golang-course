package main
import "fmt"
func main() {
/*variable local de tipo entero*/
var hora int = 10
/*Se pasa hora como variable de prueba*/
switch hora{
/*Si hora coincide con alguna de las literales especificadas*/
      case 1, 2, 3, 4: fmt.Println("Aún es temprano")
      case 5, 6, 7: fmt.Println("Está atardeciendo")
      case 8: fmt.Println("Acaba de oscurecer")
      case 9, 10, 11: fmt.Println("Ya es tarde")
      default: fmt.Println("Es demasiado tarde")
      }
}
