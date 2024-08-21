package infrastructure

import (
	"github.com/shinpr/go_project_template/go/internal/domain/entity"
	"github.com/shinpr/go_project_template/go/internal/domain/repository"
)

type hello struct {}

func NewHello() repository.Hello {
	return &hello{}
}

func (h *hello) GetMessage() (*entity.Hello, error) {
	helloEntity := &entity.Hello{
		Message:	"Hello, World!",
	}
	return helloEntity, nil
}