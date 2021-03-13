package formas

import (
	"math"
	"testing"
)

type cenariosTesteRetangulo struct {
	Retangulo
	retornoEsperado float64
}

type cenariosTesteCirculo struct {
	Circulo
	retornoEsperado float64
}

func TestArea(t *testing.T) {

	t.Run("Retângulo", TestRetangulo)

	t.Run("Círculo", func(t *testing.T) {

		circulo := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circulo.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}

	})

}

func TestRetangulo(t *testing.T) {

	cenariosTeste := []cenariosTesteRetangulo{
		{Retangulo{float64(10.0), float64(12.0)}, float64(120.0)},
		{Retangulo{float64(10.0), float64(10.0)}, float64(0.0)},
		{Retangulo{float64(10.0), float64(11.0)}, float64(110.0)},
		{Retangulo{float64(16.0), float64(11.0)}, float64(176.0)},
		{Retangulo{float64(60.0), float64(12.0)}, float64(720.0)},
		{Retangulo{float64(8.0), float64(9.0)}, float64(72.0)},
	}

	for _, cenario := range cenariosTeste {
		areaEsperada := cenario.retornoEsperado
		areaRecebida := cenario.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	}

}


