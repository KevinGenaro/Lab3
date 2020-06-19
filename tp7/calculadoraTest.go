package tp7

import (
	"fmt"

	"github.com/KevinGenaro/LAB3/tp7/operaciones"
)

type Resultados struct {
	rsuma  float64
	rresta float64
	rmulti float64
	rdiv   float64
}

func CalculadoraTest() {

	result, error := operaciones.Dividir(5.0, 4.0)
	if error != nil {
		print(error)
		return
	}
	resultado := Resultados{
		rsuma:  operaciones.Sumar(5.0, 4.0),
		rresta: operaciones.Restar(5.0, 4.0),
		rmulti: operaciones.Multiplicar(5.0, 5.0),
		rdiv:   result,
	}
	fmt.Printf("Suma: ")
	fmt.Print(resultado.rsuma)
	fmt.Printf("\n")
	fmt.Printf("Resta: ")
	fmt.Print(resultado.rresta)
	fmt.Printf("\n")
	fmt.Printf("Multiplicacion: ")
	fmt.Print(resultado.rmulti)
	fmt.Printf("\n")
	fmt.Printf("Dividir: ")
	fmt.Print(resultado.rdiv)
	fmt.Printf("\n")

}
