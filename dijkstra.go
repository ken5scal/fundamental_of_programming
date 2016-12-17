package main

import (
	"math"
	"errors"
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
func (e *EkiList) Shokika(kitenName string) {
	KitenEki := func(eki Eki) Eki {
		eki.saitan_kyori = 0
		eki.temae_list = append(eki.temae_list, kitenName)
		return eki
	}
	for i, eki := range e.eki_list {
		if eki.namae == kitenName {
			e.eki_list[i] = KitenEki(eki)
		}
	}
}

// Check if p and q are connected
// If connected, then update q's saitankyori / temae list otherwise leave as it is
func (q *Eki)Kousin1(p Eki) {
	kyori := GlobalEkikanList.GetEkikanKyori(p.namae, q.namae); if kyori == math.Inf(+1) {
		return
	}

	if q.saitan_kyori > p.saitan_kyori {
		q.saitan_kyori = p.saitan_kyori + kyori
		q.temae_list = append(q.temae_list, q.namae, p.namae)
	}
}

// Repeat kousin1 for unfixed Eki lists
func Koushin(p Eki, v *EkiList) *EkiList {
	for i, q := range v.eki_list {
		q.Kousin1(p)
		v.eki_list[i] = q
	}
	return v
}

// Returns Stations that have minimum Distance and  Other Stations
func (v *EkiList) SaitanWoBunri() (*Eki, *[]Eki, error) {
	min_kyori := math.Inf(+1)
	for i, eki := range v.eki_list {
		if min_kyori > eki.saitan_kyori {
			copy(v.eki_list[i:], v.eki_list[i+1:])
			v.eki_list = append(v.eki_list[:i], v.eki_list[i+1:]...)
			return &eki, &v.eki_list, nil
		}
	}
	return nil, nil, errors.New("At least one of SaitanKyori should be Finite Number")
}



