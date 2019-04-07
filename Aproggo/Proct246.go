    package main
    import "fmt"
    func main(){
      var alimentos = []string{"huevos", "pescado", "caraota", "maiz"}
      //Inicializar un slice con un segmento del array alimentos
      slice_alimentos:=alimentos[1:3]
      fmt.Println("Arreglo alimentos= ", alimentos)
      fmt.Printf("Arreglo alimentos len = %d cap = %d\n", len(alimentos), cap(alimentos))   //Funciones len() y cap() en acción
      fmt.Println("Slice alimentos[1:3]= ", slice_alimentos)
      fmt.Printf("slice_alimentos len = %d cap = %d\n", len(slice_alimentos), cap(slice_alimentos))
      slice_alimentos = append(slice_alimentos, "piña", "cambur") //Se añaden 2 elementos
      fmt.Println("Slice alimentos[1:3] con dos nuevos elementos= ", slice_alimentos)
      fmt.Printf("slice_alimentos len = %d cap = %d\n", len(slice_alimentos), cap(slice_alimentos))  //Capacidad de slice_alimentos = 6 porque aumenta la capacidad del array que referencia (alimentos)*/
      slice_alimentos = append(slice_alimentos, "arroz", "azucar") //Se añaden 2 elementos
      fmt.Println("Slice alimentos[1:3] con dos nuevos elementos mas = ", slice_alimentos)
      fmt.Printf("slice_alimentos len = %d cap = %d\n", len(slice_alimentos), cap(slice_alimentos))
      fmt.Println("Arreglo alimentos= ", alimentos)
      fmt.Printf("Arreglo alimentos len = %d cap = %d\n", len(alimentos), cap(alimentos))
      slice_alimentos = append(slice_alimentos, "papa", "zanaoria") //Se añaden 2 elementos
      fmt.Println("Slice alimentos[1:3] con dos nuevos elementos mas = ", slice_alimentos)
      fmt.Printf("slice_alimentos len = %d cap = %d\n", len(slice_alimentos), cap(slice_alimentos))
      fmt.Println("Arreglo alimentos= ", alimentos)
      fmt.Printf("Arreglo alimentos len = %d cap = %d\n", len(alimentos), cap(alimentos))
      fmt.Printf("\n\n\n")
      
  nums1:=make([]int, 5, 5) //Nuevo slice
  fmt.Println("nums1 = ", nums1)
  nums2:=nums1[:] //Nuevo slice referenciando nums1
  fmt.Println("Nuevo slide nums2(referenciando nums1) = ", nums2)
  nums2[1] = 99 //Modifica nums2. Los cambios hecho en nums2 se reflejan en nums1
  fmt.Println("nums2 después de cambios su elemento [2] = ",nums2)
  fmt.Println("nums1 después de cambios en nums2 = ",nums1)
  nums3:=make([]int, 5, 5) //nuevo slice
  //Ejemplo de copy()
  copy(nums3, nums1)
  fmt.Println("Nuevo slide nums3 (usando copy() de nums1) = ", nums3)
  nums1[2] = 55
  //Usando copy(), nums3 no referencia a nums1
  fmt.Println("Valores de nums1 después de cambio su elemento [2]= ", nums1)
  fmt.Println("nums3 no cambia al modificar nums1 = ", nums3)
      
    }
