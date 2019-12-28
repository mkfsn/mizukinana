package mizukinana

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// NanaPartyBiographyCollection is a collection of resources from NanaPartyBiographyUrl.
type NanaPartyBiographyCollection interface {
	// Profile returns a profile info from NanaPartyBiographyUrl
	Profile() *NanaPartyBiographyProfile
	// Voice returns a list of voice info from NanaPartyBiographyUrl
	Voice() []*NanaPartyBiographyVoice
	// Live returns a list of live info from NanaPartyBiographyUrl
	Live() []*NanaPartyBiographyLive
	// Special returns a list of special live/event info from NanaPartyBiographyUrl
	Special() []*NanaPartyBiographySpecial
	// Other returns a list of other info from NanaPartyBiographyUrl
	Other() []*NanaPartyBiographyOther
}

type NanaPartyBiographyProfile struct {
	Name     string
	Info     map[string]string
	PhotoUrl string
}

func newNanaPartyBiographyProfile(profileBlock *goquery.Selection) *NanaPartyBiographyProfile {
	nameBlock := profileBlock.Find(".name2 ruby")
	nameBlock.Find("rp").Remove()
	nameBlock.Find("rt").Remove()
	name := strings.TrimSpace(nameBlock.Text())
	name = strings.ReplaceAll(name, " ", "")
	name = strings.ReplaceAll(name, "\n", "")

	var photoUrl string
	photoBlock := profileBlock.Find(".profilePhoto")
	if src, ok := photoBlock.Find(".pPhoto").Attr("src"); ok {
		photoUrl = fmt.Sprintf("%s/%s", strings.TrimSuffix(NanaPartyBiographyUrl, "/"), src)
	}

	info := make(map[string]string)
	profileBlock.Find(".infoTxt dt").Each(func(i int, s *goquery.Selection) {
		info[s.Text()] = s.Next().Text()
	})

	return &NanaPartyBiographyProfile{
		Name:     name,
		Info:     info,
		PhotoUrl: photoUrl,
	}
}

type NanaPartyBiographyVoice struct {
	Category string
	InfoList []NanaPartyBiographyVoiceInfo
}

func makeNanaPartyBiographyVoice(voiceBlock *goquery.Selection) []*NanaPartyBiographyVoice {
	var result []*NanaPartyBiographyVoice
	voiceBlock.Find("dl.bio_list dt").Each(func(i int, dt *goquery.Selection) {
		item := &NanaPartyBiographyVoice{
			Category: dt.Text(),
			InfoList: makeNanaPartyBiographyVoiceInfo(dt.Next()),
		}
		result = append(result, item)
	})
	return result
}

type NanaPartyBiographyVoiceInfo struct {
	Title       string
	Annotations []string
}

func makeNanaPartyBiographyVoiceInfo(dd *goquery.Selection) []NanaPartyBiographyVoiceInfo {
	var result []NanaPartyBiographyVoiceInfo
	dd.Find(".binfo li ").Each(func(i int, li *goquery.Selection) {
		var annotations []string
		li.Find("span").Each(func(i int, span *goquery.Selection) {
			annotations = append(annotations, span.Text())
		})
		li.Find("span").Remove()
		result = append(result, NanaPartyBiographyVoiceInfo{
			Title:       li.Text(),
			Annotations: annotations,
		})
	})
	return result
}

type NanaPartyBiographyLive struct {
	Title         string
	Annotations   []string
	EventDay      string
	Events        []EventItemList
	EventWebsites []EventWebsite
}

func makeNanaPartyBiographyLive(liveBlock *goquery.Selection) []*NanaPartyBiographyLive {
	result := make([]*NanaPartyBiographyLive, 0)
	liveBlock.Find(".yearBlock > .info_list > li").Each(func(i int, s *goquery.Selection) {
		result = append(result, newNanaPartyBiographyLive(s))
	})
	return result
}

func newNanaPartyBiographyLive(li *goquery.Selection) *NanaPartyBiographyLive {
	var annotations []string
	li.Find(".binfo span").Each(func(i int, s *goquery.Selection) {
		annotations = append(annotations, strings.TrimSpace(s.Text()))
	})

	info := li.Find(".binfo").First()
	info.Find("span").Remove()
	title := info.Text()

	return &NanaPartyBiographyLive{
		Title:         strings.TrimSpace(title),
		Annotations:   annotations,
		EventDay:      strings.TrimSpace(li.Find(".binfo_day").Text()),
		Events:        makeEventItemList(li),
		EventWebsites: makeEventWebsite(li),
	}
}

type NanaPartyBiographySpecial struct {
	Title         string
	Annotations   []string
	EventDay      string
	Events        []EventItemList
	EventWebsites []EventWebsite
}

