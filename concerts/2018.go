package concerts

var NanaMizukiLiveIsland2018 = Concerts{
	NewConcert("NANA MIZUKI LIVE ISLAND 2018", date(2018, 6, 23), "【宮城】セキスイハイムスーパーアリーナ（グランディ・21）", clock(15, 00), clock(16, 00), GroupPrice{"指定席": NewJPY(7900, true)}, ""),
	NewConcert("NANA MIZUKI LIVE ISLAND 2018", date(2018, 7, 7), "【大阪】大阪城ホール", clock(17, 00), clock(18, 00), GroupPrice{"指定席": NewJPY(7900, true), "立見": NewJPY(7400, true)}, ""),
	NewConcert("NANA MIZUKI LIVE ISLAND 2018", date(2018, 7, 8), "【大阪】大阪城ホール", clock(15, 00), clock(16, 00), GroupPrice{"指定席": NewJPY(7900, true), "立見": NewJPY(7400, true)}, ""),
	NewConcert("NANA MIZUKI LIVE ISLAND 2018", date(2018, 7, 15), "【岐阜】長良川国際会議場 (メインホール)", clock(17, 00), clock(18, 00), GroupPrice{"指定席": NewJPY(7900, true)}, ""),
	NewConcert("NANA MIZUKI LIVE ISLAND 2018", date(2018, 7, 21), "【広島】広島サンプラザ ホール", clock(15, 00), clock(16, 00), GroupPrice{"指定席": NewJPY(7900, true), "立見": NewJPY(7400, true)}, ""),
	NewConcert("NANA MIZUKI LIVE ISLAND 2018", date(2018, 7, 28), "【高知】高知県立県民体育館", clock(15, 00), clock(16, 00), GroupPrice{"指定席": NewJPY(7900, true)}, ""),
	NewConcert("NANA MIZUKI LIVE ISLAND 2018", date(2018, 8, 4), "【熊本】熊本県野外劇場アスペクタ", clock(16, 00), clock(17, 30), GroupPrice{"立ち位置指定": NewJPY(7900, true), "ブロック自由席": NewJPY(7400, true)}, ""),
	NewConcert("NANA MIZUKI LIVE ISLAND 2018", date(2018, 8, 11), "【愛知】日本ガイシホール", clock(17, 00), clock(18, 00), GroupPrice{"指定席": NewJPY(7900, true), "立見": NewJPY(7400, true)}, ""),
	NewConcert("NANA MIZUKI LIVE ISLAND 2018", date(2018, 8, 12), "【愛知】日本ガイシホール", clock(15, 00), clock(16, 00), GroupPrice{"指定席": NewJPY(7900, true), "立見": NewJPY(7400, true)}, ""),
	NewConcert("NANA MIZUKI LIVE ISLAND 2018", date(2018, 8, 18), "【和歌山】和歌山県民文化会館 (大ホール)", clock(15, 00), clock(16, 00), GroupPrice{"指定席": NewJPY(7900, true), "立見": NewJPY(7400, true)}, ""),
	NewConcert("NANA MIZUKI LIVE ISLAND 2018", date(2018, 8, 19), "【奈良】なら100年会館", clock(17, 00), clock(18, 00), GroupPrice{"指定席": NewJPY(7900, true), "立見": NewJPY(7400, true)}, ""),
	NewConcert("NANA MIZUKI LIVE ISLAND 2018", date(2018, 9, 1), "【埼玉】メットライフドーム（西武ドーム）", clock(15, 00), clock(17, 00), GroupPrice{"指定席": NewJPY(7900, true)}, ""),

	NewConcert("NANA MIZUKI LIVE ISLAND 2018+", date(2018, 9, 29), "台湾・国立体育大学体育館（林口体育館）", clock(15, 30), clock(17, 00), Prices{NewTWD(3800), NewTWD(3400), NewTWD(2800)}, ""),
	NewConcert("NANA MIZUKI LIVE ISLAND 2018+", date(2018, 10, 13), "上海・ヒマラヤセンター大観舞台", clock(00, 00), clock(00, 00), &NilPrice{}, ""),
}
