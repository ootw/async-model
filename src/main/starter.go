package main

import "asyncmodel"

func main() {
	asyncModel := asyncmodel.AsyncModel{Version: asyncmodel.Version1, Msgid: asyncmodel.MsgidResp1112, Srcid: 1, Dstid: 1}
	buf := asyncModel.Encode()
	asyncModel.Decode(buf)
	// go udp.Listen()
	// udp.Dial()
}
