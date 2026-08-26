package cart

import (
	"errors"

	"shop/internal/cart/productclient"
)

var ErrProductNotFound = errors.New("product not found")

type Service struct {
	store         *store
	productClient *productclient.Client
}

func NewService(productClient *productclient.Client) *Service {
	return &Service{
		store:         newStore(),
		productClient: productClient,
	}
}

func (s *Service) Add(userID, sku int64, count uint16) error {
	p, found, err := s.productClient.GetProduct(sku)
	if err != nil {
		return err
	}
	if !found {
		return ErrProductNotFound
	}

	s.store.add(userID, Item{
		SKU:   p.SKU,
		Name:  p.Name,
		Price: p.Price,
		Count: count,
	})

	return nil
}

func (s *Service) Remove(userID, sku int64) {
	s.store.remove(userID, sku)
}

func (s *Service) Clear(userID int64) {
	s.store.clear(userID)
}

func (s *Service) Get(userID int64) Cart {
	items := s.store.get(userID)

	cart := Cart{
		Items: make([]Item, 0, len(items)),
	}

	for _, item := range items {
		cart.Items = append(cart.Items, item)
		cart.TotalPrice += item.Price * uint32(item.Count)
	}

	return cart
}
