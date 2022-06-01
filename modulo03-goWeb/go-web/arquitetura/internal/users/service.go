package users

type Service interface {
	Get(id int) (User, error)
	GetAll() ([]User, error)
	CreateUser(Nome string, Sobrenome string, Email string, Idade int, Ativo bool) (User, error)
	FilterUsers(Nome string, Sobrenome string) ([]User, error)
	UpdateUser(Id int, Nome string, Sobrenome string, Email string, Idade int, Ativo bool) (User, error)
	UpdateName(Id int, Nome string) (User, error)
	DeleteUser(Id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Get(id int) (User, error) {
	ulist, err := s.repository.Get(id)

	if err != nil {
		return User{}, err
	}

	return ulist, nil
}

func (s *service) FilterUsers(Nome string, Sobrenome string) ([]User, error) {
	ulist, err := s.repository.FilterUsers(Nome, Sobrenome)

	if err != nil {
		return nil, err
	}

	return ulist, nil
}

func (s *service) GetAll() ([]User, error) {
	ulist, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return ulist, nil
}

func (s *service) CreateUser(Nome string, Sobrenome string, Email string, Idade int, Ativo bool) (User, error) {
	nextId := s.repository.getNextId()

	user, err := s.repository.CreateUser(nextId, Nome, Sobrenome, Email, Idade, Ativo)

	if err != nil {
		return User{}, err
	}

	return user, nil

}

func (s *service) UpdateUser(Id int, Nome string, Sobrenome string, Email string, Idade int, Ativo bool) (User, error) {
	return s.repository.UpdateUser(Id, Nome, Sobrenome, Email, Idade, Ativo)
}

func (s *service) UpdateName(Id int, Nome string) (User, error) {
	return s.repository.UpdateName(Id, Nome)
}

func (s *service) DeleteUser(Id int) error {
	return s.repository.DeleteUser(Id)
}
