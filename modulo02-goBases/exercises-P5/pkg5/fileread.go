package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
	"strconv"
   
	"io"
)

type Product struct {
	id int
	name string
	price float64
	qty int
}

func fillProduct(id int, name string, price float64, qty int) Product {
	prod := Product{id: id, name: name, price: price, qty: qty}
	return prod
}

func printProduct(prods []Product) {
	for _, v := range prods {
    	fmt.Printf("%5d %-20s %10.2f \n", v.id, v.name, v.price * float64(v.qty))
	}
}

func readFile(){
	csvfile, err := os.Open("test.csv")
	   if err != nil {
		   log.Fatal(err)
	   }
	   r := csv.NewReader(csvfile)
		r.Comma = ';'

		products := []Product{}
   
	   for {
		   record, err := r.Read()
		   if err == io.EOF {
			   break
		   }
		   if err != nil {
			   log.Fatal(err)
		   }
		   id, _ := strconv.Atoi(record[0])
		   price, _ := strconv.ParseFloat(record[2], 64)
		   qty, _ :=strconv.Atoi(record[3])
		   products = append(products, fillProduct(id, record[1], price, qty))
	   }
	  printProduct(products)
   }