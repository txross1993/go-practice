package main

import (
	"fmt"
	"github.com/azoner/gox12"
	log "github.com/txross1993/go-practice/EdiParser/logwrapper"
	"os"
	"path/filepath"
)

func main() {
	li := log.NewLogger()
	baseDir := "C:/Users/rossth/go/src/github.com/txross1993/go-practice/EdiParser/testFiles"

	files := []string{"sample945.txt", "tys_sample_945.txt"}

	for _, f := range files {
		inFilename, err := filepath.Abs(baseDir + "/" + f)
		inFile, err := os.Open(inFilename)
		if err != nil {
			li.Fatal(err)
			os.Exit(1)
		}
		li.Infof("Reading file %v", f)
		defer inFile.Close()
		raw, err := gox12.NewRawX12FileReader(inFile)
		if err != nil {
			fmt.Println(err)
		}
		for rs := range raw.GetSegments() {
			for v := range rs.Segment.GetAllValues() {
				fmt.Printf("Path: %v, Value: %v\n", v.X12Path, v.Value)
			}
			// if rs.Segment.SegmentId == "INS" {
			// 	fmt.Println(rs)
			// 	v, _, _ := rs.Segment.GetValue("INS01")
			// 	fmt.Println(v)
			// 	for v := range rs.Segment.GetAllValues() {
			// 		fmt.Println(v.X12Path, v.Value)
			// 	}
			// 	fmt.Println()
			// }
		}
		li.Debugf("Got segments from file %v", f)
	}
}