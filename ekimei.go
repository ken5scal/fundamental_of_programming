package main

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