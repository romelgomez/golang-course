package main
  import "fmt"
  var resglobal int   // Variable global
  func main(){
    var resultado int
    resglobal = 31
    resultado = mult(20, 178)
    fmt.Printf("Resultado de multiplicación: %d\n", resultado)
}
    func mult(numero1, numero2 int ) (int){            // numero1 y numero2 son parámetros formales y su ámbito es local a la función mult
         var resultado int                          // Variable local de la función mult con el mismo nombre que una variable local de main
         resultado = numero1 * numero2                
         fmt.Printf("Valor en variable global: %d\n", resglobal)   //Accediendo a variable global desde función
    return resultado
    }
