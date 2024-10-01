package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *MockDatabase) AddSite(ctx context.Context, newSite Site) (*Site, error) {
	args := m.Called(ctx, newSite)
	return args.Get(0).(*Site), args.Error(1)
}

func (m *MockDatabase) GetRootSites(ctx context.Context) ([]Site, error) {
	args := m.Called(ctx)
	return args.Get(0).([]Site), args.Error(1)
}

func (m *MockDatabase) GetSite(ctx context.Context, id string) (bson.M, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(bson.M), args.Error(1)
}

func (m *MockDatabase) DeleteSite(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}
