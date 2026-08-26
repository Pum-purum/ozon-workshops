package cart

import "sync"

type Item struct {
	SKU   int64  `json:"sku_id"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
	Count uint16 `json:"count"`
}

type Cart struct {
	Items      []Item `json:"items"`
	TotalPrice uint32 `json:"total_price"`
}

type store struct {
	mu    sync.Mutex
	carts map[int64]map[int64]Item // userID -> skuID -> Item
}

func newStore() *store {
	return &store{
		carts: make(map[int64]map[int64]Item),
	}
}

func (s *store) add(userID int64, item Item) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.carts[userID] == nil {
		s.carts[userID] = make(map[int64]Item)
	}

	existing, ok := s.carts[userID][item.SKU]
	if ok {
		item.Count += existing.Count
	}
	s.carts[userID][item.SKU] = item
}

func (s *store) remove(userID, sku int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.carts[userID], sku)
}

func (s *store) clear(userID int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.carts, userID)
}

func (s *store) get(userID int64) map[int64]Item {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.carts[userID]
}
