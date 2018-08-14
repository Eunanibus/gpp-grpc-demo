package service

import (
	"golang.org/x/net/context"
	"time"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/Eunanibus/gpp-grpc-demo/internal/app/gpp-grpc-demo/db"
	"github.com/Eunanibus/gpp-grpc-demo/proto/auth"
)

//go:generate mockgen -package=mocks -destination=mocks/db.go github.com/Eunanibus/users-service/internal/app/users-service/db Database

type UsersService struct {
	db db.Database
}

func NewUsersService(db db.Database) *UsersService {
	return &UsersService{
		db: db,
	}
}

func (s *UsersService) CreateNewUser(ctx context.Context, req *auth.CreateUserRequest) (*empty.Empty, error) {
	log.Info("INCOMING REQUEST - Create New User")
	log.Infof("Verifying username: '%s' does not exist", req.Username)
	// Hook to your database handler here, add to DB, return
	return nil, nil
}

func (s *UsersService) GetUser(ctx context.Context, req *auth.GetUserRequest) (*auth.User, error) {
	log.Info("INCOMING REQUEST - Get User")
	return s.db.GetUser(ctx, req.Username)
}

func generateUser(req *auth.CreateUserRequest) *auth.User {

	userId, err := generateUUID()
	if err != nil {
		return nil
	}

	currentTime, err := generateTimestamp()
	if err != nil {
		return nil
	}

	return &auth.User{
		Username:     req.Username,
		UserId:       userId.String(),
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.LastName,
		Role:         "user",
		CreatedAt:    currentTime,
		UpdatedAt:    currentTime,
	}
}

func generateUUID() (uuid.UUID, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return uuid.Nil, err
	}
	return uid, nil
}

func generateTimestamp() (*timestamp.Timestamp, error) {
	currentTime, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		return nil, err
	}
	return currentTime, nil
}
