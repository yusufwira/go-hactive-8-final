package repository

import (
	"context"
	"errors"
	"final-project/service/mygram/dto"
	"final-project/service/mygram/entity"
)

func (r MyGramRepository) CreateComment(ctx context.Context, req entity.Comment) (lastID uint, err error) {
	db := r.gormDB.GetDB().WithContext(ctx)

	result := db.Create(&req) // Insert into the database
	if result.Error != nil {
		err = result.Error
		return
	}

	lastID = req.ID
	return
}

func (r MyGramRepository) ValidateComment(ctx context.Context, id string, userid string) (comment entity.Comment, err error) {
	db := r.gormDB.GetDB().WithContext(ctx)
	result := db.Raw(
		`SELECT * FROM mygram.comment WHERE id = ? and user_id = ?`, id, userid,
	).Scan(&comment)

	if result.RowsAffected == 0 {
		err = errors.New("comment not found")
		return
	}

	return
}

func (r MyGramRepository) UpdateComment(ctx context.Context, id string, req dto.RequestPostComment) (lastInsertID uint, err error) {
	db := r.gormDB.GetDB().WithContext(ctx)
	_ = db.Raw(
		`UPDATE mygram.comment SET message = ? WHERE id = ?`,
		req.Message, id,
	).Scan(&lastInsertID)

	return
}

func (r MyGramRepository) DeleteComment(ctx context.Context, id string) (err error) {
	db := r.gormDB.GetDB().WithContext(ctx)
	result := db.Exec(
		`DELETE FROM mygram.comment WHERE id = ?`, id,
	)
	if result.Error != nil {
		err = result.Error
		return
	}

	return
}

func (r MyGramRepository) GetAllComment(ctx context.Context) (res []dto.ResponseComment, err error) {
	type CommentWithDetails struct {
		ID       uint   `json:"id"`
		Message  string `json:"message"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Age      string `json:"age"`
		Title    string `json:"title"`
		PhotoURL string `json:"photo_url"`
		Caption  string `json:"caption"`
	}

	var tempComment []CommentWithDetails

	db := r.gormDB.GetDB().WithContext(ctx)
	result := db.Raw(
		`SELECT c.id, c.message, u.username, u.email, u.age, p.title, p.photo_url, p.caption 
		FROM mygram.comment as c
		INNER JOIN mygram.user as u on c.user_id = u.id
		INNER JOIN mygram.photo as p on c.photo_id = p.id`,
	).Scan(&tempComment)
	if result.RowsAffected == 0 {
		err = errors.New("comment not found")
		return
	}
	res = make([]dto.ResponseComment, len(tempComment))

	for i, x := range tempComment {
		res[i].ID = x.ID
		res[i].Message = x.Message
		res[i].User = dto.User{
			Username: x.Username,
			Email:    x.Email,
			Age:      x.Age,
		}
		res[i].Photo = dto.Photo{
			Title:    x.Title,
			PhotoURL: x.PhotoURL,
			Caption:  x.Caption,
		}
	}

	return res, nil
}

func (r MyGramRepository) GetOneComment(ctx context.Context, id string) (res dto.ResponseComment, err error) {
	type CommentWithDetails struct {
		ID       uint   `json:"id"`
		Message  string `json:"message"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Age      string `json:"age"`
		Title    string `json:"title"`
		PhotoURL string `json:"photo_url"`
		Caption  string `json:"caption"`
	}

	var tempComment CommentWithDetails

	db := r.gormDB.GetDB().WithContext(ctx)
	result := db.Raw(
		`SELECT c.id, c.message, u.username, u.email, u.age, p.title, p.photo_url, p.caption 
		FROM mygram.comment as c
		INNER JOIN mygram.user as u on c.user_id = u.id
		INNER JOIN mygram.photo as p on c.photo_id = p.id
		WHERE c.id = ?`, id,
	).Scan(&tempComment)
	if result.RowsAffected == 0 {
		err = errors.New("comment not found")
		return
	}

	res.ID = tempComment.ID
	res.Message = tempComment.Message
	res.User = dto.User{
		Username: tempComment.Username,
		Email:    tempComment.Email,
		Age:      tempComment.Age,
	}
	res.Photo = dto.Photo{
		Title:    tempComment.Title,
		PhotoURL: tempComment.PhotoURL,
		Caption:  tempComment.Caption,
	}

	return res, nil
}
