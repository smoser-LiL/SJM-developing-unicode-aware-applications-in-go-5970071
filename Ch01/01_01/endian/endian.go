package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("max byte:", math.MaxUint8) // 255

	var c uint16 = 256 // Ā
	fmt.Printf("c: %d (0x%x)\n", c, c)

	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.BigEndian, c); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("big endian   :", bytesString(buf.Bytes()))

	buf.Reset()
	if err := binary.Write(&buf, binary.LittleEndian, c); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("little endian:", bytesString(buf.Bytes()))

}

func bytesString(bs []byte) string {
	nums := make([]string, len(bs))
	for i, b := range bs {
		nums[i] = fmt.Sprintf("%02x", b)
	}

	return strings.Join(nums, " ")
}
