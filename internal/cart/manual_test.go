package cart

import (
	"testing"

	"github.com/gojuno/minimock/v3"

	"shop/internal/cart/productclient"
)

func TestService_Add_Success(t *testing.T) {
	mc := minimock.NewController(t)

	productClientMock := NewProductClientMock(mc)
	productClientMock.GetProductMock.Return(
		productclient.Product{SKU: 1, Name: "Кружка", Price: 350}, true, nil,
	)

	svc := &Service{
		store:         newStore(),
		productClient: productClientMock,
	}

	err := svc.Add(42, 1, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cart := svc.Get(42)
	if len(cart.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(cart.Items))
	}
}
