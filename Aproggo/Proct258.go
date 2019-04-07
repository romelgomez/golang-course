// fibonacci 2 con función normal
package main
    import "fmt"
    var a,s,nf int = 0,0,0   //variables globales
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
		 a = 0
		 return 0
      }
      //Otro caso base
      if(i == 1) {
		  s=1
        return 1
      }
      nf = a+s
      a=s
      s=nf
	  return nf  
    }
