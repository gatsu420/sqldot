package mockllm

type adapter struct {
	llmResp string
	err     error
}

func NewAdapter(opts ...func(*adapter)) *adapter {
	mock := &adapter{}
	for _, o := range opts {
		o(mock)
	}
	return mock
}

func WithLLMResp(resp string) func(*adapter) {
	return func(a *adapter) {
		a.llmResp = resp
	}
}

func WithErr(err error) func(*adapter) {
	return func(a *adapter) {
		a.err = err
	}
}

func (a *adapter) GetQueryStructureUsingGemini() (string, error) {
	if a.err != nil {
		return "", a.err
	}
	return a.llmResp, nil
}
