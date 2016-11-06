package main

type Ekimei struct {
	Kanji   string
	Kana    string
	Romaji  string
	Shozoku string
}

// Returns Shozoku and Kanji(Kana)
// Ex: &e{Kanji: A, Kana: B, Shozoku: C}.hyoji() -> "A, B(C)"
func (e *Ekimei) Hyoji() string {
	return ""
}