package rx

import "io/ioutil"

func File_read_file(path string) []byte {
	dst, err := ioutil.ReadFile(path)
	if nil != err {
		panic("read")
	}

	return dst
}

func File_write_file(path string, src []byte) {
	err := ioutil.WriteFile(path, src, 0644)
	if nil != err {
		panic("write")
	}
}
