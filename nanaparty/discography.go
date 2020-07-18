package nanaparty

import (
	"context"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// NanaPartyDiscographyCollection is a collection of resource from NanaPartyDiscographyUrl
type NanaPartyDiscographyCollection interface {
	// Discographies return a list of recording info.
	Discographies() []NanaPartyDiscography
}

// NanaPartyDiscography provides a brief info of the discography, and a method to get detailed info of the discography.
type NanaPartyDiscography interface {
	// Info returns the info of recording.
	Info() NanaPartyDiscographyInfo
	// Detail does the query to the recording web page and returns the details of the recording.
	Detail(ctx context.Context) (*NanaPartyDiscographyDetail, error)
}

type NanaPartyDiscographyInfo struct {
	Group     string
	Link      string
	Form      string
	Date      string
	Title     string
	JacketUrl string
}

type NanaPartyDiscographyDetail struct {
	// TODO: 収録内容
	Editions    []*NanaPartyDiscographyDetailEdition
	Special     []*NanaPartyDiscographyDetailSpecial
	Information []*NanaPartyDiscographyDetailInformation
	Movie       []*NanaPartyDiscographyDetailMovie
}

type NanaPartyDiscographyDetailSpecial struct {
	// TODO
}

type NanaPartyDiscographyDetailInformation struct {
	Title string
	Link  string
}

type NanaPartyDiscographyDetailMovie struct {
	Type  string
	Title string
	Link  string
}

type NanaPartyDiscographyDetailEdition struct {
	JacketUrl  string
	Title      string
	Attributes []string
}

type nanaPartyDiscographyDetail struct {
	doc *goquery.Document
}

func newNanaPartyDiscography(doc *goquery.Document) *nanaPartyDiscographyDetail {
	return &nanaPartyDiscographyDetail{doc: doc}
}

func (n *nanaPartyDiscographyDetail) Discographies() []NanaPartyDiscography {
	return makeNanaPartyDiscography(n.doc.Find(".jacketBlock"))
}

type nanaPartyDiscography struct {
	NanaPartyDiscographyInfo
}

func (n *nanaPartyDiscography) Info() NanaPartyDiscographyInfo {
	return n.NanaPartyDiscographyInfo
}

func (n *nanaPartyDiscography) Detail(ctx context.Context) (*NanaPartyDiscographyDetail, error) {
	doc, err := getDocumentFromUrl(ctx, n.Link)
	if err != nil {
		return nil, fmt.Errorf("cannot parse document: %w", err)
	}
	return newNanaPartyDiscographyDetail(doc), nil
}

func newNanaPartyDiscographyDetail(doc *goquery.Document) *NanaPartyDiscographyDetail {
	return &NanaPartyDiscographyDetail{
		Information: makeNanaPartyDiscographyDetailInformation(doc.Find("#discoInfoBlock")),
		Editions:    makeNanaPartyDiscographyEdition(doc.Find("#discoJacketBlock")),
	}
}

func makeNanaPartyDiscographyDetailInformation(discoInfoBlock *goquery.Selection) []*NanaPartyDiscographyDetailInformation {
	var result []*NanaPartyDiscographyDetailInformation
	discoInfoBlock.Find(".discoInfo li > a").Each(func(i int, a *goquery.Selection) {
		result = append(result, &NanaPartyDiscographyDetailInformation{
			Title: a.Text(),
			Link:  a.AttrOr("href", ""),
		})
	})
	return result
}

func makeNanaPartyDiscographyEdition(discoJacketBlock *goquery.Selection) []*NanaPartyDiscographyDetailEdition {
	var result []*NanaPartyDiscographyDetailEdition
	discoJacketBlock.Find(".discoJacketBox").Each(func(i int, box *goquery.Selection) {
		src := NanaPartyDiscographyUrl + box.Find(".jk > img").AttrOr("src", "")
		title := box.Find("li.edition").Text()
		var attributes []string
		box.Find("li").Not(".edition").Each(func(i int, li *goquery.Selection) {
			attributes = append(attributes, li.Text())
		})
		result = append(result, &NanaPartyDiscographyDetailEdition{
			Title:      title,
			JacketUrl:  src,
			Attributes: attributes,
		})
	})
	return result
}

func makeNanaPartyDiscography(jacketBlock *goquery.Selection) []NanaPartyDiscography {
	var result []NanaPartyDiscography
	jacketBlock.Find("li").Each(func(i int, li *goquery.Selection) {
		var recording nanaPartyDiscography
		recording.Group, _ = getDiscographyTypeBySelection(li)

		a := li.Find("a")
		recording.Link = NanaPartyDiscographyUrl + a.AttrOr("href", "")
		recording.JacketUrl = NanaPartyDiscographyUrl + a.Find(".jk img").AttrOr("src", "")

		caption := a.Find(".caption")
		recording.Form = caption.Find(".rform").Text()
		recording.Date = caption.Find(".rday").Text()
		recording.Title = caption.Find(".rttl").Text()
		result = append(result, &recording)
	})
	return result
}

func getDiscographyTypeBySelection(s *goquery.Selection) (string, bool) {
	var typesGroup = []string{"single", "album", "movie"}
	for _, group := range typesGroup {
		if s.HasClass(group) {
			return group, true
		}
	}
	return "", false
}
