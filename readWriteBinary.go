package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

//I was needing to pack 3 int in the packed binary formt of iih for a Python struct
//this ends up packing into the appropriate packed bytes of 4,4,2
type binData struct {
	EventID      int32
	WindGridID   int32
	WindSpeedX10 int16
}

func writeBinary() {
	f, err := os.Create("tst.bin")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	//originally I used a byte array and packed each variable in but the structure works just fine
	bd := binData{EventID: 4, WindGridID: 1000, WindSpeedX10: 102}

	binary.Write(f, binary.LittleEndian, &bd)
}

func readBinary() {
	f, err := os.Open("tst.bin")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for {
		bd := binData{}
		err = binary.Read(f, binary.LittleEndian, &bd)

		if err == io.EOF {
			break
		}
		fmt.Println(bd.EventID, bd.WindGridID, bd.WindSpeedX10)
	}
}
func main() {
	writeBinary()
	readBinary()
}
