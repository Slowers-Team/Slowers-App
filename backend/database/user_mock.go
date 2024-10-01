package database

import (
	"context"
)

func (m *MockDatabase) CountUsersWithEmail(ctx context.Context, email string) (int64, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockDatabase) CreateUser(ctx context.Context, newUser User) error {
	args := m.Called(ctx, newUser)
	return args.Error(0)
}

func (m *MockDatabase) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*User), args.Error(1)
}
