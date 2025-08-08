package handler_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gatsu420/sqldot/app/handler"
	"github.com/gatsu420/sqldot/common/tests"
	mockparser "github.com/gatsu420/sqldot/mocks/app/usecase/parser"
)

func Test_Parse(t *testing.T) {
	testCases := []struct {
		caseName      string
		dotFilePath   string
		usecaseDotMap map[string][]string
		usecaseErr    error
		expectedErr   error
	}{
		{
			caseName:      "usecase has error",
			dotFilePath:   "../../example.dot",
			usecaseDotMap: nil,
			usecaseErr:    errors.New("some error"),
			expectedErr:   fmt.Errorf("usecase error: %w", errors.New("some error")),
		},
		{
			caseName:    "dotFilePath is supplied using nonexistent path",
			dotFilePath: "this/path/is/nonexistent/example.dot",
			usecaseDotMap: map[string][]string{
				"\"events\"":    {"\"datamart-staging.fact_events.events_test\""},
				"\"trx\"":       {"\"datamart-staging.fact_trx.trx_test\""},
				"\"breakdown\"": {"\"events\"", "\"trx\""},
			},
			usecaseErr: nil,
			expectedErr: fmt.Errorf("failed to create file: %w",
				errors.New("open this/path/is/nonexistent/example.dot: no such file or directory"),
			),
		},
		{
			caseName:    "dot file is created successfully",
			dotFilePath: "../../example.dot",
			usecaseDotMap: map[string][]string{
				"\"events\"":    {"\"datamart-staging.fact_events.events_test\""},
				"\"trx\"":       {"\"datamart-staging.fact_trx.trx_test\""},
				"\"breakdown\"": {"\"events\"", "\"trx\""},
			},
			usecaseErr:  nil,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			mockParserUsecase := mockparser.NewUsecase(
				mockparser.WithDotMap(tc.usecaseDotMap),
				mockparser.WithErr(tc.usecaseErr),
			)
			handler := handler.NewHandler(mockParserUsecase)

			err := handler.Parse(tc.dotFilePath)
			tests.AssertEqualError(t, err, tc.expectedErr)
		})
	}
}
