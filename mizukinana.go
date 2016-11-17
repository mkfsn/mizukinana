package mizukinana

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"time"
)

const (
	Name                            string = "水樹奈々"
	BirthYear, BirthMonth, BirthDay int    = 1980, 1, 12
	OfficialWebsite                 string = "http://www.mizukinana.jp/news/"
	Fanclub                         string = "http://fanclub.mizukinana.jp/"
)

type FeedContent struct {
	Title string
	Date  string
	Url   string
	Site  string
	Html  string
}
type FeedContentList []FeedContent

func GetAge() int {
	age, location := 0, time.Now().Location()
	for {
		if !time.Date(age+BirthYear, time.Month(BirthMonth), BirthDay, 0, 0, 0, 0, location).Before(time.Now()) {
			break
		}
		age += 1
	}
	return age - 1
}

func FetchOfficialWebsite() (FeedContentList, error) {
	var contents FeedContentList

	doc, err := goquery.NewDocument(OfficialWebsite)
	if err != nil {
		log.Fatal(err)
		return contents, err
	}

	doc.Find("#content > div").Each(func(i int, s *goquery.Selection) {
		id, _ := s.Attr("id")
		contents = append(contents, FeedContent{
			Date:  s.Find("div.title > p.date").Text(),
			Title: s.Find("div.title > h3").Text(),
			Site:  "officialwebsite",
			Url:   fmt.Sprintf("%s#%s", OfficialWebsite, id),
			Html:  "",
		})
	})

	return contents, nil
}

func FetchFanclub() (FeedContentList, error) {
	var contents FeedContentList

	doc, err := goquery.NewDocument(Fanclub)
	if err != nil {
		log.Fatal(err)
		return contents, err
	}

	doc.Find("a.title").Each(func(i int, s *goquery.Selection) {
		contents = append(contents, FeedContent{
			Date:  s.Prev().Text(),
			Title: s.Text(),
			Site:  "fanclub",
			Url:   Fanclub,
			Html:  "",
		})
	})

	return contents, nil
}

func (f FeedContent) String() string {
	return fmt.Sprintf("{title:%s, date:%s, url:%s, site:%s}", f.Title, f.Date, f.Url, f.Site)
}

func (t FeedContent) JSON() []byte {
	b, _ := json.Marshal(t)
	return b
}

func (t FeedContentList) JSON() []byte {
	b, _ := json.Marshal(t)
	return b
}
