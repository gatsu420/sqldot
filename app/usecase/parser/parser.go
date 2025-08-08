package parser

import (
	"errors"
	"strings"
)

func (u *usecaseImpl) ParseStrToMap() (map[string][]string, error) {
	res := map[string][]string{}
	text, err := u.llmadapter.GetQueryStructureUsingGemini()
	if err != nil {
		return nil, err
	}
	rows := strings.Split(text, "\n")

	for _, r := range rows {
		parts := strings.Split(r, ":")
		if len(parts) != 2 {
			return nil, errors.New("each row must has 2 parts")
		}

		cte := "\"" + strings.TrimSpace(parts[0]) + "\""
		source := "\"" + strings.TrimSpace(parts[1]) + "\""
		res[cte] = append(res[cte], source)
	}

	return res, nil
}
