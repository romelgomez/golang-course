// Programa que determina los dias transcurrido entre dos fechas(dd/mm/yyyy) dadas.   El 14/04/2020 Rudy cumple 24000 dias de vida. hoy 1/12/2018 he vivido 23500 dias
// no se valida la entrada
package main
import (
     "fmt"
     "math"
)
func main() {
  var d1,m1,y1,d2,m2,y2,td1,td2,tm1,tm2,ty1,ty2 float64
  var td float64 = 0   // para sumar dias,meses y años  
  var Mes_Dias = [12] float64{31,28,31,30,31,30,31,31,30,31,30,31}
  var Meses = [12] string {"Enero","Febrero","Marzo","Abril","Mayo","Junio","Julio","Agosto","Septiembre","Octubre","Noviembre","Diciembre"}
    
    fmt.Print("Ingrese datos de la fecha Inicial :\n\n")
    
    fmt.Print("Ingresa día:")
	fmt.Scanln(&d1)
    fmt.Print("Ingresa mes:")
	fmt.Scanln(&m1)
	fmt.Print("Ingresa año:")
	fmt.Scanln(&y1)

    fmt.Print("Ingrese datos de la fecha Final :\n\n")
    
    fmt.Print("Ingresa día:")
	fmt.Scanln(&d2)
    fmt.Print("Ingresa mes:")
	fmt.Scanln(&m2)
	fmt.Print("Ingresa año:")
	fmt.Scanln(&y2)

    if y2==y1 && m2==m1 {       // se procede a dar el total de dias del mismo año y mes
	   td = math.Abs(d2-d1)
	   } else if y2==y1 {                                // se procede a sumar los dias cuando los año son iguales y meses distintos
		         if m2>m1 {     // las siguientes instrucciones son para ordenar los meses con sus dias de forma crecientemente para utilizar el for
					      tm2 = m2-1       // aqui los indice de los meses se resta 1 para adecuarlos al valor cero de los indices cero en los arreglos  
					      tm1 = m1-1
					      td2 = d2
					      td1 = d1
				 } else {
						  tm2 = m1-1
					      tm1 = m2-1
					      td2 = d1
					      td1 = d2
				         }
				 td =   Mes_Dias[int64(tm1)]-td1       // el total de días se inicia con el número de dias del primer mes ordenado (total dias del mes - el dia del mes )
				 if math.Floor(y1/4)*4==y1 && tm1==1 {   // si el año es bisiesto y el primer mes ordenado es febrero se suma 1 día
						td = td + 1
				     	} 
				 for im := tm1+1; im<tm2;im++ {  // se suman los dias de los demás meses intermedio, si los hubieres, menos el último 
					 td = td + Mes_Dias[int64(im)]
					 if math.Floor(y1/4)*4==y1 && im==1 {   // si el año es bisiesto y el siguiente mes intermedio ordenado es febrero se suma 1 día
						td = td + 1
				     	} 
					 }                
                 td = td + td2   // se suma  los dias del mes último
				 } else {                                          // se procede a sumar los dias para cuando los años son distintos
					     if y2>y1 {   // las siguientes instrucciones son para ordenar los años con sus meses y dias de forma crecientemente para utilizar el for
                          ty2 = y2    // se usan las variable temporales 
                          ty1 = y1 
					      tm2 = m2-1
					      tm1 = m1-1
					      td2 = d2
					      td1 = d1
				        } else {
						  ty2 = y1
						  ty1 = y2 	
						  tm2 = m1-1
					      tm1 = m2-1
					      td2 = d1
					      td1 = d2
				         }	
				       
					   td =   Mes_Dias[int64(tm1)]-td1       // el total de días se inicia con el número de dias del primer mes del primer año
				       if math.Floor(ty1/4)*4==y1 && tm1==1 {   // si el año es bisiesto y el primer mes del primer año es febrero se suma 1 día
						  td++
				       } 
				       for im := tm1+1; im<12;im++ {  // se suman los dias de los demás meses siguiente del primer año, si los hubieres. 
					       td = td + Mes_Dias[int64(im)]
					       if math.Floor(ty1/4)*4==y1 && im==1 {   // si el año es bisiesto y uno de los siguiente meses es febrero se suma 1 día
						      td++
				     	   } 
					   }                
                       for iy := ty1+1;iy<ty2;iy++ {   // se suman ahora 365(366 si es bisiesto) dias por cada año intermedio entre las fechas si los hubiren  
						   td = td + 365
						   if math.Floor(iy/4)*4==iy {
							 td++   
					        }
					   }
					   for im := 0; float64(im)<tm2;im++ {  // se suman los dias de los demás meses siguiente correspondiente al año de la otro fecha, si los hubieres. 
					       td = td + Mes_Dias[int64(im)]
					       if math.Floor(ty2/4)*4==ty2 && im==1 {   // si el año es bisiesto y uno de los siguiente meses es febrero se suma 1 día
						      td++
				     	   } 
					   }     
					   td=td+td2
					 }		  
       
	   fmt.Println("El total de días entre La fecha inicial (",d1,"de",Meses[int64(m1-1)],"del",y1,") y La fecha Final (",d2,"de",Meses[int64(m2-1)],"del",y2,") es ",td," días")
    
}
