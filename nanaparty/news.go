package nanaparty

import (
	"context"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type NanaPartyNews struct {
	Id       string
	Date     string
	Group    string
	Category string
	Title    string
	Detail   string // HTML
}

func (n *nanaParty) News(ctx context.Context) ([]*NanaPartyNews, error) {
	doc, err := getDocumentFromUrl(ctx, NanaPartyNewsUrl)
	if err != nil {
		return nil, fmt.Errorf("cannot parse document: %w", err)
	}
	return makeNanaPartyNews(doc), nil
}

func makeNanaPartyNews(doc *goquery.Document) []*NanaPartyNews {
	var result []*NanaPartyNews
	doc.Find(".info_list li.secNews").Each(func(i int, s *goquery.Selection) {
		result = append(result, newNews(s))
	})
	return result
}

func newNews(s *goquery.Selection) *NanaPartyNews {
	id := s.AttrOr("id", "")
	group, _ := findGroupBySelection(s)
	detail, _ := s.Find(".detail").Html()
	titleBlock := s.Find(".title")
	date := titleBlock.Find(".date").Text()
	title := titleBlock.Find(".titletxt").Text()
	category := titleBlock.Find(".category").Text()
	return &NanaPartyNews{
		Id:       id,
		Date:     date,
		Group:    group,
		Title:    title,
		Category: category,
		Detail:   detail,
	}
}

func findGroupBySelection(s *goquery.Selection) (string, bool) {
	var newsGroups = []string{"voice", "media", "release", "live", "other"}
	for _, group := range newsGroups {
		if s.HasClass(group) {
			return group, true
		}
	}
	return "", false
}
