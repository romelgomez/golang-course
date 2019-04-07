// Programa que genera el calendario de un mes de un año determinado. No se valida la entrada
package main
import (
     "fmt"
     "math"
    s "strings"
     "time"
)
func main() {
  var mm,yy float64
  var dd float64 = 1
  var NMESES = [] string {"ENERO","FEBRERO","MARZO","ABRIL","MAYO","JUNIO","JULIO","AGOSTO","SEPTIEMBRE","OCTUBRE","NOVIEMBRE","DICIEMBRE"}
  var NUDM = [12] int64{31,28,31,30,31,30,31,31,30,31,30,31}
  var NODS = [] string {"Domingo","Lunes","Marte","Miercoles","Jueves","Viernes","Sabado"}
  fmt.Println("Ingrese datos del mes y año del almanaque:\n\n")
    
  fmt.Print("Ingresa mes(numérico):")
  fmt.Scanln(&mm)
  fmt.Print("Ingresa año:")
  fmt.Scanln(&yy)

  a := math.Floor(float64(14-mm)/12)              // las siguientes 4 lineas determinan el día(numérico) en que se inicia el almanaque 0 lunes ...6 Domingo
  y := float64(yy)-a
  m := float64(mm)+float64(12*a)-2
  d := int32(dd + y + math.Floor(y/4) - math.Floor(y/100) + math.Floor(y/400) + math.Floor((31*m)/12))%7  
  var vipd int64
  var tdia string = NODS[d]
  for i:=0;i<7;i++ {        // for  para determinar la columna donde se inicia el mes del almanaque
	  if tdia==NODS[i] {
		  vipd=int64(i)    // valor del indice del día en se inicia el mes del almanaque
	  } 
  }
    
  var MS [8][8] int64         // matriz que almacenará los elementos del almanaque(automaticamente iniciada todos sus elementos en cero 
                              // no se utilizará la fila 0 ni la columna 0 
  var pd int64 = 1             // valores numéricos de los dias del almanaque
  var pv int64 = vipd          // posición de la matriz en el que se inicia el almanaque 
  var ban1 bool = true
  var tope int64 = NUDM[int(mm-1)]    // tope hasta donde variará los números que se corresponden los números de los dias del mes
  if math.Floor(float64(yy)/4)*4==float64(yy)&&mm==2  {       // para chequear que  si el año es bisiesto y el mes es febrero el tope se incrementa a 29 
		tope++
  } 
                                       // construcción de la matriz del mes-calendario
  for i:=1;i<=7;i++ {
	  for j:=pv+1;j<=7;j++ {
	      pv=0
	      MS[i][j]=pd    
	      pd++	  
		  if pd>tope{
			 ban1=false
			 break 
		  }
       }
	  if ban1==false {break}	  
    }                                     
                                        // Impresión de la matriz del mes calendario  
    var ban2 bool = true                // controlar las salidas cuando todos los valores de una fila son ceros  
    fmt.Println("\n\n           **",NMESES[int(mm-1)],"**\n\n")
    fmt.Println(s.Repeat("-",36))
    fmt.Println("|   D|   L|   M|   M|   J|   V|   S|")
    fmt.Println(s.Repeat("-",36))
   for i:=1;i<=6;i++ {
       for j:=1;j<=7;j++ {
		   if i>1 && j==1 && MS[i][j]==0 {
			   ban2=false
			   break
		   }	   
		   if MS[i][j]!=0 {
              fmt.Printf("|%4d",MS[i][j])  
            } else { fmt.Print("|    ")} 
        }     
    if ban2==false {break}    
    fmt.Printf("|\n") 
    fmt.Println(s.Repeat("-",36))
   }
    t := time.Now()
    hr := t.Hour()
    min := t.Minute()
    sec := t.Second()
    year, month, day := time.Now().Date()
    md := "AM"
    if hr>12{
		hr=hr-12
		md="PM" }
    mi:=int(month) // asigna mi  el entero numerico correspondiente a mes ya que month contiene el mes en inglés
    fmt.Println("Fecha de hoy  :",day,"de",NMESES[mi-1],"del",year)
    fmt.Println("Hora actual   :",hr,":",min,":",sec," ",md)
   
 }
