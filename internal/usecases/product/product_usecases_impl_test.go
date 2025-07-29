package product

import (
	"context"
	"errors"
	"testing"

	dto "post-tech-challenge-10soat/internal/dto/product"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/gateways"
	"github.com/stretchr/testify/assert"
)

type mockProductGateway struct {
	interfaces.ProductGateway
	CreateProductFunc func(ctx context.Context, product entity.Product) (entity.Product, error)
	DeleteProductFunc func(ctx context.Context, id string) error
	ListProductsFunc func(ctx context.Context, categoryId string) ([]entity.Product, error)
	UpdateProductFunc func(ctx context.Context, product entity.Product) (entity.Product, error)
}

type mockCategoryGateway struct {
	interfaces.CategoryGateway
}

func (m *mockProductGateway) CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	return m.CreateProductFunc(ctx, product)
}
func (m *mockProductGateway) DeleteProduct(ctx context.Context, id string) error {
	return m.DeleteProductFunc(ctx, id)
}
func (m *mockProductGateway) ListProducts(ctx context.Context, categoryId string) ([]entity.Product, error) {
	return m.ListProductsFunc(ctx, categoryId)
}
func (m *mockProductGateway) UpdateProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	return m.UpdateProductFunc(ctx, product)
}

func TestCreateProductUsecaseImpl_Execute_Success(t *testing.T) {
	mockProduct := &mockProductGateway{
		CreateProductFunc: func(ctx context.Context, product entity.Product) (entity.Product, error) {
			product.ID = "p1"
			return product, nil
		},
	}
	mockCategory := &mockCategoryGateway{}
	usecase := NewCreateProductUsecaseImpl(mockProduct, mockCategory)
	input := dto.CreateProductDTO{Name: "Coca", Price: 10.0, CategoryID: "c1"}
	result, err := usecase.Execute(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, "p1", result.ID)
	assert.Equal(t, "Coca", result.Name)
	assert.Equal(t, 10.0, result.Price)
	assert.Equal(t, "c1", result.CategoryID)
}

func TestCreateProductUsecaseImpl_Execute_Error(t *testing.T) {
	expectedErr := errors.New("fail")
	mockProduct := &mockProductGateway{
		CreateProductFunc: func(ctx context.Context, product entity.Product) (entity.Product, error) {
			return entity.Product{}, expectedErr
		},
	}
	mockCategory := &mockCategoryGateway{}
	usecase := NewCreateProductUsecaseImpl(mockProduct, mockCategory)
	input := dto.CreateProductDTO{Name: "Coca", Price: 10.0, CategoryID: "c1"}
	result, err := usecase.Execute(context.Background(), input)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create product")
	assert.Equal(t, entity.Product{}, result)
}

func TestDeleteProductUsecaseImpl_Execute_Success(t *testing.T) {
	mockProduct := &mockProductGateway{
		DeleteProductFunc: func(ctx context.Context, id string) error {
			return nil
		},
	}
	usecase := NewDeleteProductUsecaseImpl(mockProduct)
	err := usecase.Execute(context.Background(), "p1")
	assert.NoError(t, err)
}

func TestDeleteProductUsecaseImpl_Execute_Error(t *testing.T) {
	expectedErr := errors.New("fail")
	mockProduct := &mockProductGateway{
		DeleteProductFunc: func(ctx context.Context, id string) error {
			return expectedErr
		},
	}
	usecase := NewDeleteProductUsecaseImpl(mockProduct)
	err := usecase.Execute(context.Background(), "p1")
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
}

func TestListProductsUsecaseImpl_Execute_Success(t *testing.T) {
	mockProduct := &mockProductGateway{
		ListProductsFunc: func(ctx context.Context, categoryId string) ([]entity.Product, error) {
			return []entity.Product{{ID: "p1"}, {ID: "p2"}}, nil
		},
	}
	mockCategory := &mockCategoryGateway{}
	usecase := NewListProductsUsecaseImpl(mockProduct, mockCategory)
	products, err := usecase.Execute(context.Background(), "c1")
	assert.NoError(t, err)
	assert.Len(t, products, 2)
}

func TestListProductsUsecaseImpl_Execute_Error(t *testing.T) {
	expectedErr := errors.New("fail")
	mockProduct := &mockProductGateway{
		ListProductsFunc: func(ctx context.Context, categoryId string) ([]entity.Product, error) {
			return nil, expectedErr
		},
	}
	mockCategory := &mockCategoryGateway{}
	usecase := NewListProductsUsecaseImpl(mockProduct, mockCategory)
	products, err := usecase.Execute(context.Background(), "c1")
	assert.Error(t, err)
	assert.Nil(t, products)
}

func TestUpdateProductUsecaseImpl_Execute_Success(t *testing.T) {
	mockProduct := &mockProductGateway{
		UpdateProductFunc: func(ctx context.Context, product entity.Product) (entity.Product, error) {
			product.Name = "Fanta"
			return product, nil
		},
	}
	mockCategory := &mockCategoryGateway{}
	usecase := NewUpdateProductUsecaseImpl(mockProduct, mockCategory)
	input := dto.UpdateProductDTO{ID: "p1", Name: "Fanta"}
	result, err := usecase.Execute(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, "Fanta", result.Name)
}

func TestUpdateProductUsecaseImpl_Execute_Error(t *testing.T) {
	expectedErr := errors.New("fail")
	mockProduct := &mockProductGateway{
		UpdateProductFunc: func(ctx context.Context, product entity.Product) (entity.Product, error) {
			return entity.Product{}, expectedErr
		},
	}
	mockCategory := &mockCategoryGateway{}
	usecase := NewUpdateProductUsecaseImpl(mockProduct, mockCategory)
	input := dto.UpdateProductDTO{ID: "p1", Name: "Fanta"}
	result, err := usecase.Execute(context.Background(), input)
	assert.Error(t, err)
	assert.Equal(t, entity.Product{}, result)
}
