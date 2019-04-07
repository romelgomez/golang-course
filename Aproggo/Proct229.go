package main
import "fmt"
func main(){

min, max := minmax(12, 23, 54)   // Llamado a la función. min y max asumen el tipo asignado  por :=
fmt.Printf("Min: %d\nMax: %d\n", min, max)
 }
           /*Función para encontrar el MÍNIMO y MÁXIMO entre tres números*/
           
 func minmax(a,b,c int) (int,int) {     // Variables locales  (int , int) indica que retornará dos valores enteros. a,b y c son enteras 
   var min, max int  
   max = a
   if b > max {max=b}
   if c > max {max=c}
   min = a
   if b < min {min=b}
   if c < min {min=c}
   return min, max
 
 }
