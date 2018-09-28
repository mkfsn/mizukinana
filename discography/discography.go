package discography

//go:generate go run gen.go -input album.yaml -varname Albums -vartype DiscographyList
//go:generate go run gen.go -input single.yaml -varname Singles -vartype DiscographyList

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Edition struct {
	CatalogNumber string   `yaml:"catalog_number"`
	Price         int      `yaml:"price"`
	Tracklist     []Song   `yaml:"tracklist"`
	Extras        []string `yaml:"extras"`
}

type edition = string

type Discography struct {
	Title    string              `yaml:"title"`
	Released time.Time           `yaml:"released"`
	Editions map[edition]Edition `yaml:"editions"`
}

func (d *Discography) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var discography struct {
		Title    string              `yaml:"title"`
		Released string              `yaml:"released"`
		Editions map[edition]Edition `yaml:"editions"`
	}

	if err := unmarshal(&discography); err != nil {
		return err
	}

	var err error
	d.Title = discography.Title
	d.Editions = discography.Editions
	d.Released, err = time.ParseInLocation("2006-01-02", discography.Released, japan)
	return err
}

type DiscographyList []Discography

func (d DiscographyList) MarshalTable() ([]byte, error) {
	var buf bytes.Buffer

	table := tablewriter.NewWriter(&buf)
	table.SetHeader(d.header())
	table.AppendBulk(d.body())
	table.Render()

	return buf.Bytes(), nil
}

func (d DiscographyList) header() []string {
	return []string{
		"Title",
		"Released",
		"Edition",
		"CatalogNumber",
		"Price",
		"Tracelist",
		"Extras",
	}
}

func (d DiscographyList) body() [][]string {
	rows := make([][]string, 0)
	for _, discography := range d {
		for e, edition := range discography.Editions {
			tracklist := make([]string, len(edition.Tracklist))
			for i, track := range edition.Tracklist {
				tracklist[i] = fmt.Sprintf("%d.%s\n", i+1, track.Name)
			}

			rows = append(rows, []string{
				discography.Title,
				discography.Released.Format("2006-01-02"),
				e,
				edition.CatalogNumber,
				fmt.Sprintf("%d", edition.Price),
				strings.Join(tracklist, "\n"),
				strings.Join(edition.Extras, "\n"),
			})
		}
	}
	return rows
}
