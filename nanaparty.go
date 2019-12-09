package mizukinana

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// A URL collection in NanaParty
const (
	NanaPartyUrl            = "https://www.mizukinana.jp"
	NanaPartyNewsUrl        = NanaPartyUrl + "/news/"
	NanaPartyBiographyUrl   = NanaPartyUrl + "/biography/"
	NanaPartyBlogUrl        = NanaPartyUrl + "/blog/"
	NanaPartyBlogListUrl    = NanaPartyBlogUrl + "backnumber.html"
	NanaPartyScheduleUrl    = NanaPartyUrl + "/schedule/"
	NanaPartyDiscographyUrl = NanaPartyUrl + "/discography/"
)

// NanaPartyCollection is a collection of resources from the whole NanaParty website.
type NanaPartyCollection interface {
	// Top does the query to NanaPartyUrl and return a collection from the web page.
	Top(ctx context.Context) (NanaPartyTopCollection, error)
	// Biography does the query to NanaPartyBiographyUrl and return a collection from the web page.
	Biography(ctx context.Context) (NanaPartyBiographyCollection, error)
	// Blog does the query to the NanaPartyBlogListUrl and return a list of Blog from the web page.
	Blog(ctx context.Context) ([]NanaPartyBlog, error)
	// News does the query to NanaPartyNewsUrl and returns a list of news from the web page.
	News(ctx context.Context) ([]*NanaPartyNews, error)
	// Schedule does the query to NanaPartyScheduleUrl and returns a collection from the web page.
	Schedule(ctx context.Context) (NanaPartyScheduleCollection, error)
	// Discography does the query to NanaPartyDiscographyUrl and returns a collection from the web page.
	Discography(ctx context.Context) (NanaPartyDiscographyCollection, error)
}

// NanaPartyBiographyCollection is a collection of resources from NanaPartyUrl.
type NanaPartyTopCollection interface {
	// Main returns a list of top-main item in NanaPartyUrl
	Main() []*NanaPartyTopMainItem
	// Main returns a list of top-pickup item in NanaPartyUrl
	Pickup() []*NanaPartyTopPickupItem
	// Main returns the top-movie item in NanaPartyUrl
	Movie() *NanaPartyTopMovieItem
	// Main returns a list of top-banner item in NanaPartyUrl
	Banner() []*NanaPartyTopBannerItem
	// Main returns a list of top-topic item in NanaPartyUrl
	Topic() []*NanaPartyTopTopicItem
}

// NanaPartyScheduleCollection is a collection of resource from NanaPartyScheduleUrl
type NanaPartyScheduleCollection interface {
	// Regular returns a list of regular schedule
	Regular() []*NanaPartyScheduleRegular
	// Info returns a list of schedule information
	Info() []*NanaPartyScheduleInfo
}

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

// NanaPartyBlog provides a brief info of the blog post, and a method to get detailed info of the blog post.
type NanaPartyBlog interface {
	// Info returns a brief info of the blog post
	Info() NanaPartyBlogInfo
	// Detail does a query to the blog post link and returns the details from the blog post
	Detail(ctx context.Context) (*NanaPartyBlogDetail, error)
}

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

type NanaPartyTopMainItem struct {
	ImageUrl   string
	Annotation string
}

type NanaPartyTopPickupItem struct {
	Link       string
	ImageUrl   string
	Annotation string
}

type NanaPartyTopMovieItem struct {
	Title    string
	Link     string
	ImageUrl string
}

type NanaPartyTopBannerItem struct {
	Link       string
	PhotoUrl   string
	Annotation string
}

