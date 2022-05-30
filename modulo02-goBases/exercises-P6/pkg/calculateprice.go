package main


type Produto struct {
	nome       string
	preco      float64
	quantidade int
}

type Servicos struct {
	nome  string
	preco float64
	min   int
}

type Manutencao struct {
	nome  string
	preco float64
}

func sumProdutos(p []Produto, out chan<- float64) float64 {
	total := 0.0
	for _, pd := range p {
		total += pd.preco * float64(pd.quantidade)
	}
	out <- total

	close(out)
	return total
}

func sumServicos(s []Servicos, out chan<- float64) float64 {
	total := 0.0
	for _, pd := range s {

		if pd.min < 30 {
			total += 30
		} else {
			total += pd.preco * float64(pd.min)
		}
	}
	out <- total

	close(out)
	return total
}

func sumManutencao(m []Manutencao, out chan<- float64) float64 {
	total := 0.0
	for _, pd := range m {
		total += pd.preco
	}
	out <- total

	close(out)
	return total
}



