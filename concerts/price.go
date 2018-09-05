package concerts

import (
	"fmt"
	"strings"
)

const (
	PriceSeperator = "／"
)

type price interface {
	Price() string
}

type nilPrice struct{}

func (n nilPrice) Price() string {
	return ""
}

func (n nilPrice) MarshalJSON() ([]byte, error) {
	return []byte(`""`), nil
}

func (n nilPrice) MarshalYAML() (interface{}, error) {
	return n.Price(), nil
}

type prices []price

func (p prices) Price() string {
	var result []string
	for _, price := range p {
		result = append(result, price.Price())
	}
	return strings.Join(result, PriceSeperator)
}

type jpy struct {
	value int
	tax   bool
}

func newJPY(value int, withTax bool) price {
	return &jpy{
		value: value,
		tax:   withTax,
	}
}

func (j jpy) Price() string {
	format := "%d円"
	if !j.tax {
		format += "＋税"
	}
	return fmt.Sprintf(format, j.value)
}

func (j jpy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", j.Price())), nil
}

func (j jpy) MarshalYAML() (interface{}, error) {
	return j.Price(), nil
}

type twd struct {
	value int
}

func newTWD(value int) price {
	return &twd{
		value: value,
	}
}

func (t twd) Price() string {
	return fmt.Sprintf("NTD %d", t.value)
}

func (t twd) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", t.Price())), nil
}

func (t twd) MarshalYAML() (interface{}, error) {
	return t.Price(), nil
}

type groupPrice map[string]price

func newGroupPrice(m map[string]price) price {
	v := groupPrice(m)
	return &v
}

func (g groupPrice) Price() string {
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
