package mizukinana

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

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

type nanaPartyTop struct {
	doc *goquery.Document
}

func newNanaPartyTop(doc *goquery.Document) *nanaPartyTop {
	return &nanaPartyTop{doc: doc}
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
