package main

import (
	"bytes"
	"fmt"
	"github.com/DHowett/go-plist"
	// "io/ioutil"
	// "os"
	// "time"
)

func main() {
	var err error
	// var contents byte[]

	// パース対象の xml ファイル名を引数に受ける
	// contents, err := ioutil.ReadFile(os.Args[1])
	// if err != nil {
	// 	panic(err)
	// }
	// buf := bytes.NewReader(contents)
	buf := bytes.NewReader([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Major Version</key><integer>1</integer>
	<key>Minor Version</key><integer>1</integer>
	<key>Application Version</key><string>7.6.1</string>
	<key>Features</key><integer>5</integer>
	<key>Show Content Ratings</key><true/>
	<key>Music Folder</key><string>file://localhost/C:/mydocument/My%20Music/iTunes/iTunes%20Music/</string>
	<key>Library Persistent ID</key><string>2A0848311638E5FA</string>
	<key>Tracks</key>
	<dict>
		<key>T894</key>
		<dict>
			<key>Track ID</key><integer>894</integer>
			<key>Name</key><string>Dolce Vita</string>
			<key>Artist</key><string>情熱大陸</string>
			<key>Album</key><string>情熱大陸</string>
			<key>Genre</key><string>Classic Rock</string>
			<key>Kind</key><string>MPEG オーディオファイル</string>
			<key>Size</key><integer>9384980</integer>
			<key>Total Time</key><integer>391026</integer>
			<key>Track Number</key><integer>2</integer>
			<key>Date Modified</key><date>2005-09-17T12:49:59Z</date>
			<key>Date Added</key><date>2005-09-16T11:14:19Z</date>
			<key>Bit Rate</key><integer>192</integer>
			<key>Sample Rate</key><integer>44100</integer>
			<key>Persistent ID</key><string>2A0848311638E608</string>
			<key>Track Type</key><string>File</string>
			<key>Location</key><string>file://localhost/C:/0_mp3_touch/(MP3)%5B%E3%82%A2%E3%83%AB%E3%83%90%E3%83%A0%5D%E3%82%AA%E3%83%A0%E3%83%8B%E3%83%90%E3%82%B9%20-%20%E6%83%85%E7%86%B1%E5%A4%A7%E9%99%B8-%E8%91%89%E5%8A%A0%E7%80%AC%E5%A4%AA%E9%83%8E%20SELECTION/Dolce%20Vita.mp3</string>
			<key>File Folder Count</key><integer>-1</integer>
			<key>Library Folder Count</key><integer>-1</integer>
		</dict>
		<key>895</key>
		<dict>
			<key>Track ID</key><integer>895</integer>
			<key>Name</key><string>Dream of Wings</string>
			<key>Artist</key><string>情熱大陸</string>
			<key>Album</key><string>情熱大陸</string>
			<key>Genre</key><string>Classic Rock</string>
			<key>Kind</key><string>MPEG オーディオファイル</string>
			<key>Size</key><integer>8916030</integer>
			<key>Total Time</key><integer>371487</integer>
			<key>Track Number</key><integer>12</integer>
			<key>Date Modified</key><date>2005-09-17T12:50:00Z</date>
			<key>Date Added</key><date>2005-09-16T11:14:19Z</date>
			<key>Bit Rate</key><integer>192</integer>
			<key>Sample Rate</key><integer>44100</integer>
			<key>Persistent ID</key><string>2A0848311638E609</string>
			<key>Track Type</key><string>File</string>
			<key>Location</key><string>file://localhost/C:/0_mp3_touch/(MP3)%5B%E3%82%A2%E3%83%AB%E3%83%90%E3%83%A0%5D%E3%82%AA%E3%83%A0%E3%83%8B%E3%83%90%E3%82%B9%20-%20%E6%83%85%E7%86%B1%E5%A4%A7%E9%99%B8-%E8%91%89%E5%8A%A0%E7%80%AC%E5%A4%AA%E9%83%8E%20SELECTION/Dream%20of%20Wings.mp3</string>
			<key>File Folder Count</key><integer>-1</integer>
			<key>Library Folder Count</key><integer>-1</integer>
		</dict>
	</dict>
</dict>
</plist>`))

	var data plistData
	d := plist.NewDecoder(buf)
	err = d.Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}

type trackData struct {
	T894 playData
}

type plistData struct {
	MajorVersion        uint64    `plist:"Major Version"`
	MinorVersion        uint64    `plist:"Minor Version"`
	ApplicationVersion  string    `plist:"Application Version"`
	Features            uint64    `plist:"Features"`
	ShowContentRatings  bool      `plist:"Show Content Ratings"`
	MusicFolder         string    `plist:"Music Folder"`
	LibraryPersistentID string    `plist:"Library Persistent ID"`
	Tracks              trackData `plist:"Tracks"`
}

type playData struct {
	TrackID uint64 `plist:"Track ID"`
	Name    string `plist:"Name"`
	// Artist             string    `plist:"Artist"`
	// Album              string    `plist:"Album"`
	// Genre              string    `plist:"Genre"`
	// Kind               string    `plist:"Kind"`
	// Size               uint64    `plist:"Size"`
	// TotalTime          uint64    `plist:"Total Time"`
	// TrackNumber        uint64    `plist:"Track Number"`
	// DateModified       time.Time `plist:"Date Modified"`
	// DateAdded          time.Time `plist:"Date Added"`
	// BitRate            uint64    `plist:"Bit Rate"`
	// SampleRate         uint64    `plist:"Sample Rate"`
	// PersistentID       string    `plist:"Persistent ID"`
	// TrackType          string    `plist:"Track Type"`
	// Location           string    `plist:"Location"`
	// FileFolderCount    uint64    `plist:"File Folder Count"`
	// LibraryFolderCount uint64    `plist:"Library Folder Count"`
	// PlayCount            string `plist:"Play Count"`
	// PlayDate             string `plist:"Play Date"`
	// PlayDateUTC          string `plist:"Play Date UTC"`
	// Rating               string `plist:"Rating"`
	// AlbumRating          string `plist:"Album Rating"`
	// AlbumRatingComputed  string `plist:"Album Rating Computed"`
	// ArtworkCount         string `plist:"Artwork Count"`
	// SkipCount            string `plist:"Skip Count"`
	// SkipDate             string `plist:"Skip Date"`
	// Disabled             string `plist:"Disabled"`
	// SortAlbum            string `plist:"Sort Album"`
	// Year                 string `plist:"Year"`
	// Comments             string `plist:"Comments"`
	// SortName             string `plist:"Sort Name"`
	// SortArtist           string `plist:"Sort Artist"`
	// VolumeAdjustment     string `plist:"Volume Adjustment"`
	// AlbumArtist          string `plist:"Album Artist"`
	// DiscNumber           string `plist:"Disc Number"`
	// DiscCount            string `plist:"Disc Count"`
	// TrackCount           string `plist:"Track Count"`
	// ReleaseDate          string `plist:"Release Date"`
	// Protected            string `plist:"Protected"`
	// Purchased            string `plist:"Purchased"`
	// Compilation          string `plist:"Compilation"`
	// Composer             string `plist:"Composer"`
	// HasVideo             string `plist:"Has Video"`
	// VideoWidth           string `plist:"Video Width"`
	// VideoHeight          string `plist:"Video Height"`
	// Movie                string `plist:"Movie"`
	// Master               string `plist:"Master"`
	// PlaylistID           string `plist:"Playlist ID"`
	// PlaylistPersistentID string `plist:"Playlist Persistent ID"`
	// Visible              string `plist:"Visible"`
	// AllItems             string `plist:"All Items"`
	// PlaylistItems        string `plist:"Playlist Items"`
	// SmartCriteria        string `plist:"Smart Criteria"`
}

// type Thumb struct {
// 	Title  string `xml:"title"`
// 	Length string `xml:"length"`
// }
