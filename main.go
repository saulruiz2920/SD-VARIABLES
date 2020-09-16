package main

import "fmt"
import "math"

func main() {
  squareArea()
  triangleArea()
  circleArea()
  fahrenheitToCelcius()
}

func squareArea() {
  fmt.Println("\nSQUARE")
  var side float64
  fmt.Print("Enter value to calculate Square area: ")
  fmt.Scanf("%g", &side)
  res := math.Pow(side, 2)
  fmt.Println("AREA = ", res)
}

func triangleArea() {
  fmt.Println("\nTRIANGLE")
  var b float32
  var h float32
  fmt.Print("Enter base: ")
  fmt.Scanf("%g", &b)
  fmt.Print("Enter height : ")
  fmt.Scanf("%g", &h)
  res := (b*h) / 2
  fmt.Println("AREA = ", res)
}

func circleArea() {
  fmt.Println("\nCIRCLE")
  var r float64
  fmt.Print("Enter radius: ")
  fmt.Scan(&r)
  res := math.Pi * math.Pow(r, 2)
  fmt.Println("AREA = ", res)
}

func fahrenheitToCelcius() {
  fmt.Println("\nFahrengeit to Celcius")
  var fahrenheit float32
  fmt.Print("Enter fahrenheit: ")
  fmt.Scan(&fahrenheit)
  celcius := (fahrenheit - 32) * 5/9
  fmt.Println("Celcius = ", celcius)
}