type NanaPartyTopTopicItem struct {
	Date             string
	Category         string
	Link             string
	AbbreviatedTitle string
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

type NanaPartyBlogInfo struct {
	Date  string
	Title string
	Link  string
}

type NanaPartyBlogDetail struct {
	Content  string // HTML
	Comments []*NanaPartyBlogComment
}

type NanaPartyBlogComment struct {
	Name    string
	Date    string
	Comment string
}

type NanaPartyBiographyProfile struct {
	Name     string
	Info     map[string]string
	PhotoUrl string
}

type NanaPartyBiographyVoice struct {
	Category string
	InfoList []NanaPartyBiographyVoiceInfo
}

type NanaPartyBiographyVoiceInfo struct {
	Title       string
	Annotations []string
}

type NanaPartyBiographyLive struct {
	Title         string
	Annotations   []string
	EventDay      string
	Events        []EventItemList
	EventWebsites []EventWebsite
}

type NanaPartyBiographySpecial struct {
	Title         string
	Annotations   []string
	EventDay      string
	Events        []EventItemList
	EventWebsites []EventWebsite
}

type NanaPartyBiographyOther struct {
	Category string
	InfoList []NanaPartyBiographyOtherInfo
}

type NanaPartyBiographyOtherInfo struct {
	Title       string
	Annotations []string
}

type EventItemList []EventItem

type EventItem struct {
	Attributes []string
	Content    string
}

type EventWebsite struct {
	Title string
	Url   string
}

type NanaPartyNews struct {
	Id       string
	Date     string
	Group    string
	Category string
	Title    string
	Detail   string // HTML
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
	Editions []*NanaPartyDiscographyEdition
	// TODO
}

type NanaPartyDiscographyEdition struct {
	JacketUrl  string
	Title      string
	Attributes []string
}

type nanaParty struct{}

type nanaPartyTop struct {
	doc *goquery.Document
}

type nanaPartyBlog struct {
	NanaPartyBlogInfo
}

type nanaPartyDiscography struct {
	NanaPartyDiscographyInfo
}

type nanaPartyDiscographyDetail struct {
	doc *goquery.Document
}

type nanaPartyBiography struct {
	doc *goquery.Document
}

type nanaPartySchedule struct {
	doc *goquery.Document
}

// NanaParty returns a NanaPartyCollection.
func NanaParty() NanaPartyCollection {
	return &nanaParty{}
}

func (n *nanaParty) Top(ctx context.Context) (NanaPartyTopCollection, error) {
	doc, err := getDocumentFromUrl(ctx, NanaPartyUrl)
	if err != nil {
		return nil, fmt.Errorf("cannot parse document: %w", err)
	}
	return newNanaPartyTop(doc), nil
}

func (n *nanaParty) Biography(ctx context.Context) (NanaPartyBiographyCollection, error) {
	doc, err := getDocumentFromUrl(ctx, NanaPartyBiographyUrl)
	if err != nil {
		return nil, fmt.Errorf("cannot parse document: %w", err)
	}
	return newNanaPartyBiography(doc), nil
}

func (n *nanaParty) News(ctx context.Context) ([]*NanaPartyNews, error) {
	doc, err := getDocumentFromUrl(ctx, NanaPartyNewsUrl)
	if err != nil {
		return nil, fmt.Errorf("cannot parse document: %w", err)
	}
	return makeNanaPartyNews(doc), nil
}

func (n *nanaParty) Blog(ctx context.Context) ([]NanaPartyBlog, error) {
	doc, err := getDocumentFromUrl(ctx, NanaPartyBlogListUrl)
	if err != nil {
		return nil, fmt.Errorf("cannot parse document: %w", err)
	}
	return makeNanaPartyBlog(doc), nil
}

func (n *nanaParty) Schedule(ctx context.Context) (NanaPartyScheduleCollection, error) {
	doc, err := getDocumentFromUrl(ctx, NanaPartyScheduleUrl)
	if err != nil {
		return nil, fmt.Errorf("cannot parse document: %w", err)
	}
	return newNanaPartySchedule(doc), nil
}

func (n *nanaParty) Discography(ctx context.Context) (NanaPartyDiscographyCollection, error) {
	doc, err := getDocumentFromUrl(ctx, NanaPartyDiscographyUrl)
	if err != nil {
		return nil, fmt.Errorf("cannot parse document: %w", err)
	}
	return newNanaPartyDiscography(doc), nil
}

func (n *nanaPartyDiscographyDetail) Discographies() []NanaPartyDiscography {
	return makeNanaPartyDiscography(n.doc.Find(".jacketBlock"))
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

func (n *nanaPartySchedule) Regular() []*NanaPartyScheduleRegular {
	return makeNanaPartyScheduleRegular(n.doc.Find(".regular_list"))
}

func (n *nanaPartySchedule) Info() []*NanaPartyScheduleInfo {
	return makeNanaPartyScheduleInfo(n.doc.Find(".info_list"))
}

func (n nanaPartyBlog) Info() NanaPartyBlogInfo {
	return n.NanaPartyBlogInfo
}

func (n nanaPartyBlog) Detail(ctx context.Context) (*NanaPartyBlogDetail, error) {
	doc, err := getDocumentFromUrl(ctx, n.Link)
	if err != nil {
		return nil, fmt.Errorf("cannot parse document: %w", err)
	}
	return newNanaPartyBlogDetail(doc), nil
}

func (n *nanaPartyTop) Main() []*NanaPartyTopMainItem {
	return makeNanaPartyTopMainItem(n.doc.Find("#topMainBlock"))
}

func (n *nanaPartyTop) Pickup() []*NanaPartyTopPickupItem {
	return makeNanaPartyTopPickupItem(n.doc.Find("#topPickupBlock"))
}

func (n *nanaPartyTop) Movie() *NanaPartyTopMovieItem {
	return newNanaPartyTopMovieItem(n.doc.Find("#topMoviesBlock"))
}

func (n *nanaPartyTop) Banner() []*NanaPartyTopBannerItem {
	return makeNanaPartyTopBannerItem(n.doc.Find("#topBannerBlock"))
}

func (n nanaPartyTop) Topic() []*NanaPartyTopTopicItem {
	return makeNanaPartyTopTopicItem(n.doc.Find("#topTopicsBlock"))
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

func newNanaPartyDiscographyDetail(doc *goquery.Document) *NanaPartyDiscographyDetail {

	return &NanaPartyDiscographyDetail{}
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

func newNanaPartyDiscography(doc *goquery.Document) *nanaPartyDiscographyDetail {
	return &nanaPartyDiscographyDetail{doc: doc}
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

func newNanaPartySchedule(doc *goquery.Document) *nanaPartySchedule {
	return &nanaPartySchedule{doc: doc}
}

func makeNanaPartyTopTopicItem(topTopicsBlock *goquery.Selection) []*NanaPartyTopTopicItem {
	var result []*NanaPartyTopTopicItem
	topTopicsBlock.Find("dl > dt").Each(func(i int, dt *goquery.Selection) {
		dd := dt.Next()
		category := dt.Find(".category").Text()
		category = strings.TrimPrefix(category, "【 ")
		category = strings.TrimSuffix(category, " 】")

		dt.Find(".category").Remove()
		date := dt.Text()

		a := dd.Find("a")
		link := a.AttrOr("href", "")
		title := strings.ReplaceAll(a.Text(), "\n", "")

		result = append(result, &NanaPartyTopTopicItem{
			Date:             date,
			Category:         category,
			Link:             NanaPartyUrl + "/" + link,
			AbbreviatedTitle: title,
		})
	})
	return result
}

func makeNanaPartyTopBannerItem(topBannerBlock *goquery.Selection) []*NanaPartyTopBannerItem {
	var result []*NanaPartyTopBannerItem
	topBannerBlock.Find("li a").Each(func(i int, a *goquery.Selection) {
		link := a.AttrOr("href", "")
		img := a.Find("img")
		src := img.AttrOr("src", "")
		alt := img.AttrOr("alt", "")
		result = append(result, &NanaPartyTopBannerItem{
			Link:       link,
			PhotoUrl:   "https:" + src,
			Annotation: alt,
		})
	})
	return result
}

func newNanaPartyTopMovieItem(topMoviesBlock *goquery.Selection) *NanaPartyTopMovieItem {
	a := topMoviesBlock.Find(".moviesBox a")
	link := a.AttrOr("href", "")
	title := a.AttrOr("title", "")
	src := a.Find("img.moviesParts").AttrOr("src", "")
	return &NanaPartyTopMovieItem{
		Link:     link,
		Title:    title,
		ImageUrl: "https:" + src,
	}
}

func makeNanaPartyTopPickupItem(topPickupBlock *goquery.Selection) []*NanaPartyTopPickupItem {
	var result []*NanaPartyTopPickupItem
	topPickupBlock.Find(".bxslider li").Each(func(i int, li *goquery.Selection) {
		link := li.Find("a").AttrOr("href", "")
		src := li.Find("img").AttrOr("src", "")
		alt := li.Find("img").AttrOr("alt", "")
		result = append(result, &NanaPartyTopPickupItem{
			Link:       link,
			ImageUrl:   "https:" + src,
			Annotation: alt,
		})
	})
	return result
}

func makeNanaPartyTopMainItem(topMainBlock *goquery.Selection) []*NanaPartyTopMainItem {
	var result []*NanaPartyTopMainItem
	topMainBlock.Find(".bxslider li").Each(func(i int, li *goquery.Selection) {
		img := li.Find(".main-images")
		src := img.AttrOr("src", "")
		alt := img.AttrOr("alt", "")
		result = append(result, &NanaPartyTopMainItem{
			ImageUrl:   "https:" + src,
			Annotation: alt,
		})
	})
	return result
}

func newNanaPartyTop(doc *goquery.Document) *nanaPartyTop {
	return &nanaPartyTop{doc: doc}
}

func newNanaPartyBlogDetail(doc *goquery.Document) *NanaPartyBlogDetail {
	postBlock := doc.Find("#postBlock")
	postContent, _ := postBlock.Find(".postcontent > .message").Html()

	var comments []*NanaPartyBlogComment
	postBlock.Find(".commentBox li").Each(func(i int, li *goquery.Selection) {
		comments = append(comments, newNanaPartyBlogComment(li))
	})

	return &NanaPartyBlogDetail{
		Content:  postContent,
		Comments: comments,
	}
}

func newNanaPartyBlogComment(li *goquery.Selection) *NanaPartyBlogComment {
	footer := li.Find("span.cdate")

	published := footer.Find("abbr.published")
	date := published.AttrOr("title", "")

	published.Remove()
	name := strings.TrimSuffix(footer.Text(), " | ")

	footer.Remove()
	comment, _ := li.Html()

	return &NanaPartyBlogComment{
		Name:    name,
		Date:    date,
		Comment: comment,
	}
}

func makeNanaPartyBlog(doc *goquery.Document) []NanaPartyBlog {
	var result []NanaPartyBlog
	doc.Find("dl.backnumber_page dt").Each(func(i int, dt *goquery.Selection) {
		dd := dt.Next()
		item := &nanaPartyBlog{
			NanaPartyBlogInfo{
				Date:  dt.Text(),
				Title: dd.Find("a").Text(),
				Link:  dd.Find("a").AttrOr("href", ""),
			},
		}
		result = append(result, item)
	})
	return result
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

func newNanaPartyBiography(doc *goquery.Document) NanaPartyBiographyCollection {
	return &nanaPartyBiography{doc: doc}
}

func makeNanaPartyNews(doc *goquery.Document) []*NanaPartyNews {
	var result []*NanaPartyNews
	doc.Find(".info_list li.secNews").Each(func(i int, s *goquery.Selection) {
		result = append(result, newNews(s))
	})
	return result
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

func makeEventItemList(infoBlock *goquery.Selection) []EventItemList {
	var events []EventItemList
	infoBlock.Find(".binfo_list table tr").Each(func(i int, row *goquery.Selection) {
		events = append(events, makeEventItem(row))
	})
	return events
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

func getDiscographyTypeBySelection(s *goquery.Selection) (string, bool) {
	var typesGroup = []string{"single", "album", "movie"}
	for _, group := range typesGroup {
		if s.HasClass(group) {
			return group, true
		}
	}
	return "", false
}

func getDocumentFromUrl(ctx context.Context, url string) (*goquery.Document, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot send request: %w", err)
	}
	defer res.Body.Close()

	return goquery.NewDocumentFromReader(res.Body)
}
