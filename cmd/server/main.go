package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GabrieliPadilha/meli-bootcamp/cmd/server/handler"
	"github.com/GabrieliPadilha/meli-bootcamp/docs"
	"github.com/GabrieliPadilha/meli-bootcamp/internal/products"
	"github.com/GabrieliPadilha/meli-bootcamp/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GetDummyEndpoint(c *gin.Context) {
  resp := map[string]string{"hello":"world"}
  c.JSON(200, resp)
}

func DummyMiddleware(c *gin.Context) {
  fmt.Println("Im a dummy!")
 
  // Pass on to the next-in-chain
  c.Next()
}


// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
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
  r.GET("/dummy", GetDummyEndpoint)
  r.Use(DummyMiddleware)

  docs.SwaggerInfo.Host = os.Getenv("HOST")
  r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

  r.Run()
}
