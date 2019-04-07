// Progama para conocer la forma de determinar el tipo de variable. Esto es importante porque Go es un lenguaje fuermente tipado
// y las operaciones entre variables de tipos diferentes no se permiten por lo que muchas veces hay que realizar conversiones y 
// para ellos hay que conocer el tipo de variable con que se trabaja y ello sirva para realizar las depuraciones de errores.  
package main
import (
        "fmt"
        "reflect"  //su uso típico es tomar un valor y extraer la información de tipo de variable. la función TypeOf de "reflect" devuelve el tipo de la variable evaluada.      
        )          // Otra forma de obtener el tipo de variable o número es utilizando de "form"  Printf(imprimir con formato) el parametro %T

func main() {

 var x=3
 fmt.Println(x,"\nTipo de dato:",reflect.TypeOf(x))
 fmt.Printf("Tipo de dato: %T \n",x)                       // Printf después de imprimir no genera un salt de linea. Utilizamos \n para generar salto de linea si se quiere
 var a float32
 a=5.3
 fmt.Println(a,"\nTipo de dato:",reflect.TypeOf(a)) 
 fmt.Printf("Tipo de dato: %T \n",a)
 
 y:=x
 fmt.Println(y,"\nTipo de dato:",reflect.TypeOf(y)) 
 fmt.Printf("Tipo de dato: %T \n",y)
 
 fmt.Println(64,"\nTipo de dato:",reflect.TypeOf(64)) 
 fmt.Printf("Tipo de dato: %T",64)
 fmt.Printf("\n")
 fmt.Println(2.3,"\nTipo de dato:",reflect.TypeOf(2.3)) 
 fmt.Printf("Tipo de dato: %T \n",2.3)
 
 fmt.Println("Hola mundo \nTipo de dato:",reflect.TypeOf("Hola mundo")) 
 fmt.Printf("Tipo de dato: %T","Hola mundo")
 fmt.Printf("\n")
 fmt.Println("false \nTipo de dato:",reflect.TypeOf(false))
 fmt.Printf("Tipo de dato: %T ",false)

 fmt.Printf("\n")
 fmt.Println("true \nTipo de dato:",reflect.TypeOf(true))
 fmt.Printf("Tipo de dato: %T \n",true)

 fmt.Println("2>3 \nTipo de dato:",reflect.TypeOf(2>3))
 fmt.Printf("Tipo de dato: %T",2>3)
    
  

}




