package main

import (
	"flag"
	"log"
	"strings"

	"github.com/gatsu420/sqldot/app/handler"
	"github.com/gatsu420/sqldot/app/llm"
	"github.com/gatsu420/sqldot/app/usecase/parser"
)

func main() {
	llmAdapter := llm.NewAdapter()
	parserUsecase := parser.NewUsecase(llmAdapter)
	handler := handler.NewHandler(parserUsecase)

	dotFilePathFlag := flag.String("output", "", "path of dot file output")
	flag.Parse()
	if !strings.HasSuffix(*dotFilePathFlag, ".dot") {
		log.Fatal("output argument must has .dot suffix")
	}

	err := handler.Parse(*dotFilePathFlag)
	if err != nil {
		log.Fatal(err)
	}
}
