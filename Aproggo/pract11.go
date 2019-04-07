// Program para observar el comportamiento de las operaciones básicas matemáticas

package main
import (
        "fmt"
//        "math"
               
        )

func main() {

  var a,b float32
  a,b = 6.79,5.3

  fmt.Println("La suma de",a,"y",b,"es:",a+b) 
  fmt.Println("La resta de",a,"menos",b,"es:",a-b)
  fmt.Println("La multipliocación de",a,"por",b,"es:",a*b)
  fmt.Println("La división de",a,"entre",b,"es:",a/b)
  fmt.Println("\n") 

  var c,d int
  c,d = 32, 45
 
  fmt.Println("La suma de",c,"+",d,"es:",c+d) 
  fmt.Println("La resta de",c,"-",d,"es:",c-d)
  fmt.Println("La multipliocación de",c,"*",d,"es:",c*d)
  fmt.Println("La división de",c,"/",d,"es:",c/d,"       Resultado de convertir float32(",c,"/",d,"): ",float32(c/d),"    Resultado real, convertir float32(",c,")/","float32(",d,") : ",float32(c)/float32(d))
  fmt.Println("\n")
  
  var e = 13                                                                
  var f = 5 
 
  fmt.Println("La suma de",e,"y",f,"es:",e+f) 
  fmt.Println("La resta de",e,"menos",f,"es:",e-f)
  fmt.Println("La multipliocación de",e,"por",f,"es:",e*f)
  fmt.Println("La división de",e,"/",f,"es:",e/f,"       Resultado de convertir float32(",e,"/",f,"): ",float32(e/f),"    Resultado real, convertir float32(",e,")/","float32(",f,") : ",float32(e)/float32(f))
  fmt.Println("\n") 
                                                    
  var g int = 3            // buscar g/h imprima el cero con sus decimales
  var h int = 10  
 
  fmt.Println("La suma de",g,"+",h,"es:",g+h) 
  fmt.Println("La resta de",g,"-",h,"es:",g-h)
  fmt.Println("La multipliocación de",g,"*",h,"es:",g*h)
  fmt.Println("La división de",g,"/",h,"es:",g/h,"       Resultado de convertir float32(",g,"/",h,"): ",float32(g/h),"    Resultado real, convertir float32(",g,")/","float32(",h,") : ",float32(g)/float32(h))
  fmt.Println("\n")
 

  var t int     // buscar
  t=5
  var u int
  u=t+5

  fmt.Println("La suma de",t,"y",u,"es:",t+u) 
  fmt.Println("La resta de",t,"menos",u,"es:",t-u)
  fmt.Println("La multipliocación de",t,"por",u,"es:",t*u)
  fmt.Println("La división de",t,"/",u,"es:",t/u,"       Resultado de convertir float32(",t,"/",u,"): ",float32(t/u),"    Resultado real, convertir float32(",t,")/","float32(",u,") : ",float32(t)/float32(u))
  fmt.Println("\n") 

  fmt.P


  var v float32 
  v=8.4
  var  w float32
  w=v*2   

  fmt.Println("La suma de",v,"y",w,"es:",v+w) 
  fmt.Println("La resta de",v,"menos",w,"es:",v-w)
  fmt.Println("La multipliocación de",v,"por",w,"es:",v*w)
  fmt.Println("La división de",v,"entre",w,"es:",v/w)
  fmt.Println("\n") 


 // asignacion dinamica   utilizando := (no se puede utilizar dentro de la declaración var)

  var a1=7
  b1:=a1-2   

  fmt.Println("La suma de",a1,"+",b1,"es:",a1+b1) 
  fmt.Println("La resta de",a1,"-",b1,"es:",a1-b1)
  fmt.Println("La multipliocación de",a1,"*",b1,"es:",a1*b1)
  fmt.Println("La división de",a1,"/",b1,"es:",a1/b1,"       Resultado de convertir float32(",a1,"/",b1,"): ",float32(a1/b1),"    Resultado real, convertir float32(",a1,")/","float32(",b1,") : ",float32(a1)/float32(b1))
  fmt.Println("\n")

 
  var c1 float32 = 5.6
  d1:=c1+3  

  fmt.Println("La suma de",c1,"+",d1,"es:",c1+d1) 
  fmt.Println("La resta de",c1,"-",d1,"es:",c1-d1)
  fmt.Println("La multipliocación de",c1,"*",d1,"es:",c1*d1)
  fmt.Println("La división de",c1,"/",d1,"es:",c1/d1)
  fmt.Println("\n")


  var o,p,q,r = 3, 4.5 ,"Hola","mundo"         // por defecto la variable p toma tipo float64

  fmt.Println("La suma de float64(",float64(o),") y ",p," es:",float64(o)+p)
  fmt.Println("La concatenación de la cadena ",q," + la cadena",r,"resulta :",q+r) 


}




