package metadata

import (
	"encoding/xml"
	"io"
	"io/ioutil"

	"github.com/bpetetot/leafer/db"
)

// ComicInfo contains all ComicInfo.xml data
type ComicInfo struct {
	XMLName     xml.Name `xml:"ComicInfo"`
	Series      string   `xml:"Series"`
	Title       string   `xml:"Title"`
	Summary     string   `xml:"Summary"`
	Publisher   string   `xml:"Publisher"`
	Genre       string   `xml:"Genre"`
	LanguageISO string   `xml:"LanguageISO"`
	Number      int      `xml:"Number"`
	Volume      int      `xml:"Volume"`
}

func readComicInfo(r io.Reader) (*db.Media, error) {
	byteValue, _ := ioutil.ReadAll(r)

	var comicInfo ComicInfo
	err := xml.Unmarshal(byteValue, &comicInfo)
	if err != nil {
		return nil, err
	}

	return &db.Media{
		Title:       comicInfo.Series,
		Description: comicInfo.Summary,
		Volume:      comicInfo.Volume,
	}, nil
}
