package async

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

const async_head_len = 16

type async_head struct {
	version uint32
	msgid   uint32
	srcid   uint32
	dstid   uint32
}

func Decode(buf *bytes.Buffer) *async_head {
	head := new(async_head)
	err_version := binary.Read(buf, binary.LittleEndian, head.version)
	if err_version != nil {
		log.Println(err_version)
	}
	err_msgid := binary.Read(buf, binary.LittleEndian, head.msgid)
	if err_msgid != nil {
		log.Println(err_msgid)
	}
	err_srcid := binary.Read(buf, binary.LittleEndian, head.srcid)
	if err_srcid != nil {
		log.Println(err_srcid)
	}
	err_dstid := binary.Read(buf, binary.LittleEndian, head.dstid)
	if err_dstid != nil {
		log.Println(err_dstid)
	}
	fmt.Println(head.version)
	fmt.Println(head.msgid)
	fmt.Println(head.srcid)
	fmt.Println(head.dstid)
	return head
}

func Encode(head *async_head) *bytes.Buffer {
	buf := new(bytes.Buffer)
	err_version := binary.Write(buf, binary.LittleEndian, head.version)
	if err_version != nil {
		log.Println(err_version)
	}
	err_msgid := binary.Write(buf, binary.LittleEndian, head.msgid)
	if err_msgid != nil {
		log.Println(err_msgid)
	}
	err_srcid := binary.Write(buf, binary.LittleEndian, head.srcid)
	if err_srcid != nil {
		log.Println(err_srcid)
	}
	err_dstid := binary.Write(buf, binary.LittleEndian, head.dstid)
	if err_dstid != nil {
		log.Println(err_dstid)
	}
	return buf
}
