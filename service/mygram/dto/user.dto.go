package dto

import (
	"github.com/asaskevich/govalidator"
)

type RequestReqister struct {
	Username string `json:"username" valid:"required~Username wajib diisi"`
	Email    string `json:"email" valid:"required~Email wajib diisi,email~Format email tidak sesuai"`
	Password string `json:"password" valid:"required~Password wajib diisi,minstringlength(6)~Password minimal memiliki 6 karakter"`
	Age      int    `json:"age" valid:"required~Umur wajib diisi,range(8|1000)~umur harus melebihi dari 8 tahun"`
}

func (v *RequestReqister) ValidateRegister() (err error) {
	_, errRegis := govalidator.ValidateStruct(v)
	if errRegis != nil {
		err = errRegis
		return
	}
	err = nil
	return
}

type RequestLogin struct {
	Username string `json:"username" valid:"required~Username wajib diisi"`
	Password string `json:"password" valid:"required~Password wajib diisi,minstringlength(6)~Password minimal memiliki 6 karakter"`
}

type ResponLogin struct {
	Id          uint   `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	AccessToken string `json:"access_token"`
}

func (v *RequestLogin) Validate() (err error) {
	_, errRegis := govalidator.ValidateStruct(v)
	if errRegis != nil {
		err = errRegis
		return
	}
	err = nil
	return
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      string `json:"age"`
}
