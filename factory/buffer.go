// Author: Peter Nagy <https://peternagy.ie>
// Since: Oct, 2017
// Package: factory
// Description: wrapper for buffer for factory sample
package factory

import "bytes"

type Buffer struct {
	b *bytes.Buffer
}

func CreateBuffer() *writerImpl{
	b := &Buffer{b: bytes.NewBuffer(nil)}
	return &writerImpl{impl:b}
}

func (b *Buffer) Write(bts []byte) error{
	_, err := b.b.Write(bts)

	return err
}

func(b *Buffer) ReadAll() ([]byte, error){
	return b.b.Bytes(), nil
}