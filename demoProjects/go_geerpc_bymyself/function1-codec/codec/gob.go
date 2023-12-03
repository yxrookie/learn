package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

type GobCodec struct {
	conn io.ReadWriteCloser
	buf *bufio.Writer
	dec *gob.Decoder
	enc *gob.Encoder
}

var _ Codec = (*GobCodec)(nil)

func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec {
		conn: conn,
		buf: buf,
		dec: gob.NewDecoder(conn),
		enc: gob.NewEncoder(buf),
	}
}

func (c *GobCodec) Close() error {
	return c.conn.Close()
}

func (c *GobCodec) ReadHeader(h *Header) error {
	return c.dec.Decode(h)
}

func (c *GobCodec) Readbody(data interface{}) error {
	return c.dec.Decode(data)
}

func (c *GobCodec) Write(h *Header, data interface{}) (err error) {
	defer func() {
		_ = c.buf.Flush()
		if err != nil {
			 _ = c.conn.Close()
		}
	}()
	if err = c.enc.Encode(h); err != nil {
		log.Println("rpc codec: gob error encoding header:", err)
		return err
	}
	if err = c.enc.Encode(data); err != nil {
		log.Println("rpc codec: gob error encoding body:", err)
		return err
	}
	return nil
}
