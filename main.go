package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	r "reflect"
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

	var m map[string]string = map[string]string{}
	var shouldSet bool
	var key string
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
			if token.(xml.StartElement).Name.Local == "key" {
				shouldSet = false
			} else {
				shouldSet = true
			}
		case xml.EndElement:
			//do nothing
		case xml.CharData:
			if string(token.(xml.CharData)) == "Track ID" {
				// CharData が Track ID が来る＝今まで処理対象だったレコードの終了処理をする ＆ 次のレコードに処理を移す
				s1 := playData{}
				MapToStruct(m, &s1)
				fmt.Println(s1)
				m = map[string]string{}
			}
			if shouldSet && key != "" {
				m[key] = string(token.(xml.CharData))
				key = ""
			} else {
				key = string(token.(xml.CharData))
			}
		case xml.Comment:
			//do nothing
		case xml.ProcInst:
			//do nothing
		case xml.Directive:
			//do nothing
		default:
			panic("unknown xml token.")
		}
	}
}

func MapToStruct(mapVal map[string]string, val interface{}) (ok bool) {
	structVal := r.Indirect(r.ValueOf(val))
	for name, elem := range mapVal {
		f := structVal.FieldByName(name)
		if f.IsValid() {
			f.Set(r.ValueOf(elem))
		}
	}
	return
}

type playData struct {
	Artist               string
	Album                string
	Genre                string
	Kind                 string
	Size                 string
	TotalTime            string
	TrackNumber          string
	DateModified         string
	DateAdded            string
	BitRate              string
	SampleRate           string
	PersistentID         string
	TrackType            string
	Location             string
	FileFolderCount      string
	LibraryFolderCount   string
	PlayCount            string
	PlayDate             string
	PlayDateUTC          string
	Rating               string
	AlbumRating          string
	AlbumRatingComputed  string
	ArtworkCount         string
	SkipCount            string
	SkipDate             string
	Disabled             string
	SortAlbum            string
	Year                 string
	Comments             string
	SortName             string
	SortArtist           string
	VolumeAdjustment     string
	AlbumArtist          string
	DiscNumber           string
	DiscCount            string
	TrackCount           string
	ReleaseDate          string
	Protected            string
	Purchased            string
	Compilation          string
	Composer             string
	HasVideo             string
	VideoWidth           string
	VideoHeight          string
	Movie                string
	Master               string
	PlaylistID           string
	PlaylistPersistentID string
	Visible              string
	AllItems             string
	PlaylistItems        string
	SmartCriteria        string
}
