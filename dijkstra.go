package main

import (
	"math"
)

type Eki struct {
	namae        string // 漢字
	saitan_kyori float64
	temae_list   []string
}

type EkiList struct {
	eki_list []Eki
}

// return []Eki's temae_list(should be empty list)
// Ex: MakeEkiList(&GlobalEkimei{EkimeiList: Ekimei_List}) -> return listEki
// 		-> return Eki List with name: Kanji, saitan_kyoei inf, temae_list: empty list
func MakeEkiList(g *GlobalEkimei) *EkiList {
	var eki_list []Eki

	for _, eki := range g.EkimeiList {
		eki_list = append(
			eki_list,
			Eki{
				namae:eki.Kanji,
				saitan_kyori:math.Inf(+1)})
	}
	return &EkiList{eki_list: eki_list}
}

// kiten should be Romaji!
// kiten's saitan_kyori should be 0
// kiten's temae_list should only contain kiten itself
func (e *EkiList) Shokika(kiten string) {
	for i := range e.eki_list {
		if e.eki_list[i].namae == kiten {
			e.eki_list[i].saitan_kyori = 0
			e.eki_list[i].temae_list = append(e.eki_list[i].temae_list, kiten)
		}
	}
}

// Check if p and q are connected, and return edited q
func Kousin1(p Eki, q Eki) Eki {
	if GlobalEkikanList.GetEkikanKyori(p.namae, q.namae) == math.Inf(+1) {
		return q
	}

	if q.saitan_kyori > p.saitan_kyori {
		q.saitan_kyori = p.saitan_kyori
		q.temae_list = append(q.temae_list, q.namae, p.namae)
	}
	return q
}

func Koushin(p *Eki, v *EkiList) *EkiList{
	for _, q := range v.eki_list {
		Kousin1(p, &q)
	}
	return v
}