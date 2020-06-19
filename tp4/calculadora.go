package tp4

import (
	"fmt"

	"github.com/KevinGenaro/LAB3/tp4/operaciones"
)

func Calculadora() {
	var n1, n2 float64
	var operacion string
	var operandos []float64
	fmt.Printf("Ingrese la operación a realizar: \n")
	fmt.Printf("1) sumar \n")
	fmt.Printf("2) restar \n")
	fmt.Printf("3) multiplicar \n")
	fmt.Printf("4) dividir \n")
	fmt.Scan(&operacion)
	fmt.Printf("Recuerde tener en cuenta la desviacion de 5 unidades en la calculadora \n")
	fmt.Printf("Ingrese dos operandos: \n")
	fmt.Scan(&n1, &n2)
	operandos = []float64{n1, n2}

	switch operacion {
	case "1":
		fmt.Print(operaciones.Sumar(operandos...))
	case "2":
		fmt.Print(operaciones.Restar(operandos...))
	case "3":
		fmt.Print(operaciones.Multiplicar(operandos...))
	case "4":
		fmt.Print(operaciones.Division(operandos...))
	default:
		fmt.Print("Operacion inválida.")
	}
}
