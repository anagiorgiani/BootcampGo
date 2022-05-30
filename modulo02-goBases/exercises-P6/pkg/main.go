package main

import "fmt"

func main() {
	// p1 := people{}

	// person1 := person{id: 1, name: "Ana", sobName:"Caval", age: 28, email: "caval@gmail.com", senha: "12345"}

	// p1.addPerson(&person1)

	// p1.listPeople()

	// person2 := person{id: 2, name: "Olando", sobName:"Perrini", age: 16, email: "silva@gmail.com", senha: "7777"}

	// p1.addPerson(&person2)

	// p1.listPeople()

	// person2.name = "Lucas"

	// p1.editPerson(&person2, 2)

	// p1.listPeople()

	// user := &User{Nome: "Ana", Sobrenome: "cavalheiro", Email: "silva@gmail.com"}
	// prod := &Produtos{Nome: "tv", Preco: 2000}
	// newProduto("celular", 2500)
	// user.Adicionar(prod, user)
	// fmt.Println(user)
	// user.Delete(user)
	// fmt.Println(user)

	// var pr []Produto
	// prod := Produto{nome: "cel", preco: 40.00, quantidade: 2}
	// pr = append(pr, prod)

	// var sr []Servicos
	// serv := Servicos{nome: "arrumar", preco: 50.00, min: 30}
	// sr = append(sr, serv)

	// var mu []Manutencao
	// manut := Manutencao{nome: "pneu", preco: 80.00}
	// mu = append(mu, manut)

	// prCh := make(chan float64)
	// srCh := make(chan float64)
	// muCh := make(chan float64)

	// go sumProdutos(pr, prCh)
	// go sumServicos(sr, srCh)
	// go sumManutencao(mu, muCh)

	// total := getValue(prCh) + getValue(srCh) + getValue(muCh)

	// fmt.Println(total)

	tempoA1 := ordenaUm(A)
	tempoA2 := ordenaDois(A)
	tempoA3 := ordenaUm(B)
	tempoA4 := ordenaDois(B)
	tempoA5 := ordenaUm(C)
	tempoA6 := ordenaDois(C)

	if tempoA1 < tempoA2 {
		fmt.Println(tempoA1)
	} else {
		fmt.Println(tempoA2)
	}

	if tempoA2 < tempoA3 {
		fmt.Println(tempoA2)
	} else {
		fmt.Println(tempoA3)
	}

	if tempoA3 < tempoA4 {
		fmt.Println(tempoA3)
	} else {
		fmt.Println(tempoA4)
	}

	if tempoA4 < tempoA5 {
		fmt.Println(tempoA4)
	} else {
		fmt.Println(tempoA5)
	}

	if tempoA5 < tempoA6 {
		fmt.Println(tempoA5)
	} else {
		fmt.Println(tempoA6)
	}

}

func getValue(in <-chan float64) float64 {
	total := 0.0
	for v := range in {
		total += v
	}
	return total
}
