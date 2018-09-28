// +build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"text/template"
	"time"
)

const yamlfile = `// Code generated by go generate; DO NOT EDIT.
package discography
import (
	"gopkg.in/yaml.v2"
)

var {{ .VarName }} {{ .VarType }}

func init() {
	err := yaml.Unmarshal({{ .TempVar }}, &{{ .VarName }})
	if err != nil {
		panic(err)
	}
}

var {{ .TempVar }} = []byte{
	{{ range .RawData -}}
		{{ range . }} {{ printf "0x%02x" . }}, {{ end }}
	{{ end }}
}
`

type File struct {
	VarName string
	VarType string
	TempVar string

	inputFile string
}

func (f File) RawData() [][]byte {
	b, err := ioutil.ReadFile(f.inputFile)
	if err != nil {
		return nil
	}

	chunks := make([][]byte, 0)
	chunkSize := 8
	nChunks := len(b) / chunkSize

	for i := 0; i < nChunks; i++ {
		m := i * chunkSize
		chunks = append(chunks, b[m:m+chunkSize])
	}

	if r := len(b) % chunkSize; r > 0 {
		m := nChunks * chunkSize
		chunks = append(chunks, b[m:m+r])
	}

	return chunks
}

func (f File) Output() error {
	t := template.Must(template.New("gofile").Parse(yamlfile))

	var buf bytes.Buffer
	err := t.Execute(&buf, f)
	if err != nil {
		return err
	}

	b, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(f.inputFile+".go", b, os.FileMode(0644)); err != nil {
		return err
	}

	return nil
}

func main() {
	f := File{TempVar: fmt.Sprintf("tmp%d", time.Now().UnixNano())}
	flag.StringVar(&f.VarName, "varname", "", "variable name")
	flag.StringVar(&f.VarType, "vartype", "", "variable type")
	flag.StringVar(&f.inputFile, "input", "", "input yaml file path")
	flag.Parse()

	err := f.Output()
	if err != nil {
		fmt.Println(err)
	}
}