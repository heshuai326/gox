package gox

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/gopub/log"
)

// AES crypt data with key, iv
// key is either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
// iv is recommended 16 bytes.
func AES(data []byte, key []byte, iv []byte) error {
	n := len(key)
	if len(data) == 0 {
		return errors.New("no data")
	}

	if n != 16 && n != 24 && n != 32 {
		return errors.New("invalid key")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(data, data)
	return nil
}

// XOR xor data with key
func XOR(data []byte, key []byte) {
	n := len(key)
	if len(data) == 0 || n == 0 {
		log.Error("SimpleXORCrypt invalid or and key")
		return
	}

	for i, b := range data {
		data[i] = b ^ key[i%n]
	}
}

// MD5 returns str's md5 value which is 128 bits represented as 32 hex string
func MD5(str string) string {
	sum := md5.Sum([]byte(str))
	return hex.EncodeToString(sum[:])
}

// SHA1 returns str's sha1 value which is 160 bits represented as 40 hex string
func SHA1(str string) string {
	sha1er := sha1.New()
	b := []byte(str)
	for len(b) > 0 {
		n, err := sha1er.Write(b)
		if err != nil {
			log.Error(err)
		}
		b = b[n:]
	}
	return hex.EncodeToString(sha1er.Sum(nil))
}

// SHA256 returns str's sha256 value which is 256 bits 64 hex string
func SHA256(str string) string {
	sha256er := sha256.New()
	b := []byte(str)
	for len(b) > 0 {
		n, err := sha256er.Write(b)
		if err != nil {
			log.Error(err)
			return ""
		}
		b = b[n:]
	}
	return hex.EncodeToString(sha256er.Sum(nil))
}
