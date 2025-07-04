package main

import (
	mt "MediaUnlockTest/checks"
	"net/http"
	"sync"
	"time"
)

var (
	MUL  bool
	HK   bool
	TW   bool
	JP   bool
	KR   bool
	NA   bool
	SA   bool
	EU   bool
	AFR  bool
	SEA  bool
	OCEA bool
	AI   bool
	Conc uint64 = 0
	sem  chan struct{}
)

type TEST struct {
	Client  http.Client
	Results []*result
	Wg      *sync.WaitGroup
}

func NewTest() *TEST {
	t := &TEST{
		Client:  mt.AutoHttpClient,
		Results: make([]*result, 0),
		Wg:      &sync.WaitGroup{},
	}
	if Conc > 0 {
		sem = make(chan struct{}, Conc) // 初始化带缓冲的通道
	}
	return t
}

func (T *TEST) Check() bool {
	if MUL {
		T.Multination()
	}
	if HK {
		T.HongKong()
	}
	if TW {
		T.Taiwan()
	}
	if JP {
		T.Japan()
	}
	if KR {
		T.Korea()
	}
	if NA {
		T.NorthAmerica()
	}
	if SA {
		T.SouthAmerica()
	}
	if EU {
		T.Europe()
	}
	if AFR {
		T.Africa()
	}
	if SEA {
		T.SouthEastAsia()
	}
	if OCEA {
		T.Oceania()
	}
	if AI {
		T.AI()
	}

	ch := make(chan struct{})
	go func() {
		defer close(ch)
		T.Wg.Wait()
	}()
	select {
	case <-ch:
		return false
	case <-time.After(30 * time.Second):
		return true
	}
}

type result struct {
	Type  string
	Name  string
	Value mt.Result
}

func (T *TEST) excute(Name string, F func(client http.Client) mt.Result) {
	r := &result{Name: Name}
	T.Results = append(T.Results, r)
	T.Wg.Add(1)
	go func() {
		if Conc > 0 {
			sem <- struct{}{} // 获取一个通道资源
			defer func() {
				<-sem // 释放一个通道资源
				T.Wg.Done()
			}()
		} else {
			defer T.Wg.Done()
		}
		r.Value = F(T.Client)
	}() // 移动 T.Wg.Done() 到 defer 中
}

func (T *TEST) Multination() {
	T.excute("Dazn", mt.Dazn)
	T.excute("Disney+", mt.DisneyPlus)
	T.excute("Netflix", mt.NetflixRegion)
	T.excute("Netflix CDN", mt.NetflixCDN)
	T.excute("Youtube Premium", mt.YoutubeRegion)
	T.excute("Youtube CDN", mt.YoutubeCDN)
	T.excute("Amazon Prime Video", mt.PrimeVideo)
	T.excute("TVBAnywhere+", mt.TVBAnywhere)
	T.excute("iQiYi", mt.IQiYi)
	T.excute("Viu.com", mt.ViuCom)
	T.excute("Spotify", mt.Spotify)
	T.excute("Steam", mt.Steam)
	T.excute("Wikipedia", mt.WikipediaEditable)
	T.excute("Reddit", mt.Reddit)
	T.excute("TikTok", mt.TikTok)
	T.excute("Bing", mt.Bing)
	T.excute("Google Play Store", mt.GooglePlayStore)
	T.excute("Apple", mt.Apple)
}

func (T *TEST) HongKong() {
	T.excute("Now E", mt.NowE)
	T.excute("Viu.TV", mt.ViuTV)
	T.excute("MyTVSuper", mt.MyTvSuper)
	T.excute("Max", mt.Max)
	T.excute("Bilibili HK/MO", mt.BilibiliHKMO)
	T.excute("SonyLiv", mt.SonyLiv)
	T.excute("Bahamut Anime", mt.BahamutAnime)
	T.excute("Hoy TV", mt.HoyTV)
}

func (T *TEST) Taiwan() {
	T.excute("KKTV", mt.KKTV)
	T.excute("LiTV", mt.LiTV)
	T.excute("MyVideo", mt.MyVideo)
	T.excute("4GTV", mt.TW4GTV)
	T.excute("LineTV", mt.LineTV)
	T.excute("Hami Video", mt.HamiVideo)
	T.excute("CatchPlay+", mt.Catchplay)
	T.excute("Bahamut Anime", mt.BahamutAnime)
	T.excute("Max", mt.Max)
	T.excute("Bilibili TW", mt.BilibiliTW)
	T.excute("Ofiii", mt.Ofiii)
	T.excute("Friday Video", mt.FridayVideo)
}

