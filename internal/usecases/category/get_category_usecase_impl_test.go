package category

import (
	"context"
	"errors"
	"testing"

	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/gateways"
	"github.com/stretchr/testify/assert"
)

type mockCategoryGateway struct {
	interfaces.CategoryGateway
	GetCategoryByIdFunc func(ctx context.Context, id string) (entity.Category, error)
}

func (m *mockCategoryGateway) GetCategoryById(ctx context.Context, id string) (entity.Category, error) {
	return m.GetCategoryByIdFunc(ctx, id)
}

func TestGetCategoryUsecaseImpl_Execute_Success(t *testing.T) {
	mockGateway := &mockCategoryGateway{
		GetCategoryByIdFunc: func(ctx context.Context, id string) (entity.Category, error) {
			return entity.Category{ID: "1", Name: "Bebidas"}, nil
		},
	}
	usecase := NewGetCategoryUsecase(mockGateway)
	cat, err := usecase.Execute(context.Background(), "1")
	assert.NoError(t, err)
	assert.Equal(t, "1", cat.ID)
	assert.Equal(t, "Bebidas", cat.Name)
}

func TestGetCategoryUsecaseImpl_Execute_Error(t *testing.T) {
	expectedErr := errors.New("not found")
	mockGateway := &mockCategoryGateway{
		GetCategoryByIdFunc: func(ctx context.Context, id string) (entity.Category, error) {
			return entity.Category{}, expectedErr
		},
	}
	usecase := NewGetCategoryUsecase(mockGateway)
	cat, err := usecase.Execute(context.Background(), "2")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get category by id")
	assert.Equal(t, entity.Category{}, cat)
}
