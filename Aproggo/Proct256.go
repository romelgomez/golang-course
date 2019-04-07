    package main
    import "fmt"
    
    func main() {
      tope := 10
      fmt.Printf("La sumatoria desde 1 hasta %d es: %d\n", tope, sumatoria(tope))
    }
    func sumatoria(tope int) int {
      //Caso base
      if tope == 1 {
        return tope
      }else{
        //Suma tope a lo que regrese la recursión de tope - 1
        return tope + sumatoria(tope - 1)
      }
    }
