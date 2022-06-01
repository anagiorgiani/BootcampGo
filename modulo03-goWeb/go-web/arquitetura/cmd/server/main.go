package main

import (
	"github.com/anagiorgiani/BootcampGo/modulo03-goWeb/go-web/arquitetura/cmd/server/handler"
	"github.com/anagiorgiani/BootcampGo/modulo03-goWeb/go-web/arquitetura/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := users.NewRepository()
	service := users.NewService(repo)
	u := handler.NewUser(service)

	r := gin.Default()
	us := r.Group("/users")

	us.POST("/", u.CreateUser)
	us.GET("/", u.Get)
	us.GET("/filter", u.FilterUsers)
	us.PUT("/id", u.UpdateUser())
	us.PATCH("/id", u.UpdateName())
	us.DELETE("/id", u.DeleteUser())

	r.Run()
}