func (T *TEST) Japan() {
	T.excute("DMM", mt.DMM)
	T.excute("DMM TV", mt.DMMTV)
	T.excute("Abema", mt.Abema)
	T.excute("Niconico", mt.Niconico)
	T.excute("music.jp", mt.MusicJP)
	T.excute("Telasa", mt.Telasa)
	T.excute("U-NEXT", mt.U_NEXT)
	T.excute("Hulu Japan", mt.HuluJP)
	T.excute("VideoMarket", mt.VideoMarket)
	T.excute("FOD(Fuji TV)", mt.FOD)
	T.excute("Radiko", mt.Radiko)
	T.excute("Karaoke@DAM", mt.Karaoke)
	T.excute("J:COM On Demand", mt.J_COM_ON_DEMAND)
	T.excute("Kancolle", mt.Kancolle)
	T.excute("Pretty Derby Japan", mt.PrettyDerbyJP)
	T.excute("Princess Connect Re:Dive Japan", mt.PCRJP)
	T.excute("Project Sekai: Colorful Stage", mt.PJSK)
	T.excute("Rakuten TV", mt.RakutenTV_JP)
	T.excute("Wowow", mt.Wowow)
	T.excute("Watcha", mt.Watcha)
	T.excute("TVer", mt.TVer)
	T.excute("Lemino", mt.Lemino)
	T.excute("D Anime Store", mt.DAnimeStore)
	T.excute("Mora", mt.Mora)
	T.excute("AnimeFesta", mt.AnimeFesta)
	T.excute("EroGameSpace", mt.EroGameSpace)
	T.excute("NHK+", mt.NHKPlus)
	T.excute("Rakuten Magazine", mt.RakutenMagazine)
	T.excute("MGStage", mt.MGStage)
}

func (T *TEST) Korea() {
	T.excute("Wavve", mt.Wavve)
	T.excute("Tving", mt.Tving)
	T.excute("Watcha", mt.Watcha)
	T.excute("Coupang Play", mt.CoupangPlay)
	T.excute("Spotv Now", mt.SpotvNow)
	T.excute("Naver TV", mt.NaverTV)
	T.excute("Afreeca", mt.Afreeca)
	T.excute("KBS", mt.KBS)
	T.excute("Panda TV", mt.PandaTV)
}

func (T *TEST) NorthAmerica() {
	T.excute("A&E TV", mt.AETV)
	T.excute("FOX", mt.Fox)
	T.excute("Hulu", mt.Hulu)
	T.excute("NFL+", mt.NFLPlus)
	T.excute("ESPN+", mt.ESPNPlus)
	T.excute("MGM+", mt.MGMPlus)
	T.excute("Starz", mt.Starz)
	T.excute("Philo", mt.Philo)
	T.excute("FXNOW", mt.FXNOW)
	T.excute("TLC GO", mt.TlcGo)
	T.excute("Max", mt.Max)
	T.excute("Shudder", mt.Shudder)
	T.excute("BritBox", mt.BritBox)
	T.excute("CW TV", mt.CW_TV)
	T.excute("NBA TV", mt.NBA_TV)
	T.excute("NBC TV", mt.NBC_TV)
	T.excute("Fubo TV", mt.FuboTV)
	T.excute("Tubi TV", mt.TubiTV)
	T.excute("Sling TV", mt.SlingTV)
	T.excute("Pluto TV", mt.PlutoTV)
	T.excute("Acorn TV", mt.AcornTV)
	T.excute("SHOWTIME", mt.SHOWTIME)
	T.excute("encoreTVB", mt.EncoreTVB)
	T.excute("Discovery+", mt.DiscoveryPlus)
	T.excute("Paramount+", mt.ParamountPlus)
	T.excute("Peacock TV", mt.PeacockTV)
	T.excute("Crunchyroll", mt.Crunchyroll)
	T.excute("DirecTV Stream", mt.DirectvStream)
	T.excute("SonyLiv", mt.SonyLiv)
	T.excute("Hotstar", mt.Hotstar)
	T.excute("AMC+", mt.AMCPlus)
	T.excute("MathsSpot Roblox", mt.MathsSpotRoblox)
	T.excute("KOCOWA+", mt.KOCOWA)
	T.excute("Viaplay", mt.Viaplay)
	T.excute("CBC Gem", mt.CBCGem)
	T.excute("Crave", mt.Crave)
}

func (T *TEST) SouthAmerica() {
	T.excute("DirecTV GO", mt.DirecTVGO)
}

