// uso de los operadores aritméticos
package main
import (
        "fmt"
                )

func main() {

    var x int = 10 
    var y int = 6
    fmt.Println("La suma de",x,"mas",y,"es",x+y)  
    fmt.Println("La resta de",x,"menos",y,"es",x-y) 
    fmt.Println("El producto de",x,"por",y,"es",x*y)    
    fmt.Println("La división de",x,"entre",y,"es",x/y,",  real:",float32(x)/float32(y)) 
    fmt.Println("El resto de dividir",x,"entre",y,"es",x%y)
    
    a := x+y   // declara la variable a y le asigna el valor de la suma de los valores de x e y
    b := x-y   // declara la variable b y le asigna el valor de la resta de los valores de x e y 
    var c,d,e,f,g = 20,25,5,100,99 // declara las variables c  y d y las inicializa en 20 y 25 respectivamente
    
    fmt.Println("\n\n")
    fmt.Print("Al valor de ",c)    // fmt.Print   imprime sin salto de linea
    c += a      // la agrega al valor de la  variable c el valor de la variable a 
    fmt.Println(" se le agrega",a,"(resultado de sumar",x,"+",y,"). Resultando :",c)
    
    fmt.Println("\n\n")
    fmt.Print("Al valor de ",d)    // fmt.Print   imprime sin salto de linea al igual Printf pero hace caso omiso a los parámetros de formato
    d -= b
    fmt.Println(" se le resta",b,"(resultado de restar",x,"-",y,"). Resultando :",d)
    
    fmt.Println("\n\n")
    fmt.Print("El valor de ",e)    // fmt.Print   imprime sin salto de linea al igual Printf pero hace caso omiso a los parámetros de formato
    e *= x
    fmt.Println(" se multiplica por ",x,". Resultando :",e)

    fmt.Println("\n\n")
    fmt.Print("El valor de ",f)    // fmt.Print   imprime sin salto de linea
    f /= x    // el resultado es exacto si x divide f sino el resultado sera la parte entera de la división
    fmt.Println(" se divide por ",x,". Resultando :",f)

    fmt.Println("\n\n")
    fmt.Println("El valor de ",g)    // fmt.Print   imprime sin salto de linea al igual Printf pero hace caso omiso a los parámetros de formato
    fmt.Printf("El valor de %2d\n",g)
    fmt.Print("El valor de %2d",g)
    fmt.Println("\n\n")
    g %= x    
    fmt.Print("El valor de ",g)
    fmt.Println(" se divide por ",x,". Resultando como resto:",g)
    
}




