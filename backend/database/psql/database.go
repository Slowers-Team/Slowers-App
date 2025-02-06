package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type Database interface {
	Connect(databaseName string) error
	Disconnect() error
	// 	Clear() error
	// 	UserOwnsEntity(ctx context.Context, UserID, EntityID ObjectID, Collection string) error

	// 	CountUsersWithEmail(ctx context.Context, email string) (int64, error)
	// 	CreateUser(ctx context.Context, newUser User) (*User, error)
	// 	GetUserByEmail(ctx context.Context, email string) (*User, error)
	// 	GetUserByID(ctx context.Context, userID ObjectID) (*User, error)
	// 	SetUserRole(ctx context.Context, userID ObjectID, role string) error

	// 	GetFlowers(ctx context.Context) ([]Flower, error)
	// 	GetUserFlowers(ctx context.Context, userID ObjectID) ([]Flower, error)
	// 	GetAllFlowersRelatedToSite(ctx context.Context, siteID ObjectID, userID ObjectID) ([]Flower, error)
	// 	AddFlower(ctx context.Context, newFlower Flower) (*Flower, error)
	// 	DeleteFlower(ctx context.Context, id ObjectID) (bool, error)
	// 	ToggleFlowerVisibility(ctx context.Context, userID, flowerID ObjectID) (*bool, error)
	// 	ModifyFlower(ctx context.Context, id ObjectID, newFlower Flower) (*Flower, error)
	// 	DeleteMultipleFlowers(ctx context.Context, flowerIDs []ObjectID) error

	// 	AddSite(ctx context.Context, newSite Site) (*Site, error)
	// 	GetRootSites(ctx context.Context, userID ObjectID) ([]Site, error)
	// 	GetSite(ctx context.Context, siteID ObjectID, userID ObjectID) (bson.M, error)
	// 	DeleteSite(ctx context.Context, siteID ObjectID, userID ObjectID) (*mongo.DeleteResult, error)
	// 	AddFlowerToSite(ctx context.Context, siteID ObjectID, flowerID ObjectID) error
	// 	GetSiteByID(ctx context.Context, siteID ObjectID) (*Site, error)

	// AddImage(ctx context.Context, newImage Image) (*Image, error)
	// DeleteImage(ctx context.Context, id ObjectID) (bool, error)
	// GetImagesByEntity(ctx context.Context, entityID string) ([]Image, error)
	// SetFavoriteImage(ctx context.Context, UserID, EntityID, ImageID ObjectID, Collection string) error
	// GetImageByID(ctx context.Context, imageID ObjectID) (*Image, error)
	// ClearFavoriteImage(ctx context.Context, UserID, EntityID ObjectID, Collection string) error
}

type SQLDatabase struct {
	databaseURI string
	conn        *pgx.Conn
}

var sqlconn *pgx.Conn

func NewSQLDatabase(databaseURI string) *SQLDatabase {
	return &SQLDatabase{databaseURI, nil}
}

func (sqlDb *SQLDatabase) Connect(databaseName string) error {
	connString := fmt.Sprintf("%s/%s", sqlDb.databaseURI, databaseName)
	var err error
	conn, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		return err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return err
	}

	log.Println("Connected to PostgreSQL")

	sqlconn = conn

	return nil
}

func (sqlDb *SQLDatabase) Disconnect() error {
	return sqlconn.Close(context.Background())
}
