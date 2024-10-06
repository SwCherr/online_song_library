package app

type Song struct {
	Info
	Detail
}

type Info struct {
	Id    int    `json:"id" db:"id"`
	Group string `json:"group" db:"group_name"`
	Song  string `json:"song" db:"song"`
}

type Detail struct {
	ReleaseDate string `json:"releaseDate" db:"release_date"`
	Text        string `json:"text" db:"text"`
	Link        string `json:"link" db:"link"`
}
