package selfProto

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type DecEncer interface {
	Enc(interface{}) (*bytes.Buffer,error)
	Dec(buff *bytes.Buffer) (interface{}, error)
}


type Proto struct {
	Start uint16
	HeadLen uint32
	BodyLen uint32
	Head []byte
	Body []byte
	End   uint16
}

type privateProto struct {

}

func NewDecEncer() DecEncer {
	return &privateProto{}
}

func(p *privateProto) Enc(req interface{}) (*bytes.Buffer,error) {

	localReq, ok := req.(Proto)
	if !ok {
		return nil, fmt.Errorf("enconding failed")
	}
	var buff bytes.Buffer
	var err error
	err = binary.Write(&buff, binary.LittleEndian, localReq.Start)
	err = binary.Write(&buff, binary.LittleEndian, localReq.HeadLen)
	err = binary.Write(&buff, binary.LittleEndian, localReq.BodyLen)
	err = binary.Write(&buff, binary.LittleEndian, localReq.Head)
	err = binary.Write(&buff, binary.LittleEndian, localReq.Body)
	err = binary.Write(&buff, binary.LittleEndian, localReq.End)
	return &buff,err
}

func(p *privateProto) Dec(buff *bytes.Buffer) (interface{}, error) {
	var req Proto
	var err error
	err = binary.Read(buff, binary.LittleEndian, &req.Start)
	err = binary.Read(buff, binary.LittleEndian, &req.HeadLen)
	err = binary.Read(buff, binary.LittleEndian, &req.BodyLen)
	req.Head = make([]byte, req.HeadLen, req.HeadLen)
	req.Body = make([]byte, req.BodyLen, req.BodyLen)
	err = binary.Read(buff, binary.LittleEndian, req.Head)
	err = binary.Read(buff, binary.LittleEndian, req.Body)
	err = binary.Read(buff, binary.LittleEndian, &req.End)
	return req, err
}