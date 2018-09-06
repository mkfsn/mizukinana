package profile

import (
	"time"
)

var MizukiNana = Profile{
	Name:      "水樹奈々",
	Birthday:  time.Date(1980, 1, 21, 0, 0, 0, 0, time.FixedZone("Asia/Japan", 9*60*60)),
	Hometown:  "愛媛県 新居浜",
	BloodType: "O 型",
	age:       17,
}
