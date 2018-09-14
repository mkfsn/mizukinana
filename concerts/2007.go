package concerts

var NanaMizukiLiveMuseum2007 = Concerts{
	NewConcert("NANA MIZUKI LIVE MUSEUM 2007", date(2007, 2, 12), "横浜アリーナ", clock(15, 00), clock(16, 00), GroupPrice{"前売": NewJPY(6000, true), "当日": NewJPY(6300, true)}, ""),
}

var NanaSummerFesta2007 = Concerts{
	NewConcert("NANA SUMMER FESTA 2007", date(2007, 7, 22), "Zepp Tokyo（昼の部）", clock(11, 00), clock(12, 00), GroupPrice{"前売": NewJPY(4500, true), "当日": NewJPY(5000, true)}, ""),
	NewConcert("NANA SUMMER FESTA 2007", date(2007, 7, 22), "Zepp Tokyo（夜の部）", clock(16, 00), clock(17, 00), GroupPrice{"前売": NewJPY(4500, true), "当日": NewJPY(5000, true)}, ""),
}

var NanaMizukiLiveFormula2007 = Concerts{
	NewConcert("NANA MIZUKI LIVE FORMULA 2007-2008", date(2007, 12, 2), "名古屋国際会議場　センチュリーホール", clock(17, 30), clock(18, 30), GroupPrice{"前売": NewJPY(5900, true), "当日": NewJPY(6300, true)}, ""),
	NewConcert("NANA MIZUKI LIVE FORMULA 2007-2008", date(2007, 12, 15), "広島アステールプラザ　大ホール", clock(17, 30), clock(18, 30), GroupPrice{"前売": NewJPY(5900, true), "当日": NewJPY(6300, true)}, ""),
	NewConcert("NANA MIZUKI LIVE FORMULA 2007-2008", date(2007, 12, 16), "松山市民会館　大ホール", clock(17, 30), clock(18, 30), GroupPrice{"前売": NewJPY(5900, true), "当日": NewJPY(6300, true)}, ""),
	NewConcert("NANA MIZUKI LIVE FORMULA 2007-2008", date(2007, 12, 22), "東京厚生年金会館", clock(17, 30), clock(18, 30), GroupPrice{"前売": NewJPY(5900, true), "当日": NewJPY(6300, true)}, ""),
	NewConcert("NANA MIZUKI LIVE FORMULA 2007-2008", date(2007, 12, 24), "仙台サンプラザホール", clock(17, 00), clock(18, 00), GroupPrice{"前売": NewJPY(5900, true), "当日": NewJPY(6300, true)}, ""),
	NewConcert("NANA MIZUKI LIVE FORMULA 2007-2008", date(2007, 12, 31), "グランキューブ大阪　メインホール（カウントダウンライブ）", clock(21, 30), clock(22, 30), GroupPrice{"前売": NewJPY(5900, true), "当日": NewJPY(6300, true)}, ""),
	NewConcert("NANA MIZUKI LIVE FORMULA 2007-2008", date(2008, 1, 3), "さいたまスーパーアリーナ", clock(17, 00), clock(18, 00), GroupPrice{"前売": NewJPY(5900, true), "当日": NewJPY(6300, true)}, ""),
}
