package rx

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"io/ioutil"
	"os"
)

func Util_rand_bytes(size int) []byte {
	if 0 >= size {
		panic("size 0")
	}

	buf := make([]byte, size)

	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		panic("ReadFull")
	}

	return buf
}

func Util_to_cstr(src []byte) string {
	if 0 >= len(src) {
		return ""
	}

	zpos := 0
	for ; zpos < len(src); zpos++ {
		if 0 == src[zpos] {
			break
		}
	}

	return string(src[:zpos])
}

func Util_to_cbytes(src string) []byte {
	bsrc := []byte(src)
	return append(bsrc, []byte{0}...)
}

func Util_encode_to_base64_string(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func Util_decode_from_base64_string(src string) []byte {
	dst, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		panic("decode")
	}

	return dst
}

func Util_clear_dir(path string) {
	dir, _ := ioutil.ReadDir(path)
	for _, d := range dir {
		os.RemoveAll(path + "/" + d.Name())
	}
}
