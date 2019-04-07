// Estudio de arreglos

package main
import (
    "fmt"

)

// var nombre_de_variable [TAMAÑO] tipo_de_dato

func main() {

 var dias_semana [7] string

 dias_semana [0] = "domingo"
 dias_semana [1] = "lunes"
 dias_semana [2] = "martes"
 dias_semana [3] = "miércoles"
 dias_semana [4] = "jueves"
 dias_semana [5] = "viernes"
 dias_semana [6] = "sábado"
 
 fmt.Println(dias_semana[0])
 fmt.Println(dias_semana[2])
 fmt.Println(dias_semana[1])
 fmt.Println(dias_semana[3])
 fmt.Println(dias_semana[5])
 fmt.Println(dias_semana[4])
 fmt.Println(dias_semana[6])
 
// var nombre_de_variable = [] tipo_de_datos{datos del mismo tipo separado por comas}
      
 var vocales = [] string{"a","e","i", "o", "u"}
 
 fmt.Println(vocales[0])
 fmt.Println(vocales[1])
 fmt.Println(vocales[2])
 fmt.Println(vocales[3])
 fmt.Println(vocales[4])

// nombre_de_variable := [tamaño]tipo_de_datos{datos del mismo tipo separado por comas}
 
 numerico := [5]int{1,2,3,4,5}
 
 fmt.Println(numerico[0])
 fmt.Println(numerico[1])
 fmt.Println(numerico[2])
 fmt.Println(numerico[3])
 fmt.Println(numerico[4])
 
 
 fmt.Println("Dias de la semana:", dias_semana)
 fmt.Println("Vocales :", vocales)
 fmt.Println("Naturales menores que 6:", numerico)

//                                Arreglos multidimensionales

  suma_doce := [3][3]int{{7,0,5},{2,4,6},{3,8,1}}
	  fmt.Println("\n Cuadro mágico suma 12 filas, columnas y diagonales\n")
	  fmt.Println("   -------------")
	  fmt.Println("   | ",suma_doce[0][0],"|",suma_doce[0][1],"|",suma_doce[0][2],"|")
	  fmt.Println("   -------------")
	  fmt.Println("   | ",suma_doce[1][0],"|",suma_doce[1][1],"|",suma_doce[1][2],"|")
	  fmt.Println("   -------------")
	  fmt.Println("   | ",suma_doce[2][0],"|",suma_doce[2][1],"|",suma_doce[2][2],"|")
	  fmt.Println("   -------------")
	  


}
