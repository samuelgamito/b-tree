package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type DBActions struct{
	
}

func (DBActions) Add(id string, data string, f *os.File) error {

	_, err := f.Seek(0, io.SeekEnd)

	dataBytes := []byte(data)
	stringSize := len(dataBytes)
	nextData := -1

	if err != nil {
		panic("error to pointing to the end")
	}

	cursorPosition, err := f.Seek(0, io.SeekCurrent)

	if err != nil {
		panic("not able to get the cursor")
	}

	fmt.Println(cursorPosition)

	if err := binary.Write(f, binary.LittleEndian, int32(stringSize)); err != nil {
		log.Fatalf("Failed writing string size to file: %s", err)
	}
	if err := binary.Write(f, binary.LittleEndian, int32(0)); err != nil {
		log.Fatalf("Failed writing string size to file: %s", err)
	} 
	if err = binary.Write(f,binary.LittleEndian, dataBytes); err != nil {
		log.Fatalf("Failed writing string to file: %s", err)
	}
	if err := binary.Write(f, binary.LittleEndian, int32(nextData)); err != nil {
		log.Fatalf("Failed writing string size to file: %s", err)
	}

	return nil
}

func (DBActions) Get(pos string, f *os.File) error {
	var stringSize int32
	var deletedReg int32

	posInt,_ := strconv.Atoi(strings.ReplaceAll(pos, " ",""))

	fmt.Println(posInt)
	if _, err := f.Seek(int64(posInt), io.SeekStart); err != nil {
		log.Fatal(err)
	}


	if err := binary.Read(f, binary.LittleEndian, &stringSize); err != nil {
		log.Fatal(err)
	}
	if err := binary.Read(f, binary.LittleEndian, &deletedReg); err != nil {
		log.Fatal(err)
	}

	stringBytes := make([]byte, stringSize)
	if _, err := f.Read(stringBytes); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s, is deleted? %d\n",string(stringBytes), deletedReg)

	return nil
}