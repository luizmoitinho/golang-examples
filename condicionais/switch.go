package main

import "fmt"

func getDiaSemana(numero int) string {

	switch numero {
	case 0:
		return "Domingo"
	case 1:
		return "Segunda-feira"
	case 2:
		return "Terça-feira"
	case 3:
		return "Quarta-feira"
	case 4:
		return "Quinta-feira"
	case 5:
		return "Sexta-feira"
	case 6:
		return "Sábado-feira"
	default:
		return "Dia inválido'"

	}

}

func getDiaSemana2(numero int) string {
	var diaSemana string

	switch {
	case numero == 0:
		diaSemana = "Domingo"
		fallthrough
	case numero == 1:
		diaSemana = "Segunda-feira"
	case numero == 2:
		diaSemana = "Terça-feira"
	case numero == 3:
		diaSemana = "Quarta-feira"
	case numero == 4:
		diaSemana = "Quinta-feira"
	case numero == 5:
		diaSemana = "Sexta-feira"
	case numero == 6:
		diaSemana = "Sábado-feira"
	default:
		diaSemana = "Dia inválido'"

	}
	return diaSemana

}

func main() {
	fmt.Println("Switch")

	fmt.Println("Dia da semana: ", getDiaSemana2(0))
}
