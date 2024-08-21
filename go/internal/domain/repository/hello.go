package repository

import (
	"github.com/shinpr/go_project_template/go/internal/domain/entity"
)

type Hello interface {
	GetMessage() (*entity.Hello, error)
}
