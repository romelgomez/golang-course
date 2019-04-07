package main
  import "fmt"
  func main(){
    var var1 int = 10
    fmt.Println("Valor de la variable var1 antes de llamados a funciones: ", var1)

    sumaPorValor(var1)  //Llamado por valor
    fmt.Println("Valor de la variable var1 después de llamado por valor: ", var1)

    sumaPorReferencia(&var1)   //Llamado por referencia
    fmt.Println("Valor de la variable var1 después de llamado por referencia: ", var1)
 }
 
    func sumaPorValor(a int){
         a = a + 10
         }
      
    func sumaPorReferencia(a *int){
         *a = *a + 10
         }
