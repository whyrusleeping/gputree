package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func hashXOR(l, r, out []byte) {
	for i := range l {
		out[i] = l[i] ^ r[i]
	}
}

func main() {
	/*
		fi, err := os.Create("testfile.bin")
		if err != nil {
			panic(err)
		}

		r := rand.New(rand.NewSource(1231235))
		io.CopyN(fi, r, 2<<30)
		fi.Close()
		return
	*/

	fi, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	basebuf, err := ioutil.ReadAll(fi)
	if err != nil {
		panic(err)
	}

	for len(basebuf) > 32 {
		next := make([]byte, len(basebuf)/2)
		for i := 0; (i+1)*32 < len(basebuf); i += 2 {
			l := basebuf[i*32 : (i+1)*32]
			r := basebuf[(i+1)*32 : (i+2)*32]
			hashXOR(l, r, next[(i/2)*32:])
		}
		basebuf = next
	}

	fmt.Printf("result: %x\n", basebuf)
}
