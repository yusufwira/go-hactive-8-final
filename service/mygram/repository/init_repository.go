package repository

import (
	"context"
	"final-project/connection"
	"final-project/service/mygram/dto"
	"final-project/service/mygram/entity"
)

type MyGramRepository struct {
	gormDB *connection.GormDB
}

func InitRequestOrderRepository(gormDB *connection.GormDB) IMygramRepository {
	return &MyGramRepository{
		gormDB: gormDB,
	}
}

type IMygramRepository interface {
	// User
	CreateUser(ctx context.Context, req entity.User) (lastID uint, err error)
	FindUser(ctx context.Context, email string) (res entity.User, err error)

	// Photo
	GetAllPhoto(ctx context.Context) (res []dto.ResponsePhoto, err error)
	GetOnePhoto(ctx context.Context, id string) (res dto.ResponsePhoto, err error)
	CreatePhoto(ctx context.Context, req entity.Photo) (lastID uint, err error)
	ValidatePhoto(ctx context.Context, id string, userid string) (photo entity.Photo, err error)
	UpdatePhoto(ctx context.Context, id string, req dto.RequestPostPhoto) (lastInsertID uint, err error)
	DeletePhoto(ctx context.Context, id string) (err error)

	// Sosial media
	GetAllSosmed(ctx context.Context) (res []dto.ResponseSosmed, err error)
	GetOneSosmed(ctx context.Context, id string) (res dto.ResponseSosmed, err error)
	CreateSosmed(ctx context.Context, req entity.SocialMedia) (lastID uint, err error)
	ValidateSosmed(ctx context.Context, id string, userid string) (sosmed entity.SocialMedia, err error)
	UpdateSosmed(ctx context.Context, id string, req dto.RequestPostSosmed) (lastInsertID uint, err error)
	DeleteSosmed(ctx context.Context, id string) (err error)

	// comment
	GetAllComment(ctx context.Context) (res []dto.ResponseComment, err error)
	GetOneComment(ctx context.Context, id string) (res dto.ResponseComment, err error)
	CreateComment(ctx context.Context, req entity.Comment) (lastID uint, err error)
	ValidateComment(ctx context.Context, id string, userid string) (comment entity.Comment, err error)
	UpdateComment(ctx context.Context, id string, req dto.RequestPostComment) (lastInsertID uint, err error)
	DeleteComment(ctx context.Context, id string) (err error)
}
