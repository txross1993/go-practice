# Purpose

Use EdiParser to scan for special tokens and build structs related to each keyword token. The parser will scan EDI files character by character. When '\n' or '\r' character is encountered, it signals the end of the EdiStatement, which is then appended to an array of EdiStatements called an EdiFile.

## EDI File
`type EdiFile []*EdiStatement`

EdiFile is a slice of (array) of EdiStatements.

## EDI Statement

```
type EdiStatement struct {
	Keyword string
	Fields  []string
}
```

EDI Statement is a struct with keywords and a slice (array) of string fields, the purpose of each may be determined later. Each statement begins with a keyword, defined in ediParser/tokens. Following a keyword is any number of fields separated by the separator token, '*'. When a new line or EOF character is encountered, the statement is complete. When the statement completes, the scanner looks for EOF. If EOF is encountered the file is comlete.