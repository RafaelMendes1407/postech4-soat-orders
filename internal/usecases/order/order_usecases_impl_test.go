package order

import (
	"context"
	"errors"
	"testing"

	dto "post-tech-challenge-10soat/internal/dto/order"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/gateways"
	"github.com/stretchr/testify/assert"
)

type mockOrderGateway struct {
	interfaces.OrderGateway
	CreateOrderFunc func(ctx context.Context, order entity.Order) (entity.Order, error)
	ListOrdersFunc func(ctx context.Context, limit uint64) ([]entity.Order, error)
	UpdateOrderStatusFunc func(ctx context.Context, id string, status string) (entity.Order, error)
	GetOrderPaymentStatusFunc func(ctx context.Context, id string) (OrderPaymentStatus, error)
}

func (m *mockOrderGateway) CreateOrder(ctx context.Context, order entity.Order) (entity.Order, error) {
	return m.CreateOrderFunc(ctx, order)
}
func (m *mockOrderGateway) ListOrders(ctx context.Context, limit uint64) ([]entity.Order, error) {
	return m.ListOrdersFunc(ctx, limit)
}
func (m *mockOrderGateway) UpdateOrderStatus(ctx context.Context, id string, status string) (entity.Order, error) {
	return m.UpdateOrderStatusFunc(ctx, id, status)
}
func (m *mockOrderGateway) GetOrderPaymentStatus(ctx context.Context, id string) (OrderPaymentStatus, error) {
	return m.GetOrderPaymentStatusFunc(ctx, id)
}

func TestCreateOrderUsecaseImpl_Execute_Success(t *testing.T) {
	mockGateway := &mockOrderGateway{
		CreateOrderFunc: func(ctx context.Context, order entity.Order) (entity.Order, error) {
			order.ID = "1"
			return order, nil
		},
	}
	usecase := NewCreateOrderUsecaseImpl(mockGateway)
	input := dto.CreateOrderDTO{ClientID: "c1"}
	order, err := usecase.Execute(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, "1", order.ID)
	assert.Equal(t, "c1", order.ClientID)
}

func TestCreateOrderUsecaseImpl_Execute_Error(t *testing.T) {
	expectedErr := errors.New("fail")
	mockGateway := &mockOrderGateway{
		CreateOrderFunc: func(ctx context.Context, order entity.Order) (entity.Order, error) {
			return entity.Order{}, expectedErr
		},
	}
	usecase := NewCreateOrderUsecaseImpl(mockGateway)
	input := dto.CreateOrderDTO{ClientID: "c1"}
	order, err := usecase.Execute(context.Background(), input)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create order")
	assert.Equal(t, entity.Order{}, order)
}

func TestListOrdersUseCaseImpl_Execute_Success(t *testing.T) {
	mockGateway := &mockOrderGateway{
		ListOrdersFunc: func(ctx context.Context, limit uint64) ([]entity.Order, error) {
			return []entity.Order{{ID: "1"}, {ID: "2"}}, nil
		},
	}
	usecase := NewListOrdersUseCaseImpl(mockGateway)
	orders, err := usecase.Execute(context.Background(), 2)
	assert.NoError(t, err)
	assert.Len(t, orders, 2)
}

func TestListOrdersUseCaseImpl_Execute_Error(t *testing.T) {
	expectedErr := errors.New("fail")
	mockGateway := &mockOrderGateway{
		ListOrdersFunc: func(ctx context.Context, limit uint64) ([]entity.Order, error) {
			return nil, expectedErr
		},
	}
	usecase := NewListOrdersUseCaseImpl(mockGateway)
	orders, err := usecase.Execute(context.Background(), 2)
	assert.Error(t, err)
	assert.Nil(t, orders)
}

func TestUpdateOrderStatusUseCaseImpl_Execute_Success(t *testing.T) {
	mockGateway := &mockOrderGateway{
		UpdateOrderStatusFunc: func(ctx context.Context, id string, status string) (entity.Order, error) {
			return entity.Order{ID: id, Status: status}, nil
		},
	}
	usecase := NewUpdateOrderStatusUseCaseImpl(mockGateway)
	order, err := usecase.Execute(context.Background(), "1", "PAID")
	assert.NoError(t, err)
	assert.Equal(t, "1", order.ID)
	assert.Equal(t, "PAID", order.Status)
}

func TestUpdateOrderStatusUseCaseImpl_Execute_Error(t *testing.T) {
	expectedErr := errors.New("fail")
	mockGateway := &mockOrderGateway{
		UpdateOrderStatusFunc: func(ctx context.Context, id string, status string) (entity.Order, error) {
			return entity.Order{}, expectedErr
		},
	}
	usecase := NewUpdateOrderStatusUseCaseImpl(mockGateway)
	order, err := usecase.Execute(context.Background(), "1", "PAID")
	assert.Error(t, err)
	assert.Equal(t, entity.Order{}, order)
}

func TestGetOrderPaymentStatusUseCaseImpl_Execute_Success(t *testing.T) {
	mockGateway := &mockOrderGateway{
		GetOrderPaymentStatusFunc: func(ctx context.Context, id string) (OrderPaymentStatus, error) {
			return OrderPaymentStatus{OrderID: id, Status: "PAID"}, nil
		},
	}
	usecase := NewGetOrderPaymentStatusUseCaseImpl(mockGateway)
	status, err := usecase.Execute(context.Background(), "1")
	assert.NoError(t, err)
	assert.Equal(t, "1", status.OrderID)
	assert.Equal(t, "PAID", status.Status)
}

func TestGetOrderPaymentStatusUseCaseImpl_Execute_Error(t *testing.T) {
	expectedErr := errors.New("fail")
	mockGateway := &mockOrderGateway{
		GetOrderPaymentStatusFunc: func(ctx context.Context, id string) (OrderPaymentStatus, error) {
			return OrderPaymentStatus{}, expectedErr
		},
	}
	usecase := NewGetOrderPaymentStatusUseCaseImpl(mockGateway)
	status, err := usecase.Execute(context.Background(), "1")
	assert.Error(t, err)
	assert.Equal(t, OrderPaymentStatus{}, status)
}
