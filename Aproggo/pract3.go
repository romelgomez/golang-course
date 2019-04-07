// Aqui se estudia las diferentes formas de declarar e inicializar variables
package main
import (
        "fmt"
        )

func main() {

    // Declarar variable:  palabra clave (var) + nombre de la variable (cantidad) + tipo de datos (int)

    var cantidad int 
    var nombreAlumno string

   // Declarar varias variables del mismo tipo, separadas por comas

    var oral, escrito float64 

   // Asignamos un valor a una variable con el uso del signo (=)    

    cantidad = 2 
    nombreAlumno = "Pedro Picapiedra"

  // Podemos asignar valores a varias variables en la misma declaración

    oral, escrito = 15.5, 19 

  // Una vez asignado valores a la variables, podemos usar sus valores en cualquier contexto dentro de su ambito en que las podamos usar
    
    fmt.Println(nombreAlumno)
    fmt.Println("Ha realizado",cantidad,"exámenes")
    fmt.Println("En el oral su nota ha sido",oral)
    fmt.Println("Y en el escrito su nota ha sido",escrito)

 
 // Declaración de variables y inicialización posterior:
    var z float32
    z = 1.79
    var m int
    m = 32
    var direccion string
    direccion = "Calle Principal # 4148"
  
 
    fmt.Println("Valor de z:",z) 
    fmt.Println("Valor de m:",m) 
    fmt.Println("Direccion:",direccion) 
    fmt.Println("\n")  // Deja una linea en blanco

 
 // Inicialización de variables durante la declaración de variables, especificando o no el tipo de variable :
    var x = 3                                                                 
    var peso  = 28.5 
    var nombre = "Rita"
 
    fmt.Println("Valor de x:",x)
    fmt.Println("Peso:",peso)
    fmt.Println("Nombre:",nombre,"\n")  // se agrega "\n" para separa la salida con una linea en blanco
                                                      
    var y int = 3
    var altura float32 = 1.68 
    var apellido string = "Napolitana"
  
    fmt.Println("\n")  //     \n Deja una linea en blanco
    fmt.Println("Valor de y:",y)
    fmt.Println("Altura:",altura)
    fmt.Println("Apellido:",apellido,"\n") 
 
 
// Inicializar más de una variable en una declaración:

    var o,p,q = 3, 4.0 ,"Hola"

    fmt.Println("Valor de o:",o) 
    fmt.Println("Valor de p:",p) 
    fmt.Println("Valor de q:",q) 

    var s int
    s=5
    var t int
    t=s
    fmt.Println(s)
    fmt.Println(t)

    var c1 float32 
    c1=8.4
    var  d1 float32
    d1=c1   
    fmt.Println(c1)
    fmt.Println(d1)  


   // asignacion dinamica   utilizando el operador := (no se puede utilizar dentro de la declaración var)

    var x2=3
    y2:=x2          //  y2 asume el tipo de x2 o sea int
    fmt.Println(x2)
    fmt.Println(y2)
 
    var a1 float32 = 5.6
    b1:=a1     //  b1 asume el tipo de a1 o sea float32
    fmt.Println(a1)
    fmt.Println(b1)

    fmt.Println("\n")    //                                          \n se utiliza para generar un salto de linea
    fmt.Println("Conclusión: \n")

    fmt.Println("Da igual declarar")
    fmt.Println("\n")
    fmt.Println("var cadena = \"Ésta es una cadena\"")               //    se utiliza \" para colocar comillas entre comillas en encabezados de salida.    
    fmt.Println("var cadena string = \"Esta es una cadena\"")        //  Ojo no se puede colocar comillas entre comillas sin el caracter \
    fmt.Println("cadena := \"Esta es una cadena\"")
    
    fmt.Println("\n")
    fmt.Println("var entero = 5")     
    fmt.Println("var entero int = 5") 
    fmt.Println("entero := 5")
    fmt.Println("var entero int") 
    fmt.Println("entero = 5")    



}




