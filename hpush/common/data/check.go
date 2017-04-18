package data

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/binary"
	"hash/adler32"
	"hash/crc32"
)

const (
	Crc32 = byte(iota)
	Adler32
	Md5
	Sha1
)

func GetCheckingCode(data []byte, checktype byte) (code uint32, err error) {
	switch checktype {
	case Adler32:
		code = adler32.Checksum(data)
	case Md5:
		value1 := md5.Sum(data)
		code = binary.BigEndian.Uint32(value1[:])
	case Sha1:
		value2 := sha1.Sum(data)
		code = binary.BigEndian.Uint32(value2[:])
	case Crc32:
		fallthrough
	default:
		code = crc32.ChecksumIEEE(data)
	}
	return
}
