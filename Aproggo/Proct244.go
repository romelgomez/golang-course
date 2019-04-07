package main
import "fmt"
func main(){
  //Slice que se declara e inicializa por separado
  var nums1 []int
  nums1 = make([]int, 3, 3)
  //Slice que se declara e inicializa en la misma línea
  nums2:=make([]int, 4)
  fmt.Println(nums1)
  fmt.Println(nums2)
  
  nums1[0] = 5                                    //Modificar el elemento en el índice 0
  fmt.Println("nums1[0]=",nums1[0])                           //Imprimir el elemento en el índice 0   
  fmt.Println("nums1[1]=",nums1[1])
  fmt.Println("nums1[2]=",nums1[2])                            
  //fmt.Println("nums1[3]=",nums1[3])                         //error index out of range  
  
  nums3:=make([]int, 5, 5)                        //Inicializa en 0 los 5 elementos del slice
  nums3[3] = 77                                   //Modifica el índice 3 del slice nums1
  fmt.Println("nums3[]= ", nums3)
  nums4 := []int{9, 8, 7, 6}                      //Nuevo slice nums4 inicializado similar a un array
  fmt.Println("nums4[]= ", nums4)
  
  //Acceder a rangos (sub slices) de slices
  
  fmt.Println("nums4[0]=",nums4[0])                           //Imprimir el elemento en el índice 0   
  fmt.Println("nums4[1]=",nums4[1])
  fmt.Println("nums4[2]=",nums4[2])                            
  fmt.Println("nums4[3]=",nums4[3])
  fmt.Println("nums4[0:1]=", nums4[0:1])
  fmt.Println("nums4[0:2]=", nums4[0:2])
  fmt.Println("nums4[0:3]=", nums4[0:3])
  fmt.Println("nums4[0:4]=", nums4[0:4])
  fmt.Println("nums4[2:4]=", nums4[2:4])
  fmt.Println("nums4[1:2]=", nums4[1:2])
  fmt.Println("nums4[1:3]=", nums4[1:3])
  fmt.Println("nums4[1:4]=", nums4[1:4]) 
  fmt.Println("nums4[2:3]=", nums4[2:3]) 
  fmt.Println("nums4[2:4]=", nums4[2:4])
  fmt.Println("nums4[3:4]=", nums4[3:4])
  fmt.Println("nums4[:]=", nums4[:])
  fmt.Println("nums4[0:]=", nums4[0:])
  fmt.Println("nums4[1:]=", nums4[1:])
  fmt.Println("nums4[2:]=", nums4[2:])
  fmt.Println("nums4[3:]=", nums4[3:])
  fmt.Println("nums4[4:]=", nums4[4:])
  //fmt.Println("nums4[5:]=", nums4[5:])       error  index aut of range
  fmt.Println("nums4[:0]=", nums4[:0])
  fmt.Println("nums4[:1]=", nums4[:1])
  fmt.Println("nums4[:2]=", nums4[:2])
  fmt.Println("nums4[:3]=", nums4[:3])
  fmt.Println("nums4[:4]=", nums4[:4])
  // fmt.Println("nums4[:5]=", nums4[:5])       error index out of range
  
  
  
}
