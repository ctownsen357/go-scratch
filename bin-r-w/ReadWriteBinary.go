package main

import (
	"encoding/binary"
	"io"
	"os"
)

//I was needing to pack 3 int in the packed binary formt of iih for a Python struct
//this ends up packing into the appropriate packed bytes of 4,4,2
type BinData struct {
	EventID      int32
	WindGridID   int32
	WindSpeedX10 int16
}

func WriteBinary(bd BinData, fileName string) (err error) {
	f, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer f.Close()

	err = binary.Write(f, binary.LittleEndian, &bd)

	return err
}

func ReadBinary(fileName string) (bd BinData, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return BinData{}, err
	}
	defer f.Close()

	for {
		err = binary.Read(f, binary.LittleEndian, &bd)

		if err == io.EOF {
			err = nil
			break
		}
	}
	return bd, err
}

func main() {
}
