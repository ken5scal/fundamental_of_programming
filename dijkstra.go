package main

import (
	"math"
	"errors"
	"fmt"
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
func (q *Eki) Kousin1(p *Eki) {
	kyori := GlobalEkikanList.GetEkikanKyori(p.namae, q.namae)
	if kyori == math.Inf(+1) {
		return
	}

	if q.saitan_kyori > p.saitan_kyori {
		q.saitan_kyori = p.saitan_kyori + kyori
		q.temae_list = append(q.temae_list, q.namae)
		q.temae_list = append(q.temae_list, p.temae_list...)
	}
}

// Repeat kousin1 for unfixed Eki lists
func Koushin(p *Eki, v *EkiList) {
	for i, q := range v.eki_list {
		q.Kousin1(p)
		v.eki_list[i] = q
	}
}

// Returns Stations that have minimum Distance and  Other Stations
func SaitanWoBunri(v *EkiList) (*Eki, *EkiList, error) {
	min_kyori := math.Inf(+1)
	for i, eki := range v.eki_list {
		if min_kyori > eki.saitan_kyori {
			new_ekilist := make([]Eki, len(v.eki_list))
			copy(new_ekilist, v.eki_list)
			new_ekilist = append(new_ekilist[:i], new_ekilist[i + 1:]...)
			return &eki, &EkiList{eki_list:new_ekilist}, nil
		}
	}
	return nil, nil, errors.New("At least one of SaitanKyori should be Finite Number")
}

// Returns list with minimum distance and path
// v: Undetermined Eki list
// g: Global Ekikan
func Dijkstra_main(v *EkiList, g *GlobalEkikan) *EkiList {
	// 1) SaitanWobunri(v) -> p, v with stripped p
	// 2) Koushin(p, v) -> updated min distance for every V element based on p
	if len(v.eki_list) == 0 {
		return v
	}

	p, new_v, err := SaitanWoBunri(v); if err != nil {
		panic(err.Error())
	}

	Koushin(p, new_v)
	fmt.Printf("P: %v\n" ,p)
	Dijkstra_main(new_v, g)

	//fmt.Println(append(hoge.eki_list, p))
	return new_v
}
