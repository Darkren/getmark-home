package pricetag

import (
	"io"
)

//go:generate mockery --name Service --filename mock_service.go --output ./ --inpackage

// Service provides methods to generate price tags.
type Service interface {
	Generate(barcode, name string, cost int64) (io.Reader, error)
}
