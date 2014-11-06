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
			fmt.Println(token.(xml.StartElement))
		case xml.EndElement:
			//do something
			fmt.Println("EndElement")
			fmt.Println(token.(xml.EndElement))
		case xml.CharData:
			//do something
			fmt.Println("CharData")
			fmt.Println(string(token.(xml.CharData)))
		case xml.Comment:
			//do something
			fmt.Println("Comment")
			fmt.Println(string(token.(xml.Comment)))
		case xml.ProcInst:
			//do something
			fmt.Println("ProcInst")
			fmt.Println(token.(xml.ProcInst))
		case xml.Directive:
			//do something
			fmt.Println("Directive")
			fmt.Println(string(token.(xml.Directive)))
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
