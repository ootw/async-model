package async

type UDPHead struct {
	Version  uint16
	Len      uint16 // 整个数据包长度（UDP头+ASYNC头+用户数据）
	Check    uint16
	Reliable int
	Ack      int
	SN       uint32
}
