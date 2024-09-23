package dto

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type RequestPostComment struct {
	Message string `json:"message" valid:"required~Pesan wajib diisi"`
	PhotoId int    `json:"photo_id"`
}

func (v *RequestPostComment) Validate() (err error) {
	_, errRegis := govalidator.ValidateStruct(v)
	if errRegis != nil {
		err = errRegis
		return
	}
	err = nil
	return
}

type ResponseComment struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user"`
	Photo     Photo     `json:"photo"`
}
