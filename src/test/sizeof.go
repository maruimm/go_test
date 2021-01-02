package main

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

func(p *privateProto) Enc(req interface{}) (*bytes.Buffer,error) {

	localReq, ok := req.(Req)
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
	var req Req
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


func main() {
	var p privateProto

	head := []byte{0x33,0x33,0x33,0x33}
	headLen := uint32(len(head))
	body := []byte{0x22,0x22,0x22}
	bodyLen := uint32(len(body))
	req := Req{
		Start:0x2,
		HeadLen:headLen,
		BodyLen:bodyLen,
		Body:body,
		Head:head,
		End: 0x3,
	}
	sReq, err := p.Enc(req)
	if err != nil {
		return
	}
	fmt.Printf("en req:%+v\n", sReq.Bytes())

	r, err := p.Dec(sReq)

	if err != nil {
		return
	}
	fmt.Printf("de req:%+v\n", r.(Req))
}
