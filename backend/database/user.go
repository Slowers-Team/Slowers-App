package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
}

type LogIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Role string

const (
	Grower   = "grower"
	Retailer = "retailer"
)

func (mDb MongoDatabase) CountUsersWithEmail(ctx context.Context, email string) (int64, error) {
	filter := bson.M{"email": email}
	count, err := db.Collection("users").CountDocuments(ctx, filter)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (mDb MongoDatabase) CreateUser(ctx context.Context, newUser User) error {
	_, err := db.Collection("users").InsertOne(ctx, newUser)
	return err
}

func (mDb MongoDatabase) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	user := new(User)
	filter := bson.M{"email": email}
	err := db.Collection("users").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (mDb MongoDatabase) SetUserRole(ctx context.Context, userID primitive.ObjectID, role Role) error {
	update := bson.M{"$set": bson.M{"role": role}}
	_, err := db.Collection("users").UpdateByID(ctx, userID, update)

	return err
}
