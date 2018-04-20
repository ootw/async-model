package async

import (
	"bytes"
	"encoding/binary"
)

const async_head_len = 16

type async_head struct {
	version uint32
	msgid   uint32
	srcid   uint32
	dstid   uint32
}

func decode(buf *bytes.Buffer) *async_head {
	head := new(async_head)
	binary.Read(buf, binary.LittleEndian, head.version)
	binary.Read(buf, binary.LittleEndian, head.msgid)
	binary.Read(buf, binary.LittleEndian, head.srcid)
	binary.Read(buf, binary.LittleEndian, head.dstid)
	return head
}

func encode(head *async_head) *bytes.Buffer {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, head.version)
	binary.Write(buf, binary.LittleEndian, head.msgid)
	binary.Write(buf, binary.LittleEndian, head.srcid)
	binary.Write(buf, binary.LittleEndian, head.dstid)
	return buf
}
