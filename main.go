package main

import "golib/rx"

func main() {
	path := "sample/rand.3"

	rands := rx.Util_rand_bytes(8321573)
	rx.File_write_file(path, rands)
}
