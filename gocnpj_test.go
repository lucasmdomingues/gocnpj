package gocnpj

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	service := NewService()

	t.Run("should be able return a company passing cnpj with special characters", func(t *testing.T) {
		_, err := service.Search("31.872.495/0001-72")
		assert.Nil(t, err)
	})

	t.Run("should be able return a company passing cnpj without special characters", func(t *testing.T) {
		_, err := service.Search("18236120000158")
		assert.Nil(t, err)
	})

	t.Run("should be able return a error passing a invalid cnpj", func(t *testing.T) {
		_, err := service.Search("11111111111111")
		assert.NotNil(t, err)
	})
}
