    package main
    import "fmt"
    //Registro de materias
    type Materia struct{
      nombre string
      creditos int
    }
    func main(){
		
	/*
	  var materias map[string]Materia
	  materias = make(map[string]Materia)
	  materias["COMP"] = Materia{"Compiladores",12}
	  materias["GO"] = Materia{"Golang",8}
	  materias["INPROG"] = Materia{"Introducción a la Programación",6}
	   */ 	
			
      //Mapa con registros inicializado al declarar
      var materias = map[string]Materia{
        "COMP": Materia{
          "Compiladores", 12,
        },
        "GO": Materia{
          "Golang", 8,
        },
        "INPROG": Materia{
          "Introducción Programación", 6,
        },
      }
      //Imprime materias
      
      fmt.Println(materias)
      
      for materia := range materias{
        fmt.Println("Código:", materia, "| Nombre:", materias[materia].nombre,
          "| Créditos:", materias[materia].creditos)
      }
      fmt.Printf("\n\n")
      //Añade un elemento nuevo
      materias["POO"] = Materia{"Programación Orientada a Objetos", 7,}
      //Imprime materias
      for materia := range materias{
        fmt.Println("Código:", materia, "| Nombre:", materias[materia].nombre,
          "| Créditos:", materias[materia].creditos)
      }
      fmt.Printf("\n\n")
      //Busca materia con llave única "COMP"
      mat, ok := materias["COMP"]
      if(ok){
        fmt.Println("Nombre materia:", mat.nombre, "| Créditos:", mat.creditos)
      }else{
        fmt.Println("La materia no existe")
      }
    }
