    //un  map cuyo elementos son caracteres con llave entera
    package main
    import "fmt"
    func main(){
      var diasSemana map[int]string
      diasSemana = make(map[int]string)
      diasSemana[1] = "Domingo"
      diasSemana[2] = "Lunes"
      diasSemana[3] = "Martes"
      diasSemana[4] = "Miércoles"
      diasSemana[5] = "Jueves"
      diasSemana[6] = "Viernes"
      diasSemana[7] = "Sábado"
    /*  //Busca elemento con llave 8
      string_dia, encontrado := diasSemana[8]
      if(encontrado){
        fmt.Println("El día 2 es", string_dia)
      }else{
        fmt.Println("El día 2 no fue encontrado")
      }
      //Eliminar el elemento con llave 2
      delete(diasSemana, 2)
      fmt.Println("Tamaño del map:", len(diasSemana))
      */
      
      fmt.Println(diasSemana[2])
    }
