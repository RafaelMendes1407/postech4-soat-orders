package client

import (
	"context"
	"errors"
	"testing"

	dto "post-tech-challenge-10soat/internal/dto/client"
	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/gateways"
	"github.com/stretchr/testify/assert"
)

type mockClientGateway struct {
	interfaces.ClientGateway
	CreateClientFunc func(ctx context.Context, client entity.Client) (entity.Client, error)
}

func (m *mockClientGateway) CreateClient(ctx context.Context, client entity.Client) (entity.Client, error) {
	return m.CreateClientFunc(ctx, client)
}

func TestCreateClientUseCaseImpl_Execute_Success(t *testing.T) {
	mockGateway := &mockClientGateway{
		CreateClientFunc: func(ctx context.Context, client entity.Client) (entity.Client, error) {
			client.ID = "123"
			return client, nil
		},
	}
	usecase := NewCreateClientUsecaseImpl(mockGateway)
	input := dto.CreateClientDTO{
		Cpf:   "11122233344",
		Name:  "João Silva",
		Email: "joao@email.com",
	}
	result, err := usecase.Execute(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, "123", result.ID)
	assert.Equal(t, input.Cpf, result.Cpf)
	assert.Equal(t, input.Name, result.Name)
	assert.Equal(t, input.Email, result.Email)
}

func TestCreateClientUseCaseImpl_Execute_Error(t *testing.T) {
	expectedErr := errors.New("db error")
	mockGateway := &mockClientGateway{
		CreateClientFunc: func(ctx context.Context, client entity.Client) (entity.Client, error) {
			return entity.Client{}, expectedErr
		},
	}
	usecase := NewCreateClientUsecaseImpl(mockGateway)
	input := dto.CreateClientDTO{
		Cpf:   "11122233344",
		Name:  "João Silva",
		Email: "joao@email.com",
	}
	result, err := usecase.Execute(context.Background(), input)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create client")
	assert.Equal(t, entity.Client{}, result)
}
