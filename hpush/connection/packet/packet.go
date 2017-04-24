package packet

import (
	"HPush/hpush/common/check"
	"bytes"
	"encoding/binary"
	"errors"
)

type Packet struct {
	PacketHeader
	PacketData
}

type PacketHeader struct {
	BodyLength uint32 //表示body的长度
	Cmd        byte   //表示消息协议类型
	CheckCode  uint32 //是根据body生成的一个校验码
	Flags      byte   //表示当前包启用的特性，比如是否启用加密，是否启用压缩
	SessionID  uint32 //消息会话标识用于消息响应
	Lrc        byte   //纵向冗余校验，用于校验header
}

type PacketData struct {
	Data []byte
}

func (p *Packet) SetDataCheckCode() (err error) {
	code, err := (&check.Check{}).GetCheckingCode(p.Data, check.ADLER32)
	if err != nil {
		return
	}
	p.CheckCode = code
	return
}
func (p *Packet) SetLrcCode() {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, p.BodyLength)
	binary.Write(buf, binary.LittleEndian, p.Cmd)
	binary.Write(buf, binary.LittleEndian, p.CheckCode)
	binary.Write(buf, binary.LittleEndian, p.Flags)
	binary.Write(buf, binary.LittleEndian, p.SessionID)
	lrc := byte(0x0)
	for _, b := range buf.Bytes() {
		lrc = lrc ^ b
	}
	p.Lrc = lrc
}

func (p *Packet) check(data []byte) (err error) {
	// data长度检查
	var header PacketHeader
	headersize := binary.Size(header)
	datasize := uint32(len(data) - headersize)
	header.BodyLength = binary.LittleEndian.Uint32(data[:4])
	if datasize != header.BodyLength {
		err = errors.New("数据长度不正确")
		return
	}
	// header垂直校验
	lrc := byte(0x0)
	for i := 0; i < headersize; i++ {
		lrc = lrc ^ data[0]
	}
	if lrc != byte(0x0) {
		err = errors.New("header垂直校验不正确")
		return
	}
	// data校验
	code, err := (&check.Check{}).GetCheckingCode(data[headersize:], check.ADLER32)
	if err != nil {
		return
	}
	checkcode := binary.LittleEndian.Uint32(data[5:9])
	if checkcode != code {
		err = errors.New("数据Hash校验不正确")
		return
	}
	return
}
func NewPacket(buffer []byte) (p *Packet, err error) {
	p = new(Packet)
	err = p.check(buffer)
	if err != nil {
		return
	}
	p.BodyLength = binary.LittleEndian.Uint32(buffer[:4])
	p.Cmd = buffer[4]
	p.CheckCode = binary.LittleEndian.Uint32(buffer[5:9])
	p.Flags = buffer[9]
	p.SessionID = binary.LittleEndian.Uint32(buffer[10:14])
	p.Flags = buffer[14]
	p.Data = buffer[15:]
	return
}
