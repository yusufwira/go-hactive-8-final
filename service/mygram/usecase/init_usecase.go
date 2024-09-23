package usecase

import (
	"context"
	"final-project/connection"
	"final-project/service/mygram/dto"
	repo "final-project/service/mygram/repository"
)

type MyGramUsecase struct {
	MyGramRepository repo.IMygramRepository
	GormDB           *connection.GormDB
}

func InitMyGramUsecase(mygramRepository repo.IMygramRepository, gormDb *connection.GormDB) IMyGramUsecase {
	return &MyGramUsecase{
		MyGramRepository: mygramRepository,
		GormDB:           gormDb,
	}
}

type IMyGramUsecase interface {
	RegisterUser(ctx context.Context, req dto.RequestReqister) (lastID uint, err error)
	LoginUser(ctx context.Context, req dto.RequestLogin) (res dto.ResponLogin, err error)

	// photos
	GetAllPhoto(ctx context.Context) (photos []dto.ResponsePhoto, err error)
	GetOnePhoto(ctx context.Context, id string) (photo dto.ResponsePhoto, err error)
	PostPhoto(ctx context.Context, req dto.RequestPostPhoto, userid string) (lastID uint, err error)
	UpdatePhoto(ctx context.Context, id string, userid string, req dto.RequestPostPhoto) (lastID uint, err error)
	DeletePhoto(ctx context.Context, id string, userid string) (lastID uint, err error)

	//sosmed
	GetAllSosmed(ctx context.Context) (sosmeds []dto.ResponseSosmed, err error)
	GetOneSosmed(ctx context.Context, id string) (sosmed dto.ResponseSosmed, err error)
	PostSosmed(ctx context.Context, req dto.RequestPostSosmed, userid string) (lastID uint, err error)
	UpdateSosmed(ctx context.Context, id string, userid string, req dto.RequestPostSosmed) (lastID uint, err error)
	DeleteSosmed(ctx context.Context, id string, userid string) (lastID uint, err error)

	// comment
	PostComment(ctx context.Context, req dto.RequestPostComment, userid string) (lastID uint, err error)
	UpdateComment(ctx context.Context, id string, userid string, req dto.RequestPostComment) (lastID uint, err error)
	DeleteComment(ctx context.Context, id string, userid string) (lastID uint, err error)
	GetAllComment(ctx context.Context) (comments []dto.ResponseComment, err error)
	GetOneComment(ctx context.Context, id string) (comment dto.ResponseComment, err error)
}
