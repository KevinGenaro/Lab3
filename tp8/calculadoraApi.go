package tp8

import (
"github.com/KevinGenaro/LAB3/tp8/controller"
"github.com/gin-gonic/gin"
)

func CalculadoraApi() {
	r := gin.Default()
	r.GET("/calc/sum", controller.Sumar)
	r.GET("/calc/res", controller.Restar)
	r.GET("/calc/mult", controller.Multiplicar)
	r.GET("/calc/div", controller.Dividir)
	r.Run(":8080")
}