    package main
    import "fmt"
    //Definición de struct Pelicula
    type Pelicula struct{
      titulo string
      director string
      protagonistas string
      año int // Go permite utilizar 'ñ' en identificadores
    }
    func main(){
      /*Declaración de variables tipo Pelicula*/
      var peli01 Pelicula
      var peli02 Pelicula
      /*Llenamos cada variable con la información
      correspondiente para cada película*/
      
      peli01.titulo = "Titanic"
      peli01.director = "James Cameron"
      peli01.protagonistas = "Leonardo DiCaprio,Kate Winslet"
      peli01.año = 1997
      
      peli02.titulo = "Soy leyenda"
      peli02.director = "Francis Lawrence"
      peli02.protagonistas = "Will Smith"
      peli02.año = 2007
      
      /*Imprimimos la información de las películas*/
      
      fmt.Printf("%-20s%-20s%-35s%-6s\n","NOMBRE","DIRECTOR","PROTAGONISTAS","AÑO")
      
      fmt.Printf("%-20s", peli01.titulo)
      fmt.Printf("%-20s", peli01.director)
      fmt.Printf("%-35s", peli01.protagonistas)
      fmt.Printf("%-6d\n", peli01.año)
      
      fmt.Printf("%-20s", peli02.titulo)
      fmt.Printf("%-20s", peli02.director)
      fmt.Printf("%-35s", peli02.protagonistas)
      fmt.Printf("%-6d\n", peli02.año)
      
   }
