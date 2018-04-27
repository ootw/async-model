package udp

import (
	"fmt"
	"log"
	"net"
)

const host = "127.0.0.1"
const port = 60030
const pageSize = 1024

func Listen() {
	ip := net.ParseIP(host)
	udpAddr := &net.UDPAddr{IP: ip, Port: port}
	server, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("服务端:<%s>启动成功...", server.LocalAddr().String())
	data := make([]byte, pageSize)
	for {
		n, remoteAddr, err := server.ReadFromUDP(data)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("接收到客户端:<%s>发送的消息【%s】", remoteAddr, data[:n])
		server.WriteToUDP([]byte("已经收到信息"), remoteAddr)
	}
}
