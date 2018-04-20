package async

import "io"

type async_head struct {
	version uint32
	msgid   uint32
	srcid   uint32
	dstid   uint32
}

func decode(src io.Reader) *async_head {
	head := new(async_head)

	return head
}
