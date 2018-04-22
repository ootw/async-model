package asyncmodel

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type UDPHeadReliable struct {
	Version      uint16
	Len          uint16 // 整个数据包长度（UDP头+ASYNC头+用户数据）
	CheckSum     uint16
	ReliableFlag int
	ACkFlag      int
	SN           uint32
}

type UDPHeadUnreliable struct {
	Version      uint16
	Len          uint16 // 整个数据包长度（UDP头+ASYNC头+用户数据）
	CheckSum     uint16
	ReliableFlag int
	Padding      int
}

func (udpHeadRelia UDPHeadReliable) Encode() *bytes.Buffer {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, udpHeadRelia.Version)
	binary.Write(buf, binary.LittleEndian, udpHeadRelia.Len)
	binary.Write(buf, binary.LittleEndian, udpHeadRelia.CheckSum)
	binary.Write(buf, binary.LittleEndian, udpHeadRelia.ReliableFlag)
	binary.Write(buf, binary.LittleEndian, udpHeadRelia.ACkFlag)
	binary.Write(buf, binary.LittleEndian, udpHeadRelia.SN)
	return buf
}

func (udpHeadRelia UDPHeadReliable) Decode(buf *bytes.Buffer) {
	binary.Read(buf, binary.LittleEndian, &udpHeadRelia.Version)
	binary.Read(buf, binary.LittleEndian, &udpHeadRelia.Len)
	binary.Read(buf, binary.LittleEndian, &udpHeadRelia.CheckSum)
	binary.Read(buf, binary.LittleEndian, &udpHeadRelia.ReliableFlag)
	binary.Read(buf, binary.LittleEndian, &udpHeadRelia.ACkFlag)
	binary.Read(buf, binary.LittleEndian, &udpHeadRelia.SN)
	printUDPHeadReliable(&udpHeadRelia)
}

func printUDPHeadReliable(udpHeadRelia *UDPHeadReliable) {
	fmt.Println(udpHeadRelia.Version)
	fmt.Println(udpHeadRelia.Len)
	fmt.Println(udpHeadRelia.CheckSum)
	fmt.Println(udpHeadRelia.ReliableFlag)
	fmt.Println(udpHeadRelia.ACkFlag)
	fmt.Println(udpHeadRelia.SN)
}
