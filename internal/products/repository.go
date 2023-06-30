package products

import (
	"github.com/GabrieliPadilha/meli-bootcamp/pkg/store"
	"fmt"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category     string  `json:"category"`
	Count    int     `json:"count"`
	Price    float64 `json:"price"`
}

type Repository interface{
	GetAll() ([]Product, error)
	Store(product Product) (Product, error)
	Update(product Product) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

var ps []Product

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

func (r *repository) Store(p Product) (Product, error) {
	var produtos []Product
	r.db.Read(&produtos)
	produtos = append(produtos, p)
	err := r.db.Write(produtos)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) Update(p Product) (Product, error) {
	// p := Product{Name: name, Category: productType, Count: count, Price: price} // Instância de "p" para Update
	updated := false    // Atribuição false para Updated - não foi realizado nenhum update até aqui
	for i := range ps { // Este For percorrerá a lista dos elementos criados no array para buscar o elemento com o Id que já existe
		if ps[i].ID == p.ID { // Caso encontre esse Id ...
			ps[i].ID = p.ID    // ... o Id do novo produto será o mesmo do já existente (basicamente, o Id que passamos substituirá o já existente, só que são iguais)...
			ps[i] = p      // ... e aqui, irá atualizar (neste Id), todos os valores dos elementos que enviarmos no Put...
			updated = true // ... alterando o seu status para "True"
		}
	}
	if !updated { // Caso não tenha havido esse update, ou seja, se continuar como 'false'...
		return Product{}, fmt.Errorf("produto %d não encontrado", p.ID) // ... nos será enviada uma mensagem de erro
	}
	return p, nil // Retorno do novo produto com um erro do tipo 'nil'
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