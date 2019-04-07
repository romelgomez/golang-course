// conversiones de cadenas a numerico
package main
import "strconv"
import "fmt"
func main() {
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println("La conversión(utilizando strconv.ParseFloat) de la cadena \"1.234\" y su resultado asignado a la variable f es:",f)
    fmt.Println("El producto del valor de f:",f,"por",2,"es",f*2)
    fmt.Println() // Println sin argumento deja linea en blanco
    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println("La conversión(utilizando strconv.ParseInt) de la cadena \"1234\" y su resultado asignado a la variable i es:",i)
    fmt.Println("El producto del valor de f:",i,"por",3,"es",i*2)
    fmt.Println()
    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println("La conversión(utilizando strconv.ParseUint) de la cadena \"789\" y su resultado asignado a la variable u es:",u)
    fmt.Println("El producto del valor de u:",u,"por",2,"es",u*2)
    fmt.Println()
    k, _ := strconv.Atoi("135")
    fmt.Println("La conversión(utilizando strconv.Atoi) de la cadena \"135\" y su resultado asignado a la variable k es:",k)
    fmt.Println("El producto del valor de k:",k,"+",3,"es",k+3)
    fmt.Println()
    z, e := strconv.Atoi("135")
    fmt.Println("La conversión(utilizando strconv.Atoi) de la cadena \"135\" y su resultado asignado a la variable i es:",z)
    fmt.Println("El producto del valor de z:",z,"por",3,"es",z*3)
    fmt.Println()
    ventrada := "12"
    v,e := strconv.Atoi(ventrada)
    fmt.Println("La conversión del valor de variable de cadena ventrada:",ventrada," y su resultado asignado a la variable v es:",v)
    fmt.Println("El producto del valor de v:",v,"por",3,"es",v*3)
    fmt.Println("La variable(e) que acompaña a la variable v, se utiliza para detectar errores de conversión:",e)
    
    
    
}





