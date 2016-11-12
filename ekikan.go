package main

import (
	"math"
	"fmt"
)

var GlobalEkikanList = &GlobalEkikan{ekikanList: Ekikan_List}

// Return distance between two directly connected stations
// Ex: KyoriWoHyoji("myogadani", "shinotsuka") -> "茗荷谷駅と新大塚駅までは1.2kmです"
// Ex: KyoriWoHyoji("myogadani", "Not directly connected") -> "茗荷谷駅と新大塚駅はつながっていません”
// Ex: KyoriWoHyoji("myogadani", "I dont exist") -> "I dont existという駅は存在しません”
func KyoriWoHyoji(station1, station2 string) string {
	station1_jp := GlobalEkimeiList.RomajiToKanji(station1)
	station2_jp := GlobalEkimeiList.RomajiToKanji(station2)

	if station1_jp == ""  {
		return fmt.Sprintf("%vという駅は存在しません", station1)
	} else if station2_jp == "" {
		return fmt.Sprintf("%vという駅は存在しません", station2)
	}

	kyori := GlobalEkikanList.GetEkikanKyori(station1_jp, station2_jp)
	if kyori != math.Inf(+1) {
		return fmt.Sprintf("%v駅と%v駅までは%vkmです", station1_jp, station2_jp, kyori)
	}
	return fmt.Sprintf("%v駅と%v駅はつながっていません", station1_jp, station2_jp)
}

type Ekikan struct {
	kiten  string
	shuten string
	keiyu  string
	kyori  float64
	jikan  int // in Minute
}

// Get Distance between two stations with Kanji Input
// Ex: &g{list: Ekikan_List}.GetEkikanKyori("茗荷谷", "新大塚") -> 1.2
// Ex: &g{list: Ekikan_List}.GetEkikanKyori("新大塚", "茗荷谷") -> 1.2
// Ex: &g{list: Ekikan_List}.GetEkikanKyori("茗荷谷", "**Non Direct Connection**") -> infinity
func (g *GlobalEkikan) GetEkikanKyori(station1, station2 string) float64 {
	for _, ekikan := range g.ekikanList {
		if (ekikan.kiten == station1  && ekikan.shuten == station2) ||
			(ekikan.kiten == station2 && ekikan.shuten == station1) {
			return ekikan.kyori
		}
	}

	return math.Inf(+1)

	// Recursive Method on Below
	//if cap(g.ekikanList) == 0 {
	//	return math.Inf(+1)
	//}
	//
	//ekikan := g.ekikanList[0]
	//
	//if ekikan.kiten == station1 && ekikan.shuten == station2 {
	//	return ekikan.kyori
	//} else if ekikan.kiten == station2 && ekikan.shuten == station1 {
	//	return ekikan.kyori
	//} else {
	//	g.ekikanList = g.ekikanList[1:]
	//	return g.GetEkikanKyori(station1, station2)
	//}
}

type GlobalEkikan struct {
	ekikanList []Ekikan
}

