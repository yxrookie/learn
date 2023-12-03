package function1codec

import (
	"encoding/json"
	"go_geerpc_bymyself/function1-codec/codec"
	"io"
	"log"
	"net"
)

const MagicNumber = 0x3bef5c

type Option struct {
	MagicNumber int
	CodecType   codec.Type
}

var DefaultOption = &Option{
	MagicNumber: OptionMagicNumber,
	CodecType:   codec.GobType,
}

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

func (server *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("rpc server: accept error:", err)
			return 
		}
		go server.ServeConn(conn)
	}
}

func Accept(lis net.Listener) {
	DefaultServer.Accept(lis)
}

func (server *Server) ServeConn(conn io.ReadWriteCloser) {
	defer func() {
		_ = conn.Close()
	}()
	var opt Option
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		log.Println("rpc server: options error: ", err)
		return 
	}
	if opt.MagicNumber != MagicNumber {
		log.Printf("rpc server: invalid magic number %x", opt.MagicNumber)
		return 
	}
	if _, exist := codec.NewCodecFuncMap[opt.CodecType]; !exist {
		log.Printf("rpc server: invalid codec type %s", opt.CodecType)
		return
	}
	server.ServeCodec(codec.NewCodecFuncMap[opt.CodecType](conn))
}

func (server *Server) ServeCodec(cc codec.Codec) {


}
