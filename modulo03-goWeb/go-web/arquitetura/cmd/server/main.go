package main

import (
	"os"

	"github.com/anagiorgiani/BootcampGo/modulo03-goWeb/go-web/arquitetura/cmd/server/docs"
	"github.com/anagiorgiani/BootcampGo/modulo03-goWeb/go-web/arquitetura/cmd/server/handler"
	"github.com/anagiorgiani/BootcampGo/modulo03-goWeb/go-web/arquitetura/internal/users"
	"github.com/anagiorgiani/BootcampGo/modulo03-goWeb/go-web/arquitetura/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	_ = godotenv.Load()

	db := store.New(store.FileType, "./users.json")

	repo := users.NewRepository(db)

	service := users.NewService(repo)
	u := handler.NewUser(service)

	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	us := r.Group("/users")

	us.POST("/", u.CreateUser)
	us.GET("/", u.Get)
	us.GET("/filter", u.FilterUsers)
	us.PUT("/id", u.UpdateUser())
	us.PATCH("/id", u.UpdateName())
	us.DELETE("/id", u.DeleteUser())

	r.Run()
}
