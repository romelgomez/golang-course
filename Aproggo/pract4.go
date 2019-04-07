// Go ofrece un excelente soporte para el formateo de cadenas en la tradición printf. A continuación se muestran algunos ejemplos de tareas de formato de cadena comunes.
package main

	

import "fmt"
import "os"

	

type punto struct {
    x, y int
}

	

func main() {

    p := punto{1, 2}
    fmt.Printf("%v\n", p)  //  Esto imprime una instancia de la estructura punto 

    fmt.Printf("%+v\n", p) // Si el valor es una estructura, la variante %+v incluirá los nombres de campo de la estructura.

    fmt.Printf("%#v\n", p) // La variante %#v imprime una representación de sintaxis go del valor, es decir, el fragmento de código fuente que produciría ese valor.
	
	fmt.Printf("%T\n", p)  // Para imprimir el tipo de un valor, utilice %T
	
	fmt.Printf("%t\n", true)  // El formato booleano es directo.

	fmt.Printf("%d\n", 123)  // Hay muchas opciones para formatear enteros. Utilice %d para formato estándar, base-10.
	
	fmt.Printf("%b\n", 14)  // Esto imprime una representación binaria.
	
	fmt.Printf("%c\n", 33)  // Esto imprime el carácter correspondiente al entero dado.
	
	fmt.Printf("%x\n", 456)  //% x proporciona codificación hexadecimal.
	
	fmt.Printf("%f\n", 78.9)  //También hay varias opciones de formato para los flotxx. Para el uso de formato decimal básico %f.
	
	fmt.Printf("%e\n", 123400000.0) //  %e y %E formatean los flotxx en (versiones ligeramente diferentes de) notación científica.
	fmt.Printf("%E\n", 123400000.0)

    fmt.Printf("%s\n", "\"string\"") // Para la impresión básica de cadenas Utilice %s.

    fmt.Printf("%q\n", "\"string\"") // To double-quote strings as in Go source, use %q. Para doble comillas en cadenas como en el codigo de Go, utilice %q.

    fmt.Printf("%x\n", "hex this") // Al igual que con los enteros vistos anteriormente, %x representa la cadena en la base-16, con dos caracteres de salida por byte de entrada.

    fmt.Printf("%p\n", &p) //Para imprimir una representación de un puntero, utilice %p.

    fmt.Printf("|%6d|%6d|\n", 12,345) // Al formatear números, a menudo querrá controlar el ancho y la precisión de la expresión resultante.
                                       // Para especificar el ancho de un entero, utilice un número después de % en el verbo.
                                       // Por defecto el resultado será justificado a la derecha y rellenado con espacios.
	
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45) //También puede especificar el ancho de los flotxx, aunque por lo general también desea 
                                             //restringir la precisión de los decimales al mismo tiempo con la sintaxis de precisión de ancho.
	
    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) //Para justificar la izquierda, utilice el indicador bandera -.

    fmt.Printf("|%6s|%6s|\n", "foo", "b") // También es posible que desee controlar el ancho al formatear cadenas,
                                          // especialmente para asegurarse de que se alineen en una salida similar a una tabla. Para la anchura justificada a la derecha.

    fmt.Printf("|%-6s|%-6s|\n", "foo", "b") // Par justificar a la  izquierda utilice el indicador bandera - como con los números.

    s := fmt.Sprintf("a %s", "string")  // Hasta ahora hemos visto printf, que imprime la cadena formateada en el os.Stdout. Sprintf formatea y devuelve una cadena sin imprimirla en ningún sitio.
    fmt.Println(s)

    fmt.Fprintf(os.Stderr, "an %s\n", "error") // You can format+print to io.Writers other than os.Stdout using Fprintf. Puede formatear la impresión en IO.impresoras que no sean os.Stdout usando fprintf.
}
