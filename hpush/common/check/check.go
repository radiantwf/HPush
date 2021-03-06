package check

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/binary"
	"hash/adler32"
	"hash/crc32"
)

const (
	CRC32 = byte(iota)
	ADLER32
	MD5
	SHA1
)
const DEFAULT_CHECK_TYPE = CRC32

type Check struct{}

func (c *Check) GetCheckingCode(data []byte, checktype byte) (code uint32, err error) {
	switch checktype {
	case ADLER32:
		code = adler32.Checksum(data)
	case MD5:
		value1 := md5.Sum(data)
		code = binary.LittleEndian.Uint32(value1[:])
	case SHA1:
		value2 := sha1.Sum(data)
		code = binary.LittleEndian.Uint32(value2[:])
	case CRC32:
		fallthrough
	default:
		code = crc32.ChecksumIEEE(data)
	}
	return
}
