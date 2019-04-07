
// programa para  determinar la posición que ocupa un(a) elemento(palabra) dentro de un arreglos de los mismos 

package main

import (
	"fmt"
	//"strings"
)

func main() {
    var j  int32
    var tdia string = "Viernes"
    // var NMESESI = [] string {"January","February","March","April","May","June","July","August","September","October","November","December"}
	var dia = [] string {"Lunes","Marte","Miercoles","Jueves","Viernes","Sabado","Domingo"}
	for i:=0;i<7;i++ {
	    if tdia==dia[i] {j=int32(i)}        // guarda en j la posición que ocupa tdia(viernes) en el arreglo de dia 
    }
    fmt.Printf("El dia %s acupa la posición %d de la semana",tdia,j+1)
   
}

