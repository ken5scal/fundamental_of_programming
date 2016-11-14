package main

type Eki struct {
	namae        string
	saitan_kyori float64
	temae_list   []string
}

// return []Eki's temae_list should be empty list
// Ex: MakeEkiList(&GlobalEkimei{EkimeiList: Ekimei_List}) -> return Eki
func MakeEkiList(g GlobalEkimei) []Eki {
	return &Eki{}
}

// kiten should be Romaji!
// kiten's saitan_kyori should be 0
// kiten's temae_list should only contain kiten itself
func Shokika(eki_list []Eki, kiten string) {

}