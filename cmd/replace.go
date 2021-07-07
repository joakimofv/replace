package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/joakimofv/find"
	"github.com/spf13/viper"
)

func replace(args []string) (e error) {
	oldPattern, newPattern := tweakPattern(args[0]), tweakPattern(args[1])
	if args[1] == "REMOVE" {
		// Don't replace line, remove completely
		newPattern = "REMOVE"
	}
	newPattern = strings.Replace(newPattern, "\\n", "\n", -1)
	candidateFiles := find.Files(find.LongestFixedPart(oldPattern), args[2:])
	for _, file := range candidateFiles {
		if err := findAndReplace(file, oldPattern, newPattern); err != nil {
			if e != nil {
				// Combine the error strings.
				e = fmt.Errorf("%v\n%v", e, err)
			} else {
				e = err
			}
		}
	}
	return
}

func tweakPattern(pat string) string {
	if viper.GetBool("exact") {
		return pat
	}
	return "*" + strings.Trim(pat, "*") + "*"
}

func findAndReplace(filename string, oldPattern, newPattern string) error {
	writeToFile := false
	var newContents bytes.Buffer
	defer func() {
		if writeToFile {
			if err := ioutil.WriteFile(filename, newContents.Bytes(), 0777); err != nil {
				// Can the error be returned?
				fmt.Printf("Error writing to file %v: %v\n", filename, err)
			}
		}
	}()
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close() // The Close should execute before the WriteFile.

	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		line := scanner.Text()
		modLine, madeChange := find.Replace(line, oldPattern, newPattern)
		if madeChange {
			writeToFile = true
			line = modLine
			if newPattern == "REMOVE" {
				continue
			}
		}
		newContents.WriteString(line + "\n")
	}

	return nil
}
