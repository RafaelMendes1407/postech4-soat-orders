package payment

import (
	"context"
	"errors"
	"testing"

	dto "post-tech-challenge-10soat/internal/dto/payment"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/gateways"
	"github.com/stretchr/testify/assert"
)

type mockPaymentGateway struct {
	interfaces.PaymentGateway
	CreatePaymentFunc func(ctx context.Context, payment entity.Payment) (entity.Payment, error)
}

func (m *mockPaymentGateway) CreatePayment(ctx context.Context, payment entity.Payment) (entity.Payment, error) {
	return m.CreatePaymentFunc(ctx, payment)
}

func TestPaymentCheckoutUseCaseImpl_Execute_Success(t *testing.T) {
	mockGateway := &mockPaymentGateway{
		CreatePaymentFunc: func(ctx context.Context, payment entity.Payment) (entity.Payment, error) {
			payment.ID = "p1"
			return payment, nil
		},
	}
	usecase := NewPaymentCheckoutUsecaseImpl(mockGateway)
	input := dto.CreatePaymentDTO{OrderID: "o1", Amount: 100.0}
	result, err := usecase.Execute(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, "p1", result.ID)
	assert.Equal(t, "o1", result.OrderID)
	assert.Equal(t, 100.0, result.Amount)
}

func TestPaymentCheckoutUseCaseImpl_Execute_Error(t *testing.T) {
	expectedErr := errors.New("fail")
	mockGateway := &mockPaymentGateway{
		CreatePaymentFunc: func(ctx context.Context, payment entity.Payment) (entity.Payment, error) {
			return entity.Payment{}, expectedErr
		},
	}
	usecase := NewPaymentCheckoutUsecaseImpl(mockGateway)
	input := dto.CreatePaymentDTO{OrderID: "o1", Amount: 100.0}
	result, err := usecase.Execute(context.Background(), input)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to checkout payment")
	assert.Equal(t, entity.Payment{}, result)
}
