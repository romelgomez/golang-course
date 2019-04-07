    package main
    import "fmt"
    type Pelicula struct{
      titulo string
      director string
      protagonistas string
      año int
    }
    func main(){

      /*Declaración de variables tipo Pelicula*/
      var pelis[2] Pelicula

      /*Llenamos cada variable con la información correspondiente para cada película*/
      pelis[0].titulo = "Titanic"
      pelis[0].director = "James Cameron"
      pelis[0].protagonistas = "Leonardo DiCaprio,Kate Winslet"
      pelis[0].año = 1997
      pelis[1].titulo = "Soy leyenda"
      pelis[1].director = "Francis Lawrence"
      pelis[1].protagonistas = "Will Smith"
      pelis[1].año = 2007
      
      /*Imprimimos la información de las películas*/
      imprimePelicula(pelis[0])
      imprimePelicula(pelis[1])
    }
    
    //Función que imprime la información de una película
    func imprimePelicula(peli Pelicula){
      fmt.Printf("NOMBRE: %s\n", peli.titulo)
      fmt.Printf("DIRECTOR: %s\n", peli.director)
      fmt.Printf("PROTAGONISTAS: %s\n", peli.protagonistas)
      fmt.Printf("AÑO: %d\n\n", peli.año)
    }
