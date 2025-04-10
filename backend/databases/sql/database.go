package sql

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database interface {
	Connect(databaseName string, testEnv bool, prodEnv bool) error
	Disconnect() error
	Clear() error

	// 	CountUsersWithEmail(ctx context.Context, email string) (int64, error)
	CreateUser(ctx context.Context, newUser User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, userID string) (*User, error)

	CreateBusiness(ctx context.Context, newBusiness Business) (*Business, error)
	GetBusinessByUserID(ctx context.Context, userID string) (*Business, error)

	AddMembership(ctx context.Context, newMembership Membership) (*Membership, error)
	GetMembershipByUserId(ctx context.Context, userID string) (*Membership, error)
	GetAllMembersInBusiness(ctx context.Context, businessID int) ([]Membership, error)
	DeleteMembership(ctx context.Context, userEmail string, businessId int) error
}

type SQLDatabase struct {
	databaseURI string
	pool        *pgxpool.Pool
}

func NewSQLDatabase(databaseURI string) *SQLDatabase {
	return &SQLDatabase{databaseURI, nil}
}

func (sqlDb *SQLDatabase) Connect(databaseName string, testEnv bool, prodEnv bool) error {
	var connString string
	var err error

	if prodEnv {
		connString = sqlDb.databaseURI
	} else {
		connString = fmt.Sprintf("%s/%s", sqlDb.databaseURI, databaseName)
	}

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
		filepathToSqlFiles = "../../../databases/sql/schema.sql"
	} else {
		filepathToSqlFiles = "databases/sql/schema.sql"
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
	_, err := sqlDb.pool.Exec(context.Background(), "DELETE FROM memberships; DELETE FROM businesses; DELETE FROM users;")
	if err != nil {
		return err
	}
	return err
}

// func ParseID(id string) (string, error) {
// 	// TODO: Tämä logiikka
// 	return id, nil
// }
