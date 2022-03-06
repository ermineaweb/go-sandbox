package data

type Product struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type Products []Product

var ProductList = []Product{
	{Id: "1", Name: "banane", Price: 2.2},
	{Id: "2", Name: "orange", Price: 1.5},
	{Id: "3", Name: "stylo", Price: 0.5},
}
