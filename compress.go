package gox

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"fmt"
	"io/ioutil"
)

type Compressor interface {
	Compress(data []byte) ([]byte, error)
}

func GetCompressor(name string) Compressor {
	switch name {
	case "gzip":
		return gzipCompressionInstance
	case "flate":
		return flateCompressionInstance
	default:
		return nil
	}
}

type Decompressor interface {
	Decompress(data []byte) ([]byte, error)
}

func GetDecompressor(name string) Decompressor {
	switch name {
	case "gzip":
		return gzipCompressionInstance
	case "deflate":
		return flateCompressionInstance
	default:
		return nil
	}
}

var flateCompressionInstance = &flateCompression{}

type flateCompression struct {
}

func (g *flateCompression) Compress(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, nil
	}

	buffer := new(bytes.Buffer)
	writer, err := flate.NewWriter(buffer, flate.DefaultCompression)
	if err != nil {
		return nil, err
	}
	//Make sure writer is closed before calling buffer.Bytes()!!!
	err = WriteAll(writer, data)
	if err != nil {
		writer.Close()
		return nil, err
	}
	writer.Close()
	return buffer.Bytes(), nil
}

func (g *flateCompression) Decompress(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, nil
	}

	buffer := bytes.NewBuffer(data)
	reader := flate.NewReader(buffer)
	defer reader.Close()
	return ioutil.ReadAll(reader)
}

var gzipCompressionInstance = &gzipCompression{}

type gzipCompression struct {
}

func (g *gzipCompression) Compress(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, nil
	}

	buffer := new(bytes.Buffer)
	writer := gzip.NewWriter(buffer)
	//Make sure writer is closed before calling buffer.Bytes()!!!
	err := WriteAll(writer, data)
	if err != nil {
		return nil, fmt.Errorf("write failed: %w", err)
	}
	if err = writer.Close(); err != nil {
		return nil, fmt.Errorf("close failed: %w", err)
	}
	return buffer.Bytes(), nil
}

func (g *gzipCompression) Decompress(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, nil
	}

	buffer := bytes.NewBuffer(data)
	reader, err := gzip.NewReader(buffer)
	if err != nil {
		return nil, err
	}

	defer reader.Close()
	return ioutil.ReadAll(reader)
}
