package products

type Service interface {
	GetAll() ([]Product, error)
	Store(name string, category string, count int, price float64) (Product, error)
	Update(id int, name, category string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}
type service struct {
	repository Repository
}
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil{
		return nil, err
	}
	return ps, nil
}

func (s *service) Store(name, category string, count int, price float64) (Product, error) {
	product := Product{Name: name, Category: category, Count: count, Price: price}
	product, err := s.repository.Store(product)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s service) Update(id int, name, productType string, count int, price float64) (Product, error) {
	product := Product{id, name, productType, count, price}

	product, err := s.repository.Update(product)

	return product, err
}

func (s *service) UpdateName(id int, name string) (Product, error) {

	return s.repository.UpdateName(id, name)
} 
 
func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
 