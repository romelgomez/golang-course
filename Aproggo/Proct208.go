// Programa que determina el día de la semana de cualquier fecha(dd/mm/yyyy) no se valida la entrada
package main
import (
     "fmt"
     "math"
)
func main() {

  var dia = [] string {"Domingo","Lunes","Marte","Miercoles","Jueves","Viernes","Sabado"}
    fmt.Print("Ingrese datos de la fecha de la que desea saber qué día fué:\n\n")
    var dd,mm,yy float64
    fmt.Print("Ingresa día:")
	fmt.Scanln(&dd)
    fmt.Print("Ingresa mes:")
	fmt.Scanln(&mm)
	fmt.Print("Ingresa año:")
	fmt.Scanln(&yy)
    a := math.Floor(float64(14-mm)/12) 
    y := float64(yy)-a
    m := float64(mm)+float64(12*a)-2
    d := int32(dd + y + math.Floor(y/4) - math.Floor(y/100) + math.Floor(y/400) + math.Floor((31*m)/12))%7
    fmt.Println(d)
    fmt.Println("La fecha (",dd,"/",mm,"/",yy,") le corresponde el día :",dia[d])
nota := `
Tomado del resumen de La congruencia de Zeller, un algoritmo ideado por Julius Christian Johannes Zeller 
para calcular el día de la semana de cualquier fecha del calendario. 
a = (14 - Mes) / 12
y = Año - a
m = Mes + 12 * a - 2

Para el calendario Gregoriano:
 d = (día + y + y/4 - y/100 + y/400 + (31*m)/12) mod 7
El resultado es un cero (0) para el domingo, 1 para el lunes… 6 para el sábado

Para el calendario Juliano:
d = (5 + dia + y + y/4 + (31*m)/12) mod 7
---------------------------------------------------------------
Ejemplo, ¿En qué día de la semana cae el 2 de agosto de 1953?? cálculo gregoriano.
 a = (14 - 8) / 12 = 0
 y = 1953 - 0 = 1953
 m = 8 + 12 * 0 - 2 = 6
 d = (2 + 1953 + 1953 / 4 - 1953 / 100 + 1953 / 400 + (31 * 6) / 12) Mod 7
   = (2 + 1953 +  488   -    19   +     4    +    15    ) mod 7
   = 2443 mod 7
   = 0
  El valor cero(0) corresponde al domingo. `
  fmt.Println(nota)

}
