package main

func main() {
	kiten := "shibuya"
	syuten := "mitsukoshimae"

	ekimei_list := &GlobalEkimei{EkimeiList: Ekimei_List}
	ekimei_list.Seiretsu()
	kiten = ekimei_list.RomajiToKanji(kiten)
	syuten = ekimei_list.RomajiToKanji(syuten)

	eki_list := MakeEkiList(ekimei_list)
	eki_list.Shokika(kiten)

}
