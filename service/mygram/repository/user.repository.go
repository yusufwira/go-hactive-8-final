package repository

import (
	"context"
	"errors"
	"final-project/service/mygram/entity"
)

func (r MyGramRepository) CreateUser(ctx context.Context, req entity.User) (lastID uint, err error) {
	db := r.gormDB.GetDB().WithContext(ctx)

	result := db.Create(&req) // Insert into the database
	if result.Error != nil {
		err = result.Error
		return
	}

	lastID = req.ID
	return
}

func (r MyGramRepository) FindUser(ctx context.Context, username string) (res entity.User, err error) {
	db := r.gormDB.GetDB().WithContext(ctx)

	result := db.Raw(
		`SELECT * FROM mygram.user where email = ? or username = ?`, username, username,
	).Scan(&res)
	if result.Error != nil {
		err = result.Error
		return
	}

	if result.RowsAffected == 0 {
		err = errors.New("user not found")
		return
	}

	return res, nil
}
