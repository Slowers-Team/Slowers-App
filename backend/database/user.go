package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	ID       ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Role     string   `json:"role"`
}

type LogIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (mDb MongoDatabase) CountUsersWithEmail(ctx context.Context, email string) (int64, error) {
	filter := bson.M{"email": email}
	count, err := db.Collection("users").CountDocuments(ctx, filter)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (mDb MongoDatabase) CreateUser(ctx context.Context, newUser User) (*User, error) {
	result, err := db.Collection("users").InsertOne(ctx, newUser)
	if err != nil {
		return nil, err
	}

	newUser.ID = result.InsertedID.(ObjectID)
	return &newUser, nil
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

func (mDb MongoDatabase) SetUserRole(ctx context.Context, userID ObjectID, role string) error {
	update := bson.M{"$set": bson.M{"role": role}}
	_, err := db.Collection("users").UpdateByID(ctx, userID, update)

	return err
}

func (mDb MongoDatabase) GetUserByID(ctx context.Context, userID ObjectID) (*User, error) {
	user := new(User)
	filter := bson.M{"_id": userID}
	err := db.Collection("users").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return user, nil
}
