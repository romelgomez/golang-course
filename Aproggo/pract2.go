//* Todo programa o archivo de Go se inicia con la línea paquete(package) principal(main) */

package main

/* Luego de la linea package main viene la carga(importación) de los  paquetes(grupos)  de programas(funciones), preelaborados. 
   En esta oportunidad se carga el paquete  fmt que contiene el grupo de funciones y métodos  de entrada y salida */

import "fmt"          // se importa los métodos del paquete "fmt", entre ellos: Print, Println, Printf

/* Se procede con la edición o construción de la  Función principal que dirige la  ejecución del programa */

func main() {   //  Llave de inicio que agrupa las instrucciones que controlaran la ejecucion del  programa

/*Del paquete fmt se toma el método Println para mostrar el texto entre comillas en la pantalla y que produce un salto de linea */

fmt.Println("           Hola Mundo")
fmt.Println("¡Has instalado Go correctamente!")
fmt.Println("Println imprime el texto entre comillas")
fmt.Println("y luego ejecuta un salto de línea")
fmt.Println("como lo habrás observado")

 }   // Llave de cierre del grupo de instruciones que dirige la ejecucion del  programa


