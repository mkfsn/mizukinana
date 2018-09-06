package concerts_test

import (
	"testing"
	"time"

	"github.com/mkfsn/mizukinana/concerts"
)

func TestNewLive(t *testing.T) {
	open := time.Hour*15 + time.Minute*0
	start := time.Hour*17 + time.Minute*0
	title := "NANA MIZUKI LIVE PARK 2016"
	location := "阪神甲子園球場"
	date := time.Date(2016, 9, 22, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))
	price := concerts.NewJPY(7777, true)
	note := ""

	concert := concerts.NewConcert(title, date, location, open, start, price, note)

	if concert.Title != title {
		t.Error("title does not match")
	}

	if concert.Open != date.Add(open) {
		t.Error("open does not match")
	}

	if concert.Start != date.Add(start) {
		t.Error("start does not match")
	}

	if concert.Location != location {
		t.Error("location does not match")
	}

	if concert.Date != date {
		t.Error("date does not match")
	}

	if concert.Price != price {
		t.Error("price does not match")
	}

	if concert.Note != note {
		t.Error("note does not match")
	}
}
