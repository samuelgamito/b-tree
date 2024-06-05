package main

import (
	"b-tree/database"
	"b-tree/disk"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"

	"golang.org/x/term"
)

func main() {
	fDatabaseFile, _ := disk.GetDatabaseFile()
	dIndexDatabaseFile, _ := disk.GetIndexFile()

	defer disk.CloseAll(fDatabaseFile, dIndexDatabaseFile)

	if err := chat(fDatabaseFile); err != nil {
		log.Fatal(err)
	}
}

func getCommand(line string) string {
	return strings.ReplaceAll(strings.Split(line, ",")[0], " ", "")
}

func getCommandParams(line string, f *os.File) []reflect.Value {

	var paramsReflected []reflect.Value

	params := strings.Split(line, ",")[1:]

	for _, p := range params {
		paramsReflected = append(paramsReflected, reflect.ValueOf(p))
	}

	paramsReflected = append(paramsReflected, reflect.ValueOf(f))

	return paramsReflected
}

func chat(fDatabase *os.File) error {
	dbActions := database.DBActions{}

	if !term.IsTerminal(0) || !term.IsTerminal(1) {
		return fmt.Errorf("stdin/stdout should be terminal")
	}
	oldState, err := term.MakeRaw(0)
	if err != nil {
		return err
	}
	defer term.Restore(0, oldState)
	screen := struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}
	term := term.NewTerminal(screen, "")
	term.SetPrompt(string(term.Escape.Red) + "> " + string(term.Escape.Reset))

	databaseResponsePrefix := string(term.Escape.Cyan) + "Database Response:" + string(term.Escape.Reset)

	for {
		line, err := term.ReadLine()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if line == "" {
			continue
		}

		cmd := getCommand(line)

		action := reflect.ValueOf(dbActions).MethodByName(cmd)

		if action.IsValid() {

			action.Call(getCommandParams(line, fDatabase))
		} else {
			fmt.Fprintln(term, databaseResponsePrefix, cmd, " is not a valid command")
		}

	}
}
