package main

import (
	"fmt"
	"math"
)

//Forma ... interface
type Forma interface {
	area() float64
}

func escreverArea(f Forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

//Rantagulo ...
type Retangulo struct {
	altura  float64
	largura float64
}

func (r Retangulo) area() float64 {
	return r.altura * r.largura
}

func (c Circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

//Circulo ...
type Circulo struct {
	raio float64
}

func main() {
	fmt.Println("Interfaces")

	r := Retangulo{10, 15}
	escreverArea(r)
	c := Circulo{10}
	escreverArea(c)
}
