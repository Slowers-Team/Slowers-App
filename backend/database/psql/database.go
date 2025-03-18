package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database interface {
	Connect(databaseName string, testEnv bool) error
	Disconnect() error
	Clear() error
	// 	UserOwnsEntity(ctx context.Context, UserID, EntityID ObjectID, Collection string) error

	// 	CountUsersWithEmail(ctx context.Context, email string) (int64, error)
	CreateUser(ctx context.Context, newUser User) (*User, error)
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
	pool        *pgxpool.Pool
}

//var sqlpool *pgxpool.Pool

func NewSQLDatabase(databaseURI string) *SQLDatabase {
	return &SQLDatabase{databaseURI, nil}
}

func (sqlDb *SQLDatabase) Connect(databaseName string, testEnv bool) error {
	connString := fmt.Sprintf("%s/%s", sqlDb.databaseURI, databaseName)
	var err error
	pool, err := pgxpool.New(context.Background(), connString)

	if err != nil {
		return err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return err
	}

	log.Println("Connected to PostgreSQL")

	var filepathToSqlFiles string
	if testEnv {
		filepathToSqlFiles = "../../../database/psql/schema.sql"
	} else {
		filepathToSqlFiles = "database/psql/schema.sql"
	}
	sqlQuery, err := os.ReadFile(filepathToSqlFiles)
	if err != nil {
		return err
	}

	_, err = pool.Exec(context.Background(), string(sqlQuery))
	if err != nil {
		return err
	}

	sqlFunctions, err := os.ReadFile(filepathToSqlFiles)
	if err != nil {
		return err
	}

	_, err = pool.Exec(context.Background(), string(sqlFunctions))
	if err != nil {
		return err
	}

	sqlDb.pool = pool

	return nil
}

func (sqlDb *SQLDatabase) Disconnect() error {
	sqlDb.pool.Close()
	return nil
}

func (sqlDb *SQLDatabase) Clear() error {
	fmt.Println("Clearin onnistuminen")
	_, err := sqlDb.pool.Exec(context.Background(), "DROP slowerstest;")
	return err
}

// func ParseID(id string) (string, error) {
// 	// TODO: Tämä logiikka
// 	return id, nil
// }
