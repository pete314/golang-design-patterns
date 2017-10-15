// Author: Peter Nagy <https://peternagy.ie>
// Since: Oct, 2017
// Package: factory
// Description: holds interface
package factory

import (
	"github.com/pkg/errors"
)

type writer interface{
	Write(b []byte) error
	ReadAll() ([]byte, error)
}

type writerImpl struct{
	impl interface{}
}

func NewWriter(t string) (*writerImpl, error){
	switch t {
	case "memory":
		return CreateBuffer(), nil
	}

	return nil, errors.New("Writer Type not yet implemented")
}
