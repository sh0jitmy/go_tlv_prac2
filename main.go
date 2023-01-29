package main

import (
   "bytes"
   "encoding/binary"
   "encoding/hex"
   "fmt"
)

type TLV struct {
	ID uint16
	Length uint16
	//Value []byte 
}

func has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}

func main() {
	Value1:="0x000000000001"
	if !has0xPrefix(Value1) {
		return
	}
	binval , _ :=  hex.DecodeString(Value1[2:])
	fmt.Println(binval)	
	dLen := uint16(len(binval))

	tlv := &TLV { ID: 0x0001, Length: dLen} 
	buf := &bytes.Buffer{}
	
	fmt.Printf("tlv: %v\n",tlv)	
	err := binary.Write(buf,binary.BigEndian,tlv)
	if err != nil {
		fmt.Println(err)
	}
	err = binary.Write(buf,binary.BigEndian,binval)
	fmt.Printf("encoding binary length: %d\n",buf.Len())	
} 
