package asyncmodel

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const udpReliaLen, udpUnreliaLen = 12, 8

type UDPReliableModel struct {
	Version      uint16
	Len          uint16 // 整个数据包长度（UDP头+ASYNC头+用户数据）
	CheckSum     uint16
	ReliableFlag int
	ACkFlag      int
	SN           uint32
}

type UDPUnreliableModel struct {
	Version      uint16
	Len          uint16 // 整个数据包长度（UDP头+ASYNC头+用户数据）
	CheckSum     uint16
	ReliableFlag int
	Padding      int
}

func (udpRelia UDPReliableModel) Encode() *bytes.Buffer {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, udpRelia.Version)
	binary.Write(buf, binary.LittleEndian, udpRelia.Len)
	binary.Write(buf, binary.LittleEndian, udpRelia.CheckSum)
	binary.Write(buf, binary.LittleEndian, udpRelia.ReliableFlag)
	binary.Write(buf, binary.LittleEndian, udpRelia.ACkFlag)
	binary.Write(buf, binary.LittleEndian, udpRelia.SN)
	return buf
}

func (udpRelia UDPReliableModel) Decode(buf *bytes.Buffer) {
	binary.Read(buf, binary.LittleEndian, &udpRelia.Version)
	binary.Read(buf, binary.LittleEndian, &udpRelia.Len)
	binary.Read(buf, binary.LittleEndian, &udpRelia.CheckSum)
	binary.Read(buf, binary.LittleEndian, &udpRelia.ReliableFlag)
	binary.Read(buf, binary.LittleEndian, &udpRelia.ACkFlag)
	binary.Read(buf, binary.LittleEndian, &udpRelia.SN)
	printUDPReliableModel(&udpRelia)
}

func (udpUnrelia UDPUnreliableModel) Encode() *bytes.Buffer {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, udpUnrelia.Version)
	binary.Write(buf, binary.LittleEndian, udpUnrelia.Len)
	binary.Write(buf, binary.LittleEndian, udpUnrelia.CheckSum)
	binary.Write(buf, binary.LittleEndian, udpUnrelia.ReliableFlag)
	binary.Write(buf, binary.LittleEndian, udpUnrelia.Padding)
	return buf
}

func (udpUnrelia UDPUnreliableModel) Decode(buf *bytes.Buffer) {
	binary.Read(buf, binary.LittleEndian, &udpUnrelia.Version)
	binary.Read(buf, binary.LittleEndian, &udpUnrelia.Len)
	binary.Read(buf, binary.LittleEndian, &udpUnrelia.CheckSum)
	binary.Read(buf, binary.LittleEndian, &udpUnrelia.ReliableFlag)
	binary.Read(buf, binary.LittleEndian, &udpUnrelia.Padding)
	printUDPUnreliableModel(&udpUnrelia)
}

func printUDPReliableModel(udpRelia *UDPReliableModel) {
	fmt.Println(udpRelia.Version)
	fmt.Println(udpRelia.Len)
	fmt.Println(udpRelia.CheckSum)
	fmt.Println(udpRelia.ReliableFlag)
	fmt.Println(udpRelia.ACkFlag)
	fmt.Println(udpRelia.SN)
}

func printUDPUnreliableModel(udpUnrelia *UDPUnreliableModel) {
	fmt.Println(udpUnrelia.Version)
	fmt.Println(udpUnrelia.Len)
	fmt.Println(udpUnrelia.CheckSum)
	fmt.Println(udpUnrelia.ReliableFlag)
	fmt.Println(udpUnrelia.Padding)
}
