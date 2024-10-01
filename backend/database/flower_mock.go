package database

import (
	"context"
)

func (m *MockDatabase) GetFlowers(ctx context.Context) ([]Flower, error) {
	args := m.Called(ctx)
	return args.Get(0).([]Flower), args.Error(1)
}

func (m *MockDatabase) AddFlower(ctx context.Context, newFlower Flower) (*Flower, error) {
	args := m.Called(ctx, newFlower)
	return args.Get(0).(*Flower), args.Error(1)
}

func (m *MockDatabase) DeleteFlower(ctx context.Context, id string) (bool, error) {
	args := m.Called(ctx, id)
	return args.Bool(0), args.Error(1)
}

