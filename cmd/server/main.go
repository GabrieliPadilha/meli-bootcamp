package main

import (
  "log"
	"github.com/GabrieliPadilha/meli-bootcamp/cmd/server/handler"
	"github.com/GabrieliPadilha/meli-bootcamp/internal/products"
  "github.com/GabrieliPadilha/meli-bootcamp/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load("../../.env")
    if err != nil {
      log.Fatal("error ao carregar o arquivo .env")
    }
    store := store.Factory("arquivo", "products.json")
    if store == nil {
      log.Fatal("Não foi possivel criar a store")
    }

    repo := products.NewRepository(store)
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
