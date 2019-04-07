// un mapa con elementos perteneciente a una estructura y con llave tipo string(cadeana)
    package main
    import "fmt"
        
    type Materia struct{                   //Registro o estructura de materias
      nombre string
      creditos int
    }
        
    
    func main(){
	
	   var materias map[string]Materia
	   materias = make(map[string]Materia)
	   materias["COMP"] = Materia{"Compiladores",12}
	   materias["GO"] = Materia{"Golang",8}
	   materias["INPROG"] = Materia{"Introducción a la Programación",6}
      //Imprime materias
      for m := range materias{
        fmt.Println("Código:", m, "| Nombre:", materias[m].nombre,
          "| Créditos:", materias[m].creditos)
      }
      fmt.Printf("\n\n")
      //Añade un elemento nuevo
      materias["POO"] = Materia{"Programación Orientada a Objetos", 7}
      //Imprime materias
      for m := range materias{
        fmt.Println("Código:", m, "| Nombre:", materias[m].nombre,
          "| Créditos:", materias[m].creditos)
      }
      fmt.Printf("\n\n")
      
      //Busca materia con llave única "COMP"
      mat, ok := materias["COMP"]
      if(ok){
        fmt.Println("Nombre materia:", mat.nombre, "| Créditos:", mat.creditos)
      }else{
        fmt.Println("La materia no existe")
      }
      
      //Busca materia con llave única "PEPE"
      mat, ok = materias["PEPE"]
      if(ok){
        fmt.Println("Nombre materia:", mat.nombre, "| Créditos:", mat.creditos)
      }else{
        fmt.Println("La materia no existe")
      }
      
    }
