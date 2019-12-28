package mizukinana

import (
	"context"
	"fmt"
	"net/http"

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

type nanaParty struct{}

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
