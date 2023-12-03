package codec

import "io"

type Header struct {
	ServiceMethod string
	Seq int
	Error string
}

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	Readbody(interface{}) error
	Write(*Header, interface{}) error 
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}

