package repository

import (
	"context"
	"errors"
	"final-project/service/mygram/dto"
	"final-project/service/mygram/entity"
)

func (r MyGramRepository) CreatePhoto(ctx context.Context, req entity.Photo) (lastID uint, err error) {
	db := r.gormDB.GetDB().WithContext(ctx)

	result := db.Create(&req) // Insert into the database
	if result.Error != nil {
		err = result.Error
		return
	}

	lastID = req.ID
	return
}

func (r MyGramRepository) ValidatePhoto(ctx context.Context, id string, userid string) (photo entity.Photo, err error) {
	db := r.gormDB.GetDB().WithContext(ctx)
	result := db.Raw(
		`SELECT * FROM mygram.photo WHERE id = ? and user_id = ?`, id, userid,
	).Scan(&photo)

	if result.RowsAffected == 0 {
		err = errors.New("photo not found")
		return
	}

	return
}

func (r MyGramRepository) UpdatePhoto(ctx context.Context, id string, req dto.RequestPostPhoto) (lastInsertID uint, err error) {
	db := r.gormDB.GetDB().WithContext(ctx)
	_ = db.Raw(
		`UPDATE mygram.photo SET title = ?, photo_url = ?, caption = ? WHERE id = ?`,
		req.Title, req.PhotoUrl, req.Caption, id,
	).Scan(&lastInsertID)

	return
}

func (r MyGramRepository) DeletePhoto(ctx context.Context, id string) (err error) {
	db := r.gormDB.GetDB().WithContext(ctx)
	result := db.Exec(
		`DELETE FROM mygram.photo WHERE id = ?`, id,
	)
	if result.Error != nil {
		err = result.Error
		return
	}

	return
}

func (r MyGramRepository) GetAllPhoto(ctx context.Context) (res []dto.ResponsePhoto, err error) {
	type PhotoWithUser struct {
		ID       uint   `json:"id"`
		Title    string `json:"title"`
		PhotoURL string `json:"photo_url"`
		Caption  string `json:"caption"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Age      string `json:"age"`
	}

	var tempPhoto []PhotoWithUser

	db := r.gormDB.GetDB().WithContext(ctx)
	result := db.Raw(
		`SELECT p.id, p.title, p.photo_url, p.caption, u.username, u.email, u.age
        FROM mygram.photo p
        JOIN mygram.user u ON p.user_id = u.id
        WHERE p.user_id IS NOT NULL AND p.user_id != 0`,
	).Scan(&tempPhoto)
	if result.RowsAffected == 0 {
		err = errors.New("photo not found")
		return
	}
	res = make([]dto.ResponsePhoto, len(tempPhoto))

	for i, x := range tempPhoto {
		res[i].ID = x.ID
		res[i].Title = x.Title
		res[i].PhotoURL = x.PhotoURL
		res[i].Caption = x.Caption
		res[i].User = dto.User{
			Username: x.Username,
			Email:    x.Email,
			Age:      x.Age,
		}
	}

	return res, nil
}

func (r MyGramRepository) GetOnePhoto(ctx context.Context, id string) (res dto.ResponsePhoto, err error) {
	type PhotoWithUser struct {
		ID       uint   `json:"id"`
		Title    string `json:"title"`
		PhotoURL string `json:"photo_url"`
		Caption  string `json:"caption"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Age      string `json:"age"`
	}

	var tempPhoto PhotoWithUser

	db := r.gormDB.GetDB().WithContext(ctx)
	result := db.Raw(
		`SELECT p.id, p.title, p.photo_url, p.caption, u.username, u.email, u.age
        FROM mygram.photo p
        JOIN mygram.user u ON p.user_id = u.id
        WHERE p.id=?`, id,
	).Scan(&tempPhoto)
	if result.RowsAffected == 0 {
		err = errors.New("photo not found")
		return
	}

	res.ID = tempPhoto.ID
	res.Title = tempPhoto.Title
	res.PhotoURL = tempPhoto.PhotoURL
	res.Caption = tempPhoto.Caption
	res.User = dto.User{
		Username: tempPhoto.Username,
		Email:    tempPhoto.Email,
		Age:      tempPhoto.Age,
	}

	return res, nil
}
