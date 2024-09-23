package usecase

import (
	"context"
	"errors"
	"final-project/service/mygram/dto"
	"final-project/service/mygram/entity"
	"strconv"
)

func (u *MyGramUsecase) PostComment(ctx context.Context, req dto.RequestPostComment, userid string) (lastID uint, err error) {

	userIDInt, err := strconv.Atoi(userid)
	if err != nil {
		return
	}
	userID := int64(userIDInt)

	strIdPhoto := strconv.Itoa(req.PhotoId)

	_, err = u.MyGramRepository.GetOnePhoto(ctx, strIdPhoto)
	if err != nil {
		return lastID, errors.New("maaf id foto tidak ditemukan")
	}

	dataEntity := entity.Comment{
		Message: req.Message,
		PhotoID: req.PhotoId,
		UserID:  userID,
	}

	lastId, err := u.MyGramRepository.CreateComment(ctx, dataEntity)
	if err != nil {
		return
	}
	lastID = lastId
	return
}

func (u *MyGramUsecase) UpdateComment(ctx context.Context, id string, userid string, req dto.RequestPostComment) (lastID uint, err error) {
	_, err = u.MyGramRepository.GetOneComment(ctx, id)
	if err != nil {
		return lastID, errors.New("maaf id komen tidak ditemukan")
	}
	_, err = u.MyGramRepository.ValidateComment(ctx, id, userid)
	if err != nil {
		return lastID, errors.New("akun kamu tidak punya akses untuk merubah pesan ini")
	}

	lastId, err := u.MyGramRepository.UpdateComment(ctx, id, req)
	if err != nil {
		return
	}
	lastID = lastId
	return
}

func (u *MyGramUsecase) DeleteComment(ctx context.Context, id string, userid string) (lastID uint, err error) {
	_, err = u.MyGramRepository.GetOneComment(ctx, id)
	if err != nil {
		return lastID, errors.New("maaf id komen tidak ditemukan")
	}
	_, err = u.MyGramRepository.ValidateComment(ctx, id, userid)
	if err != nil {
		return lastID, errors.New("akun kamu tidak punya akses untuk menghapus pesan ini")
	}

	err = u.MyGramRepository.DeleteComment(ctx, id)
	if err != nil {
		return
	}
	return
}

func (u *MyGramUsecase) GetAllComment(ctx context.Context) (comments []dto.ResponseComment, err error) {
	comments, err = u.MyGramRepository.GetAllComment(ctx)
	if err != nil {
		return
	}
	return
}

func (u *MyGramUsecase) GetOneComment(ctx context.Context, id string) (comment dto.ResponseComment, err error) {
	comment, err = u.MyGramRepository.GetOneComment(ctx, id)
	if err != nil {
		return
	}
	return
}
