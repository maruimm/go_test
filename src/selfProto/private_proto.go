package selfProto

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

type DecEncer interface {
	Enc(interface{}) (*bytes.Buffer,error)
	Dec(buff *bytes.Buffer) (interface{}, error)
	StartLen() uint32
}


type Proto struct {
	Start uint16
	Head []byte
	Body []byte
	End   uint16
}

type privateProto struct {

}

func NewDecEncer() DecEncer {
	return &privateProto{}
}

func(p *privateProto) StartLen() uint32 {

	return uint32(unsafe.Sizeof(uint16(1)) + unsafe.Sizeof(uint32(1)) + unsafe.Sizeof(uint32(1)))
}


func(p *privateProto) Enc(req interface{}) (*bytes.Buffer,error) {

	localReq, ok := req.(Proto)
	if !ok {
		return nil, fmt.Errorf("enconding failed")
	}
	var buff bytes.Buffer
	var err error
	var headLen, bodyLen uint32
	headLen = uint32(len(localReq.Head))
	bodyLen = uint32(len(localReq.Body))
	err = binary.Write(&buff, binary.LittleEndian, localReq.Start)
	err = binary.Write(&buff, binary.LittleEndian, headLen)
	err = binary.Write(&buff, binary.LittleEndian, bodyLen)
	err = binary.Write(&buff, binary.LittleEndian, localReq.Head)
	err = binary.Write(&buff, binary.LittleEndian, localReq.Body)
	err = binary.Write(&buff, binary.LittleEndian, localReq.End)
	return &buff,err
}

func(p *privateProto) Dec(buff *bytes.Buffer) (interface{}, error) {
	var req Proto
	var err error
	var HeadLen,BodyLen uint32
	err = binary.Read(buff, binary.LittleEndian, &req.Start)
	err = binary.Read(buff, binary.LittleEndian, &HeadLen)
	err = binary.Read(buff, binary.LittleEndian, &BodyLen)
	req.Head = make([]byte, HeadLen, HeadLen)
	req.Body = make([]byte, BodyLen, BodyLen)
	err = binary.Read(buff, binary.LittleEndian, req.Head)
	err = binary.Read(buff, binary.LittleEndian, req.Body)
	err = binary.Read(buff, binary.LittleEndian, &req.End)
	return req, err
}