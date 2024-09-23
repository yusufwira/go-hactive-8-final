package usecase

import (
	"context"
	"errors"
	"final-project/service/mygram/dto"
	"final-project/service/mygram/entity"
	"strconv"
)

func (u *MyGramUsecase) PostSosmed(ctx context.Context, req dto.RequestPostSosmed, userid string) (lastID uint, err error) {

	userIDInt, err := strconv.Atoi(userid)
	if err != nil {
		return
	}
	userID := int64(userIDInt)

	dataEntity := entity.SocialMedia{
		Name:           req.Name,
		SocialMediaURL: req.SocialMediaUrl,
		UserID:         userID,
	}

	lastId, err := u.MyGramRepository.CreateSosmed(ctx, dataEntity)
	if err != nil {
		return
	}
	lastID = lastId
	return
}

func (u *MyGramUsecase) UpdateSosmed(ctx context.Context, id string, userid string, req dto.RequestPostSosmed) (lastID uint, err error) {
	_, err = u.MyGramRepository.GetOneSosmed(ctx, id)
	if err != nil {
		return lastID, errors.New("maaf id sosial media tidak ditemukan")
	}

	_, err = u.MyGramRepository.ValidateSosmed(ctx, id, userid)
	if err != nil {
		return lastID, errors.New("akun kamu tidak punya akses untuk merubah social media ini")
	}

	lastId, err := u.MyGramRepository.UpdateSosmed(ctx, id, req)
	if err != nil {
		return
	}
	lastID = lastId
	return
}

func (u *MyGramUsecase) DeleteSosmed(ctx context.Context, id string, userid string) (lastID uint, err error) {
	_, err = u.MyGramRepository.GetOneSosmed(ctx, id)
	if err != nil {
		return lastID, errors.New("maaf id sosial media tidak ditemukan")
	}

	_, err = u.MyGramRepository.ValidateSosmed(ctx, id, userid)
	if err != nil {
		return lastID, errors.New("akun kamu tidak punya akses untuk menghapus sosial media ini")
	}

	err = u.MyGramRepository.DeleteSosmed(ctx, id)
	if err != nil {
		return
	}
	return
}

func (u *MyGramUsecase) GetAllSosmed(ctx context.Context) (sosmeds []dto.ResponseSosmed, err error) {
	sosmeds, err = u.MyGramRepository.GetAllSosmed(ctx)
	if err != nil {
		return
	}
	return
}

func (u *MyGramUsecase) GetOneSosmed(ctx context.Context, id string) (sosmed dto.ResponseSosmed, err error) {
	sosmed, err = u.MyGramRepository.GetOneSosmed(ctx, id)
	if err != nil {
		return
	}
	return
}
