// Programa donde se observan la variación de los indices i y j que controlan las posiciones donde se colocan los valores crecientes en una unidad a partir de cero de los cuadros magicos impares
package main
import (
     "fmt"
)
func main() {
  var i,j int   
  var filcol int = 9 //Declarar variable que guarda el número impar de filas y columnas
	i=0   
	j=(filcol-1)/2
	for m:=0 ; m < filcol; m++ {
	    for k:=0 ; k < filcol ; k++ {
	        fmt.Println("  K=",k,"  m=",m,"  (i=",i,",j=",j,")")
	     	i=i-1
	    	j=j+1
	    	if i<0 && k!=(filcol-1) { i=filcol-1 }  // i se inicia en 8 cuando i toma un valor negativo y no transcurrido un ciclo completo (k#8)
	    	if j==(filcol) && k!=(filcol-1) { j=0 } // j se inicia en 0 cuando j toma un valor  
		}	
        i=i+2
        j=j-1    
     }
 nota := `1. En este ejemplo un cuadro de (9x9) los indices de i y j variarán 81 veces. El comportamiento de las variaciones 
 es regular (los valores de i disminuyen cada ves una unidad y los de j aumenta una unidad) excepto después de 9 variaciones, por lo que 
 se divide la variaciones en 9 ciclos de 9 variaciones cada una. Por ello se implementa dos for. El for interno varia 9 veces a i y j.    
 2. Se observa en la 1ra secuencia que el primer valor de i es 0 y j toma el valor 4 [(núm de filas-1)/2 caso general (9-1)/2]
 3. En la cada secuencia de 9, en general los valores de i disminuyen 1 unidad mientras que los de j aumentan una unidad
 4. cuando i pasa a un valor menor a 0, se cambia su valor a 8(num. de filas-1) [9-1]
 5. Cuando j pasa al valor de 9(núm. de columnas), su valor se cambia a 0
 6. Cada vez que transcurre un ciclo de 9(núm. de filas o columnas)[cuando m cambia de valor] variaciones de i y j, sus valores serán:
    6.1 Al salir de un ciclo de 9 vueltas, como a i se le ha restado 1, y se entra en nuevo ciclo de 9 vueltas, se le incrementa 2 unidades 
    6.2 Al salir de un ciclo de 9 vueltas, como a j se le había sumado 1, y se entra en un nuevo ciclo de 9 vueltas, se le resta una unidad 
 7. observe que los valores que toman i y j estan entre el rango de 0 a 8, o sea, 9 valores.
 8. Puede observar las variaciones para un número de filas o columnas diferente cambiando el valor de la variable filcol 
 por otro numero impar en la línea 8 del programa.
 `
 fmt.Println(nota)
  
}
