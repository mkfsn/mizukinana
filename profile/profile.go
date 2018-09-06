package profile

import (
	"bytes"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Profile struct {
	Name      string    `json:"name"      yaml:"name"`
	Birthday  time.Time `json:"birthday"  yaml:"birthday"`
	Hometown  string    `json:"hometown"  yaml:"hometown"`
	BloodType string    `json:"bloodtype" yaml:"bloodtype"`
	age       int
}

func (p Profile) MarshalTable() ([]byte, error) {
	var buf bytes.Buffer

	table := tablewriter.NewWriter(&buf)
	// table.SetHeader(p.header())
	table.AppendBulk(p.table())
	table.Render()

	return buf.Bytes(), nil
}

func (p Profile) table() [][]string {
	return [][]string{
		[]string{"Name", p.Name},
		[]string{"Birthday", p.Birthday.Format("2006-01-02")},
		[]string{"Hometown", p.Hometown},
		[]string{"BloodType", p.BloodType},
	}
}
