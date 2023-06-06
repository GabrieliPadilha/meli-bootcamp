package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/GabrieliPadilha/meli-bootcamp/internal/products"
	"github.com/GabrieliPadilha/meli-bootcamp/pkg/web"
	"github.com/gin-gonic/gin"
)
type request struct {
    Name  string  `json:"name"`
    Category  string  `json:"category"`
    Count int     `json:"count"`
    Price float64 `json:"price"`
}
type Product struct {
	service products.Service
}
func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}


// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description Get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (c *Product) GetAll() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        token := ctx.Request.Header.Get("token")
        if token != os.Getenv("TOKEN") {
           ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
           return
        }
   
        p, err := c.service.GetAll()
        if err != nil {
           ctx.JSON(404, web.NewResponse(404, nil, "não há produtudos armazenados"))
           return
        }
      ctx.JSON(200, web.NewResponse(200, p, ""))
   }
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (c *Product) Store() gin.HandlerFunc {
   return func(ctx *gin.Context) {
      token := ctx.Request.Header.Get("token")
      if token != os.Getenv("TOKEN") {
         ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
         return
      }
      var req request
      err := ctx.Bind(&req)
      if  err != nil {
         ctx.JSON(404, gin.H{
            "error": err.Error(),
         })
         return
      }

      if req.Name == "" {
         ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "o nome do produto é obrigatório"))
         return
      }

      if req.Category == "" {
         ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "a categoria do produto é obrigatório"))
         return
      }

      if req.Count == 0 {
         ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "o quantidade é necessária"))
         return
      }

      if req.Price == 0 {
         ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "o preço é obrigatório"))
         return
      }

      p, err := c.service.Store(req.Name, req.Category, req.Count, req.Price)
      if err != nil {
         ctx.JSON(404, gin.H{ "error": err.Error() })
         return
      }
      ctx.JSON(200, p)
   }
}

// Update godoc
// @Summary Update products
// @Tags Products
// @Description Update products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [put]
func (c *Product) Update() gin.HandlerFunc {
   return func(ctx *gin.Context) {
      token := ctx.GetHeader("token")
      if token != os.Getenv("TOKEN") {
         ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
         return
      }

      id, err := strconv.ParseInt(ctx.Param("id"),10, 64) //verifica o tamanho e o tipo do id
      if err != nil {
         ctx.JSON(400, web.NewResponse(400, nil, "ID inválido"))
         return
      }

      var req request
      if err := ctx.ShouldBindJSON(&req); err != nil {
         ctx.JSON(400, gin.H{ "error": err.Error() })
         return
      }
      if req.Name == "" {
         ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "o nome do produto é obrigatório"))
         return
      }
      if req.Category == "" {
         ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "o tipo do produto é obrigatório"))
         return
      }
      if req.Count == 0 {
         ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil,"o quantidade é obrigatória"))
         return
      }
      if req.Price == 0 {
         ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil,"o preço é obrigatório"))

         return
      }
      p, err := c.service.Update(int(id), req.Name, req.Category, req.Count, req.Price)
      if err != nil {
         ctx.JSON(404, gin.H{ "error": err.Error() })
         return
      }
      ctx.JSON(200, p)
   }
}

// UpdateName godoc
// @Summary Update Name
// @Tags Products
// @Description Update products name
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [patch]
func (c *Product) UpdateName() gin.HandlerFunc {
   return func(ctx *gin.Context) {
       token := ctx.GetHeader("token")
       if token != os.Getenv("TOKEN") {
           ctx.JSON(401, gin.H{ "error": "token inválido" })
           return
       }
       id, err := strconv.ParseInt(ctx.Param("id"),10, 64)
       if err != nil {
           ctx.JSON(400, gin.H{ "error":  "invalid ID"})
           return
       }
       var req request
       if err := ctx.ShouldBindJSON(&req); err != nil {
           ctx.JSON(400, gin.H{ "error": err.Error() })
           return
       }
       if req.Name == "" {
           ctx.JSON(400, gin.H{ "error": "O nome do produto é obrigatório"})
           return
       }
       p, err := c.service.UpdateName(int(id), req.Name)
       if err != nil {
           ctx.JSON(404, gin.H{ "error": err.Error() })
           return
       }
       ctx.JSON(200, p)
   }
}

// Delete godoc
// @Summary Delete product
// @Tags Products
// @Description Delete product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [delete]
func (c *Product) Delete() gin.HandlerFunc {
   return func(ctx *gin.Context) {
       token := ctx.GetHeader("token")
       if token != os.Getenv("TOKEN") {
           ctx.JSON(401, gin.H{ "error": "token inválido" })
           return
       }
       id, err := strconv.ParseInt(ctx.Param("id"),10, 64)
       if err != nil {
           ctx.JSON(400, gin.H{ "error":  "invalid ID"})
           return
       }
       err = c.service.Delete(int(id))
       if err != nil {
           ctx.JSON(404, gin.H{ "error": err.Error() })
           return
       }
       ctx.JSON(200, gin.H{ "data": fmt.Sprintf("O produto %d foi removido", id) })
   }
}