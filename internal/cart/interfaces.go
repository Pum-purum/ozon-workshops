package cart

import "shop/internal/cart/productclient"

// ProductClient умеет получать информацию о товаре по SKU.
// Обрати внимание: сигнатура метода один-в-один совпадает с методом
// productclient.Client.GetProduct — поэтому *productclient.Client
// автоматически подходит под этот интерфейс, ничего в нём менять не нужно.
type ProductClientInterface interface {
	GetProduct(sku int64) (productclient.Product, bool, error)
}

// Repository — абстракция над хранилищем корзин.
// store (in-memory) уже реализует все эти методы "как есть".
type RepositoryInterface interface {
	add(userID int64, item Item)
	remove(userID, sku int64)
	clear(userID int64)
	get(userID int64) map[int64]Item
}
