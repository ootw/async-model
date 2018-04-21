package async

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

const asyncHeadLen = 16

const Version1 uint32 = 1
const Version3 uint32 = 3

const MsgidReq1118 uint32 = 1118
const MsgidResp1112 uint32 = 1112
const MsgidReq2000 uint32 = 2000
const MsgidResp2001 uint32 = 2001

type AsyncHead struct {
	Version uint32
	Msgid   uint32
	Srcid   uint32
	Dstid   uint32
}

func Decode(buf *bytes.Buffer) *AsyncHead {
	if buf.Len() != asyncHeadLen {
		log.Printf("buf len is error")
	} else {
		log.Printf("buf len is %d", buf.Len())
	}
	head := new(AsyncHead)
	errVersion := binary.Read(buf, binary.LittleEndian, &head.Version)
	if errVersion != nil {
		log.Println(errVersion)
	}
	errMsgid := binary.Read(buf, binary.LittleEndian, &head.Msgid)
	if errMsgid != nil {
		log.Println(errMsgid)
	}
	errSrcid := binary.Read(buf, binary.LittleEndian, &head.Srcid)
	if errSrcid != nil {
		log.Println(errSrcid)
	}
	errDstid := binary.Read(buf, binary.LittleEndian, &head.Dstid)
	if errDstid != nil {
		log.Println(errDstid)
	}
	return head
}

func printAsyncHead(head *AsyncHead) {
	fmt.Println(head.Version)
	fmt.Println(head.Msgid)
	fmt.Println(head.Srcid)
	fmt.Println(head.Dstid)
}

func Encode(head *AsyncHead) *bytes.Buffer {
	printAsyncHead(head)
	buf := new(bytes.Buffer)
	err_version := binary.Write(buf, binary.LittleEndian, head.Version)
	if err_version != nil {
		log.Println(err_version)
	}
	err_msgid := binary.Write(buf, binary.LittleEndian, head.Msgid)
	if err_msgid != nil {
		log.Println(err_msgid)
	}
	err_srcid := binary.Write(buf, binary.LittleEndian, head.Srcid)
	if err_srcid != nil {
		log.Println(err_srcid)
	}
	err_dstid := binary.Write(buf, binary.LittleEndian, head.Dstid)
	if err_dstid != nil {
		log.Println(err_dstid)
	}
	return buf
}
