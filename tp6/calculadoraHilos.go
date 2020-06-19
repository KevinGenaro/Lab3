package tp6

import (
	"fmt"
	"time"
)

type Base struct {
	Operandos []float64
	Operacion string
	Resultado float64
}

var channel1 chan Base = make(chan Base)
var channel2 chan Base = make(chan Base)
var channel3 chan Base = make(chan Base)
var channel4 chan Base = make(chan Base)

func CalculadoraHilos() {

	var num1, num2 float64
	var operandos []float64

	fmt.Printf("Ingrese dos operandos: \n")
	fmt.Scan(&num1, &num2)
	operandos = []float64{num1, num2}

	go func() {

		resultado := operandos[0] + operandos[1]

		channel1 <- Base{

			Operandos: operandos,
			Operacion: "suma",
			Resultado: resultado,
		}
		time.Sleep(time.Second * 2)

	}()

	go func() {
		resultado := operandos[0] - operandos[1]

		channel2 <- Base{

			Operandos: operandos,
			Operacion: "resta",
			Resultado: resultado,
		}
		time.Sleep(time.Second * 2)
	}()

	go func() {

		resultado := operandos[0] * operandos[1]

		channel3 <- Base{

			Operandos: operandos,
			Operacion: "multiplicacion",
			Resultado: resultado,
		}
		time.Sleep(time.Second * 2)

	}()

	go func() {

		resultado := operandos[0] / operandos[1]

		channel4 <- Base{

			Operandos: operandos,
			Operacion: "Division",
			Resultado: resultado,
		}
		time.Sleep(time.Second * 2)

	}()

	go func() {
		for {
			select {
			case suma := <-channel1:
				fmt.Println(suma)
			case resta := <-channel2:
				fmt.Println(resta)
			case multiplicar := <-channel3:
				fmt.Println(multiplicar)
			case dividir := <-channel4:
				fmt.Println(dividir)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
