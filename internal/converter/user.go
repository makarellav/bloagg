package converter

import (
	db "github.com/makarellav/bloagg/internal/db/sqlc"
	"github.com/makarellav/bloagg/internal/entity"
)

func ToUserFromDB(dbUser db.User) *entity.User {
	return &entity.User{
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt.Time,
		UpdatedAt: dbUser.UpdatedAt.Time,
		ID:        dbUser.ID.Bytes,
		ApiKey:    dbUser.ApiKey,
	}
}
