package nanaparty

import (
	"context"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// NanaPartyBlog provides a brief info of the blog post, and a method to get detailed info of the blog post.
type NanaPartyBlog interface {
	// Info returns a brief info of the blog post
	Info() NanaPartyBlogInfo
	// Detail does a query to the blog post link and returns the details from the blog post
	Detail(ctx context.Context) (*NanaPartyBlogDetail, error)
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

type NanaPartyBlogComment struct {
	Name    string
	Date    string
	Comment string
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

type nanaPartyBlog struct {
	NanaPartyBlogInfo
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
