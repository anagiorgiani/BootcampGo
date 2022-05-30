package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
)

type Client struct {
	Arquivo   string
	Nome      string
	Sobrenome string
	Rg        string
	Telefone  string
	Endereco  string
}

func generateId() int {
	val := rand.Intn(1)

	return val
}

func trataErro() {
	err := recover()

	if err != nil {
		fmt.Println("erroii: ", err)
	}
}

func register(c Client) {
	defer trataErro()

	clients := readFile("customers.txt")
	fmt.Println("inicio")
	verifyClient(c, clients)

}
func verifyClient(c Client, clients []Client) {
	for _, client := range clients {
		if c == client {
			panic("Cliente já existe")
		}
	}
}

func readFile(file string) []Client {
	csvfile, err := os.Open(file)
	if err != nil {
		panic("arquivo n encontrado")
	}
	r := csv.NewReader(csvfile)
	r.Comma = ';'

	clients := []Client{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		arquivo := record[0]
		nome := record[1]
		sobrenome := record[2]
		rg := record[3]
		telefone := record[4]
		endereco := record[5]

		cli := Client{}
		cli.Arquivo = arquivo
		cli.Nome = nome
		cli.Sobrenome = sobrenome
		cli.Rg = rg
		cli.Telefone = telefone
		cli.Endereco = endereco
		clients = append(clients, cli)
	}
	return clients
}

func openFiles(file string) []byte {
	data, err := os.ReadFile("customers.txt")
	if err != nil {
		panic("Arquivo nao encontrado")
	}
	return data
}
