package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/makarellav/bloagg/internal/converter"
	db "github.com/makarellav/bloagg/internal/db/sqlc"
	"github.com/makarellav/bloagg/internal/entity"
	"time"
)

type UserService struct {
	db  *db.Queries
	ctx context.Context
}

func NewUserService(db *db.Queries, ctx context.Context) *UserService {
	return &UserService{
		db:  db,
		ctx: ctx,
	}
}

func (us *UserService) Create(user *entity.User) (*entity.User, error) {
	userParams := db.CreateUserParams{
		ID: pgtype.UUID{
			Bytes: uuid.New(),
			Valid: true,
		},
		Name: user.Name,
		CreatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
	}

	dbUser, err := us.db.CreateUser(us.ctx, userParams)

	if err != nil {
		return nil, err
	}

	newUser := converter.ToUserFromDB(dbUser)

	return newUser, nil
}

func (us *UserService) GetByKey(apiKey string) (*entity.User, error) {
	dbUser, err := us.db.GetUserByKey(us.ctx, apiKey)

	if err != nil {
		return nil, err
	}

	return converter.ToUserFromDB(dbUser), nil
}
