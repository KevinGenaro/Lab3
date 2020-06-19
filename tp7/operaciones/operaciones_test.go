package operaciones

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumarOk(t *testing.T) {
	// GIVEN
	a := 5.0
	b := 5.0
	// WHEN
	resultado := Sumar(a, b)
	// THEN
	assert.Equal(t, 10.0, resultado, fmt.Sprintf("Suma incorrecta!"))
}

func TestRestarOk(t *testing.T) {
	// GIVEN
	a := 5.0
	b := 5.0
	// WHEN
	resultado := Restar(a, b)
	// THEN
	assert.Equal(t, 0.0, resultado, fmt.Sprintf("Resta incorrecta!"))
}

func TestMultiplicarOk(t *testing.T) {
	// GIVEN
	a := 5.0
	b := 5.0
	// WHEN
	resultado := Multiplicar(a, b)
	// THEN
	assert.Equal(t, 25.0, resultado, fmt.Sprintf("Multiplicacion incorrecta!"))
}

func TestDividirOk(t *testing.T) {
	// GIVEN
	a := 5.0
	b := 5.0
	// WHEN
	resultado, err := Dividir(a, b)
	// THEN
	assert.Equal(t, 1.0, resultado, fmt.Sprintf("Division incorrecta!"))
	assert.Nil(t, err, fmt.Sprintf("Fallo - error no nulo!!"))
}

/*func TestDividirPorCero(t *testing.T) {
	// GIVEN
	a := 5.0
	b := 0.0
	// WHEN
	resultado, err := Dividir(a, b)
	// THEN
	assert.Equal(t, -1.0, resultado, fmt.Sprintf("Division incorrecta!"))
	assert.NotNil(t, err, fmt.Sprintf("Division incorrecta  - error nil!!"))
	assert.Equal(t, "no se divide por 0", err.Error(), fmt.Sprintf("Mensaje de error fallido!!"))
}*/
