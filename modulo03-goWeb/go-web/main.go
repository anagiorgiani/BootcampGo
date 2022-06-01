package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type User struct {
	Id          string `json:"id"`
	Nome        string `json:"nome" binding:"required"`
	Sobrenome   string `json:"sobrenome" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Idade       int    `json:"idade" binding:"required"`
	Ativo       bool   `json:"ativo" binding:"required"`
	DataCriacao string `json:"dataCriacao" binding:"required"`
}

var usersCreate = []User{}

func main() {

	_ = godotenv.Load()

	router := gin.Default()
	router.POST("/users/create", CreateUser)
	router.GET("/users", Get)
	router.GET("/filter", FilterUsers)
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": " Hello, Ana"})
	})
	router.Run()

}

func FilterUsers(c *gin.Context) {
	users := []User{}

	users = readFile()

	paramNome := c.Query("nome")
	paramSobrenome := c.Query("sobrenome")

	userFiltered := []User{}

	for _, u := range users {
		if strings.Contains(u.Nome, paramNome) && strings.Contains(u.Sobrenome, paramSobrenome) {
			userFiltered = append(userFiltered, u)
		}
	}

	c.JSON(http.StatusOK, userFiltered)
}

func Get(c *gin.Context) {
	users := []User{}

	users = readFile()

	if c.Query("id") != "" {
		for _, u := range users {
			if u.Id == c.Query("id") {
				c.JSON(http.StatusOK, u)
				return
			}
		}
	}

	c.JSON(http.StatusOK, users)
}

func readFile() []User {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	byteValueJSON, _ := ioutil.ReadAll(jsonFile)

	users := []User{}
	json.Unmarshal(byteValueJSON, &users)

	return users
}

func CreateUser(c *gin.Context) {

	if validateToken(c) == false {
		c.JSON(401, gin.H{
			"ërror": "token invalido",
		})
		return
	}

	var req User

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Id = strconv.Itoa(getNextId())
	usersCreate = append(usersCreate, req)
	fmt.Println(usersCreate)
}

func getNextId() int {
	if len(usersCreate) == 0 {
		return 1
	}
	codInt, _ := strconv.Atoi(usersCreate[len(usersCreate)-1].Id)

	return codInt + 1
}

func validateToken(c *gin.Context) bool {
	fmt.Println("Token do ENV: ", os.Getenv("TOKEN"))
	token := c.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		return false
	}
	return true
}
