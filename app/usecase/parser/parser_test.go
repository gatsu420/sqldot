package parser_test

import (
	"errors"
	"testing"

	"github.com/gatsu420/sqldot/app/usecase/parser"
	"github.com/gatsu420/sqldot/common/tests"
	mockllm "github.com/gatsu420/sqldot/mocks/app/llm"
)

func Test_ParseStrToMap(t *testing.T) {
	testCases := []struct {
		caseName       string
		llmAdapterText string
		llmAdapterErr  error
		expectedMap    map[string][]string
		expectedErr    error
	}{
		{
			caseName:       "llm adapter has error",
			llmAdapterText: "",
			llmAdapterErr:  errors.New("some error"),
			expectedMap:    nil,
			expectedErr:    errors.New("some error"),
		},
		{
			caseName: "text has row with less than 2 parts",
			llmAdapterText: `events: datamart-staging.fact_events.events_test
				trx
				breakdown: events
				breakdown: trx`,
			llmAdapterErr: nil,
			expectedMap:   nil,
			expectedErr:   errors.New("each row must has 2 parts"),
		},
		{
			caseName: "text has row with more than 2 parts",
			llmAdapterText: `events: datamart-staging.fact_events.events_test
				trx: datamart-staging.fact_trx.trx_test: this part messed up the row
				breakdown: events
				breakdown: trx`,
			llmAdapterErr: nil,
			expectedMap:   nil,
			expectedErr:   errors.New("each row must has 2 parts"),
		},
		{
			caseName: "text is parsed successfully",
			llmAdapterText: `events: datamart-staging.fact_events.events_test
				trx: datamart-staging.fact_trx.trx_test
				breakdown: events
				breakdown: trx`,
			llmAdapterErr: nil,
			expectedMap: map[string][]string{
				"\"events\"":    {"\"datamart-staging.fact_events.events_test\""},
				"\"trx\"":       {"\"datamart-staging.fact_trx.trx_test\""},
				"\"breakdown\"": {"\"events\"", "\"trx\""},
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			mockLLMAdapter := mockllm.NewAdapter(
				mockllm.WithLLMResp(tc.llmAdapterText),
				mockllm.WithErr(tc.llmAdapterErr),
			)
			usecase := parser.NewUsecase(mockLLMAdapter)

			dotMap, err := usecase.ParseStrToMap()
			tests.AssertEqualObject(t, dotMap, tc.expectedMap)
			tests.AssertEqualObject(t, err, tc.expectedErr)
		})
	}
}
