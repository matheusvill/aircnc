// +build unit

package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckEmail(t *testing.T) {
	tests := map[string]bool{
		"valido@gmai.com": true,
		"invalido":        false,
	}

	for value, resultExpected := range tests {
		resultActual := CheckEmail(value)
		assert.Exactly(t, resultExpected, resultActual, "Expected to be equal.")
	}
}
