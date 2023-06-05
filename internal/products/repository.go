package products

import "fmt"

//repository de produto coma estrutuda de produto
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category     string  `json:"type"`
	Count    int     `json:"count"`
	Price    float64 `json:"price"`
}

//variavel de tipo array de Produtos onde irá persistir os dados
var ps []Product
// variavel para guardar o valor do ultimo ID que será usado para criar novos produtos
var lastID int

//interface com os metodos de buscar todos, inserir produto e obtem o valor do ultimo id
type Repository interface{
	GetAll() ([]Product, error)
	Store(id int, name, category string, count int, price float64) (Product, error)
	LastID() (int, error)
	Update(id int,  name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

//estrutura repository criada vazia
type repository struct {}

// retorna o endereço de memeoria da estrutura vazia repository
func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Product, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}
//Recebe os campos de produto pro parametro e inser atribuindo eles a Product e depois insere na variavel ps que tem um array vazio de product
func (r *repository) Store(id int, name, category string, count int, price float64) (Product, error) {
	p := Product{id, name, category, count, price}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}
func (r *repository) Update(id int,  name, category string, count int, price float64) (Product, error) {
	p := Product{Name: name, Category: category, Count: count,Price: price}
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("producto %d no encontrado", id)
	}
	return p, nil
}

func (r *repository) UpdateName(id int, name string) (Product, error) {
	var p Product
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			ps[i].Name = name
			updated = true
			p = ps[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Produto %d no encontrado", id)
	}
	return p, nil
}

 func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range ps {
		if ps[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("produto %d nao encontrado", id)
	}
	ps = append(ps[:index], ps[index+1:]...)
	return nil
 }
 