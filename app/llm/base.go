package llm

type Adapter interface {
	GetQueryStructureUsingGemini() (string, error)
}

type adapterImpl struct{}

func NewAdapter() Adapter {
	return &adapterImpl{}
}
