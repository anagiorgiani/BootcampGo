package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name string
	Sobrenome string
	Rg string
	DataAdmission time.Time
}


func (a Student) details() {
	fmt.Println("Nome:", a.Name)
	fmt.Println("Sobrenome:", a.Sobrenome)
	fmt.Println("RG:", a.Rg)
	fmt.Println("DataAdmissao:", a.DataAdmission)
}