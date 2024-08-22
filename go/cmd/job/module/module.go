package module

import (
	"github.com/shinpr/go_project_template/go/internal/domain/repository"
	hellousecase "github.com/shinpr/go_project_template/go/internal/usecase/hello"
)

type Module struct {
	// repositories
	helloRepository repository.Hello

	// usecases
	HelloInteractor *hellousecase.Interactor
}

func NewModule() (*Module, error) {
	m := &Module{}
	m.initRepository()
	m.initUsecase()

	return m, nil
}
