package compress

import (
	"bytes"
	"compress/gzip"
	"io"
)

const (
	GZIP = byte(iota)
	BZIP2
)

const DEFAULT_COMPRESS_TYPE = GZIP

type Compress struct{}

func (c *Compress) Compress(srcdata []byte, compresstype byte) (desdata []byte, err error) {
	var buf bytes.Buffer
	switch compresstype {
	case BZIP2:
	case GZIP:
		fallthrough
	default:
		zw := gzip.NewWriter(&buf)
		_, err = zw.Write(srcdata)
		zw.Close()
		if err != nil {
			return
		}
	}
	desdata = buf.Bytes()
	return
}

func (c *Compress) Decompress(srcdata []byte, compresstype byte) (desdata []byte, err error) {
	var buf bytes.Buffer
	switch compresstype {
	case BZIP2:
	case GZIP:
		fallthrough
	default:
		zr, err1 := gzip.NewReader(bytes.NewBuffer(srcdata))
		if err1 != nil {
			err = err1
			return
		}
		_, err = io.Copy(&buf, zr)
		zr.Close()
		if err != nil {
			return
		}
	}
	desdata = buf.Bytes()
	return
}
