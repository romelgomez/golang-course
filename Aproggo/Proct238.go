    package main
    import "fmt"
    //Definición
    type Pelicula struct{
      titulo string
      director string
      protagonistas string
      año int
    }
    func main(){

      /*Declaración de arreglo tipo Pelicula*/
      var pelis[2] Pelicula

      /*Llenamos cada índice con la información correspondiente para cada película*/
      pelis[0].titulo = "Titanic"
      pelis[0].director = "James Cameron"
      pelis[0].protagonistas = "Leonardo DiCaprio,Kate Winslet"
      pelis[0].año = 1997
      pelis[1].titulo = "Soy leyenda"
      pelis[1].director = "Francis Lawrence"
      pelis[1].protagonistas = "Will Smith"
      pelis[1].año = 2007

      /*Imprimimos la información de las películas*/
      
      fmt.Printf("%-20s%-20s%-35s%-6s\n","NOMBRE","DIRECTOR","PROTAGONISTAS","AÑO")

      fmt.Printf("%-20s", pelis[0].titulo)
      fmt.Printf("%-20s", pelis[0].director)
      fmt.Printf("%-35s", pelis[0].protagonistas)
      fmt.Printf("%-6d\n", pelis[0].año)
      
      fmt.Printf("%-20s", pelis[1].titulo)
      fmt.Printf("%-20s", pelis[1].director)
      fmt.Printf("%-35s", pelis[1].protagonistas)
      fmt.Printf("%-6d\n", pelis[1].año)

    }
