// operadores lógicos a nivel de bit
package main
import (
        "fmt"
                )

func main() {
	
	
	nota:=`Lógica proposicional/Proposiciones. ... En la lógica proposicional nos interesan los enunciados aseverativos y se 
les llama proposiciones. La lógica se interesa por este tipo de enunciados porque se les puede asignar un valor 
de verdad, ya sea falso (la información es incorrecta) o verdadero (la información es correcta). La proposición
"los pajaros vuelan" tiene un valor de verdad verdadero(true). "Las lombrises vuelan" tienen un valor de verdad
falso(false). La proposición: 2>3 es false. La proposición: 2<3 es true `

	fmt.Println(nota,"\n\n")
	

//Imprime el valor de verdad de una comparación con los operadores de relación
    var x int = 10 
    var y int = 6
    fmt.Println("\n")
    fmt.Println("El resultado de la comparación de si",x,"es menor a",y,"  (",x,"<",y,")","es",x<y)
    fmt.Println("El resultado de la comparación de si",x,"es mayor a",y,"  (",x,">",y,")","es",x>y)
    fmt.Println("El resultado de la comparación de si",x,"es igual a",y,"  (",x,"==",y,")","es",x==y)
    fmt.Println("El resultado de la comparación de si",x,"es distinto a",y,"  (",x,"!=",y,")","es",x!=y)  
    fmt.Println("El resultado de la comparación de si",x,"es menor o igual a",y,"  (",x,"<=",y,")","es",x<=y)
    fmt.Println("El resultado de la comparación de si",x,"es mayor o igual a",y,"  (",x,">=",y,")","es",x>=y)
    fmt.Println("El resultado de la comparación de si",x,"-",4,"es menor o igual a",y,"  (",x,"-",4,"<=",y,") es",x-4<=y)
    fmt.Println("El resultado de la comparación de si",x,"es mayor o igual a",y,"+",4,"  (",x,">=",y,"+",4,") es",x>=y+4)



 // Imprime los resultados de los operadores lógicos a nivel de bit
    fmt.Println(" Conjunción(&) : 'A & B' puede ser interpretado como 'A y B' ")
    fmt.Println("\n El resultado de ",0,"&",0,"es",0&0,"    false")  
    fmt.Println("\n El resultado de ",0,"&",1,"es",0&1,"    false")
    fmt.Println("\n El resultado de ",1,"&",0,"es",1&0,"    false")
    fmt.Println("\n El resultado de ",1,"&",1,"es",1&1,"    true")
    fmt.Println("\n Disyunción)(|) : | 'A & B' puede ser interpretado como 'A o B' ")
    fmt.Println("\n El resultado de ",0,"|",0,"es",0|0,"    false")  
    fmt.Println("\n El resultado de ",0,"|",1,"es",0|1,"    true")
    fmt.Println("\n El resultado de ",1,"|",0,"es",1|0,"    true")
    fmt.Println("\n El resultado de ",1,"|",1,"es",1|1,"    true")
    fmt.Println("\n O-exclusiva (XOR)(^) : 'A ^ B' puede ser interpretado como o A, o B, pero no ninguno ni ambos ")
    fmt.Println("\n El resultado de ",0,"^",0,"es",0^0,"    false")  
    fmt.Println("\n El resultado de ",0,"^",1,"es",0^1,"    true")
    fmt.Println("\n El resultado de ",1,"^",0,"es",1^0,"    true")
    fmt.Println("\n El resultado de ",1,"^",1,"es",1^1,"    false")


}
