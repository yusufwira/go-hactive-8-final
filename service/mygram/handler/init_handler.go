package handler

import (
	"final-project/service/mygram/usecase"
)

type MyGramHandler struct {
	MyGramUsecase usecase.IMyGramUsecase
}

func InitMyGramHandler(mygramUsecase usecase.IMyGramUsecase) *MyGramHandler {
	return &MyGramHandler{
		MyGramUsecase: mygramUsecase,
	}
}
