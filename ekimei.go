package main

var GlobalEkimeiList = &GlobalEkimei{EkimeiList: Ekimei_List}

type Ekimei struct {
	Kanji   string
	Kana    string
	Romaji  string
	Shozoku string
}

// Returns Shozoku and Kanji(Kana)
// Ex: &e{Kanji: A, Kana: B, Shozoku: C}.hyoji() -> "C, A(B)"
func (e *Ekimei) Hyoji() string {
	return e.Shozoku + ", " + e.Kanji + "(" + e.Kana + ")"
}

type GlobalEkimei struct {
	EkimeiList []Ekimei
}

// Sort EkimeiList in ascending order and remove any duplicated station
// ex: &g{list: global_ekimei_list}.Seiretsu() -> sorted global_ekimei_list
func (g *GlobalEkimei) Seiretsu() []Ekimei {
	return []Ekimei{}
}

// Convert ローマ字 to 漢字
// Ex: &g{list: global_ekimei_list}.RomajiToKanji("myogadani") ->　茗荷谷
func (g *GlobalEkimei) RomajiToKanji(romaji string) string {
	for _, ekimei := range g.EkimeiList {
		if ekimei.Romaji == romaji {
			return ekimei.Kanji
		}
	}
	return ""
	// Recursive Method on Below
	//if cap(g.EkimeiList) == 0 || g.EkimeiList == nil {
	//	return ""
	//}
	//
	//if g.EkimeiList[0].Romaji == romaji {
	//	return g.EkimeiList[0].Kanji
	//} else {
	//	g.EkimeiList = g.EkimeiList[1:]
	//	return g.RomajiToKanji(romaji)
	//}
}

