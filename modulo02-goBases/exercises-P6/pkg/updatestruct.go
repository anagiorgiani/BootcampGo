package main
import "fmt"

type person struct {
	id   int
	name string
	sobName string
	age  int
	email string
	senha string
  }
  
  type people []*person
  
  func (p *people) addPerson(person *person) {
	*p = append(*p, person)
  }
  
  func (p *people) editPerson(person *person, id int) {
	for _, personEdit := range *p {
	  if personEdit.id == id {
		personEdit.age = person.age
		personEdit.name = person.name
		personEdit.sobName = person.sobName
		personEdit.email  = person.email
		personEdit.senha = person.senha
	  } else {
		fmt.Println("não foi possível achar essa pessoa")
	  }
	}
  }
  
  func (p people) listPeople() {
	for _, person := range p {
	  fmt.Println("Nome: ", person.name)
	  fmt.Println("Sobrenome: ", person.sobName)
	  fmt.Println("Idade: ", person.age)
	  fmt.Println("Email: ", person.email)
	}
  }
  
  