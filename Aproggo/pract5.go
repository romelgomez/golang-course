// Elementos para formatos de salida
package main
import (
        "fmt"

        )

func main() {

    
    fmt.Printf("Un punto flotante: %f\n",3.141516)
    fmt.Printf("Un entero: %d\n",64)
    fmt.Printf("Una cadena de caracteres: %s\n","Lograré mis proyectos")
    fmt.Printf("Un booleano: %t\n",true)
    fmt.Printf("formato apropiado segun tipo valores: %v %v %v %v\n",85,3.14,"Josefina",true)
    fmt.Printf("formato como aparece en el código: %#v %#v %#v %#v\n",85,3.14,"Josefina",true)
    fmt.Printf("Tipo de valor proporcionado: %T %T %T %T\n",85,3.14,"Josefina",true)
    fmt.Printf("Signo de porcentaje: %%\n")
    fmt.Printf("El costo de las %s es de %d bolivares cada uno.\n","mantequillas",250)
    fmt.Printf("%f cajas necesarias\n",3.141516)
    fmt.Printf("%2.3f cajas necesarias\n",3.141516)
    fmt.Printf("%.2f cajas necesarias\n",3.141516)
    fmt.Printf("\n\n\n\n\n")
    fmt.Printf("%12s | %s\n","Producto","Costo en Bolivares")
    fmt.Printf("------------------------------------------\n")
    fmt.Printf("%12s | %d\n","Abono",172)
    fmt.Printf("%12s | %d\n","Herramientas",143)
    fmt.Printf("%12s | %d\n","Madera",55)
    
 // acento invertido   
        
     var acento_inv string = `
 
          Esto es una Prueba 
        
         del uso del acento 
         
        invertido para construir
           
            encabezados, menu, etc `
     
     fmt.Println(acento_inv)
     
// para colocar comillas entre comillas
     
     var comilla_entre_com string = "Muestra la combinación de caracteres para \"colocar comillas en las salidas\" "
     fmt.Println("\n",comilla_entre_com)


    var encabezado string  
	encabezado = 
`                        // acento invertido
	Bienvenido
	[ 1 ] zapato
	[ 2 ] corbata
	¿Qué prefieres?
`                        // acento invertido
	fmt.Println(encabezado)
     
}
