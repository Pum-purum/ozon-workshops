package product

type Product struct {
	SKU   int64  `json:"sku_id"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

var store = []Product{
	{SKU: 1, Name: "Кружка", Price: 350},
	{SKU: 2, Name: "Футболка", Price: 1200},
	{SKU: 3, Name: "Кепка", Price: 800},
}

func All() []Product {
	return store
}

func FindBySKU(sku int64) (Product, bool) {
	for _, p := range store {
		if p.SKU == sku {
			return p, true
		}
	}
	return Product{}, false
}
