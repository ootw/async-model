package udp

import (
	"fmt"
	"log"
	"net"
)

const dstHost = "127.0.0.1"
const dstPort = 60030

func Dial() {
	ip := net.ParseIP(dstHost)
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: dstPort}
	client, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	fmt.Printf("客户端监听:<%s>连接服务端:<%s>成功\n", client.LocalAddr().String(), client.RemoteAddr().String())
	client.Write([]byte(""))
}
