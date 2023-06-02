package products

type Product struct {
    ID       int     `json:"id"`
    Name     string  `json:"name"`
    Type     string  `json:"type"`
    Count    int     `json:"count"`
    Price    float64 `json:"price"`
}

var ps []Product
var lastID int

type Repository interface{
    GetAll() ([]Product, error)
    Store(id int, name, type string, count int, price float64) (Product, error)
    LastID() (int, error)
}

type repository struct {}

func NewRepository() Repository {
    return &repository{}
}

func (r *repository) GetAll() ([]Product, error) {
    return ps, nil
}

func (r *repository) LastID() (int, error) {
    return lastID, nil
}

func (r *repository) Store(id int, name, type string, count int, price float64) (Product, error) {
    p := Product{id, name, type, count, price}
    ps = append(ps, p)
    lastID = p.ID
    return p, nil
}
