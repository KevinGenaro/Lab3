package operaciones

import "errors"

func Sumar( first float64, second float64 ) ( float64 ) {
	return first + second
}

func Restar( first float64, second float64 ) ( float64 ) {
	return first - second
}

func Multiplicar( first float64, second float64 ) ( float64 ) {
	return first * second
}

func Dividir( first float64, second float64 ) ( float64, error ) {
	if second == 0 {
		return -1, errors.New("no se puede dividir por cero")
	}
	return first / second, nil
}
