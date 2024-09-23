package repository

import (
	"context"
	"errors"
	"final-project/service/mygram/dto"
	"final-project/service/mygram/entity"
)

func (r MyGramRepository) CreateSosmed(ctx context.Context, req entity.SocialMedia) (lastID uint, err error) {
	db := r.gormDB.GetDB().WithContext(ctx)

	result := db.Create(&req) // Insert into the database
	if result.Error != nil {
		err = result.Error
		return
	}

	lastID = req.ID
	return
}

func (r MyGramRepository) ValidateSosmed(ctx context.Context, id string, userid string) (sosmed entity.SocialMedia, err error) {
	db := r.gormDB.GetDB().WithContext(ctx)
	result := db.Raw(
		`SELECT * FROM mygram.social_media WHERE id = ? and user_id = ?`, id, userid,
	).Scan(&sosmed)

	if result.RowsAffected == 0 {
		err = errors.New("sosmed not found")
		return
	}

	return
}

func (r MyGramRepository) UpdateSosmed(ctx context.Context, id string, req dto.RequestPostSosmed) (lastInsertID uint, err error) {
	db := r.gormDB.GetDB().WithContext(ctx)
	_ = db.Raw(
		`UPDATE mygram.social_media SET name = ?, social_media_url = ? WHERE id = ?`,
		req.Name, req.SocialMediaUrl, id,
	).Scan(&lastInsertID)

	return
}

func (r MyGramRepository) DeleteSosmed(ctx context.Context, id string) (err error) {
	db := r.gormDB.GetDB().WithContext(ctx)
	result := db.Exec(
		`DELETE FROM mygram.social_media WHERE id = ?`, id,
	)
	if result.Error != nil {
		err = result.Error
		return
	}

	return
}

func (r MyGramRepository) GetAllSosmed(ctx context.Context) (res []dto.ResponseSosmed, err error) {
	type SosmedWithUser struct {
		ID             uint   `json:"id"`
		Name           string `json:"name"`
		SocialMediaURL string `json:"social_media_url"`
		Username       string `json:"username"`
		Email          string `json:"email"`
		Age            string `json:"age"`
	}

	var tempSosmed []SosmedWithUser

	db := r.gormDB.GetDB().WithContext(ctx)
	result := db.Raw(
		`SELECT p.id, p.name, p.social_media_url, u.username, u.email, u.age
        FROM mygram.social_media p
        JOIN mygram.user u ON p.user_id = u.id
        WHERE p.user_id IS NOT NULL AND p.user_id != 0`,
	).Scan(&tempSosmed)
	if result.RowsAffected == 0 {
		err = errors.New("social media not found")
		return
	}
	res = make([]dto.ResponseSosmed, len(tempSosmed))

	for i, x := range tempSosmed {
		res[i].ID = x.ID
		res[i].Name = x.Name
		res[i].SocialMediaUrl = x.SocialMediaURL
		res[i].User = dto.User{
			Username: x.Username,
			Email:    x.Email,
			Age:      x.Age,
		}
	}

	return res, nil
}

func (r MyGramRepository) GetOneSosmed(ctx context.Context, id string) (res dto.ResponseSosmed, err error) {
	type SosmedWithUser struct {
		ID             uint   `json:"id"`
		Name           string `json:"name"`
		SocialMediaURL string `json:"social_media_url"`
		Username       string `json:"username"`
		Email          string `json:"email"`
		Age            string `json:"age"`
	}

	var tempSosmed SosmedWithUser

	db := r.gormDB.GetDB().WithContext(ctx)
	result := db.Raw(
		`SELECT p.id, p.name, p.social_media_url, u.username, u.email, u.age
        FROM mygram.social_media p
        JOIN mygram.user u ON p.user_id = u.id
        WHERE p.id=?`, id,
	).Scan(&tempSosmed)
	if result.RowsAffected == 0 {
		err = errors.New("sosial media not found")
		return
	}

	res.ID = tempSosmed.ID
	res.Name = tempSosmed.Name
	res.SocialMediaUrl = tempSosmed.SocialMediaURL
	res.User = dto.User{
		Username: tempSosmed.Username,
		Email:    tempSosmed.Email,
		Age:      tempSosmed.Age,
	}

	return res, nil
}
