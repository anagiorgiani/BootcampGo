package main



const (
	pequeno = 1
	medio  =  2
	grande =  3
)


type Produto struct { 
	nome string
	preco float64
	categoria int 
}

type Loja struct {
	Items []iProduto
}


type iProduto interface {
	CalcularCusto() float64
} 

type Ecommerce interface {
	Total() float64
	Adicionar(p iProduto) 

}

func (box *Loja) Adicionar(p iProduto)  {
	box.Items = append(box.Items, p)
}

func newProd(cat int, nom string, prc float64) iProduto  {
	return &Produto{nome: nom, preco: prc, categoria: cat}
}

func (box *Loja) Total() float64 {
	total := 0.0
	for _ , v := range box.Items {
		total += v.CalcularCusto()
	}
	return total
}

func novaLoja() Ecommerce {
	nvLoja := &Loja{}
	return nvLoja
}

func calcPequeno(valueProd float64 ) float64 {
	return valueProd

}

func calcMedio(valueProd float64 ) float64 {
	result := valueProd * 1.03 
	return result

}

func calcGrande(valueProd float64 ) float64 {
	result := valueProd * 1.06 + 2500 
	return result

}


func (a *Produto)CalcularCusto() float64 {
	switch a.categoria {
	case 1:
		return calcPequeno(a.preco)
	case 2:
		return calcMedio(a.preco)
	case 3:
		return calcGrande(a.preco)
	default:
		return 0.0
	}
	
}