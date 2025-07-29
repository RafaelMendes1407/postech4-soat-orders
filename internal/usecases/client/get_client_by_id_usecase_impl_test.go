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
	GetClientByIdFunc func(ctx context.Context, id string) (entity.Client, error)
}

func (m *mockClientGateway) GetClientById(ctx context.Context, id string) (entity.Client, error) {
	return m.GetClientByIdFunc(ctx, id)
}

func TestGetClientByIdUseCaseImpl_Execute_Success(t *testing.T) {
	mockGateway := &mockClientGateway{
	   GetClientByIdFunc: func(ctx context.Context, id string) (entity.Client, error) {
			   return entity.Client{Id: id, Cpf: "12345678900", Name: "Ana", Email: "ana@email.com"}, nil
	   },
	}
	usecase := NewGetClientByIdUseCaseImpl(mockGateway)
	client, err := usecase.Execute(context.Background(), "42")
	assert.NoError(t, err)
	   assert.Equal(t, "42", client.Id)
	assert.Equal(t, "12345678900", client.Cpf)
	assert.Equal(t, "Ana", client.Name)
	assert.Equal(t, "ana@email.com", client.Email)
}

func TestGetClientByIdUseCaseImpl_Execute_Error(t *testing.T) {
	expectedErr := errors.New("not found")
	mockGateway := &mockClientGateway{
		GetClientByIdFunc: func(ctx context.Context, id string) (entity.Client, error) {
			return entity.Client{}, expectedErr
		},
	}
	usecase := NewGetClientByIdUseCaseImpl(mockGateway)
	client, err := usecase.Execute(context.Background(), "99")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get client by id")
	assert.Equal(t, entity.Client{}, client)
}
