package main

import (
	"log"
	"os"
)

const (
	DatabaseFileName = "database.bin"
)


func getDatabaseFile() (*os.File, error)  {

	var (
		f *os.File
		err error
	)
	
	f, err = os.OpenFile(DatabaseFileName, os.O_RDWR|os.O_CREATE, 0644)
	

	if err != nil {
		log.Fatal("Error on creating database file")
		return nil, err
	}


	return f, nil
}

func closeAll(f *os.File){
	f.Close()
}