func (T *TEST) Europe() {
	T.excute("Rakuten TV", mt.RakutenTV_EU)
	T.excute("Setanta Sports", mt.SetantaSports)
	T.excute("Sky Show Time", mt.SkyShowTime)
	T.excute("Max", mt.Max)
	T.excute("SonyLiv", mt.SonyLiv)
	T.excute("BBC iPlayer", mt.BBCiPlayer)
	T.excute("Channel 4", mt.Channel4)
	T.excute("Channel 5", mt.Channel5)
	T.excute("Sky Go", mt.SkyGo)
	T.excute("ITVX", mt.ITVX)
	T.excute("Rai Play", mt.RaiPlay)
	T.excute("Canal+", mt.CanalPlus)
	T.excute("ZDF", mt.ZDF)
	T.excute("Joyn", mt.Joyn)
	T.excute("Sky DE", mt.Sky_DE)
	T.excute("Molotov", mt.Molotov)
	T.excute("NPO Start Plus", mt.NPOStartPlus)
	T.excute("Video Land", mt.VideoLand)
	T.excute("NLZIET", mt.NLZIET)
	T.excute("Movistar Plus+", mt.MoviStarPlus)
	T.excute("Eurosport RO", mt.EurosportRO)
	T.excute("Sky CH", mt.Sky_CH)
	T.excute("Amediateka", mt.Amediateka)
	T.excute("Hotstar", mt.Hotstar)
	T.excute("MathsSpot Roblox", mt.MathsSpotRoblox)
	T.excute("KOCOWA+", mt.KOCOWA)
	T.excute("France TV", mt.FranceTV)
	T.excute("Viaplay", mt.Viaplay)
	T.excute("Discovery+ UK", mt.DiscoveryPlus_UK)
	T.excute("TNTSports", mt.TNTSports)
}

func (T *TEST) Africa() {
	T.excute("DSTV", mt.DSTV)
	T.excute("Showmax", mt.Showmax)
}

func (T *TEST) SouthEastAsia() {
	T.excute("Bilibili SEA Only", mt.BilibiliSEA)
	T.excute("SonyLiv", mt.SonyLiv)
	T.excute("MeWatch", mt.MeWatch)
	T.excute("Bilibili TH Only", mt.BilibiliTH)
	T.excute("AIS Play", mt.AISPlay)
	T.excute("TrueID", mt.TrueID)
	T.excute("Bilibili ID Only", mt.BilibiliID)
	T.excute("Bilibili VN Only", mt.BilibiliVN)
	T.excute("Hotstar", mt.Hotstar)
	T.excute("CatchPlay+", mt.Catchplay)
	T.excute("Max", mt.Max)
	T.excute("Clip TV", mt.ClipTV)
	T.excute("Galaxy Play", mt.GalaxyPlay)
	T.excute("K+", mt.KPlus)
	T.excute("Sooka", mt.Sooka)
	T.excute("Zee5", mt.Zee5)
	T.excute("Tata Play", mt.TataPlay)
	T.excute("NBA TV", mt.NBA_TV)
	T.excute("Jio Cinema", mt.JioCinema)
	T.excute("MX Player", mt.MXPlayer)
}

func (T *TEST) Oceania() {
	T.excute("NBA TV", mt.NBA_TV)
	T.excute("Acorn TV", mt.AcornTV)
	T.excute("BritBox", mt.BritBox)
	T.excute("Paramount+", mt.ParamountPlus)
	T.excute("SonyLiv", mt.SonyLiv)
	T.excute("Stan", mt.Stan)
	T.excute("Binge", mt.Binge)
	T.excute("Doc Play", mt.DocPlay)
	T.excute("7Plus", mt.SevenPlus)
	T.excute("Channel 9", mt.Channel9)
	T.excute("10 Play", mt.Channel10)
	T.excute("ABC iView", mt.ABCiView)
	T.excute("Optus Sports", mt.OptusSports)
	T.excute("SBS on Demand", mt.SBSonDemand)
	T.excute("Kayo Sports", mt.KayoSports)
	T.excute("Neon TV", mt.NeonTV)
	T.excute("Three Now", mt.ThreeNow)
	T.excute("Maori TV", mt.MaoriTV)
	T.excute("Sky Go NZ", mt.SkyGo_NZ)
	T.excute("AMC+", mt.AMCPlus)
	T.excute("KOCOWA+", mt.KOCOWA)
}

func (T *TEST) AI() {
	T.excute("ChatGPT", mt.ChatGPT)
	T.excute("Claude", mt.Claude)
	T.excute("Copilot", mt.Copilot)
	T.excute("Google Gemini", mt.Gemini)
	T.excute("Meta AI", mt.MetaAI)
	T.excute("Sora", mt.Sora)
}
