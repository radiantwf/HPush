package data

import (
	"bytes"
	"compress/gzip"
	"io"
)

const (
	Gzip = byte(iota)
	Bzip2
)

func Compress(srcdata []byte, compresstype byte) (desdata []byte, err error) {
	var buf bytes.Buffer
	switch compresstype {
	case Bzip2:
	case Gzip:
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

func Decompress(srcdata []byte, compresstype byte) (desdata []byte, err error) {
	var buf bytes.Buffer
	switch compresstype {
	case Bzip2:
	case Gzip:
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
