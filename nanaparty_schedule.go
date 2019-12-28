package mizukinana

import (
	"github.com/PuerkitoBio/goquery"
)

// NanaPartyScheduleCollection is a collection of resource from NanaPartyScheduleUrl
type NanaPartyScheduleCollection interface {
	// Regular returns a list of regular schedule
	Regular() []*NanaPartyScheduleRegular
	// Info returns a list of schedule information
	Info() []*NanaPartyScheduleInfo
}

type NanaPartyScheduleRegular struct {
	Category string
	InfoList []*NanaPartyScheduleRegularInfo
}

type NanaPartyScheduleRegularInfo struct {
	Info   string
	Detail string // HTML
}

type NanaPartyScheduleInfo struct {
	Date     string
	Group    string
	Category string
	Title    string
	Link     string
}

type nanaPartySchedule struct {
	doc *goquery.Document
}

func newNanaPartySchedule(doc *goquery.Document) *nanaPartySchedule {
	return &nanaPartySchedule{doc: doc}
}

func (n *nanaPartySchedule) Regular() []*NanaPartyScheduleRegular {
	return makeNanaPartyScheduleRegular(n.doc.Find(".regular_list"))
}

func (n *nanaPartySchedule) Info() []*NanaPartyScheduleInfo {
	return makeNanaPartyScheduleInfo(n.doc.Find(".info_list"))
}

func makeNanaPartyScheduleRegular(regularList *goquery.Selection) []*NanaPartyScheduleRegular {
	var result []*NanaPartyScheduleRegular
	regularList.Find("dl").Each(func(i int, dl *goquery.Selection) {
		category := dl.Find("dt").Text()
		result = append(result, &NanaPartyScheduleRegular{
			Category: category,
			InfoList: makeNanaPartyScheduleRegularInfo(dl),
		})
	})
	return result
}

func makeNanaPartyScheduleRegularInfo(dl *goquery.Selection) []*NanaPartyScheduleRegularInfo {
	var result []*NanaPartyScheduleRegularInfo
	dl.Find("dd").Each(func(i int, dd *goquery.Selection) {
		info := dd.Find(".list_day").Text()
		detail, _ := dd.Find(".list_info").Html()
		result = append(result, &NanaPartyScheduleRegularInfo{
			Info:   info,
			Detail: detail,
		})
	})
	return result
}

func makeNanaPartyScheduleInfo(infoList *goquery.Selection) []*NanaPartyScheduleInfo {
	var result []*NanaPartyScheduleInfo
	infoList.Find("li").Each(func(i int, li *goquery.Selection) {
		group, _ := findGroupBySelection(li)
		a := li.Find(".schListTxt > a")
		result = append(result, &NanaPartyScheduleInfo{
			Date:     li.Find(".date").Text(),
			Group:    group,
			Category: li.Find(".category").Text(),
			Title:    a.Text(),
			Link:     NanaPartyUrl + a.AttrOr("href", ""),
		})
	})
	return result
}
