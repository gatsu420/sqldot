package mockparser

type usecase struct {
	dotMap map[string][]string
	err    error
}

func NewUsecase(opts ...func(*usecase)) *usecase {
	mock := &usecase{}
	for _, o := range opts {
		o(mock)
	}
	return mock
}

func WithDotMap(dotMap map[string][]string) func(*usecase) {
	return func(u *usecase) {
		u.dotMap = dotMap
	}
}

func WithErr(err error) func(*usecase) {
	return func(u *usecase) {
		u.err = err
	}
}

func (u *usecase) ParseStrToMap() (map[string][]string, error) {
	if u.err != nil {
		return nil, u.err
	}
	return u.dotMap, nil
}
