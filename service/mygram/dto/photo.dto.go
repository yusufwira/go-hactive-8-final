package dto

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type RequestPostPhoto struct {
	Title    string `json:"title" valid:"required~Judul Foto wajib diisi"`
	PhotoUrl string `json:"photo_url" valid:"required~Photo Url wajib diisi,url~Photo Url harus berupa URL"`
	Caption  string `json:"caption"`
}

func (v *RequestPostPhoto) Validate() (err error) {
	_, errRegis := govalidator.ValidateStruct(v)
	if errRegis != nil {
		err = errRegis
		return
	}
	err = nil
	return
}

type ResponsePhoto struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user"`
}

type Photo struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
}