func makeNanaPartyBiographySpecial(spliveBlock *goquery.Selection) []*NanaPartyBiographySpecial {
	result := make([]*NanaPartyBiographySpecial, 0)
	spliveBlock.Find(".yearBlock > .info_list > li").Each(func(i int, li *goquery.Selection) {
		result = append(result, newNanaPartyBiographySpecial(li))
	})
	return result
}

func newNanaPartyBiographySpecial(li *goquery.Selection) *NanaPartyBiographySpecial {
	var annotations []string
	li.Find(".binfo span").Each(func(i int, s *goquery.Selection) {
		annotations = append(annotations, strings.TrimSpace(s.Text()))
	})

	info := li.Find(".binfo").First()
	info.Find("span").Remove()
	title := info.Text()

	return &NanaPartyBiographySpecial{
		Title:         strings.TrimSpace(title),
		Annotations:   annotations,
		EventDay:      strings.TrimSpace(li.Find(".binfo_day").Text()),
		Events:        makeEventItemList(li),
		EventWebsites: makeEventWebsite(li),
	}
}

type NanaPartyBiographyOther struct {
	Category string
	InfoList []NanaPartyBiographyOtherInfo
}

func makeNanaPartyBiographyOther(otherBlock *goquery.Selection) []*NanaPartyBiographyOther {
	var result []*NanaPartyBiographyOther
	otherBlock.Find("dl.bio_list dt").Each(func(i int, dt *goquery.Selection) {
		item := &NanaPartyBiographyOther{
			Category: dt.Text(),
			InfoList: makeNanaPartyBiographyOtherInfo(dt.Next()),
		}
		result = append(result, item)
	})
	return result
}

type NanaPartyBiographyOtherInfo struct {
	Title       string
	Annotations []string
}

func makeNanaPartyBiographyOtherInfo(dd *goquery.Selection) []NanaPartyBiographyOtherInfo {
	var result []NanaPartyBiographyOtherInfo
	dd.Find(".binfo li ").Each(func(i int, li *goquery.Selection) {
		var annotations []string
		li.Find("span").Each(func(i int, span *goquery.Selection) {
			annotations = append(annotations, span.Text())
		})
		li.Find("span").Remove()
		result = append(result, NanaPartyBiographyOtherInfo{
			Title:       li.Text(),
			Annotations: annotations,
		})
	})
	return result
}

type EventItemList []EventItem

func makeEventItemList(infoBlock *goquery.Selection) []EventItemList {
	var events []EventItemList
	infoBlock.Find(".binfo_list table tr").Each(func(i int, row *goquery.Selection) {
		events = append(events, makeEventItem(row))
	})
	return events
}

type EventItem struct {
	Attributes []string
	Content    string
}

func makeEventItem(row *goquery.Selection) []EventItem {
	event := []EventItem{
		newEventItem(row.Find("td.num").Text(), "num"),
	}
	row.Find("td li").Each(func(i int, li *goquery.Selection) {
		item := newEventItem(li.Text(), strings.Split(li.AttrOr("class", ""), " ")...)
		event = append(event, item)
	})
	return event
}

func newEventItem(content string, attributes ...string) EventItem {
	return EventItem{Content: content, Attributes: attributes}
}

type EventWebsite struct {
	Title string
	Url   string
}

func makeEventWebsite(infoBlock *goquery.Selection) []EventWebsite {
	var websites []EventWebsite
	infoBlock.Find(".binfo a").Each(func(i int, s *goquery.Selection) {
		websites = append(websites, EventWebsite{
			Title: strings.TrimSpace(s.Text()),
			Url:   s.AttrOr("href", ""),
		})
	})
	return websites
}

type nanaPartyBiography struct {
	doc *goquery.Document
}

func newNanaPartyBiography(doc *goquery.Document) NanaPartyBiographyCollection {
	return &nanaPartyBiography{doc: doc}
}

func (n *nanaPartyBiography) Profile() *NanaPartyBiographyProfile {
	return newNanaPartyBiographyProfile(n.doc.Find(".profileBlock"))
}

func (n *nanaPartyBiography) Voice() []*NanaPartyBiographyVoice {
	return makeNanaPartyBiographyVoice(n.doc.Find(".voiceBlock"))
}

func (n *nanaPartyBiography) Live() []*NanaPartyBiographyLive {
	return makeNanaPartyBiographyLive(n.doc.Find(".liveBlock"))
}

func (n *nanaPartyBiography) Special() []*NanaPartyBiographySpecial {
	return makeNanaPartyBiographySpecial(n.doc.Find(".spliveBlock"))
}

func (n *nanaPartyBiography) Other() []*NanaPartyBiographyOther {
	return makeNanaPartyBiographyOther(n.doc.Find(".otherBlock"))
}
