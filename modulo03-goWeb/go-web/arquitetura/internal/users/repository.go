package users

import (
	"fmt"
	"strings"
	"time"

	"github.com/anagiorgiani/BootcampGo/modulo03-goWeb/go-web/arquitetura/pkg/store"
)

type User struct {
	Id          int    `json:"id"`
	Nome        string `json:"nome" binding:"required"`
	Sobrenome   string `json:"sobrenome" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Idade       int    `json:"idade" binding:"required"`
	Ativo       bool   `json:"ativo" binding:"required"`
	DataCriacao string `json:"dataCriacao"`
}

var userEnt = []User{}
var nextId int

type Repository interface {
	Get(id int) (User, error)
	GetAll() ([]User, error)
	CreateUser(Id int, Nome string, Sobrenome string, Email string, Idade int, Ativo bool) (User, error)
	FilterUsers(Nome string, Sobrenome string) ([]User, error)
	getNextId() int
	UpdateUser(Id int, Nome string, Sobrenome string, Email string, Idade int, Ativo bool) (User, error)
	UpdateName(Id int, Nome string) (User, error)
	DeleteUser(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {

	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]User, error) {

	var ulist []User
	r.db.Read(&ulist)

	return ulist, nil
}

func (r *repository) Get(id int) (User, error) {
	for _, u := range userEnt {
		if u.Id == id {
			return u, nil
		}
	}
	return User{}, nil
}

func (r *repository) getNextId() int {
	var ulist []User

	if err := r.db.Read(&ulist); err != nil {
		return 1
	}

	if len(ulist) == 0 {
		return 1
	}

	return ulist[len(ulist)-1].Id + 1
}

func (r *repository) CreateUser(Id int, Nome string, Sobrenome string, Email string, Idade int, Ativo bool) (User, error) {

	u := User{Id, Nome, Sobrenome, Email, Idade, Ativo, time.Now().Format("2006-01-02 15:04:05")}

	var ulist []User
	r.db.Read(&ulist)

	ulist = append(ulist, u)

	if err := r.db.Write(ulist); err != nil {
		return User{}, err
	}

	return u, nil
}

func (r *repository) FilterUsers(Nome string, Sobrenome string) ([]User, error) {
	userFiltered := []User{}

	for _, u := range userEnt {
		if strings.Contains(u.Nome, Nome) && strings.Contains(u.Sobrenome, Sobrenome) {
			userFiltered = append(userFiltered, u)
		}
	}

	return userFiltered, nil
}

func (r *repository) UpdateUser(Id int, Nome string, Sobrenome string, Email string, Idade int, Ativo bool) (User, error) {
	u := User{Nome: Nome, Sobrenome: Sobrenome, Email: Email, Idade: Idade, Ativo: Ativo}
	updated := false
	for i := range userEnt {
		if userEnt[i].Id == Id {
			u.Id = Id
			userEnt[i] = u
			updated = true
		}
	}
	if !updated {
		return User{}, fmt.Errorf("user %d nao encontrado", Id)
	}
	return u, nil
}

func (r *repository) UpdateName(Id int, Nome string) (User, error) {
	u := User{Nome: Nome}
	updated := false
	for i := range userEnt {
		if userEnt[i].Id == Id {
			u.Nome = Nome
			userEnt[i] = u
			updated = true
		}
	}
	if !updated {
		return User{}, fmt.Errorf("user %d nao encontrado", Id)
	}
	return u, nil
}

func (r *repository) DeleteUser(Id int) error {
	deleted := false
	var index int
	for i := range userEnt {
		if userEnt[i].Id == Id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("user %d nao encontrado", Id)
	}
	userEnt = append(userEnt[:index], userEnt[index+1:]...)
	return nil
}
