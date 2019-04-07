// cálculo del área de un polígono regular (longitud del perímetro por el valor de apotema entre dos) pentágono( 5 lados)
// cálculo conocido el radio de la circunferencia que lo circunscribe y el número de lados 
// http://www.electrontools.com/Home/WP/2017/12/05/apotema-de-un-poligono-regular/

package main
import (
        "fmt"
        "math"
        )

func main() {

    var radio,num_lados = 150,5
    
    aalfa := 360/num_lados
    fmt.Println(aalfa)         
    abeta := aalfa/2
    fmt.Println(abeta)
    adelta := 90-abeta
    fmt.Println(adelta)
    angrad:=float64(adelta)*math.Pi/180    // convertir adelta(grados) a angrad(radianes)
    fmt.Println(adelta,"grados equivalen a ",angrad,"radianes")          // el argumento del coseno ha de ser en radianes 
    long_lado := 2*float64(radio)*math.Cos(float64(angrad))              // calculo de la longitud del lado
    fmt.Println("Longitud del lado:",long_lado)
    apotema := math.Sqrt(float64(radio)*float64(radio)-(long_lado*long_lado/4))
    fmt.Println("Longitud del apotema:",apotema)
    area := float64(num_lados)*long_lado*apotema/2
    fmt.Println("Area del polígono regular de",num_lados,",lados, circunscrita en la circunferencia de radio",radio,"cm es",area,"cm cuadrados")
    
    fmt.Println("\n")
    fmt.Println("\n")

// Cálculo conocido el tamaño del lado y número de lados
// https://www.universoformulas.com/matematicas/geometria/apotema/
   
   lado := long_lado    // se asigna el tamaño del lado calculado anteriormente
   atita := 360/num_lados
   mitad_atita := atita/2
   atita2:=float64(mitad_atita)*math.Pi/180    // convertir mitad_atita(grados) a atita2(radianes)
   apotema = float64(lado)/(2*math.Tan((atita2)))
   area = float64(num_lados)*float64(lado)*apotema/2
   fmt.Println("Area del polígono regular de",num_lados,",lados, cada uno de ellos de largo",lado,"cm es",area,"cm cuadrados")
   
 
  
}    
