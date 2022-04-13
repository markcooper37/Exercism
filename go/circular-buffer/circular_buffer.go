package circular

import "errors"

type Buffer struct {
	buffer  []byte
	maxSize int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{buffer: make([]byte, 0, size), maxSize: size}
}

func (b *Buffer) ReadByte() (byte, error) {
	if len(b.buffer) == 0 {
		return 0, errors.New("buffer contains no bytes")
	} else {
		readByte := b.buffer[0]
		b.buffer = b.buffer[1:]
		return readByte, nil
	}
}

func (b *Buffer) WriteByte(c byte) error {
	if len(b.buffer) == b.maxSize {
		return errors.New("buffer is full")
	} else {
		b.buffer = append(b.buffer, c)
		return nil
	}
}

func (b *Buffer) Overwrite(c byte) {
	if len(b.buffer) < b.maxSize {
		b.buffer = append(b.buffer, c)
	} else {
		b.buffer[0] = c
		b.buffer = append(b.buffer[1:], b.buffer[0])
	}
}

func (b *Buffer) Reset() {
	b.buffer = make([]byte, 0, b.maxSize)
}
