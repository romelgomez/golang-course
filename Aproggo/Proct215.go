package main

import (
  "fmt"
  "reflect"
  "time"
)

func main() {
  var now time.Time = time.Now().UTC()
  fmt.Println("now is a type of: ", reflect.TypeOf(now))
  var name string = "Carl Johannes"
  fmt.Println("name is a type of: ", reflect.TypeOf(name))
  var age int = 5
  fmt.Println("age is a type of: ", reflect.TypeOf(age))
  
  t := time.Now()
  hr := t.Hour()
  min := t.Minute()
  sec := t.Second()
  year, month, day := time.Now().Date()
  fmt.Println("t is a type of: ", reflect.TypeOf(t))
  fmt.Println("hr is a type of: ", reflect.TypeOf(hr))
  fmt.Println("min is a type of: ", reflect.TypeOf(min))
  fmt.Println("sec is a type of: ", reflect.TypeOf(sec))
  fmt.Println("year is a type of: ", reflect.TypeOf(year))
  fmt.Println("month is a type of: ", reflect.TypeOf(month))
  fmt.Println("day is a type of: ", reflect.TypeOf(day))
  m := time.Now().Month()
  fmt.Println(m, int(m))  // int(m)  convierte al mes tipo "tiempo" al numérico correspondiente 
  
}
