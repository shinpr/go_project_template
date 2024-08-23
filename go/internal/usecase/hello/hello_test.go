package hello

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/shinpr/go_project_template/go/internal/infrastructure"
)

func TestShowMessage(t *testing.T) {
	interactor := NewInteractor(
		infrastructure.NewHello(),
	)
	err := interactor.ShowMessage()
	require.NoError(t, err)
}
