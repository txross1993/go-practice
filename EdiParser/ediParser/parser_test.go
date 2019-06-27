package ediParser

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

var testParser = map[string]struct {
	in       string
	outFile  EdiFile
	outError error
}{
	"PositiveExample": {
		in: "ISA*00* *00* *08*9338945000 *08*9393939394*021104*1405*U*00501*000000031*0*T*>~",
		outFile: EdiFile{&EdiStatement{
			Keyword: "ISA",
			Fields:  []string{"00", " ", "00", "08", "9338945000", "08", "9393939394", "021104", "1405", "U", "00501", "000000031", "0", "T", ">"},
		}},
		outError: nil,
	},
	"EOF": {
		in:       string(eof),
		outFile:  nil,
		outError: fmt.Errorf("Found %v, expected keyword", ""),
	},
	"NewLineMid": {
		in:       "ISA*00* *00* *08*9338945000 *08*9393939394*021104*1405*U*00501*000000031*0*T*>~",
		outFile:  nil,
		outError: &KeywordError{"\n"},
	},
	"NoKeyword": {
		in:       "*00* *00* *08*9338945000 *08*9393939394\n*021104*1405*U*00501*000000031*0*T*>~",
		outFile:  nil,
		outError: fmt.Errorf("Found %v, expected keword", ""),
	},
	"NoTilde": {
		in:       "ISA*00* *00* *08*9338945000 *08*9393939394*021104*1405*U*00501*000000031*0*T*>",
		outFile:  EdiFile{},
		outError: nil,
	},
}

func compareExpectedActualError(expectedErr error, actualError error) bool {
	return reflect.TypeOf(expectedErr) == reflect.TypeOf(actualError)
}

func compareSlices(expectedSlice EdiFile, actualSlice EdiFile) bool {
	return reflect.DeepEqual(expectedSlice, actualSlice)
}

func TestParser(t *testing.T) {
	for name, test := range testParser {
		t.Logf("Running test case: %s", name)

		r := strings.NewReader(test.in)

		p := NewParser(r)

		file, err := p.Parse()

		if compareSlices(test.outFile, file) != true && compareExpectedActualError(test.outError, err) != true {
			t.Errorf("Test case failure: %v \n\t Parse(%v) => GOT (%v, %v), WANT (%v, %v)", name, test.in, file, err, test.outFile, test.outError)
		}
	}
}
