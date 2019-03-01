package gocnpj

import (
	"testing"
)

func TestBuscaCNPJ(t *testing.T) {

	_, err := BuscaCNPJ("06.813.678/0001-70")
	if err != nil {
		t.Error()
	}
}
