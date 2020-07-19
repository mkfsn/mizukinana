package table

import (
	"bytes"
	"errors"
	"reflect"

	"github.com/olekukonko/tablewriter"
)

var (
	marshalerType = reflect.TypeOf((*Marshaler)(nil)).Elem()
)

type Marshaler interface {
	MarshalTableHeader() []string
	MarshalTableBody() [][]string
}

type MarshalerWithColWidth interface {
	MarshalTableColWidth() int
}

func Marshal(v interface{}) ([]byte, error) {
	value := reflect.ValueOf(v)
	typ := value.Type()
	if !typ.Implements(marshalerType) {
		return nil, errors.New("unsupported type: " + typ.String())
	}
	m, ok := value.Interface().(Marshaler)
	if !ok {
		return []byte(""), nil
	}

	var buf bytes.Buffer
	table := tablewriter.NewWriter(&buf)
	if m, ok := value.Interface().(MarshalerWithColWidth); ok {
		table.SetColWidth(m.MarshalTableColWidth())
	}
	table.SetHeader(m.MarshalTableHeader())
	table.AppendBulk(m.MarshalTableBody())
	table.Render()
	return buf.Bytes(), nil
}
