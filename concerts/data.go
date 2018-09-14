package concerts

var PersonalConcerts = concat(
	MizukiNana20thBirthdayAnniversaryLive,
	MizukiNana20thBirthdayAnniversaryLive,
	MizukiNana21thAnniversaryConcertHappy,
	MizukiNanaXmaxLiveSupersonicGirl,
	NanaMizukiLiveAttraction2002,
	NanaMizukiLiveSensation2003,
	NanaMizukiLiveSkipperCountdown20032004,
	NanaMizukiLiveSpark2004Summer,
	NanaMizukiLiveRainbow20042005,
	NanaMizukiLiveRocket2005Summer,
	NanaMizukiLivedom2006Birth,
	NanaMizukiLiveUniverse2006Summer,
	NanaMizukiLiveMuseum2007,
	NanaSummerFesta2007,
	NanaMizukiLiveFormula2007,
	NanaMizukiLiveFighter2008,
	NanaMizukiLiveFever2009,
	NanaMizukiLiveDiamond2009,
	NanaMizukiLiveAcademy2010,
	NanaMizukiLiveGames2010,
	NanaMizukiLiveGrace2011Orchestra,
	NanaMizukiLiveJourney2011,
	NanaMizukiLiveCastle2011,
	NanaMizukiLiveUnion2012,
	MizukiNanaHeianJinguHounouKouenSougetsuNoUtage,
	NanaMizukiLiveGrace2013,
	NanaMizukiLiveCircus2013,
	NanaWinterFesta2014,
	NanaMizukiLiveFlight,
	NanaMizukiLiveTheater2015,
	NanaMizukiLiveAdventure2015,
	NanaMizukiLiveGalaxy2016,
	NanaMizukiLivePark2016,
	NanaMizukiLiveZipangu2017,
	NanaMizukiLiveIsland2018,
)

func concat(concerts ...Concerts) Concerts {
	var result Concerts
	for _, c := range concerts {
		result = append(result, c...)
	}
	return result
}
