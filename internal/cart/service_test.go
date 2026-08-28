package cart

import (
	"errors"
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

func TestService_Add_CallsRepositoryWithCorrectItem(t *testing.T) {
	mc := minimock.NewController(t)

	productClientMock := NewProductClientMock(mc)
	productClientMock.GetProductMock.Return(
		productclient.Product{SKU: 1, Name: "Кружка", Price: 350}, true, nil,
	)

	var gotUserID int64
	var gotItem Item
	repositoryMock := NewRepositoryMock(mc)
	repositoryMock.addMock.Set(func(userID int64, item Item) {
		gotUserID = userID
		gotItem = item
	})

	svc := &Service{
		store:         repositoryMock,
		productClient: productClientMock,
	}

	err := svc.Add(42, 1, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotUserID != 42 {
		t.Errorf("expected userID 42, got %d", gotUserID)
	}

	want := Item{SKU: 1, Name: "Кружка", Price: 350, Count: 3}
	if gotItem != want {
		t.Errorf("expected item %+v, got %+v", want, gotItem)
	}
}

func TestService_Add_ProductNotFound(t *testing.T) {
	mc := minimock.NewController(t)

	productClientMock := NewProductClientMock(mc)
	productClientMock.GetProductMock.Return(productclient.Product{}, false, nil)

	repositoryMock := NewRepositoryMock(mc)
	// AddMock намеренно не настраиваем — если Service.Add всё же вызовет
	// store.add, мок сам упадёт с понятной ошибкой ("addMock вызван без .Set()").

	svc := &Service{
		store:         repositoryMock,
		productClient: productClientMock,
	}

	err := svc.Add(42, 999, 1)
	if !errors.Is(err, ErrProductNotFound) {
		t.Fatalf("expected ErrProductNotFound, got %v", err)
	}
}

func TestService_Add_ProductClientError(t *testing.T) {
	mc := minimock.NewController(t)

	wantErr := errors.New("network error")
	productClientMock := NewProductClientMock(mc)
	productClientMock.GetProductMock.Return(productclient.Product{}, false, wantErr)

	repositoryMock := NewRepositoryMock(mc)
	// addMock снова не настраиваем — store.add не должен вызываться.

	svc := &Service{
		store:         repositoryMock,
		productClient: productClientMock,
	}

	err := svc.Add(42, 1, 1)
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected %v, got %v", wantErr, err)
	}
}

func TestService_Remove(t *testing.T) {
	mc := minimock.NewController(t)

	repositoryMock := NewRepositoryMock(mc)
	var gotUserID, gotSKU int64
	repositoryMock.removeMock.Set(func(userID, sku int64) {
		gotUserID = userID
		gotSKU = sku
	})

	svc := &Service{
		store:         repositoryMock,
		productClient: NewProductClientMock(mc),
	}

	svc.Remove(42, 7)

	if gotUserID != 42 || gotSKU != 7 {
		t.Errorf("expected Remove(42, 7), got Remove(%d, %d)", gotUserID, gotSKU)
	}
}

func TestService_Clear(t *testing.T) {
	mc := minimock.NewController(t)

	repositoryMock := NewRepositoryMock(mc)
	var gotUserID int64
	repositoryMock.clearMock.Set(func(userID int64) {
		gotUserID = userID
	})

	svc := &Service{
		store:         repositoryMock,
		productClient: NewProductClientMock(mc),
	}

	svc.Clear(42)

	if gotUserID != 42 {
		t.Errorf("expected Clear(42), got Clear(%d)", gotUserID)
	}
}

func TestService_Get_Empty(t *testing.T) {
	mc := minimock.NewController(t)

	repositoryMock := NewRepositoryMock(mc)
	repositoryMock.getMock.Return(nil)

	svc := &Service{
		store:         repositoryMock,
		productClient: NewProductClientMock(mc),
	}

	cart := svc.Get(42)

	if len(cart.Items) != 0 {
		t.Errorf("expected empty cart, got %+v", cart)
	}
	if cart.TotalPrice != 0 {
		t.Errorf("expected total price 0, got %d", cart.TotalPrice)
	}
}
