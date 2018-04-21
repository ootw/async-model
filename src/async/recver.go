package async

import (
	"fmt"
	"log"
	"net"
)

const host = "127.0.0.1"
const port = 60030
const bufInitSize = 1024

//http://blog.jobbole.com/107004/
func recv() {
	ip := net.ParseIP(host)
	udpAddr := &net.UDPAddr{IP: ip, Port: port}
	server, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("服务端:<%s>启动成功...", server.LocalAddr().String())
	data := make([]byte, bufInitSize)
	for {
		n, remoteAddr, err := server.ReadFromUDP(data)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("接收到客户端:<%s>发送的消息【%s】", remoteAddr, data[:n])
		server.WriteToUDP([]byte("以及"), remoteAddr)
	}
}
