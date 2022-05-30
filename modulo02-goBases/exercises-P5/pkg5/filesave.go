package main

import (
	"fmt"
    "log"
    "os"
	"io/ioutil"
)


func saveFile(id int, name string, value float64, qt int){
	content, err := ioutil.ReadFile("test.csv") 
	if err != nil {
		content = []byte(nil)
	}
	d1:=[]byte(fmt.Sprint("\n", id , ";" , name , ";" , value, ";" , qt))
	d1 = append(content, d1...)
	erro := os.WriteFile("test.csv", d1, 0644)
	if erro != nil {
		log.Println("erro ao escrever arquivo:", erro)
	}

   }