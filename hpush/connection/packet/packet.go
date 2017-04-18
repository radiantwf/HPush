package packet

type Packet struct {
	PacketHeader
	PacketData
}

type PacketHeader struct {
}

type PacketData struct {
	data []byte
}
