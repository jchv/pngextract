package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-restruct/restruct"
)

type pngchunk struct {
	Len  uint32 `struct:"sizeof=Data"`
	Type string `struct:"[4]byte"`
	Data []byte
	CRC  uint32
}

var pngsig = []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}

func extractpng(filename string, b []byte) error {
	// Start at first chunk.
	i := uint32(8)

	for {
		c := pngchunk{}
		cb := b[i:]
		if err := restruct.Unpack(cb, binary.BigEndian, &c); err != nil {
			return fmt.Errorf("error decoding embedded PNG structure: %w", err)
		}
		i += c.Len + 12
		if c.Type == "IEND" {
			break
		}
	}
	if err := ioutil.WriteFile(filename, b[:i], 0644); err != nil {
		return fmt.Errorf("error writing file %q: %v", filename, err)
	}
	return nil
}

func extract(filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file %q: %w", filename, err)
	}
	s := b
	for x, d := bytes.Index(s, pngsig), 0; x > -1; x, d = bytes.Index(s, pngsig), d+x+1 {
		if err := extractpng(fmt.Sprintf("%s-%08x.png", filename, x+d), b[x+d:]); err != nil {
			log.Printf("error extracting PNG at %08x: %v", x+d, err)
		}
		s = s[x+1:]
	}
	return nil
}

func main() {
	for _, filename := range os.Args[1:] {
		if err := extract(filename); err != nil {
			log.Fatalf("Could not extract from file %q: %v", filename, err)
		}
	}
}
