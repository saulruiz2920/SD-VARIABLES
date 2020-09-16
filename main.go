package main

import "fmt"
import "math"

func main() {
  getSquareArea()
  getTriangleArea()
  getCircleArea()
}

func getSquareArea() {
  fmt.Println("SQUARE")
  var side float64
  fmt.Print("Enter value to calculate Square area: ")
  fmt.Scanf("%g", &side)
  res := math.Pow(side, 2)
  fmt.Println("AREA = ", res)
}

func getTriangleArea() {
  fmt.Println("TRIANGLE")
  var b float32
  var h float32
  fmt.Print("Enter base: ")
  fmt.Scanf("%g", &b)
  fmt.Print("Enter height : ")
  fmt.Scanf("%g", &h)
  res := (b*h) / 2
  fmt.Println("AREA = ", res)
}

func getCircleArea() {
  fmt.Println("CIRCLE")
  var r float64
  fmt.Print("Enter radius: ")
  fmt.Scanf("%g", &r)
  res := math.Pi * math.Pow(r, 2)
  fmt.Println("AREA = ", res)
}