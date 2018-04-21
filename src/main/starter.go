package main

import (
	"async"
)

func main() {
	asyncModel := &async.AsyncHead{Version: async.Version1, Msgid: async.MsgidResp1112, Srcid: 1, Dstid: 1}
	buf := async.Encode(asyncModel)
	async.Decode(buf)
}
