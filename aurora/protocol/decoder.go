package protocol

import (
	"bufio"
	"bytes"
	"crypto/cipher"
	"fmt"
	"io"

	errs "errors"
)

type Decoder struct {
	reader *Reader

	temp *bytes.Buffer
}

const (
	MaxPacket = 1 << 21
)

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		reader: NewReader(r),
	}
}

func (d *Decoder) DecodePacket() ([]byte, error) {
	pLen, err := d.reader.ReadVarInt()
	if err != nil {
		return nil, fmt.Errorf("error reading packet length: %v", err)
	}

	if pLen > MaxPacket || pLen == 0 {
		return nil, fmt.Errorf("invalid packet length: %v", pLen)
	}

	pl, err := d.reader.Next(pLen)
	if err != nil {
		if errs.Is(err, bufio.ErrBufferFull) {
			buf := d.buffer(pLen)
			return buf.Bytes()[:pLen], d.reader.readFull(buf, pLen)
		}
	}

	if d.temp != nil {
		PutBuffer(d.temp)
		d.temp = nil
	}

	return pl, err
}

func (d *Decoder) buffer(size int) *bytes.Buffer {
	if d.temp != nil {
		d.temp.Grow(size)
		d.temp.Reset()
		return d.temp
	}

	d.temp = GetBuffer(size)
	return d.temp
}

func (d *Decoder) EnableDecryption(block cipher.Block, iv []byte) {
	d.reader.EnableDecryption(block, iv)
}
