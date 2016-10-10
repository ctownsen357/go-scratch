package main

import (
	"os"
	"testing"
)

func TestWriteReadBinary(t *testing.T) {
	fileName := "test.bin"
	_ = os.Remove(fileName)

	bd := BinData{EventID: 4, WindGridID: 1000, WindSpeedX10: 102}
	err := WriteBinary(bd, fileName)
	if err != nil {
		t.Error(err)
	}

	bdResult, err := ReadBinary(fileName)
	if err != nil {
		t.Error(err)
	}

	if bd.EventID != bdResult.EventID || bd.WindGridID != bdResult.WindGridID || bd.WindSpeedX10 != bdResult.WindSpeedX10 {
		t.Error("binary data written doesn't match read binary data: ", bd, bdResult)
	}
}
