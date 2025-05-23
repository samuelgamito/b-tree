package disk

import (
	"log"
	"os"
)

const (
	DatabaseFileName = "database.bin"
	IndexFileName    = "index.bin"
)

func GetDatabaseFile() (*os.File, error) {

	var (
		f   *os.File
		err error
	)

	f, err = os.OpenFile(DatabaseFileName, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal("Error on creating database file")
		return nil, err
	}

	return f, nil
}

func GetIndexFile() (*os.File, error) {

	var (
		f   *os.File
		err error
	)

	f, err = os.OpenFile(IndexFileName, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal("Error on creating database file")
		return nil, err
	}

	return f, nil
}

func CloseAll(files ...*os.File) {

	for _, f := range files {

		err := f.Close()
		if err != nil {
			log.Fatalf("Error on closing %s file", f.Name())
		}
	}

}
