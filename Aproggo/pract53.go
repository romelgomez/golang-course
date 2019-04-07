package main

import (
	"fmt"
	//"strings"
)

func main() {

    var j  int32
    var tdia string = "Viernes"
	var dia = [] string {"Lunes","Marte","Miercoles","Jueves","Viernes","Sabado","Domingo"}
	for i:=0;i<7;i++ {
	    if tdia==dia[i] {j=int32(i)}
    }
   fmt.Printf("El dia %s acupa la posición %d de la semana",tdia,j+1)
   fmt.Printf("\n...\n")
                          // imprime valores desde 0 a 9  
   var p = 0
   for p < 10 {
      fmt.Println("Valor de p:", p)
      p++
   }
   fmt.Printf("\n...\n")

                          // imprime valores desde 0  a 9
   var q int
   for q=0 ; q < 10; q++ {
       fmt.Println("Valor de q:", q)
   }
   fmt.Printf("\n...\n")
    
                                         //  imprime los valores de i desde 0  hasta completar el número de elementos de arreglo que son 9  
   arreglo:=[9]int{0,1,4,6,10,9}
   for i:= range arreglo{
       fmt.Printf("Valor de i: %d \n", i)
   }
   fmt.Printf("\n...\n")

                                        //    imprime los valores del arreglo al recorrerlo y en cada vuelta el valor del indice i  
   for i, j:= range arreglo{
   fmt.Printf("Valor de j: %d en vuelta # %d \n", j,i)
   }
   fmt.Printf("\n...\n")

                                       //  imprime el valor del indice i desde cero y cuando este toma el valor de 7 se sale del for. 
   for i:=0 ; i < 10; i++ {
      fmt.Printf("Valor de i: %d \n", i)
      if i == 7{
         fmt.Printf(" así que saldremos del ciclo...\n")
         break
      }
  }
  fmt.Printf("\n...\n")
                                       //   imprime el valor de la variable i inicializada en 0 y cuando esta toma el valor de 4 le suma 3 y continua hasta alcansar   
  var i = 0
  for i < 10 {
     fmt.Println("Valor de i: %d", i)
     if i == 6{
        fmt.Printf(" sumaremos 3\n")
        i = i + 3
        continue
     }
  i++
  }
  fmt.Printf("\n...\n")
                                        // imprime los elementos del arreglo dia y la posición que ocupa cada elemento
   arreglo2:=dia
   for i, j:= range arreglo2{
   fmt.Printf("A %s le corresponde el número %d \n", j,i)
   }
    
}
