package main
import "fmt"
import "time"


func main() {
	pStudent := Student{
		Name: "Ana",
		Sobrenome: "Cavalheiro",
		Rg : "12312323",
		DataAdmission: time.Date(2016,01,15,0,0,0, 0, time.UTC),
	}

	pStudent.details()

	lj := novaLoja()

	p1 := newProd(1, "cel", 2000)
	p2 := newProd(2, "tv", 4000)
	p3 := newProd(3, "radio", 5000)

	lj.Adicionar(p1)
	lj.Adicionar(p2)
	lj.Adicionar(p3)
	fmt.Println(lj.Total())
}

