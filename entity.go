package goqrcode

import (
	"errors"
	"image"
)

var (
	ErrorWidth   = errors.New("width must be greater than zero")
	ErrorHeight  = errors.New("height must be greater than zero")
	ErrorContent = errors.New("content cannot be empty")
)

type QR struct {
	width               int
	height              int
	content             string
	encode              string
	errorCorretionLevel string
	margin              int
	URL                 string
	Image               image.Image
}

func (qr *QR) Validate() error {
	if qr.width == 0 {
		return ErrorWidth
	}

	if qr.height == 0 {
		return ErrorHeight
	}

	if qr.content == "" {
		return ErrorContent
	}

	return nil
}
