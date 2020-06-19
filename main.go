package main

import (
	"fmt"

	"github.com/KevinGenaro/LAB3/tp3"
	"github.com/KevinGenaro/LAB3/tp4"
	"github.com/KevinGenaro/LAB3/tp5"
	"github.com/KevinGenaro/LAB3/tp6"
	"github.com/KevinGenaro/LAB3/tp7"
	"github.com/KevinGenaro/LAB3/tp8"
)

func main() {
	var numeroTP int64
	fmt.Printf("Ingrese el numero del TP que quiere ejecutar: \n")
	fmt.Printf("3): Introduccion a Go \n")
	fmt.Printf("4): Variables y estructuras en Go \n")
	fmt.Printf("5): Consumo de APIs en Go \n")
	fmt.Printf("6): Concurrencia y Paralelirmos en Go \n")
	fmt.Printf("7): Testing \n")
	fmt.Printf("8): Creacion de una API \n\n ")
	fmt.Scan(&numeroTP)
	switch numeroTP {
	case 3:
		tp3.HelloWord()
		break
	case 4:
		tp4.Calculadora()
		break
	case 5:
		tp5.ConsumoApi()
		break
	case 6:
		tp6.CalculadoraHilos()
		var input string
		fmt.Scanln(&input)
		break
	case 7:
		tp7.CalculadoraTest()
		break
	case 8:
		tp8.CalculadoraApi()
		break
	}
}
