package module

import (
        "github.com/shinpr/go_project_template/go/internal/infrastructure"
        hellousecase "github.com/shinpr/go_project_template/go/internal/usecase/hello"
)

func (m *Module) initRepository() {
	m.helloRepository = infrastructure.NewHello()
}

func (m *Module) initUsecase() {
	m.HelloInteractor = hellousecase.NewInteractor(
		m.helloRepository,
	)
}