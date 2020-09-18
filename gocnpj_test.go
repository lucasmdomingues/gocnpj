package gocnpj

import (
	"testing"
)

func TestSearch(t *testing.T) {
	_, err := Search("27865757000102")
	if err != nil {
		t.Error(err)
		return
	}
}
