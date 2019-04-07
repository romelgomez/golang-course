    package main
    import "fmt"
    
   func main() {
	  var tope int
	  tope = 15	
      for n := 0; n < tope; n++ {
        fmt.Printf("| %d ", fibonacci(n))
      }
      fmt.Printf("|\n")
    }

    
    func fibonacci(i int) int {
      //Caso base
      if(i == 0) {
        return 0
      }
      //Otro caso base
      if(i == 1) {
        return 1
      }
      return fibonacci(i-1) + fibonacci(i-2)
    }
