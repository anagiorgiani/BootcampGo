package main

type User struct {
	Nome string
	Sobrenome string
	Email string
	Itens []Produtos
}

type Produtos struct {
	Nome string
	Preco float64
	qt int
}



func (usr *User) Adicionar(p *Produtos, u *User)  {
	usr.Itens = append(usr.Itens, *p)
	usr.Nome = u.Nome
	usr.Sobrenome = u.Sobrenome
	usr.Email = u.Email

}

func newProduto(Nome string, Preco float64) *Produtos {
	return &Produtos{Nome: Nome, Preco: Preco}
}

func (usr *User) Delete(u *User) {
   usr.Itens = nil
}

