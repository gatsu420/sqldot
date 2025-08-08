package llm

func (g *adapterImpl) GetQueryStructureUsingGemini() (string, error) {
	return `events: public.datamart.test_table
		trx: public.datamart.another_test_table
		summary: events
		summary: trx`, nil
}
