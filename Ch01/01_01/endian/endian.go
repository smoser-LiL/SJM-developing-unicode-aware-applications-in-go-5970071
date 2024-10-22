package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func main() {
	fmt.Println("max byte:", math.MaxUint8) // 255

	var c uint16 = 1023
	fmt.Printf("c: %d (0x%04x)\n", c, c)

	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.BigEndian, c); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("big endian   : %x\n", buf.Bytes())

	buf.Reset()
	if err := binary.Write(&buf, binary.LittleEndian, c); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("little endian: %x\n", buf.Bytes())
}
