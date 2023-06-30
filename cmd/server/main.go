package main

import (
	"log"
	"os"

	"github.com/GabrieliPadilha/meli-bootcamp/cmd/server/handler"
	"github.com/GabrieliPadilha/meli-bootcamp/config"
	"github.com/GabrieliPadilha/meli-bootcamp/docs"
	"github.com/GabrieliPadilha/meli-bootcamp/internal/products"
	"github.com/GabrieliPadilha/meli-bootcamp/pkg/store"
	"github.com/GabrieliPadilha/meli-bootcamp/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func respondWithError(c *gin.Context, code int, message string) {
  c.AbortWithStatusJSON(code, web.NewResponse(code, nil, message))
}

func TokenAuthMiddleware() gin.HandlerFunc {
  requiredToken := os.Getenv("TOKEN")
 
  // We want to make sure the token is set, bail if not
  if requiredToken == "" {
      log.Fatal("por favor configure a variável de ambiente TOKEN")
  }
 
  return func(c *gin.Context) {
      token := c.GetHeader("token")
     
      if token == "" {
          respondWithError(c, 401, "API token obrigatório")
          return
      }
     
      if token != requiredToken {
          respondWithError(c, 401, "token inválido")
          return
      }
     
      c.Next()
  }
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
  config.InitConfig()

  store := store.Factory("arquivo", "products.json")
  if store == nil {
    log.Fatal("Não foi possivel criar a store")
  }

  repo := products.NewRepository(store)
  service := products.NewService(repo)
  p := handler.NewProduct(service)
  r := gin.Default()

  pr := r.Group("/products")
  {
    pr.Use(TokenAuthMiddleware())

    pr.POST("/", p.Store())
    pr.GET("/", p.GetAll())
    pr.PUT("/:id", p.Update())
    pr.PATCH("/:id", p.UpdateName())
    pr.DELETE("/:id", p.Delete())
  }
  
  docs.SwaggerInfo.Host = os.Getenv("HOST")
  r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    
  r.Run()
}