var Ekikan_List = [] Ekikan{
	{kiten:"代々木上原", shuten:"代々木公園", keiyu:"千代田線", kyori:1.0, jikan:2},
	{kiten:"代々木公園", shuten:"明治神宮前", keiyu:"千代田線", kyori:1.2, jikan:2},
	{kiten:"明治神宮前", shuten:"表参道", keiyu:"千代田線", kyori:0.9, jikan:2},
	{kiten:"表参道", shuten:"乃木坂", keiyu:"千代田線", kyori:1.4, jikan:3},
	{kiten:"乃木坂", shuten:"赤坂", keiyu:"千代田線", kyori:1.1, jikan:2},
	{kiten:"赤坂", shuten:"国会議事堂前", keiyu:"千代田線", kyori:0.8, jikan:1},
	{kiten:"国会議事堂前", shuten:"霞ヶ関", keiyu:"千代田線", kyori:0.7, jikan:1},
	{kiten:"霞ヶ関", shuten:"日比谷", keiyu:"千代田線", kyori:1.2, jikan:2},
	{kiten:"日比谷", shuten:"二重橋前", keiyu:"千代田線", kyori:0.7, jikan:1},
	{kiten:"二重橋前", shuten:"大手町", keiyu:"千代田線", kyori:0.7, jikan:1},
	{kiten:"大手町", shuten:"新御茶ノ水", keiyu:"千代田線", kyori:1.3, jikan:2},
	{kiten:"新御茶ノ水", shuten:"湯島", keiyu:"千代田線", kyori:1.2, jikan:2},
	{kiten:"湯島", shuten:"根津", keiyu:"千代田線", kyori:1.2, jikan:2},
	{kiten:"根津", shuten:"千駄木", keiyu:"千代田線", kyori:1.0, jikan:2},
	{kiten:"千駄木", shuten:"西日暮里", keiyu:"千代田線", kyori:0.9, jikan:1},
	{kiten:"西日暮里", shuten:"町屋", keiyu:"千代田線", kyori:1.7, jikan:2},
	{kiten:"町屋", shuten:"北千住", keiyu:"千代田線", kyori:2.6, jikan:3},
	{kiten:"北千住", shuten:"綾瀬", keiyu:"千代田線", kyori:2.5, jikan:3},
	{kiten:"綾瀬", shuten:"北綾瀬", keiyu:"千代田線", kyori:2.1, jikan:4},
	{kiten:"浅草", shuten:"田原町", keiyu:"銀座線", kyori:0.8, jikan:2},
	{kiten:"田原町", shuten:"稲荷町", keiyu:"銀座線", kyori:0.7, jikan:1},
	{kiten:"稲荷町", shuten:"上野", keiyu:"銀座線", kyori:0.7, jikan:2},
	{kiten:"上野", shuten:"上野広小路", keiyu:"銀座線", kyori:0.5, jikan:2},
	{kiten:"上野広小路", shuten:"末広町", keiyu:"銀座線", kyori:0.6, jikan:1},
	{kiten:"末広町", shuten:"神田", keiyu:"銀座線", kyori:1.1, jikan:2},
	{kiten:"神田", shuten:"三越前", keiyu:"銀座線", kyori:0.7, jikan:1},
	{kiten:"三越前", shuten:"日本橋", keiyu:"銀座線", kyori:0.6, jikan:2},
	{kiten:"日本橋", shuten:"京橋", keiyu:"銀座線", kyori:0.7, jikan:2},
	{kiten:"京橋", shuten:"銀座", keiyu:"銀座線", kyori:0.7, jikan:1},
	{kiten:"銀座", shuten:"新橋", keiyu:"銀座線", kyori:0.9, jikan:2},
	{kiten:"新橋", shuten:"虎ノ門", keiyu:"銀座線", kyori:0.8, jikan:2},
	{kiten:"虎ノ門", shuten:"溜池山王", keiyu:"銀座線", kyori:0.6, jikan:2},
	{kiten:"溜池山王", shuten:"赤坂見附", keiyu:"銀座線", kyori:0.9, jikan:2},
	{kiten:"赤坂見附", shuten:"青山一丁目", keiyu:"銀座線", kyori:1.3, jikan:2},
	{kiten:"青山一丁目", shuten:"外苑前", keiyu:"銀座線", kyori:0.7, jikan:2},
	{kiten:"外苑前", shuten:"表参道", keiyu:"銀座線", kyori:0.7, jikan:1},
	{kiten:"表参道", shuten:"渋谷", keiyu:"銀座線", kyori:1.3, jikan:1},
	{kiten:"渋谷", shuten:"表参道", keiyu:"半蔵門線", kyori:1.3, jikan:2},
	{kiten:"表参道", shuten:"青山一丁目", keiyu:"半蔵門線", kyori:1.4, jikan:2},
	{kiten:"青山一丁目", shuten:"永田町", keiyu:"半蔵門線", kyori:1.3, jikan:2},
	{kiten:"永田町", shuten:"半蔵門", keiyu:"半蔵門線", kyori:1.0, jikan:2},
	{kiten:"半蔵門", shuten:"九段下", keiyu:"半蔵門線", kyori:1.6, jikan:2},
	{kiten:"九段下", shuten:"神保町", keiyu:"半蔵門線", kyori:0.4, jikan:1},
	{kiten:"神保町", shuten:"大手町", keiyu:"半蔵門線", kyori:1.7, jikan:3},
	{kiten:"大手町", shuten:"三越前", keiyu:"半蔵門線", kyori:0.7, jikan:1},
	{kiten:"三越前", shuten:"水天宮前", keiyu:"半蔵門線", kyori:1.3, jikan:2},
	{kiten:"水天宮前", shuten:"清澄白河", keiyu:"半蔵門線", kyori:1.7, jikan:3},
	{kiten:"清澄白河", shuten:"住吉", keiyu:"半蔵門線", kyori:1.9, jikan:3},
	{kiten:"住吉", shuten:"錦糸町", keiyu:"半蔵門線", kyori:1., jikan:2},
	{kiten:"錦糸町", shuten:"押上", keiyu:"半蔵門線", kyori:1.4, jikan:2},
	{kiten:"中目黒", shuten:"恵比寿", keiyu:"日比谷線", kyori:1., jikan:2},
	{kiten:"恵比寿", shuten:"広尾", keiyu:"日比谷線", kyori:1.5, jikan:3},
	{kiten:"広尾", shuten:"六本木", keiyu:"日比谷線", kyori:1.7, jikan:3},
	{kiten:"六本木", shuten:"神谷町", keiyu:"日比谷線", kyori:1.5, jikan:3},
	{kiten:"神谷町", shuten:"霞ヶ関", keiyu:"日比谷線", kyori:1.3, jikan:2},
	{kiten:"霞ヶ関", shuten:"日比谷", keiyu:"日比谷線", kyori:1.2, jikan:2},
	{kiten:"日比谷", shuten:"銀座", keiyu:"日比谷線", kyori:0.4, jikan:1},
	{kiten:"銀座", shuten:"東銀座", keiyu:"日比谷線", kyori:0.4, jikan:1},
	{kiten:"東銀座", shuten:"築地", keiyu:"日比谷線", kyori:0.6, jikan:2},
	{kiten:"築地", shuten:"八丁堀", keiyu:"日比谷線", kyori:1., jikan:2},
	{kiten:"八丁堀", shuten:"茅場町", keiyu:"日比谷線", kyori:0.5, jikan:1},
	{kiten:"茅場町", shuten:"人形町", keiyu:"日比谷線", kyori:0.9, jikan:2},
	{kiten:"人形町", shuten:"小伝馬町", keiyu:"日比谷線", kyori:0.6, jikan:1},
	{kiten:"小伝馬町", shuten:"秋葉原", keiyu:"日比谷線", kyori:0.9, jikan:2},
	{kiten:"秋葉原", shuten:"仲御徒町", keiyu:"日比谷線", kyori:1., jikan:1},
	{kiten:"仲御徒町", shuten:"上野", keiyu:"日比谷線", kyori:0.5, jikan:1},
	{kiten:"上野", shuten:"入谷", keiyu:"日比谷線", kyori:1.2, jikan:2},
	{kiten:"入谷", shuten:"三ノ輪", keiyu:"日比谷線", kyori:1.2, jikan:2},
	{kiten:"三ノ輪", shuten:"南千住", keiyu:"日比谷線", kyori:0.8, jikan:2},
	{kiten:"南千住", shuten:"北千住", keiyu:"日比谷線", kyori:1.8, jikan:3},
	{kiten:"池袋", shuten:"新大塚", keiyu:"丸ノ内線", kyori:1.8, jikan:3},
	{kiten:"新大塚", shuten:"茗荷谷", keiyu:"丸ノ内線", kyori:1.2, jikan:2},
	{kiten:"茗荷谷", shuten:"後楽園", keiyu:"丸ノ内線", kyori:1.8, jikan:2},
	{kiten:"後楽園", shuten:"本郷三丁目", keiyu:"丸ノ内線", kyori:0.8, jikan:1},
	{kiten:"本郷三丁目", shuten:"御茶ノ水", keiyu:"丸ノ内線", kyori:0.8, jikan:1},
	{kiten:"御茶ノ水", shuten:"淡路町", keiyu:"丸ノ内線", kyori:0.8, jikan:1},
	{kiten:"淡路町", shuten:"大手町", keiyu:"丸ノ内線", kyori:0.9, jikan:2},
	{kiten:"大手町", shuten:"東京", keiyu:"丸ノ内線", kyori:0.6, jikan:1},
	{kiten:"東京", shuten:"銀座", keiyu:"丸ノ内線", kyori:1.1, jikan:2},
	{kiten:"銀座", shuten:"霞ヶ関", keiyu:"丸ノ内線", kyori:1.0, jikan:2},
	{kiten:"霞ヶ関", shuten:"国会議事堂前", keiyu:"丸ノ内線", kyori:0.7, jikan:1},
	{kiten:"国会議事堂前", shuten:"赤坂見附", keiyu:"丸ノ内線", kyori:0.9, jikan:2},
	{kiten:"赤坂見附", shuten:"四ツ谷", keiyu:"丸ノ内線", kyori:1.3, jikan:2},
	{kiten:"四ツ谷", shuten:"四谷三丁目", keiyu:"丸ノ内線", kyori:1.0, jikan:2},
	{kiten:"四谷三丁目", shuten:"新宿御苑前", keiyu:"丸ノ内線", kyori:0.9, jikan:1},
	{kiten:"新宿御苑前", shuten:"新宿三丁目", keiyu:"丸ノ内線", kyori:0.7, jikan:1},
	{kiten:"新宿三丁目", shuten:"新宿", keiyu:"丸ノ内線", kyori:0.3, jikan:1},
	{kiten:"新宿", shuten:"西新宿", keiyu:"丸ノ内線", kyori:0.8, jikan:1},
	{kiten:"西新宿", shuten:"中野坂上", keiyu:"丸ノ内線", kyori:1.1, jikan:2},
	{kiten:"中野坂上", shuten:"新中野", keiyu:"丸ノ内線", kyori:1.1, jikan:2},
	{kiten:"新中野", shuten:"東高円寺", keiyu:"丸ノ内線", kyori:1.0, jikan:1},
	{kiten:"東高円寺", shuten:"新高円寺", keiyu:"丸ノ内線", kyori:0.9, jikan:1},
	{kiten:"新高円寺", shuten:"南阿佐ヶ谷", keiyu:"丸ノ内線", kyori:1.2, jikan:2},
	{kiten:"南阿佐ヶ谷", shuten:"荻窪", keiyu:"丸ノ内線", kyori:1.5, jikan:2},
	{kiten:"中野坂上", shuten:"中野新橋", keiyu:"丸ノ内線", kyori:1.3, jikan:2},
	{kiten:"中野新橋", shuten:"中野富士見町", keiyu:"丸ノ内線", kyori:0.6, jikan:1},
	{kiten:"中野富士見町", shuten:"方南町", keiyu:"丸ノ内線", kyori:1.3, jikan:2},
	{kiten:"市ヶ谷", shuten:"四ツ谷", keiyu:"南北線", kyori:1.0, jikan:2},
	{kiten:"四ツ谷", shuten:"永田町", keiyu:"南北線", kyori:1.3, jikan:3},
	{kiten:"永田町", shuten:"溜池山王", keiyu:"南北線", kyori:0.9, jikan:1},
	{kiten:"溜池山王", shuten:"六本木一丁目", keiyu:"南北線", kyori:0.9, jikan:2},
	{kiten:"六本木一丁目", shuten:"麻布十番", keiyu:"南北線", kyori:1.2, jikan:2},
	{kiten:"麻布十番", shuten:"白金高輪", keiyu:"南北線", kyori:1.3, jikan:2},
	{kiten:"白金高輪", shuten:"白金台", keiyu:"南北線", kyori:1.0, jikan:2},
	{kiten:"白金台", shuten:"目黒", keiyu:"南北線", kyori:1.3, jikan:2},
	{kiten:"市ヶ谷", shuten:"飯田橋", keiyu:"南北線", kyori:1.1, jikan:2},
	{kiten:"飯田橋", shuten:"後楽園", keiyu:"南北線", kyori:1.4, jikan:2},
	{kiten:"後楽園", shuten:"東大前", keiyu:"南北線", kyori:1.3, jikan:3},
	{kiten:"東大前", shuten:"本駒込", keiyu:"南北線", kyori:0.9, jikan:2},
	{kiten:"本駒込", shuten:"駒込", keiyu:"南北線", kyori:1.4, jikan:2},
	{kiten:"駒込", shuten:"西ヶ原", keiyu:"南北線", kyori:1.4, jikan:2},
	{kiten:"西ヶ原", shuten:"王子", keiyu:"南北線", kyori:1.0, jikan:2},
	{kiten:"王子", shuten:"王子神谷", keiyu:"南北線", kyori:1.2, jikan:2},
	{kiten:"王子神谷", shuten:"志茂", keiyu:"南北線", kyori:1.6, jikan:3},
	{kiten:"志茂", shuten:"赤羽岩淵", keiyu:"南北線", kyori:1.1, jikan:2},
	{kiten:"西船橋", shuten:"原木中山", keiyu:"東西線", kyori:1.9, jikan:3},
	{kiten:"原木中山", shuten:"妙典", keiyu:"東西線", kyori:2.1, jikan:2},
	{kiten:"妙典", shuten:"行徳", keiyu:"東西線", kyori:1.3, jikan:2},
	{kiten:"行徳", shuten:"南行徳", keiyu:"東西線", kyori:1.5, jikan:2},
	{kiten:"南行徳", shuten:"浦安", keiyu:"東西線", kyori:1.2, jikan:2},
	{kiten:"浦安", shuten:"葛西", keiyu:"東西線", kyori:1.9, jikan:2},
	{kiten:"葛西", shuten:"西葛西", keiyu:"東西線", kyori:1.2, jikan:2},
	{kiten:"西葛西", shuten:"南砂町", keiyu:"東西線", kyori:2.7, jikan:2},
	{kiten:"南砂町", shuten:"東陽町", keiyu:"東西線", kyori:1.2, jikan:2},
	{kiten:"東陽町", shuten:"木場", keiyu:"東西線", kyori:0.9, jikan:1},
	{kiten:"木場", shuten:"門前仲町", keiyu:"東西線", kyori:1.1, jikan:1},
	{kiten:"門前仲町", shuten:"茅場町", keiyu:"東西線", kyori:1.8, jikan:2},
	{kiten:"茅場町", shuten:"日本橋", keiyu:"東西線", kyori:0.5, jikan:1},
	{kiten:"日本橋", shuten:"大手町", keiyu:"東西線", kyori:0.8, jikan:1},
	{kiten:"大手町", shuten:"竹橋", keiyu:"東西線", kyori:1.0, jikan:2},
	{kiten:"竹橋", shuten:"九段下", keiyu:"東西線", kyori:1.0, jikan:1},
	{kiten:"九段下", shuten:"飯田橋", keiyu:"東西線", kyori:0.7, jikan:1},
	{kiten:"飯田橋", shuten:"神楽坂", keiyu:"東西線", kyori:1.2, jikan:2},
	{kiten:"神楽坂", shuten:"早稲田", keiyu:"東西線", kyori:1.2, jikan:2},
	{kiten:"早稲田", shuten:"高田馬場", keiyu:"東西線", kyori:1.7, jikan:3},
	{kiten:"高田馬場", shuten:"落合", keiyu:"東西線", kyori:1.9, jikan:3},
	{kiten:"落合", shuten:"中野", keiyu:"東西線", kyori:2.0, jikan:3},
	{kiten:"新木場", shuten:"辰巳", keiyu:"有楽町線", kyori:1.5, jikan:2},
	{kiten:"辰巳", shuten:"豊洲", keiyu:"有楽町線", kyori:1.7, jikan:2},
	{kiten:"豊洲", shuten:"月島", keiyu:"有楽町線", kyori:1.4, jikan:2},
	{kiten:"月島", shuten:"新富町", keiyu:"有楽町線", kyori:1.3, jikan:2},
	{kiten:"新富町", shuten:"銀座一丁目", keiyu:"有楽町線", kyori:0.7, jikan:1},
	{kiten:"銀座一丁目", shuten:"有楽町", keiyu:"有楽町線", kyori:0.5, jikan:1},
	{kiten:"有楽町", shuten:"桜田門", keiyu:"有楽町線", kyori:1.0, jikan:1},
	{kiten:"桜田門", shuten:"永田町", keiyu:"有楽町線", kyori:0.9, jikan:2},
	{kiten:"永田町", shuten:"麹町", keiyu:"有楽町線", kyori:0.9, jikan:1},
	{kiten:"麹町", shuten:"市ヶ谷", keiyu:"有楽町線", kyori:0.9, jikan:1},
	{kiten:"市ヶ谷", shuten:"飯田橋", keiyu:"有楽町線", kyori:1.1, jikan:2},
	{kiten:"飯田橋", shuten:"江戸川橋", keiyu:"有楽町線", kyori:1.6, jikan:3},
	{kiten:"江戸川橋", shuten:"護国寺", keiyu:"有楽町線", kyori:1.3, jikan:2},
	{kiten:"護国寺", shuten:"東池袋", keiyu:"有楽町線", kyori:1.1, jikan:2},
	{kiten:"東池袋", shuten:"池袋", keiyu:"有楽町線", kyori:2.0, jikan:2},
	{kiten:"池袋", shuten:"要町", keiyu:"有楽町線", kyori:1.2, jikan:2},
	{kiten:"要町", shuten:"千川", keiyu:"有楽町線", kyori:1.0, jikan:1},
	{kiten:"千川", shuten:"小竹向原", keiyu:"有楽町線", kyori:1.0, jikan:2},
	{kiten:"小竹向原", shuten:"氷川台", keiyu:"有楽町線", kyori:1.5, jikan:2},
	{kiten:"氷川台", shuten:"平和台", keiyu:"有楽町線", kyori:1.4, jikan:2},
	{kiten:"平和台", shuten:"営団赤塚", keiyu:"有楽町線", kyori:1.8, jikan:2},
	{kiten:"営団赤塚", shuten:"営団成増", keiyu:"有楽町線", kyori:1.5, jikan:2},
	{kiten:"営団成増", shuten:"和光市", keiyu:"有楽町線", kyori:2.1, jikan:3},
}