package goqrcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	service := New()

	t.Run("should be able create qr code", func(t *testing.T) {
		qr := &QR{
			width:   300,
			height:  300,
			content: "https://lucasmdomingues.dev",
		}

		err := service.Create(qr)
		assert.Nil(t, err)
		assert.NotEmpty(t, qr.URL)
		assert.NotNil(t, qr.Image)
	})

	t.Run("should be able return error width on create qr code", func(t *testing.T) {
		qr := &QR{
			height:  300,
			content: "https://lucasmdomingues.dev",
		}

		err := service.Create(qr)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ErrorWidth)
	})

	t.Run("should be able return error height on create qr code", func(t *testing.T) {
		qr := &QR{
			width:   300,
			content: "https://lucasmdomingues.dev",
		}

		err := service.Create(qr)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ErrorHeight)
	})

	t.Run("should be able return error content on create qr code", func(t *testing.T) {
		qr := &QR{
			width:  300,
			height: 300,
		}

		err := service.Create(qr)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ErrorContent)
	})
}