var Ekimei_List = []Ekimei{
	{Kanji:"代々木上原", Kana:"よよぎうえはら", Romaji:"yoyogiuehara", Shozoku:"千代田線"},
	{Kanji:"代々木公園", Kana:"よよぎこうえん", Romaji:"yoyogikouen", Shozoku:"千代田線"},
	{Kanji:"明治神宮前", Kana:"めいじじんぐうまえ", Romaji:"meijijinguumae", Shozoku:"千代田線"},
	{Kanji:"表参道", Kana:"おもてさんどう", Romaji:"omotesandou", Shozoku:"千代田線"},
	{Kanji:"乃木坂", Kana:"のぎざか", Romaji:"nogizaka", Shozoku:"千代田線"},
	{Kanji:"赤坂", Kana:"あかさか", Romaji:"akasaka", Shozoku:"千代田線"},
	{Kanji:"国会議事堂前", Kana:"こっかいぎじどうまえ", Romaji:"kokkaigijidoumae", Shozoku:"千代田線"},
	{Kanji:"霞ヶ関", Kana:"かすみがせき", Romaji:"kasumigaseki", Shozoku:"千代田線"},
	{Kanji:"日比谷", Kana:"ひびや", Romaji:"hibiya", Shozoku:"千代田線"},
	{Kanji:"二重橋前", Kana:"にじゅうばしまえ", Romaji:"nijuubasimae", Shozoku:"千代田線"},
	{Kanji:"大手町", Kana:"おおてまち", Romaji:"otemachi", Shozoku:"千代田線"},
	{Kanji:"新御茶ノ水", Kana:"しんおちゃのみず", Romaji:"shin-ochanomizu", Shozoku:"千代田線"},
	{Kanji:"湯島", Kana:"ゆしま", Romaji:"yushima", Shozoku:"千代田線"},
	{Kanji:"根津", Kana:"ねづ", Romaji:"nedu", Shozoku:"千代田線"},
	{Kanji:"千駄木", Kana:"せんだぎ", Romaji:"sendagi", Shozoku:"千代田線"},
	{Kanji:"西日暮里", Kana:"にしにっぽり", Romaji:"nishinippori", Shozoku:"千代田線"},
	{Kanji:"町屋", Kana:"まちや", Romaji:"machiya", Shozoku:"千代田線"},
	{Kanji:"北千住", Kana:"きたせんじゅ", Romaji:"kitasenjyu", Shozoku:"千代田線"},
	{Kanji:"綾瀬", Kana:"あやせ", Romaji:"ayase", Shozoku:"千代田線"},
	{Kanji:"北綾瀬", Kana:"きたあやせ", Romaji:"kitaayase", Shozoku:"千代田線"},
	{Kanji:"浅草", Kana:"あさくさ", Romaji:"asakusa", Shozoku:"銀座線"},
	{Kanji:"田原町", Kana:"たわらまち", Romaji:"tawaramachi", Shozoku:"銀座線"},
	{Kanji:"稲荷町", Kana:"いなりちょう", Romaji:"inaricho", Shozoku:"銀座線"},
	{Kanji:"上野", Kana:"うえの", Romaji:"ueno", Shozoku:"銀座線"},
	{Kanji:"上野広小路", Kana:"うえのひろこうじ", Romaji:"uenohirokoji", Shozoku:"銀座線"},
	{Kanji:"末広町", Kana:"すえひろちょう", Romaji:"suehirocho", Shozoku:"銀座線"},
	{Kanji:"神田", Kana:"かんだ", Romaji:"kanda", Shozoku:"銀座線"},
	{Kanji:"三越前", Kana:"みつこしまえ", Romaji:"mitsukoshimae", Shozoku:"銀座線"},
	{Kanji:"日本橋", Kana:"にほんばし", Romaji:"nihonbashi", Shozoku:"銀座線"},
	{Kanji:"京橋", Kana:"きょうばし", Romaji:"kyobashi", Shozoku:"銀座線"},
	{Kanji:"銀座", Kana:"ぎんざ", Romaji:"ginza", Shozoku:"銀座線"},
	{Kanji:"新橋", Kana:"しんばし", Romaji:"shinbashi", Shozoku:"銀座線"},
	{Kanji:"虎ノ門", Kana:"とらのもん", Romaji:"toranomon", Shozoku:"銀座線"},
	{Kanji:"溜池山王", Kana:"ためいけさんのう", Romaji:"tameikesannou", Shozoku:"銀座線"},
	{Kanji:"赤坂見附", Kana:"あかさかみつけ", Romaji:"akasakamitsuke", Shozoku:"銀座線"},
	{Kanji:"青山一丁目", Kana:"あおやまいっちょうめ", Romaji:"aoyamaicchome", Shozoku:"銀座線"},
	{Kanji:"外苑前", Kana:"がいえんまえ", Romaji:"gaienmae", Shozoku:"銀座線"},
	{Kanji:"表参道", Kana:"おもてさんどう", Romaji:"omotesando", Shozoku:"銀座線"},
	{Kanji:"渋谷", Kana:"しぶや", Romaji:"shibuya", Shozoku:"銀座線"},
	{Kanji:"渋谷", Kana:"しぶや", Romaji:"shibuya", Shozoku:"半蔵門線"},
	{Kanji:"表参道", Kana:"おもてさんどう", Romaji:"omotesandou", Shozoku:"半蔵門線"},
	{Kanji:"青山一丁目", Kana:"あおやまいっちょうめ", Romaji:"aoyama-itchome", Shozoku:"半蔵門線"},
	{Kanji:"永田町", Kana:"ながたちょう", Romaji:"nagatacho", Shozoku:"半蔵門線"},
	{Kanji:"半蔵門", Kana:"はんぞうもん", Romaji:"hanzomon", Shozoku:"半蔵門線"},
	{Kanji:"九段下", Kana:"くだんした", Romaji:"kudanshita", Shozoku:"半蔵門線"},
	{Kanji:"神保町", Kana:"じんぼうちょう", Romaji:"jinbocho", Shozoku:"半蔵門線"},
	{Kanji:"大手町", Kana:"おおてまち", Romaji:"otemachi", Shozoku:"半蔵門線"},
	{Kanji:"三越前", Kana:"みつこしまえ", Romaji:"mitsukoshimae", Shozoku:"半蔵門線"},
	{Kanji:"水天宮前", Kana:"すいてんぐうまえ", Romaji:"suitengumae", Shozoku:"半蔵門線"},
	{Kanji:"清澄白河", Kana:"きよすみしらかわ", Romaji:"kiyosumi-shirakawa", Shozoku:"半蔵門線"},
	{Kanji:"住吉", Kana:"すみよし", Romaji:"sumiyoshi", Shozoku:"半蔵門線"},
	{Kanji:"錦糸町", Kana:"きんしちょう", Romaji:"kinshicho", Shozoku:"半蔵門線"},
	{Kanji:"押上", Kana:"おしあげ", Romaji:"oshiage", Shozoku:"半蔵門線"},
	{Kanji:"中目黒", Kana:"なかめぐろ", Romaji:"nakameguro", Shozoku:"日比谷線"},
	{Kanji:"恵比寿", Kana:"えびす", Romaji:"ebisu", Shozoku:"日比谷線"},
	{Kanji:"広尾", Kana:"ひろお", Romaji:"hiro", Shozoku:"日比谷線"},
	{Kanji:"六本木", Kana:"ろっぽんぎ", Romaji:"roppongi", Shozoku:"日比谷線"},
	{Kanji:"神谷町", Kana:"かみやちょう", Romaji:"kamiyacho", Shozoku:"日比谷線"},
	{Kanji:"霞ヶ関", Kana:"かすみがせき", Romaji:"kasumigaseki", Shozoku:"日比谷線"},
	{Kanji:"日比谷", Kana:"ひびや", Romaji:"hibiya", Shozoku:"日比谷線"},
	{Kanji:"銀座", Kana:"ぎんざ", Romaji:"ginza", Shozoku:"日比谷線"},
	{Kanji:"東銀座", Kana:"ひがしぎんざ", Romaji:"higashiginza", Shozoku:"日比谷線"},
	{Kanji:"築地", Kana:"つきじ", Romaji:"tsukiji", Shozoku:"日比谷線"},
	{Kanji:"八丁堀", Kana:"はっちょうぼり", Romaji:"hacchobori", Shozoku:"日比谷線"},
	{Kanji:"茅場町", Kana:"かやばちょう", Romaji:"kayabacho", Shozoku:"日比谷線"},
	{Kanji:"人形町", Kana:"にんぎょうちょう", Romaji:"ningyomachi", Shozoku:"日比谷線"},
	{Kanji:"小伝馬町", Kana:"こでんまちょう", Romaji:"kodemmacho", Shozoku:"日比谷線"},
	{Kanji:"秋葉原", Kana:"あきはばら", Romaji:"akihabara", Shozoku:"日比谷線"},
	{Kanji:"仲御徒町", Kana:"なかおかちまち", Romaji:"nakaokachimachi", Shozoku:"日比谷線"},
	{Kanji:"上野", Kana:"うえの", Romaji:"ueno", Shozoku:"日比谷線"},
	{Kanji:"入谷", Kana:"いりや", Romaji:"iriya", Shozoku:"日比谷線"},
	{Kanji:"三ノ輪", Kana:"みのわ", Romaji:"minowa", Shozoku:"日比谷線"},
	{Kanji:"南千住", Kana:"みなみせんじゅ", Romaji:"minamisenju", Shozoku:"日比谷線"},
	{Kanji:"北千住", Kana:"きたせんじゅ", Romaji:"kitasenju", Shozoku:"日比谷線"},
	{Kanji:"池袋", Kana:"いけぶくろ", Romaji:"ikebukuro", Shozoku:"丸ノ内線"},
	{Kanji:"新大塚", Kana:"しんおおつか", Romaji:"shinotsuka", Shozoku:"丸ノ内線"},
	{Kanji:"茗荷谷", Kana:"みょうがだに", Romaji:"myogadani", Shozoku:"丸ノ内線"},
	{Kanji:"後楽園", Kana:"こうらくえん", Romaji:"korakuen", Shozoku:"丸ノ内線"},
	{Kanji:"本郷三丁目", Kana:"ほんごうさんちょうめ", Romaji:"hongosanchome", Shozoku:"丸ノ内線"},
	{Kanji:"御茶ノ水", Kana:"おちゃのみず", Romaji:"ochanomizu", Shozoku:"丸ノ内線"},
	{Kanji:"淡路町", Kana:"あわじちょう", Romaji:"awajicho", Shozoku:"丸ノ内線"},
	{Kanji:"大手町", Kana:"おおてまち", Romaji:"otemachi", Shozoku:"丸ノ内線"},
	{Kanji:"東京", Kana:"とうきょう", Romaji:"tokyo", Shozoku:"丸ノ内線"},
	{Kanji:"銀座", Kana:"ぎんざ", Romaji:"ginza", Shozoku:"丸ノ内線"},
	{Kanji:"霞ヶ関", Kana:"かすみがせき", Romaji:"kasumigaseki", Shozoku:"丸ノ内線"},
	{Kanji:"国会議事堂前", Kana:"こっかいぎじどうまえ", Romaji:"kokkaigijidomae", Shozoku:"丸ノ内線"},
	{Kanji:"赤坂見附", Kana:"あかさかみつけ", Romaji:"akasakamitsuke", Shozoku:"丸ノ内線"},
	{Kanji:"四ツ谷", Kana:"よつや", Romaji:"yotsuya", Shozoku:"丸ノ内線"},
	{Kanji:"四谷三丁目", Kana:"よつやさんちょうめ", Romaji:"yotsuyasanchome", Shozoku:"丸ノ内線"},
	{Kanji:"新宿御苑前", Kana:"しんじゅくぎょえんまえ", Romaji:"shinjuku-gyoemmae", Shozoku:"丸ノ内線"},
	{Kanji:"新宿三丁目", Kana:"しんじゅくさんちょうめ", Romaji:"shinjuku-sanchome", Shozoku:"丸ノ内線"},
	{Kanji:"新宿", Kana:"しんじゅく", Romaji:"shinjuku", Shozoku:"丸ノ内線"},
	{Kanji:"西新宿", Kana:"にししんじゅく", Romaji:"nishi-shinjuku", Shozoku:"丸ノ内線"},
	{Kanji:"中野坂上", Kana:"なかのさかうえ", Romaji:"nakano-sakaue", Shozoku:"丸ノ内線"},
	{Kanji:"新中野", Kana:"しんなかの", Romaji:"shin-nakano", Shozoku:"丸ノ内線"},
	{Kanji:"東高円寺", Kana:"ひがしこうえんじ", Romaji:"higashi-koenji", Shozoku:"丸ノ内線"},
	{Kanji:"新高円寺", Kana:"しんこうえんじ", Romaji:"shin-koenji", Shozoku:"丸ノ内線"},
	{Kanji:"南阿佐ヶ谷", Kana:"みなみあさがや", Romaji:"minami-asagaya", Shozoku:"丸ノ内線"},
	{Kanji:"荻窪", Kana:"おぎくぼ", Romaji:"ogikubo", Shozoku:"丸ノ内線"},
	{Kanji:"中野新橋", Kana:"なかのしんばし", Romaji:"nakano-shimbashi", Shozoku:"丸ノ内線"},
	{Kanji:"中野富士見町", Kana:"なかのふじみちょう", Romaji:"nakano-fujimicho", Shozoku:"丸ノ内線"},
	{Kanji:"方南町", Kana:"ほうなんちょう", Romaji:"honancho", Shozoku:"丸ノ内線"},
	{Kanji:"四ツ谷", Kana:"よつや", Romaji:"yotsuya", Shozoku:"南北線"},
	{Kanji:"永田町", Kana:"ながたちょう", Romaji:"nagatacho", Shozoku:"南北線"},
	{Kanji:"溜池山王", Kana:"ためいけさんのう", Romaji:"tameikesanno", Shozoku:"南北線"},
	{Kanji:"六本木一丁目", Kana:"ろっぽんぎいっちょうめ", Romaji:"roppongiitchome", Shozoku:"南北線"},
	{Kanji:"麻布十番", Kana:"あざぶじゅうばん", Romaji:"azabujuban", Shozoku:"南北線"},
	{Kanji:"白金高輪", Kana:"しろかねたかなわ", Romaji:"shirokanetakanawa", Shozoku:"南北線"},
	{Kanji:"白金台", Kana:"しろかねだい", Romaji:"shirokanedai", Shozoku:"南北線"},
	{Kanji:"目黒", Kana:"めぐろ", Romaji:"meguro", Shozoku:"南北線"},
	{Kanji:"市ヶ谷", Kana:"いちがや", Romaji:"ichigaya", Shozoku:"南北線"},
	{Kanji:"飯田橋", Kana:"いいだばし", Romaji:"idabashi", Shozoku:"南北線"},
	{Kanji:"後楽園", Kana:"こうらくえん", Romaji:"korakuen", Shozoku:"南北線"},
	{Kanji:"東大前", Kana:"とうだいまえ", Romaji:"todaimae", Shozoku:"南北線"},
	{Kanji:"本駒込", Kana:"ほんこまごめ", Romaji:"honkomagome", Shozoku:"南北線"},
	{Kanji:"駒込", Kana:"こまごめ", Romaji:"komagome", Shozoku:"南北線"},
	{Kanji:"西ヶ原", Kana:"にしがはら", Romaji:"nishigahara", Shozoku:"南北線"},
	{Kanji:"王子", Kana:"おうじ", Romaji:"oji", Shozoku:"南北線"},
	{Kanji:"王子神谷", Kana:"おうじかみや", Romaji:"ojikamiya", Shozoku:"南北線"},
	{Kanji:"志茂", Kana:"しも", Romaji:"shimo", Shozoku:"南北線"},
	{Kanji:"赤羽岩淵", Kana:"あかばねいわぶち", Romaji:"akabaneiwabuchi", Shozoku:"南北線"},
	{Kanji:"西船橋", Kana:"にしふなばし", Romaji:"nishi-funabashi", Shozoku:"東西線"},
	{Kanji:"原木中山", Kana:"ばらきなかやま", Romaji:"baraki-nakayama", Shozoku:"東西線"},
	{Kanji:"妙典", Kana:"みょうでん", Romaji:"myoden", Shozoku:"東西線"},
	{Kanji:"行徳", Kana:"ぎょうとく", Romaji:"gyotoku", Shozoku:"東西線"},
	{Kanji:"南行徳", Kana:"みなみぎょうとく", Romaji:"minami-gyotoku", Shozoku:"東西線"},
	{Kanji:"浦安", Kana:"うらやす", Romaji:"urayasu", Shozoku:"東西線"},
	{Kanji:"葛西", Kana:"かさい", Romaji:"kasai", Shozoku:"東西線"},
	{Kanji:"西葛西", Kana:"にしかさい", Romaji:"nishi-kasai", Shozoku:"東西線"},
	{Kanji:"南砂町", Kana:"みなみすなまち", Romaji:"minami-sunamachi", Shozoku:"東西線"},
	{Kanji:"東陽町", Kana:"とうようちょう", Romaji:"touyoucho", Shozoku:"東西線"},
	{Kanji:"木場", Kana:"きば", Romaji:"kiba", Shozoku:"東西線"},
	{Kanji:"門前仲町", Kana:"もんぜんなかちょう", Romaji:"monzen-nakacho", Shozoku:"東西線"},
	{Kanji:"茅場町", Kana:"かやばちょう", Romaji:"kayabacho", Shozoku:"東西線"},
	{Kanji:"日本橋", Kana:"にほんばし", Romaji:"nihonbashi", Shozoku:"東西線"},
	{Kanji:"大手町", Kana:"おおてまち", Romaji:"otemachi", Shozoku:"東西線"},
	{Kanji:"竹橋", Kana:"たけばし", Romaji:"takebashi", Shozoku:"東西線"},
	{Kanji:"九段下", Kana:"くだんした", Romaji:"kudanshita", Shozoku:"東西線"},
	{Kanji:"飯田橋", Kana:"いいだばし", Romaji:"iidabashi", Shozoku:"東西線"},
	{Kanji:"神楽坂", Kana:"かぐらざか", Romaji:"kagurazaka", Shozoku:"東西線"},
	{Kanji:"早稲田", Kana:"わせだ", Romaji:"waseda", Shozoku:"東西線"},
	{Kanji:"高田馬場", Kana:"たかだのばば", Romaji:"takadanobaba", Shozoku:"東西線"},
	{Kanji:"落合", Kana:"おちあい", Romaji:"ochiai", Shozoku:"東西線"},
	{Kanji:"中野", Kana:"なかの", Romaji:"nakano", Shozoku:"東西線"},
	{Romaji:"shinkiba", Kana:"しんきば", Kanji:"新木場", Shozoku:"有楽町線"},
	{Romaji:"tatsumi", Kana:"たつみ", Kanji:"辰巳", Shozoku:"有楽町線"},
	{Romaji:"toyosu", Kana:"とよす", Kanji:"豊洲", Shozoku:"有楽町線"},
	{Romaji:"tsukishima", Kana:"つきしま", Kanji:"月島", Shozoku:"有楽町線"},
	{Romaji:"shintomityou", Kana:"しんとみちょう", Kanji:"新富町", Shozoku:"有楽町線"},
	{Romaji:"ginzaittyoume", Kana:"ぎんざいっちょうめ", Kanji:"銀座一丁目", Shozoku:"有楽町線"},
	{Romaji:"yuurakutyou", Kana:"ゆうらくちょう", Kanji:"有楽町", Shozoku:"有楽町線"},
	{Romaji:"sakuradamon", Kana:"さくらだもん", Kanji:"桜田門", Shozoku:"有楽町線"},
	{Romaji:"nagatacho", Kana:"ながたちょう", Kanji:"永田町", Shozoku:"有楽町線"},
	{Romaji:"koujimachi", Kana:"こうじまち", Kanji:"麹町", Shozoku:"有楽町線"},
	{Romaji:"ichigaya", Kana:"いちがや", Kanji:"市ヶ谷", Shozoku:"有楽町線"},
	{Romaji:"iidabashi", Kana:"いいだばし", Kanji:"飯田橋", Shozoku:"有楽町線"},
	{Kanji:"江戸川橋", Kana:"えどがわばし", Romaji:"edogawabasi", Shozoku:"有楽町線"},
	{Kanji:"護国寺", Kana:"ごこくじ", Romaji:"gokokuji", Shozoku:"有楽町線"},
	{Kanji:"東池袋", Kana:"ひがしいけぶくろ", Romaji:"higasiikebukuro", Shozoku:"有楽町線"},
	{Kanji:"池袋", Kana:"いけぶくろ", Romaji:"ikebukuro", Shozoku:"有楽町線"},
	{Kanji:"要町", Kana:"かなめちょう", Romaji:"kanametyou", Shozoku:"有楽町線"},
	{Kanji:"千川", Kana:"せんかわ", Romaji:"senkawa", Shozoku:"有楽町線"},
	{Kanji:"小竹向原", Kana:"こたけむかいはら", Romaji:"kotakemukaihara", Shozoku:"有楽町線"},
	{Kanji:"氷川台", Kana:"ひかわだい", Romaji:"hikawadai", Shozoku:"有楽町線"},
	{Kanji:"平和台", Kana:"へいわだい", Romaji:"heiwadai", Shozoku:"有楽町線"},
	{Kanji:"営団赤塚", Kana:"えいだんあかつか", Romaji:"eidanakakuka", Shozoku:"有楽町線"},
	{Kanji:"営団成増", Kana:"えいだんなります", Romaji:"eidannarimasu", Shozoku:"有楽町線"},
	{Kanji:"和光市", Kana:"わこうし", Romaji:"wakousi", Shozoku:"有楽町線"},
}