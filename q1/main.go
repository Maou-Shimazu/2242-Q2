package main

import(
  "fmt"
  "math"
)

type calculation interface {
  volume()       float64
  surfaceArea()  float64
  faceDiagonal() float64
}

type Square struct {
  edge float64
}

func (s Square) volume() float64 {
  return math.Pow(s.edge, 3)
}

func (s Square) surfaceArea() float64 {
  return 6 * math.Pow(s.edge, 2)
}

func (s Square) faceDiagonal() float64 {
  return math.Sqrt(2) * s.edge
}

func main() {
  square := Square { edge: 2 }
  fmt.Println(square.volume())
  fmt.Println(square.surfaceArea())
  fmt.Println(square.faceDiagonal())
}
