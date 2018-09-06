package concerts

import (
	"bytes"
	// "encoding/json"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

const (
	minTitleWidth = 44
)

type Concert struct {
	Title    string    `json:"title"    yaml:"title"`
	Date     time.Time `json:"date"     yaml:"date"`
	Open     time.Time `json:"open"     yaml:"open"`
	Start    time.Time `json:"start"    yaml:"start"`
	Location string    `json:"location" yaml:"location"`
	Price    price     `json:"price"    yaml:"price"`
	Note     string    `json:"note"     yaml:"note"`
}

func NewConcert(title string, date time.Time, location string, open, start time.Duration, price price, note string) Concert {
	return Concert{
		Title:    title,
		Date:     date,
		Location: location,
		Open:     date.Add(open),
		Start:    date.Add(start),
		Price:    price,
		Note:     note,
	}
}

type concerts []Concert

func (c concerts) MarshalTable() ([]byte, error) {
	var buf bytes.Buffer

	table := tablewriter.NewWriter(&buf)
	table.SetColWidth(minTitleWidth)
	table.SetHeader(c.header())
	table.AppendBulk(c.body())
	table.Render()

	return buf.Bytes(), nil
}

func (c concerts) header() []string {
	return []string{
		"Date",
		"Open",
		"Start",
		"Title",
		"Location",
		"Price",
		// "Note",
	}
}

func (c concerts) body() [][]string {
	rows := make([][]string, len(c))
	for i, concert := range c {
		rows[i] = []string{
			concert.Date.Format("2006-01-02"),
			concert.Open.Format("15:04"),
			concert.Start.Format("15:04"),
			concert.Title,
			concert.Location,
			strings.Replace(concert.Price.Price(), PriceSeperator, "\n", -1),
			// concert.Note,
		}
	}
	return rows
}
