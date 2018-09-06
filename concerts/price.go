package concerts

import (
	"fmt"
	"strings"
)

const (
	priceSeperator = "／"
)

type Price interface {
	Price() string
}

type NilPrice struct{}

func (n NilPrice) Price() string {
	return ""
}

func (n NilPrice) MarshalJSON() ([]byte, error) {
	return []byte(`""`), nil
}

func (n NilPrice) MarshalYAML() (interface{}, error) {
	return n.Price(), nil
}

type Prices []Price

func (p Prices) Price() string {
	var result []string
	for _, price := range p {
		result = append(result, price.Price())
	}
	return strings.Join(result, priceSeperator)
}

type JPY struct {
	value int
	tax   bool
}

func NewJPY(value int, withTax bool) Price {
	return &JPY{
		value: value,
		tax:   withTax,
	}
}

func (j JPY) Price() string {
	format := "%d円"
	if !j.tax {
		format += "＋税"
	}
	return fmt.Sprintf(format, j.value)
}

func (j JPY) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", j.Price())), nil
}

func (j JPY) MarshalYAML() (interface{}, error) {
	return j.Price(), nil
}

type TWD struct {
	value int
}

func NewTWD(value int) Price {
	return &TWD{
		value: value,
	}
}

func (t TWD) Price() string {
	return fmt.Sprintf("NTD %d", t.value)
}

func (t TWD) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", t.Price())), nil
}

func (t TWD) MarshalYAML() (interface{}, error) {
	return t.Price(), nil
}

type GroupPrice map[string]Price

func NewGroupPrice(m map[string]Price) Price {
	v := GroupPrice(m)
	return &v
}

func (g GroupPrice) Price() string {
	var result []string
	for k, v := range g {
		if k != "" {
			result = append(result, k+" "+v.Price())
		} else {
			result = append(result, v.Price())
		}
	}
	return strings.Join(result, "／")
}
