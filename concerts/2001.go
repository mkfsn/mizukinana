package concerts

var MizukiNana21thAnniversaryConcertHappy = Concerts{
	NewConcert(`水樹奈々 21th ANNIVERSARY CONCERT "HAPPY"`, date(2001, 1, 21), "ヤクルトホール", clock(17, 30), clock(18, 00), GroupPrice{"前売": NewJPY(3800, true)}, ""),
}

var MizukiNanaXmaxLiveSupersonicGirl = Concerts{
	NewConcert(`水樹奈々 X'mas LIVE "supersonic girl"`, date(2001, 12, 23), "原宿アストロホール（昼公演）", clock(14, 00), clock(15, 00), GroupPrice{"前売": NewJPY(3990, true)}, ""),
	NewConcert(`水樹奈々 X'mas LIVE "supersonic girl"`, date(2001, 12, 23), "原宿アストロホール（夜公演）", clock(17, 30), clock(18, 30), GroupPrice{"前売": NewJPY(3990, true)}, ""),
}
