package mizukinana_test

import (
	"github.com/mkfsn/mizukinana"
	"os"
	"testing"
)

func TestFetchOfficialWebsite(t *testing.T) {
	result, err := mizukinana.FetchOfficialWebsite()
	if err != nil {
		t.Errorf("Error occurs")
	}
	if len(result) == 0 {
		t.Errorf("Nothing fetched")
	}
}

func TestFetchFanclub(t *testing.T) {
	result, err := mizukinana.FetchFanclub()
	if err != nil {
		t.Errorf("Error occurs")
	}
	if len(result) == 0 {
		t.Errorf("Nothing fetched")
	}
}

func TestGetAge(t *testing.T) {
	age := 36
	result := mizukinana.GetAge()
	if result != age {
		t.Errorf("Except\t%+v\nGet\t%+v\n", age, result)
	}
}

func TestJSONer(t *testing.T) {
	os.Stdout.Write([]byte("[OfficialWebsite]\n"))
	contents, _ := mizukinana.FetchOfficialWebsite()
	b := contents.JSON()
	os.Stdout.Write(b)
	os.Stdout.Write([]byte("\n\n"))

	os.Stdout.Write([]byte("[Fanclub]\n"))
	contents, _ = mizukinana.FetchFanclub()
	b = contents.JSON()
	os.Stdout.Write(b)
	os.Stdout.Write([]byte("\n\n"))
}
