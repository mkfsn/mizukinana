package concerts

var NanaMizukiLiveGrace2013 = Concerts{
	NewConcert("NANA MIZUKI LIVE GRACE 2013 -OPUSⅡ-", date(2013, 1, 19), "さいたまスーパーアリーナ", clock(16, 30), clock(18, 30), GroupPrice{"前売": NewJPY(7700, true)}, ""),
	NewConcert("NANA MIZUKI LIVE GRACE 2013 -OPUSⅡ-", date(2013, 1, 20), "さいたまスーパーアリーナ", clock(14, 30), clock(16, 30), GroupPrice{"前売": NewJPY(7700, true)}, ""),
}

var NanaMizukiLiveCircus2013 = Concerts{
	NewConcert("NANA MIZUKI LIVE CIRCUS 2013", date(2013, 7, 7), "【愛媛】 愛媛県武道館", clock(16, 00), clock(17, 00), GroupPrice{"前売": NewJPY(6500, true), "当日": NewJPY(6800, true)}, ""),
	NewConcert("NANA MIZUKI LIVE CIRCUS 2013", date(2013, 7, 14), "【大阪】大阪城ホール", clock(16, 00), clock(17, 00), GroupPrice{"前売": NewJPY(6500, true), "当日": NewJPY(6800, true)}, ""),
	NewConcert("NANA MIZUKI LIVE CIRCUS 2013", date(2013, 7, 15), "【大阪】大阪城ホール", clock(16, 00), clock(17, 00), GroupPrice{"前売": NewJPY(6500, true), "当日": NewJPY(6800, true)}, ""),
	NewConcert("NANA MIZUKI LIVE CIRCUS 2013", date(2013, 7, 20), "【岩手】 北上市文化交流センター さくらホール", clock(17, 00), clock(18, 00), GroupPrice{"前売": NewJPY(6500, true), "当日": NewJPY(6800, true)}, ""),
	NewConcert("NANA MIZUKI LIVE CIRCUS 2013", date(2013, 7, 21), "【福島】 須賀川市文化センター", clock(16, 00), clock(17, 00), GroupPrice{"前売": NewJPY(6500, true), "当日": NewJPY(6800, true)}, ""),
	NewConcert("NANA MIZUKI LIVE CIRCUS 2013", date(2013, 8, 3), "【崎玉】 西武ドーム", clock(17, 30), clock(18, 30), GroupPrice{"前売": NewJPY(6500, true), "当日": NewJPY(6800, true)}, ""),
	NewConcert("NANA MIZUKI LIVE CIRCUS 2013", date(2013, 8, 4), "【崎玉】 西武ドーム", clock(17, 00), clock(18, 00), GroupPrice{"前売": NewJPY(6500, true), "当日": NewJPY(6800, true)}, ""),
	NewConcert("NANA MIZUKI LIVE CIRCUS 2013", date(2013, 8, 10), "【宮崎】 宮崎市民文化ホール", clock(17, 00), clock(18, 00), GroupPrice{"前売": NewJPY(6500, true), "当日": NewJPY(6800, true)}, ""),
	NewConcert("NANA MIZUKI LIVE CIRCUS 2013", date(2013, 8, 11), "【鹿児島】 鹿児島市民文化ホール 第1ホール", clock(15, 00), clock(16, 00), GroupPrice{"前売": NewJPY(6500, true), "当日": NewJPY(6800, true)}, ""),
	NewConcert("NANA MIZUKI LIVE CIRCUS 2013", date(2013, 8, 17), "【愛知】 日本ガイシホール", clock(17, 30), clock(18, 30), GroupPrice{"前売": NewJPY(6500, true), "当日": NewJPY(6800, true)}, ""),
	NewConcert("NANA MIZUKI LIVE CIRCUS 2013", date(2013, 8, 18), "【愛知】 日本ガイシホール", clock(17, 00), clock(18, 00), GroupPrice{"前売": NewJPY(6500, true), "当日": NewJPY(6800, true)}, ""),

	NewConcert("NANA MIZUKI LIVE CIRCUS 2013+", date(2013, 11, 23), "【台湾】 Legacy Taipei", clock(19, 45), clock(20, 00), &NilPrice{}, ""),
	NewConcert("NANA MIZUKI LIVE CIRCUS 2013+", date(2013, 11, 24), "【台湾】 Legacy Taipei", clock(16, 45), clock(17, 00), &NilPrice{}, ""),
}
