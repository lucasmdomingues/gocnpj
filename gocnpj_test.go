package gocnpj

import (
	"testing"
)

func TestBuscaCNPJ(t *testing.T) {

	_, err := BuscaCNPJ("1111111111")
	if err != nil {
		t.Error()
	}
}
