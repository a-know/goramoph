package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func main() {
	var fp *os.File
	var err error

	// パース対象の xml ファイル名を引数に受ける
	fp, err = os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 4096)

	d := xml.NewDecoder(reader)
	for {
		token, err := d.Token()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			panic(err)
		}
		switch token.(type) {
		case xml.StartElement:
			//do something
			fmt.Println("StartElement")
		case xml.EndElement:
			//do something
			fmt.Println("EndElement")
		case xml.CharData:
			//do something
			fmt.Println("CharData")
		case xml.Comment:
			//do something
			fmt.Println("Comment")
		case xml.ProcInst:
			//do something
			fmt.Println("ProcInst")
		case xml.Directive:
			//do something
			fmt.Println("Directive")
		default:
			panic("unknown xml token.")
		}
	}
}

// type Dict struct {
// 	Thumb Thumb `xml:"thumb"`
// }

// type Thumb struct {
// 	Title  string `xml:"title"`
// 	Length string `xml:"length"`
// }
