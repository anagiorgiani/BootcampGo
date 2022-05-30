package main

import "os"

func openFile(file string) []byte {
	data, err := os.ReadFile("customers.txt")
	if err != nil {
		panic("Arquivo nao encontrado")
	}
	return data
}
