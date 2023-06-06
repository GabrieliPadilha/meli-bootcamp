package products

import (
	"github.com/GabrieliPadilha/meli-bootcamp/pkg/store"
	"fmt"
)

//repository de produto coma estrutuda de produto
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category     string  `json:"category"`
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
	Update(id int,  name, category string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

// retorna o endereço de memeoria da estrutura vazia repository
func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Product, error) {
	var produtos []Product
	err := r.db.Read(&produtos)

	if err != nil {
		return nil, err
	}
	return produtos, nil
}

func (r *repository) LastID() (int, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}

	if len(ps) == 0 {
		return 0, nil
	}
	ultimoProduto := ps[len(ps)-1]
	return ultimoProduto.ID, nil
}

func (r *repository) Store(id int, name, category string, count int, price float64) (Product, error) {
	var produtos []Product
	r.db.Read(&produtos)
	p := Product{id, name, category, count, price}
	produtos = append(produtos, p)
	err := r.db.Write(produtos)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) Update(id int, name, category string, count int, price float64) (Product, error) {
	updated := false
	var produtos []Product
	r.db.Read(&produtos)
	p := Product{Name: name, Category: category, Count: count, Price: price}

	for indice := range produtos{
		if produtos[indice].ID == id {
			p.ID = id 
			produtos[indice] = p 
			err := r.db.Write(produtos)
			if err != nil {
				return Product{}, err
			}
			updated = true
		}
	}

	if !updated {
		return Product{}, fmt.Errorf("produto %d no encontrado", id)
	}
	return p, nil
}

func (r *repository) UpdateName(id int, name string) (Product, error) {
	var produtos []Product
	r.db.Read(&produtos)
	var produto Product
	updated := false
	for i := range produtos {
		if produtos[i].ID == id {
			produtos[i].Name = name
			produto = produtos[i]
			err := r.db.Write(produtos)
			if err != nil {
				return Product{}, err
			}
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("produto %d no encontrado", id)
	}
	return produto, nil
}

func (r *repository) Delete(id int) error {
	var produtos []Product
	r.db.Read(&produtos)
	deleted := false
	var index int
	for i := range produtos {
		if produtos[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("produto %d nao encontrado", id)
	}
	produtos = append(produtos[:index], produtos[index+1:]...)
	r.db.Write(produtos)
	return nil
}
 