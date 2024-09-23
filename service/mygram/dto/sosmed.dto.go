package dto

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type RequestPostSosmed struct {
	Name           string `json:"name" valid:"required~Nama wajib diisi"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~Sosial Media Url wajib diisi,url~Sosial Media Url harus berupa URL"`
}

func (v *RequestPostSosmed) Validate() (err error) {
	_, errRegis := govalidator.ValidateStruct(v)
	if errRegis != nil {
		err = errRegis
		return
	}
	err = nil
	return
}

type ResponseSosmed struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           User      `json:"user"`
}
