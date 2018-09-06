package concerts

import (
	"bytes"
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

func (c Concert) GetDate() string {
	return c.Date.Format("2006-01-02")
}

func (c Concert) GetOpenTime() string {
	return c.Open.Format("15:04")
}

func (c Concert) GetStartTime() string {
	return c.Start.Format("15:04")
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

func (c Concert) filter(key, value string) bool {
	if value == "" {
		return false
	}

	switch key {
	case "title":
		return strings.Contains(c.Title, value)
	case "date":
		return strings.Contains(c.GetDate(), value)
	case "open":
		return strings.Contains(c.GetOpenTime(), value)
	case "start":
		return strings.Contains(c.GetStartTime(), value)
	case "location":
		return strings.Contains(c.Location, value)
	case "price":
		return strings.Contains(c.Price.Price(), value)
	case "note":
		return strings.Contains(c.Note, value)
	}

	return false
}

type Concerts []Concert

func (c Concerts) Filter(filter string) Concerts {
	args := strings.SplitN(filter, "=", 2)
	if len(args) != 2 {
		return nil
	}

	var result Concerts
	filterKey, filterValue := args[0], args[1]
	for _, concert := range c {
		if concert.filter(filterKey, filterValue) {
			result = append(result, concert)
		}
	}
	return result
}

func (c Concerts) MarshalTable() ([]byte, error) {
	var buf bytes.Buffer

	table := tablewriter.NewWriter(&buf)
	table.SetColWidth(minTitleWidth)
	table.SetHeader(c.header())
	table.AppendBulk(c.body())
	table.Render()

	return buf.Bytes(), nil
}

func (c Concerts) header() []string {
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

func (c Concerts) body() [][]string {
	rows := make([][]string, len(c))
	for i, concert := range c {
		rows[i] = []string{
			concert.GetDate(),
			concert.GetOpenTime(),
			concert.GetStartTime(),
			concert.Title,
			concert.Location,
			strings.Replace(concert.Price.Price(), PriceSeperator, "\n", -1),
			// concert.Note,
		}
	}
	return rows
}
