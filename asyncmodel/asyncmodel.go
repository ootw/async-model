package asyncmodel

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

const asyncModelLen = 16

const Version1 uint32 = 1
const Version3 uint32 = 3

const MsgidReq1118 uint32 = 1118
const MsgidResp1112 uint32 = 1112
const MsgidReq2000 uint32 = 2000
const MsgidResp2001 uint32 = 2001

type AsyncModel struct {
	Version uint32
	Msgid   uint32
	Srcid   uint32
	Dstid   uint32
}

func (async *AsyncModel) Decode(buf *bytes.Buffer) {
	if buf.Len() != asyncModelLen {
		log.Printf("buf len is error")
	} else {
		log.Printf("buf len is %d", buf.Len())
	}
	errVersion := binary.Read(buf, binary.LittleEndian, &async.Version)
	if errVersion != nil {
		log.Println(errVersion)
	}
	errMsgid := binary.Read(buf, binary.LittleEndian, &async.Msgid)
	if errMsgid != nil {
		log.Println(errMsgid)
	}
	errSrcid := binary.Read(buf, binary.LittleEndian, &async.Srcid)
	if errSrcid != nil {
		log.Println(errSrcid)
	}
	errDstid := binary.Read(buf, binary.LittleEndian, &async.Dstid)
	if errDstid != nil {
		log.Println(errDstid)
	}
	printAsyncModel(async)
}

func printAsyncModel(head *AsyncModel) {
	fmt.Println(head.Version)
	fmt.Println(head.Msgid)
	fmt.Println(head.Srcid)
	fmt.Println(head.Dstid)
}

func (async *AsyncModel) Encode() *bytes.Buffer {
	printAsyncModel(async)
	buf := new(bytes.Buffer)
	errVersion := binary.Write(buf, binary.LittleEndian, async.Version)
	if errVersion != nil {
		log.Println(errVersion)
	}
	errMsgid := binary.Write(buf, binary.LittleEndian, async.Msgid)
	if errMsgid != nil {
		log.Println(errMsgid)
	}
	errSrcid := binary.Write(buf, binary.LittleEndian, async.Srcid)
	if errSrcid != nil {
		log.Println(errSrcid)
	}
	errDstid := binary.Write(buf, binary.LittleEndian, async.Dstid)
	if errDstid != nil {
		log.Println(errDstid)
	}
	return buf
}
