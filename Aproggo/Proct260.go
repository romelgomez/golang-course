    package main
    import "fmt"
    /* Declaración de una interfaz tipo FiguraGeométrica*/
    type FiguraGeometrica interface {
      area() float64
      //Aquí podrían ir más métodos
    }
    /* Declaración de registro tipo Rectangulo */
    type Rectangulo struct {
      base, altura float64
    }
    /* Declaración de registro tipo Trapecio */
    type Trapecio struct {
      base_mayor, base_menor, altura float64
    }
    
    func main() {
    
      rectangulo := Rectangulo {base: 4, altura: 7.5}
    
      trapecio := Trapecio {base_mayor: 5, base_menor: 2, altura: 3}
    
      /* Llamado a función dameArea() */
      fmt.Printf("Area del rectángulo: %f\n", dameArea(rectangulo))
   
      /* Uso del registro para llamar a area() */
      fmt.Printf("Area del trapecio: %f\n", trapecio.area())
    }
    
    /* Definición de un método para Figura */
    func dameArea(fig FiguraGeometrica) float64 {
      return fig.area()
    }
    
    
    /* Método para los registros Rectangulo implementando la interfaz */
    func(figura Rectangulo) area() float64 {
      return figura.base * figura.altura
    }
    /* Método para los registros Trapecio implementando la interfaz */
    func(figura Trapecio) area() float64 {
      return (figura.base_mayor + figura.base_menor) * figura.altura / 2
    }
    
