package db

import (
	"github.com/eunanibus/gpp-grpc-demo/proto/auth"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"github.com/go-redis/redis"
)

// This class serves as an example of your database layer
// We use Redis here to simulate database persistence for this exercise
// This is a dummy project. The DB layer is for illustration, and has not been implemented.

type Database interface {
	CreateUser(ctx context.Context, u *auth.User) error
	GetUser(ctx context.Context, username string) (*auth.User, error)
}

type DatabaseManager struct {
	dbClient *redis.Client
}

func NewDatabaseManager() *DatabaseManager {
	// The error management here sucks. But it's just an example. Always handle your errors kids.
	dbClient := newClient()
	return &DatabaseManager{
		dbClient: dbClient}
}

func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	log.Println(pong, err)
	return client
}

func (db *DatabaseManager) CreateUser(ctx context.Context, u *auth.User) error {
	log.Infof("Creating new user with Username: '%s' and User ID: '%s'", u.Username, u.UserId)
	if err := db.dbClient.Set(u.UserId, u, 0); err != nil {
		log.Println(err)
		return err.Err()
	}
	log.Infof("Successfully created new user with Username: '%s' and User ID: '%s'", u.Username, u.UserId)
	return nil
}

func (db *DatabaseManager) GetUser(ctx context.Context, username string) (*auth.User, error) {
	log.Infof("Attempting to fetch User with Username: '%s'", username)
	return nil, nil
}
