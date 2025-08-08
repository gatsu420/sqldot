package parser

import "github.com/gatsu420/sqldot/app/llm"

type Usecase interface {
	ParseStrToMap() (map[string][]string, error)
}

type usecaseImpl struct {
	llmadapter llm.Adapter
}

func NewUsecase(llmadapter llm.Adapter) Usecase {
	return &usecaseImpl{
		llmadapter: llmadapter,
	}
}
