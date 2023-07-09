package pricetag

import (
	"io"
)

// Service provides methods to generate price tags.
type Service interface {
	Generate(barcode, name string, cost int64) (io.Reader, error)
}
