package client

import (
	"context"
	"errors"
	"testing"

	entity "post-tech-challenge-10soat/internal/entities"
	interfaces "post-tech-challenge-10soat/internal/interfaces/gateways"
	"github.com/stretchr/testify/assert"
)

type mockClientGateway struct {
	interfaces.ClientGateway
	GetClientByCpfFunc func(ctx context.Context, cpf string) (entity.Client, error)
}

func (m *mockClientGateway) GetClientByCpf(ctx context.Context, cpf string) (entity.Client, error) {
	return m.GetClientByCpfFunc(ctx, cpf)
}

func TestGetClientByCpfUseCaseImpl_Execute_Success(t *testing.T) {
	mockGateway := &mockClientGateway{
		GetClientByCpfFunc: func(ctx context.Context, cpf string) (entity.Client, error) {
			return entity.Client{ID: "1", Cpf: cpf, Name: "Maria", Email: "maria@email.com"}, nil
		},
	}
	usecase := NewGetClientByCpfUseCaseImpl(mockGateway)
	client, err := usecase.Execute(context.Background(), "12345678900")
	assert.NoError(t, err)
	assert.Equal(t, "1", client.ID)
	assert.Equal(t, "12345678900", client.Cpf)
	assert.Equal(t, "Maria", client.Name)
	assert.Equal(t, "maria@email.com", client.Email)
}

func TestGetClientByCpfUseCaseImpl_Execute_Error(t *testing.T) {
	expectedErr := errors.New("not found")
	mockGateway := &mockClientGateway{
		GetClientByCpfFunc: func(ctx context.Context, cpf string) (entity.Client, error) {
			return entity.Client{}, expectedErr
		},
	}
	usecase := NewGetClientByCpfUseCaseImpl(mockGateway)
	client, err := usecase.Execute(context.Background(), "00000000000")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get client by cpf")
	assert.Equal(t, entity.Client{}, client)
}
