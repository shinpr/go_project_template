package hello

import (
	"fmt"

	"github.com/shinpr/go_project_template/go/internal/domain/repository"
)

type Interactor struct {
	helloRepository repository.Hello
}

func NewInteractor(
	helloRepository repository.Hello,
) *Interactor {
	return &Interactor{
		helloRepository: helloRepository,
	}
}

func (i *Interactor) ShowMessage() error {
	message, err := i.helloRepository.GetMessage()
	if err != nil {
		return err
	}
	fmt.Println(message.Message) // TODO: usecaseから出す
	return nil
}
