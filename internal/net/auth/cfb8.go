package auth

import (
	"crypto/cipher"
	"errors"
)

type cfb8 struct {
	block         cipher.Block
	sr            []byte
	srEnc         []byte
	srPos         int
	shouldDecrypt bool
}

func NewCFB8Encrypter(block cipher.Block, iv []byte) (cipher.Stream, error) {
	if len(iv) != block.BlockSize() {
		return nil, errors.New("cfb8::NewCFB8Encrypter: IV length must equal block size")
	}
	return newCFB8(block, iv, false), nil
}

func NewCFB8Decrypter(block cipher.Block, iv []byte) (cipher.Stream, error) {
	if len(iv) != block.BlockSize() {
		return nil, errors.New("cfb8::NewCFB8Decrypter: IV length must equal block size")
	}
	return newCFB8(block, iv, true), nil
}

func newCFB8(block cipher.Block, iv []byte, decrypt bool) cipher.Stream {
	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		return nil
	}

	x := &cfb8{
		block:         block,
		sr:            make([]byte, blockSize*4),
		srEnc:         make([]byte, blockSize),
		srPos:         0,
		shouldDecrypt: decrypt,
	}
	copy(x.sr, iv)
	return x
}

func (x *cfb8) XORKeyStream(dst, src []byte) {
	blockSize := x.block.BlockSize()

	for i := 0; i < len(src); i++ {
		x.block.Encrypt(x.srEnc, x.sr[x.srPos:x.srPos+blockSize])

		var c byte
		if x.shouldDecrypt {
			c = src[i]
			dst[i] = c ^ x.srEnc[0]
		} else {
			c = src[i] ^ x.srEnc[0]
			dst[i] = c
		}

		x.sr[x.srPos+blockSize] = c
		x.srPos++

		if x.srPos+blockSize == len(x.sr) {
			copy(x.sr, x.sr[x.srPos:])
			x.srPos = 0
		}
	}
}
