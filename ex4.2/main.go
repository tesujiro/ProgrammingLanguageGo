package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var bits int

func main() {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.IntVar(&bits, "b", 256, "hash bits (256,384,512)")
	f.Parse(os.Args[1:])

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var fnc func([]byte) []byte
	switch bits {
	case 256:
		fnc = func(b []byte) []byte {
			hash := sha256.Sum256(b)
			return hash[:]
		}
	case 384:
		fnc = func(b []byte) []byte {
			hash := sha512.Sum384(b)
			return hash[:]
		}
	case 512:
		fnc = func(b []byte) []byte {
			hash := sha512.Sum512(b)
			return hash[:]
		}
	default:
		log.Fatal("hash bits error")
	}

	fmt.Printf("%x\n", fnc(b))
}
