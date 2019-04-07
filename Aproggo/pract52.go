// Estudio del for
   
package main

import (

     "fmt"
)
func main() {

   // Ciclo for con una condición de paro

 var i = 0
 for i < 10 {
 fmt.Println("Valor de i:", i)
 i++
 }
 
 // Ciclo for con inicialización, condición e incremento
 for i=0 ; i < 10; i++ {
 fmt.Println("Valor de i:", i)
 }
 
// Ciclo for con un rango

 Valores_rango := [7]int{0,1,4,6,10,9}
 fmt.Printf("\n") 
 for i:= range Valores_rango {
	
   fmt.Printf("Valor de i (tomado del rango): %d\n", i)
 }
 
 for i, j:= range Valores_rango {
 
 fmt.Printf("Valor de j: %d en vuelta #%d\n", j,i)

 }
 
 
 Valores_rangoS := [7]string{"pedro","pablo","chucho","Jacinto","Jose"}
 fmt.Printf("\n") 
 for i:= range Valores_rangoS {
	
   fmt.Printf("Valor de i (tomado del rango): %d\n", i)
 }
 
 for i, j:= range Valores_rangoS {
 
 fmt.Printf("Valor de j: %s en vuelta #%d\n", j,i)

 }
 
nums := []int{2, 3, 4}  // Aquí usamos rango para sumar los números en un arreglo. Las matrices funcionan así también.
    sum := 0
    for _, num := range nums {
        sum += num
    } 
  fmt.Println("sum:", sum)
 
 // rango en matrices y divisiones proporciona tanto el índice como el valor de cada entrada.
 // Arriba no necesitábamos el índice, así que lo ignoramos con el identificador en blanco. Sin embargo a veces realmente queremos los índices.
 
 for i, num := range nums {      
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
 
 // el rango en el mapa itera sobre pares clave/valor.
 
 kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
    
    // rango también puede iterar sólo las claves de un mapa.
    
    for k := range kvs {
        fmt.Println("key:", k)
    }

    // el intervalo en cadenas itera sobre puntos de código Unicode. El primer valor es el índice del octeto que comienza de la runa y el segundo la runa sí mismo.
    
    for i, c := range "go" {
        fmt.Println(i, c)
    }    
    
    
 }

 
 
 



