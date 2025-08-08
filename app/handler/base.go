package handler

import (
	"github.com/gatsu420/sqldot/app/usecase/parser"
)

type Handler interface {
	Parse(dotFileName string) error
}

type handlerImpl struct {
	parserUsecase parser.Usecase
}

func NewHandler(parserUsecase parser.Usecase) Handler {
	return &handlerImpl{
		parserUsecase: parserUsecase,
	}
}
