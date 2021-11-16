package goqrcode

import (
	"fmt"
	"image/png"
	"net/http"
	"net/url"
)

type service struct {
	BaseURL *url.URL
}

type Service interface {
	Create(qr *QR) error
}

func New() Service {
	return &service{
		BaseURL: &url.URL{
			Scheme: "https",
			Host:   "chart.googleapis.com",
			Path:   "chart",
		},
	}
}

func (s *service) Create(qr *QR) error {
	if err := qr.Validate(); err != nil {
		return err
	}

	rel, err := s.BaseURL.Parse("")
	if err != nil {
		return err
	}

	q := make(url.Values)
	q.Set("cht", "qr")
	q.Set("chs", fmt.Sprintf("%dx%d", qr.width, qr.height))
	q.Set("chl", qr.content)

	if qr.encode != "" {
		q.Set("choe", qr.encode)
	}

	if qr.errorCorretionLevel != "" && qr.margin != 0 {
		q.Set("chld", fmt.Sprintf("%s|%d", qr.errorCorretionLevel, qr.margin))
	}

	rel.RawQuery = q.Encode()

	resp, err := http.Get(rel.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	qr.URL = rel.String()

	img, err := png.Decode(resp.Body)
	if err != nil {
		return err
	}

	qr.Image = img

	return nil
}
