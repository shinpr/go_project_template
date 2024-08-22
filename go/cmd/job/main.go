package main

import (
	"log/slog"

	"github.com/shinpr/go_project_template/go/cmd/job/module"
)

func main() {
	m, err := module.NewModule()
	if err != nil {
		slog.Error("Module create error", slog.Any("error", err))
	}
	err = m.HelloInteractor.ShowMessage()
	if err != nil {
		slog.Error("ShowMessage error", slog.Any("error", err))
	}
}
