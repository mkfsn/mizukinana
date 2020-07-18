package nanaparty_test

import (
	"context"
	"log"
	"testing"

	"github.com/mkfsn/mizukinana/nanaparty"
)

func TestNanaPartyNews(t *testing.T) {
	news, err := nanaparty.New().News(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Println(len(news))

	table := make(map[string]int)
	for _, data := range news {
		switch {
		case data.Id == "":
			log.Printf("empty id: %v\n", data)
		case data.Group == "":
			log.Printf("empty group: %v\n", data)
		case data.Title == "":
			log.Printf("empty title: %v\n", data)
		case data.Date == "":
			log.Printf("empty date: %v\n", data)
		case data.Category == "":
			log.Printf("empty category: %v\n", data)
		case data.Detail == "":
			log.Printf("empty detail: %v\n", data)
		}
		table[data.Group]++
	}
	log.Printf("table: %+v\n", table)
}

func TestNanaPartyBlogInfo(t *testing.T) {
	blogs, err := mizukinana.NanaParty().Blog(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, blog := range blogs {
		log.Printf("Blog Info: %+v\n", blog.Info())
	}
}

func TestNanaPartyBiographyProfile(t *testing.T) {
	biography, err := mizukinana.NanaParty().Biography(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	profile := biography.Profile()
	log.Printf("Profile: %+v\n", profile)
}

func TestNanaPartyBiographyLive(t *testing.T) {
	biography, err := mizukinana.NanaParty().Biography(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	lives := biography.Live()
	for _, live := range lives {
		log.Printf("Live: %+v\n", live)
	}
}

func TestNanaPartyBiographyVoice(t *testing.T) {
	biography, err := mizukinana.NanaParty().Biography(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	voices := biography.Voice()
	for _, voice := range voices {
		log.Printf("Voice Category: %s\n", voice.Category)
		for _, info := range voice.InfoList {
			log.Printf("Voice Info: %+v\n", info)
		}
	}
}

func TestNanaPartyBiographySpecial(t *testing.T) {
	biography, err := mizukinana.NanaParty().Biography(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	specials := biography.Special()
	for _, special := range specials {
		log.Printf("Special: %+v\n", special)
	}
}

func TestNanaPartyBiographyOther(t *testing.T) {
	biography, err := mizukinana.NanaParty().Biography(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	others := biography.Other()
	for _, other := range others {
		log.Printf("Other Category: %s\n", other.Category)
		for _, info := range other.InfoList {
			log.Printf("Other Info: %+v\n", info)
		}
	}
}

func TestNanaPartyTopMain(t *testing.T) {
	top, err := mizukinana.NanaParty().Top(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range top.Main() {
		log.Printf("Top Main Item: %+v\n", item)
	}
}

func TestNanaPartyTopPickup(t *testing.T) {
	top, err := mizukinana.NanaParty().Top(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range top.Pickup() {
		log.Printf("Top Main Item: %+v\n", item)
	}
}

func TestNanaPartyTopMovie(t *testing.T) {
	top, err := mizukinana.NanaParty().Top(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("Top Movie: %+v\n", top.Movie())
}

func TestNanaPartyTopBanner(t *testing.T) {
	top, err := mizukinana.NanaParty().Top(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range top.Banner() {
		log.Printf("Top Banner Item: %+v\n", item)
	}
}

func TestNanaPartyTopTopic(t *testing.T) {
	top, err := mizukinana.NanaParty().Top(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range top.Topic() {
		log.Printf("Top Topic Item: %+v\n", item)
	}
}

func TestNanaPartyScheduleInfo(t *testing.T) {
	schedule, err := mizukinana.NanaParty().Schedule(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, info := range schedule.Info() {
		log.Printf("Schedule Info: %+v\n", info)
	}
}

func TestNanaPartyDiscographyInfo(t *testing.T) {
	discography, err := mizukinana.NanaParty().Discography(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, recording := range discography.Discographies() {
		log.Printf("Recording Info: %+v\n", recording.Info())
	}
}

func TestNanaPartyDiscographyDetail(t *testing.T) {
	discographyCollection, err := mizukinana.NanaParty().Discography(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	found := make(map[string]bool)

	for _, discography := range discographyCollection.Discographies() {
		info := discography.Info()
		if _, ok := found[info.Form]; ok {
			continue
		}
		found[info.Form] = true

		log.Printf("Discography Info: %+v\n", info)
		detail, err := discography.Detail(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		for _, edition := range detail.Editions {
			log.Printf("Discography Detail Edition: %+v\n", edition)
		}
	}
}
