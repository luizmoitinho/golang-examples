package formas

import (
	"testing"
)

type cenarioTesteRetangulo struct {
	Retangulo
	retornoEsperado float64
}

type cenarioTesteCirculo struct {
	Circulo
	retornoEsperado float64
}

func TestArea(t *testing.T) {

	t.Run("Retângulo", TestRetangulo)

	t.Run("Círculo", TestCirculo)

}

func TestCirculo(t *testing.T) {
	cenariosTesteCirculo := []cenarioTesteCirculo{
		{Circulo{float64(10)}, float64(314.1592653589793)},
		{Circulo{float64(-10)}, float64(0)},
		{Circulo{float64(1)}, float64(3.141592653589793)},
		{Circulo{float64(250)}, float64(196349.54084936206)},
		{Circulo{float64(36)}, float64(4071.5040790523717)},
	}

	for _, cenario := range cenariosTesteCirculo {
		areaRecebida := cenario.Area()
		if cenario.retornoEsperado != areaRecebida {
			t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, cenario.retornoEsperado)
		}
	}

}

func TestRetangulo(t *testing.T) {

	cenariosTeste := []cenarioTesteRetangulo{
		{Retangulo{float64(10.0), float64(12.0)}, float64(120.0)},
		{Retangulo{float64(10.0), float64(10.0)}, float64(0.0)},
		{Retangulo{float64(10.0), float64(11.0)}, float64(110.0)},
		{Retangulo{float64(16.0), float64(11.0)}, float64(176.0)},
		{Retangulo{float64(60.0), float64(12.0)}, float64(720.0)},
		{Retangulo{float64(8.0), float64(9.0)}, float64(72.0)},
		{Retangulo{float64(-8.0), float64(9.0)}, float64(-72.0)},
	}

	for _, cenario := range cenariosTeste {
		areaEsperada := cenario.retornoEsperado
		areaRecebida := cenario.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	}

}
