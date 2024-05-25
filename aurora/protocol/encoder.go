package protocol

import (
	"bytes"
	"crypto/cipher"

	"github.com/MinimixMC/AuroraAPI/aurora/protocol/crypto"
)

type Encoder struct {
	buf *bytes.Buffer

	cipher *crypto.CFB8

	threshold int

	headerSize int
}

func NewEncoder(b *bytes.Buffer) *Encoder {
	return &Encoder{
		buf:        b,
		threshold:  -1,
		headerSize: 3, //max pk length in bytes
	}
}

func (enc *Encoder) EncodePacket(data []byte) error {
	pkLen := len(data)
	dataLength := -1

	enc.buf.Grow(enc.headerSize) //ensures the max header can fit

	start := enc.buf.Len()
	//makes room for the header
	copy(enc.buf.Bytes()[start+enc.headerSize:enc.buf.Cap()], data)

	//changes the writers position to write at the space we created
	enc.buf.Truncate(start)

	enc.writeHeader(pkLen, dataLength)

	start += enc.headerSize //updates the position where the data starts

	enc.buf.Write(enc.buf.Bytes()[start : start+len(data)])

	return nil
}

func (enc *Encoder) Flush() []byte {
	data := enc.buf.Bytes()
	enc.buf.Reset()
	if enc.cipher != nil {
		enc.cipher.XORKeyStream(data, data)
	}

	return data
}

func writeVarInt(b *bytes.Buffer, n int) {
	ux := uint32(n)

	for ux >= 0x80 {
		b.WriteByte(byte(ux&0x7F) | 0x80)
		ux >>= 7
	}

	b.WriteByte(byte(ux))
}

func (enc *Encoder) writeHeader(pkLen, dataLen int) {
	writeVarInt(enc.buf, pkLen)
	if dataLen != -1 {
		writeVarInt(enc.buf, dataLen)
	}
}

// VarIntSize returns the number of bytes n takes up
func VarIntSize[T int | int32](n T) (i int) {
	for n >= 0x80 {
		n >>= 7
		i++
	}
	i++
	return
}

func (enc *Encoder) EnableEncryption(block cipher.Block, iv []byte) {
	enc.cipher = crypto.NewCFB8(block, iv, false)
}
