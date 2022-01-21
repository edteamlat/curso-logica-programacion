package main

import "fmt"

var disponible = map[int]int{
	100: 3,
	50:  6,
	20:  10,
	10:  50,
	1:   50,
}

var denominaciones = []int{100, 50, 20, 10, 1}

func main() {
	valorARetirar := 1034

	if valorARetirar > calcularTotalDisponible() {
		fmt.Println("Error, no tenemos esa cantidad para entregar")
		return
	}

	saldoARetirar := valorARetirar
	for saldoARetirar > 0 {
		for _ , denominacion := range denominaciones {
			cantidad := obtenerCantidadPorDenominacion(denominacion, saldoARetirar)
			saldoARetirar -= cantidad * denominacion
			fmt.Printf("Denominación %d, cantidad %d\n", denominacion, cantidad)
		}

	}
}




func calcularTotalDisponible() int {
	total := 0
	for billete, cantidad := range disponible {
		total += billete * cantidad
	}

	return total
}


func obtenerCantidadPorDenominacion(denominacion, saldoARetirar int) int {
	cantidad := saldoARetirar / denominacion
	if cantidad > disponible[denominacion] {
		cantidad = disponible[denominacion]
	}
	disponible[denominacion] -= cantidad

	return cantidad
}


