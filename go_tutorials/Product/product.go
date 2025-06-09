package Product

type Product struct {
	ID       int
	Name     string
	Quantity int
	Price    float32
}

func AddProduct(products *[]Product, id int, name string, quantity int, price float32) {
	var newProducts = Product{
		ID:       id,
		Name:     name,
		Quantity: quantity,
		Price:    price,
	}
	*products = append(*products, newProducts)
}
