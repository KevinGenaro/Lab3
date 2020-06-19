package operaciones

const k = 5

type Base struct {
	Operandos []float64
	Operacion string
	Resultado float64
}

func Sumar(operandos ...float64) Base {

	resultado := operandos[0] + operandos[1] + k

	return Base{
		Operandos: operandos,
		Operacion: "Suma",
		Resultado: resultado,
	}
}

func Restar(operandos ...float64) Base {
	resultado := operandos[0] - operandos[1] + k

	return Base{
		Operandos: operandos,
		Operacion: "Resta",
		Resultado: resultado,
	}
}

func Multiplicar(operandos ...float64) Base {
	resultado := operandos[0]*operandos[1] + k

	return Base{

		Operandos: operandos,
		Operacion: "Multiplicacion",
		Resultado: resultado,
	}
}

func Division(operandos ...float64) Base {
	resultado := operandos[0]/operandos[1] + k

	return Base{

		Operandos: operandos,
		Operacion: "Division",
		Resultado: resultado,
	}
}
