package main
import "fmt"
func main(){

min, max := minmax(12, 23, 54)          // Llamado a la función. min y max asumen el tipo asignado  por :=
fmt.Printf("Min: %d\nMax: %d\n", min, max)
 }

 /*Función para encontrar el MÍNIMO y MÁXIMO entre tres números*/

 func minmax(a, b, c int) (int, int){     //Variables locales
 var min, max int
 
 if a > b && a > c{                  // Primero verifica si "a" es el MÁXIMO
    max = a
    if b < c{
       min = b
       } else {
         min = c
         }
    } else if b > a && b > c {        //Si "a" no fue el MÁXIMO, prueba con "b"
              max = b
              if a < c{
                 min = a
              } else {
                      min = c
                }
                        
            } else {                 // Si "a" y "b" no fueron los MÁXIMOS, entonces "c" es el MÁsimo
              max = c
              if a < b{
                  min = a
              } else {
                min = b
                }
               }
 return min, max
 }
