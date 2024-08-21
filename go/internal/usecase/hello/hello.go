package hello

import (
	"github.com/shinpr/go_project_template/go/internal/domain/repository"
	"fmt"
)

type Interactor struct {
	helloRepository	repository.Hello
}

func NewInteractor(
	helloRepository	repository.Hello,
) *Interactor {
	return &Interactor{
		helloRepository:	helloRepository,
	}
}

func (i *Interactor) ShowMessage() {
	message, _ := i.helloRepository.GetMessage()
	fmt.Println(message.Message)
}