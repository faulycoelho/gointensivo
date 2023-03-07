package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_Get_An_Error_If_Id_Is_Blank(t *testing.T) {
	order := Order{}
	err := order.Validate()
	if err == nil {
		t.Error("Expected error, got nil")
	}

	assert.Error(t, order.Validate(), "invalid id")
}
