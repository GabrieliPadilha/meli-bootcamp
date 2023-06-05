package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/GabrieliPadilha/meli-bootcamp/internal/products"
    "strconv"
    "fmt"
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

func (c *Product) GetAll() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        token := ctx.Request.Header.Get("token")
        if token != "123456" {
           ctx.JSON(401, gin.H{
              "error": "token inválido",
           })
           return
        }
   
        p, err := c.service.GetAll()
        if err != nil {
           ctx.JSON(404, gin.H{
              "error": err.Error(),
           })
           return
        }
        ctx.JSON(200, p)
     }
   }

func (c *Product) Store() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        token := ctx.Request.Header.Get("token")
        if token != "123456" {
           ctx.JSON(401, gin.H{ "error": "token inválido" })
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
        p, err := c.service.Store(req.Name, req.Category, req.Count, req.Price)
        if err != nil {
           ctx.JSON(404, gin.H{ "error": err.Error() })
           return
        }
        ctx.JSON(200, p)
     }
   }

func (c *Product) Update() gin.HandlerFunc {
   return func(ctx *gin.Context) {
      token := ctx.GetHeader("token")
      if token != "123456" {
         ctx.JSON(401, gin.H{ "error": "token inválido" })
         return
      }

      id, err := strconv.ParseInt(ctx.Param("id"),10, 64) //verifica o tamanho e o tipo do id
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
      if req.Category == "" {
         ctx.JSON(400, gin.H{ "error": "O tipo do produto é obrigatório"})
         return
      }
      if req.Count == 0 {
         ctx.JSON(400, gin.H{ "error": "A quantidade é obrigatória"})
         return
      }
      if req.Price == 0 {
         ctx.JSON(400, gin.H{ "error": "O preço é obrigatório"})
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

func (c *Product) UpdateName() gin.HandlerFunc {
   return func(ctx *gin.Context) {
       token := ctx.GetHeader("token")
       if token != "123456" {
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

func (c *Product) Delete() gin.HandlerFunc {
   return func(ctx *gin.Context) {
       token := ctx.GetHeader("token")
       if token != "123456" {
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