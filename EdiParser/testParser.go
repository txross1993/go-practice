package main

import (
	"fmt"
	p "github.com/txross1993/go-practice/EdiParser/ediParser"
	log "github.com/txross1993/go-practice/EdiParser/logwrapper"
	"os"
	"path/filepath"
	"time"
)

func main() {
	li := log.NewLogger()

	folderName := "testFiles"
	folder, err := filepath.Abs(folderName)

	if err != nil {
		li.Fatalf("Could not read folder %v", folderName)
	}

	err = filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			li.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if info.Mode().IsRegular() {
			// Test - Open, read, and scan the entire file for tokenized content 100 times
			for i := 1; i < 101; i++ {
				t0 := time.Now()
				fullPath, err := filepath.Abs(path)

				if err != nil {
					return err
				}

				f, err := os.Open(fullPath)

				if err != nil {
					return err
				}

				parser := p.NewParser(f)

				_, err = parser.Parse()
				if err != nil {
					return err
				}

				f.Close()
				//fmt.Printf("%v\n", output)
				t1 := time.Now()
				duration := t1.Sub(t0)
				fmt.Printf("Iteration %v duration: %v\n", i, duration)
			}
		}
		return nil
	})

	if err != nil {
		li.Fatalf("error walking the path %q: %v\n", folder, err)
		return
	}
}
