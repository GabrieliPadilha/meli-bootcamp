package main

import (
    "github.com/gin-gonic/gin"
    "github.com/GabrieliPadilha/meli-bootcamp/cmd/server/handler"
    "github.com/GabrieliPadilha/meli-bootcamp/internal/products"
)

func main() {
    repo := products.NewRepository()
    service := products.NewService(repo)
    p := handler.NewProduct(service)

    r := gin.Default()
    pr := r.Group("/products")
    pr.POST("/", p.Store())
    pr.GET("/", p.GetAll())
    pr.PUT("/:id", p.Update())
    pr.PATCH("/:id", p.UpdateName())
    pr.DELETE("/:id", p.Delete())
    r.Run()
}
