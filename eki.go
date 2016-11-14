package main

type Eki struct {
	namae string
	saitan_kyori float64
	temae_list []string
}

func New(ekimei []Ekimei) Eki {
	return &Eki{}
}