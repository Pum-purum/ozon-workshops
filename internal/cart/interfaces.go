package cart

import "shop/internal/cart/productclient"

//go:generate minimock -i shop/internal/cart.ProductClientInterface -o ./product_client_mock_test.go -n ProductClientMock -p cart
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
