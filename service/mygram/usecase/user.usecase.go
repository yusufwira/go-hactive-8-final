package usecase

import (
	"context"
	"errors"
	"final-project/helpers"
	"final-project/service/mygram/dto"
	"final-project/service/mygram/entity"
)

func (u *MyGramUsecase) RegisterUser(ctx context.Context, req dto.RequestReqister) (lastID uint, err error) {
	dataEntity := entity.User{
		Username: req.Username,
		Email:    req.Email,
		Age:      req.Age,
		Password: helpers.MakeHash(req.Password),
	}

	lastId, err := u.MyGramRepository.CreateUser(ctx, dataEntity)
	if err != nil {
		return
	}
	lastID = lastId
	return
}

func (u *MyGramUsecase) LoginUser(ctx context.Context, req dto.RequestLogin) (res dto.ResponLogin, err error) {
	dataUser, err := u.MyGramRepository.FindUser(ctx, req.Username)
	if err != nil {
		return res, errors.New("invalid email/username")
	}
	validPassword := helpers.ValidatePassword([]byte(dataUser.Password), []byte(req.Password))
	if !validPassword {
		return res, errors.New("invalid password")
	}
	token := helpers.GenerateToken(dataUser.ID, dataUser.Email)
	res.Id = dataUser.ID
	res.Username = dataUser.Username
	res.Email = dataUser.Email
	res.Age = dataUser.Age
	res.AccessToken = token
	return

}
