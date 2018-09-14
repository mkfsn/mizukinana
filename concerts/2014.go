package concerts

var NanaWinterFesta2014 = Concerts{
	NewConcert("NANA WINTER FESTA 2014", date(2014, 1, 18), "有明コロシアム", clock(16, 00), clock(17, 00), NewJPY(5800, false), ""),
	NewConcert("NANA WINTER FESTA 2014", date(2014, 1, 19), "有明コロシアム", clock(15, 00), clock(16, 00), NewJPY(5800, false), ""),
}

var NanaMizukiLiveFlight = Concerts{
	NewConcert("NANA MIZUKI LIVE FLIGHT 2014", date(2014, 6, 1), "【山梨】富士急ハイランド コニファーフォレスト", clock(15, 00), clock(16, 30), GroupPrice{"前売": NewJPY(7500, true)}, ""),
	NewConcert("NANA MIZUKI LIVE FLIGHT 2014", date(2014, 6, 8), "【長野】長野ビッグハット", clock(17, 00), clock(18, 00), GroupPrice{"前売": NewJPY(7500, true)}, ""),
	NewConcert("NANA MIZUKI LIVE FLIGHT 2014", date(2014, 6, 14), "【三重】三重県営サンアリーナ", clock(16, 00), clock(17, 00), GroupPrice{"前売": NewJPY(7500, true)}, ""),
	NewConcert("NANA MIZUKI LIVE FLIGHT 2014", date(2014, 6, 15), "【三重】三重県営サンアリーナ", clock(15, 00), clock(16, 00), GroupPrice{"前売": NewJPY(7500, true)}, ""),
	NewConcert("NANA MIZUKI LIVE FLIGHT 2014", date(2014, 6, 21), "【福岡】北九州メディアドーム", clock(17, 00), clock(18, 00), GroupPrice{"前売": NewJPY(7500, true)}, ""),
	NewConcert("NANA MIZUKI LIVE FLIGHT 2014", date(2014, 6, 22), "【福岡】北九州メディアドーム", clock(15, 00), clock(16, 00), GroupPrice{"前売": NewJPY(7500, true)}, "水樹奈々さんが急性声帯炎のため中止"),
	NewConcert("NANA MIZUKI LIVE FLIGHT 2014", date(2014, 6, 28), "【山口】防府市公会堂", clock(15, 00), clock(16, 00), GroupPrice{"前売": NewJPY(7500, true)}, "水樹奈々さんが急性声帯炎のため中止"),
	NewConcert("NANA MIZUKI LIVE FLIGHT 2014", date(2014, 6, 29), "【鳥取】米子コンベンションセンター BiG SHiP", clock(17, 00), clock(18, 00), GroupPrice{"前売": NewJPY(7500, true)}, "水樹奈々さんが急性声帯炎のため中止"),
	NewConcert("NANA MIZUKI LIVE FLIGHT 2014", date(2014, 7, 5), "【山形】鹿児島市民文化ホール 第1ホール", clock(17, 00), clock(18, 00), GroupPrice{"前売": NewJPY(7500, true)}, ""),
	NewConcert("NANA MIZUKI LIVE FLIGHT 2014", date(2014, 7, 6), "【宮城】東京エレクトロンホール宮城", clock(16, 30), clock(17, 30), GroupPrice{"前売": NewJPY(7500, true)}, ""),
	NewConcert("NANA MIZUKI LIVE FLIGHT 2014", date(2014, 7, 11), "【大阪】大阪市中央体育館", clock(17, 30), clock(18, 30), GroupPrice{"前売": NewJPY(7500, true)}, ""),
	NewConcert("NANA MIZUKI LIVE FLIGHT 2014", date(2014, 7, 12), "【大阪】大阪市中央体育館", clock(15, 00), clock(16, 00), GroupPrice{"前売": NewJPY(7500, true)}, ""),
	NewConcert("NANA MIZUKI LIVE FLIGHT 2014", date(2014, 7, 20), "【香川】さぬき市野外音楽広場テアトロン", clock(16, 00), clock(17, 30), GroupPrice{"前売": NewJPY(7500, true)}, ""),
	NewConcert("NANA MIZUKI LIVE FLIGHT 2014", date(2014, 8, 3), "【神奈川】横浜スタジアム", clock(15, 00), clock(17, 00), GroupPrice{"前売": NewJPY(7500, true)}, ""),

	NewConcert("NANA MIZUKI LIVE FLIGHT 2014+", date(2014, 9, 27), "【新加坡】Resorts World Sentosa Official Website | Singapore Attractions", clock(17, 30), clock(18, 30), &NilPrice{}, ""),
	NewConcert("NANA MIZUKI LIVE FLIGHT 2014+", date(2014, 10, 4), "【台灣】臺北市立大學天母校區體育館", clock(18, 00), clock(19, 00), &NilPrice{}, ""),
	NewConcert("NANA MIZUKI LIVE FLIGHT 2014+", date(2014, 10, 5), "【台灣】臺北市立大學天母校區體育館", clock(17, 00), clock(18, 00), &NilPrice{}, ""),
}
