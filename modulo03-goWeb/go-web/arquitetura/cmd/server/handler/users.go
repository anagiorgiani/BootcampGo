package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/anagiorgiani/BootcampGo/modulo03-goWeb/go-web/arquitetura/internal/users"
	"github.com/anagiorgiani/BootcampGo/modulo03-goWeb/go-web/arquitetura/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Id          int    `json:"id"`
	Nome        string `json:"nome" binding:"required"`
	Sobrenome   string `json:"sobrenome" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Idade       int    `json:"idade" binding:"required"`
	Ativo       bool   `json:"ativo" binding:"required"`
	DataCriacao string `json:"dataCriacao"`
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (u *User) FilterUsers(c *gin.Context) {
	paramNome := c.Query("nome")
	paramSobrenome := c.Query("sobrenome")

	userFiltered, err := u.service.FilterUsers(paramNome, paramSobrenome)

	if err != nil {
		c.JSON(404, web.NewResponse(404, nil, err.Error()))
	}

	c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, userFiltered, ""))
}

// StoreUsers godoc
// @Sumary Store Users
// @Tags Users
// @Description Store Users
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param user body request true "User to store"
// @Sucess 200 {object} web.Response
// @Router /users [get]
func (u *User) Get(c *gin.Context) {

	if c.Query("id") != "" {
		id, errConv := strconv.Atoi(c.Query("id"))

		if errConv != nil {
			c.JSON(404, web.NewResponse(404, nil, errConv.Error()))

		}

		users, err := u.service.Get(id)

		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
		}

		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, users, ""))
	} else {
		users, err := u.service.GetAll()
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, users, ""))
	}

}

// StoreUsers godoc
// @Sumary Store Users
// @Tags Users
// @Description Store Users
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param user body request true "User to store"
// @Sucess 200 {object} web.Response
// @Router /users [post]
func (u *User) CreateUser(c *gin.Context) {

	if validateToken(c) == false {
		c.JSON(401, web.NewResponse(401, nil, "token invalido"))
		return
	}

	var req request

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	addedUser, err := u.service.CreateUser(req.Nome, req.Sobrenome, req.Email, req.Idade, req.Ativo)

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(202, web.NewResponse(202, addedUser, ""))

}

func validateToken(c *gin.Context) bool {
	token := c.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		return false
	}
	return true
}

func (c *User) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "1234" {
			ctx.JSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "invalido id"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Nome == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "o nome do user e obrigatorio "))
		}
		if req.Sobrenome == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "o sobrenome do user e obrigatorio "))
		}
		if req.Email == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "o email do user e obrigatorio "))
		}
		if req.Idade == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "a idade do user e obrigatorio "))
		}

		u, err := c.service.UpdateUser(int(id), req.Nome, req.Sobrenome, req.Email, req.Idade, req.Ativo)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}

}

func (c *User) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "1234" {
			ctx.JSON(401, gin.H{"error": "token invalido"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalido id"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Nome == " " {
			ctx.JSON(400, gin.H{"error": "o nome do user e obrigatorio "})
		}

		u, err := c.service.UpdateName(int(id), req.Nome)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}

}

func (c *User) DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "1234" {
			ctx.JSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalido id"})
			return
		}
		err = c.service.DeleteUser(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("O user %d foi deletado", id), ""))
	}

}
