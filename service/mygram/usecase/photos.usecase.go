package usecase

import (
	"context"
	"errors"
	"final-project/service/mygram/dto"
	"final-project/service/mygram/entity"
	"strconv"
)

func (u *MyGramUsecase) PostPhoto(ctx context.Context, req dto.RequestPostPhoto, userid string) (lastID uint, err error) {

	userIDInt, err := strconv.Atoi(userid)
	if err != nil {
		return
	}
	userID := int64(userIDInt)

	dataEntity := entity.Photo{
		Title:    req.Title,
		PhotoURL: req.PhotoUrl,
		Caption:  req.Caption,
		UserID:   userID,
	}

	lastId, err := u.MyGramRepository.CreatePhoto(ctx, dataEntity)
	if err != nil {
		return
	}
	lastID = lastId
	return
}

func (u *MyGramUsecase) UpdatePhoto(ctx context.Context, id string, userid string, req dto.RequestPostPhoto) (lastID uint, err error) {

	_, err = u.MyGramRepository.GetOnePhoto(ctx, id)
	if err != nil {
		return lastID, errors.New("maaf id foto tidak ditemukan")
	}
	_, err = u.MyGramRepository.ValidatePhoto(ctx, id, userid)
	if err != nil {
		return lastID, errors.New("akun kamu tidak punya akses untuk merubah photo ini")
	}

	lastId, err := u.MyGramRepository.UpdatePhoto(ctx, id, req)
	if err != nil {
		return
	}
	lastID = lastId
	return
}

func (u *MyGramUsecase) DeletePhoto(ctx context.Context, id string, userid string) (lastID uint, err error) {

	_, err = u.MyGramRepository.GetOnePhoto(ctx, id)
	if err != nil {
		return lastID, errors.New("maaf id foto tidak ditemukan")
	}
	_, err = u.MyGramRepository.ValidatePhoto(ctx, id, userid)
	if err != nil {
		return lastID, errors.New("akun kamu tidak punya akses untuk menghapus photo ini")
	}

	err = u.MyGramRepository.DeletePhoto(ctx, id)
	if err != nil {
		return
	}
	return
}

func (u *MyGramUsecase) GetAllPhoto(ctx context.Context) (photos []dto.ResponsePhoto, err error) {
	photos, err = u.MyGramRepository.GetAllPhoto(ctx)
	if err != nil {
		return
	}
	return
}

func (u *MyGramUsecase) GetOnePhoto(ctx context.Context, id string) (photo dto.ResponsePhoto, err error) {
	photo, err = u.MyGramRepository.GetOnePhoto(ctx, id)
	if err != nil {
		return
	}
	return
}
