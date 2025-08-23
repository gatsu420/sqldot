package handler

import (
	"fmt"
	"os"
)

func (h *handlerImpl) Parse(dotFilePath string) error {
	dotMap, err := h.parserUsecase.ParseStrToMap()
	if err != nil {
		return fmt.Errorf("usecase error: %w", err)
	}

	dotFile, err := os.Create(dotFilePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer dotFile.Close()

	dotFile.Write([]byte("digraph {\n"))
	for cte, sources := range dotMap {
		for _, s := range sources {
			dag := fmt.Sprintf("\t%v -> %v\n", s, cte)
			dotFile.Write([]byte(dag))
		}
	}
	dotFile.Write([]byte("}\n"))

	return nil
}
