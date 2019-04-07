package main
 import "fmt"
 func main(){
 var val1 = 100
 var apuntval1 *int
 apuntval1 = &val1
 fmt.Printf("Valor de la variable val1: %d \n", val1)
 fmt.Printf("Valor de Apuntval1, variable que apunta a la variable val1: %d \n", *apuntval1)
 fmt.Printf("Dirección de almacenamiento de val1 : %x \n", &val1)
 fmt.Printf("Dirección que ocupa el apuntador apuntval1: %x \n", &apuntval1)
 
 fmt.Printf("\n\n")
 
 var val2 = 200
 var apuntval2 *int
 apuntval2 = &val2  // Asigna una dirección válida
 fmt.Printf("Valor de la variable val2: %d \n", val2)
 fmt.Printf("Valor de Apuntval2, variable que apunta a la variable val2: %d \n", *apuntval2)
 fmt.Printf("Dirección de almacenamiento de val2 : %x \n", &val2)
 fmt.Printf("Dirección que ocupa el apuntador apuntval1: %x \n", &apuntval2)
 if(apuntval2 != nil){
    fmt.Println("El apuntador apuntval2 No apunta a nil")
    } else { fmt.Println("El apuntador apuntval2 apunta a nil")
           }
 fmt.Printf("\n\n")
 
 var apuntval3 *int
 if(apuntval3 != nil){
    fmt.Println("El apuntador apuntval3 No apunta a nil")
    } else { fmt.Println("El apuntador apuntval3 apunta a nil")
           }
 fmt.Printf("\n\n")
           
 num := []int{5, 10, 33, 42, 59}
 fmt.Println("Elementos del arreglo num:",num)
 var apuntnum[5]*int
 fmt.Println("El arreglo apuntnum de apuntadores para el arreglo num: ",apuntnum)
                          
 for i:=0; i < 5 ; i++ {
     apuntnum[i] = &num[i]                                    //Asigna las direcciones de los elementos a los apuntadores
     fmt.Printf("-----------------------------\n")
     fmt.Printf("Contenido de num[%d]: %d \n", i, *apuntnum[i])    //Se accede al arreglo num por medio de apuntadores
     fmt.Printf("Dirección de num[%d]: %x \n", i, &apuntnum[i])
 }

 fmt.Printf("\n\n")
 var val4 int = 100
 var apuntval4 *int
 var apunt_apuntval4 **int
 apuntval4 = &val4      //apuntval4 obtiene la dirección de la val4
 apunt_apuntval4 = &apuntval4    //apunt_apuntval4 obtiene la dirección del apuntador apuntval4
 fmt.Printf("Valor original de la variable val4: %d \n", val4)
 fmt.Printf("Valor del apuntador apuntval4 de la variable val4: %d \n", *apuntval4)
 fmt.Printf("Valor del apuntador apunt_apuntval4 del apuntador apuntval4 de la variable val4: %d \n", **apunt_apuntval4)


 }           
  
