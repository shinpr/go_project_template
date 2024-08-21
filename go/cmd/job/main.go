package main

import (
	"github.com/shinpr/go_project_template/go/cmd/job/module"
)

func main() {
	m, _ := module.NewModule()
	m.HelloInteractor.ShowMessage()
}