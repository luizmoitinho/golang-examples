package formas

import (
	"math"
)

//Forma ... interface
type Forma interface {
	area() float64
}

//Rantagulo ...
type Retangulo struct {
	Altura  float64
	Largura float64
}

//Area ... cálculo da area do retângulo
func (r Retangulo) Area() float64 {
	if r.Altura == r.Largura || (r.Altura <= 0 && r.Largura <= 0) {
		return float64(0.0)
	}
	return r.Altura * r.Largura
}

//Area ... cálculo da área do círculo
func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

//Circulo ...
type Circulo struct {
	Raio float64
}

func main() {

}
