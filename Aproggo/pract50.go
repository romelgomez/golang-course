package main

import(
	"fmt"
)

func main() {
	
	menu := 
`
	Bienvenido
	[ 1 ] Pizza
	[ 2 ] Tacos
	¿Qué prefieres?
`
	fmt.Print(menu)

	var eleccion int //Declarar variable y tipo antes de escanear, esto es obligatorio
	fmt.Scanln(&eleccion)
	//fmt.Println(eleccion)

	switch eleccion{
		case 1:
			fmt.Println("Prefieres pizza")
		case 2:
			fmt.Println("Prefieres tacos")
		default:
			fmt.Println("No prefieres ninguno de ellos")
	}

}
