// Programa que construye cuadros mágicos con un número filas y columnas impares menores o iguales 15 (tamaño maximo de la capacidad del arreglo)
package main
import (
     "fmt"
     "math/rand"
     "time"
     "math"
)
func main() {
  var magico [15][15] int
  var i,j int   
  seg := time.Now().Unix()
  rand.Seed(seg)  // inicializar la raiz de los números rand a partir del valor de la variable seg
  vm := rand.Intn(10)		// se genera un número aleatorio entre 0 y 10 que inicia los valores del cuadro mágico
  var filcol int //Declarar variable que guarda el número impar de filas y columnas
  
  fmt.Printf("Dinos de cuantas filas o columnas (impares) quieres que construya el cuadro mágico: ")
  fmt.Scanln(&filcol)
  
if math.Floor(float64(filcol)/2)*2!=float64(filcol) && filcol<=15 && filcol>=3 {       // para chequear que la entrada para el número de filas y columnas sea impar 
  
	i=0   
	j=(filcol-1)/2
	for m:=0 ; m < filcol; m++ {
	    for k:=0 ; k < filcol ; k++ {
	        magico[i][j]=vm
	     	i=i-1
	    	j=j+1
	    	if i<0 && k!=(filcol-1) { i=filcol-1 }
	    	if j==(filcol) && k!=(filcol-1) { j=0 }
	    	vm=vm+1			
		}	
        i=i+2
        j=j-1    
     }
 fmt.Println("  CUADRO MAGICO ")    // IMPRESIÓN DEL CUADRO MAGICO

 for v:=0; v<filcol ; v++ {
	 fmt.Printf("----")
 }
 for u:=0; u<filcol ; u++ {
	 fmt.Printf("\n")
	 for w:=0; w<filcol ; w++ {
         fmt.Printf("|%3d",magico[u][w])
         }
     fmt.Printf("|\n")    
     for v:=0; v<filcol ; v++ {
	     fmt.Printf("----")
     }
 }
 var suma int = 0
 for z:=0 ; z<filcol; z++ {
     suma=suma+magico[z][0]
 }
     
 fmt.Printf("\n\n La suma de los elementos de cualquier fila, columna o diagoles principales suman:%5d",suma)
 var f int
 f=(filcol-1)/2        
 fmt.Printf("\n\n Los elementos de la fila central son:\n\n")
 for z:=0 ; z<filcol; z++ {
     fmt.Printf("|%3d",magico[f][z])
    }
  fmt.Printf("| y  suman:%4d",suma)

    fmt.Printf("\n\n Los elementos de la columna central son:\n\n")
 for z:=0 ; z<filcol; z++ {
     fmt.Printf("|%3d",magico[z][f])
  }
  fmt.Printf("| y  suman:%4d",suma)
  
  fmt.Printf("\n\n Los elementos de la diagonal principal son:\n\n")
 for z:=0 ; z<filcol; z++ {
     fmt.Printf("|%3d",magico[z][z])
  }
  fmt.Printf("| y  suman:%4d",suma)

  fmt.Printf("\n\n Los elementos de la diagonal secundaria son:\n\n")
 for z:=0 ; z<filcol; z++ {
     fmt.Printf("|%3d",magico[filcol-z-1][z])
  }
  fmt.Printf("| y  suman:%4d",suma)
 } else {
	fmt.Printf("El número de filas y columnas debe ser un numero, impar mayor igual a 3 y menor o igual a 15")
        }
}
