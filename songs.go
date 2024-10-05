package app

type Song struct {
	Info
	Detail
}

type Info struct {
	Id    int    `json:"id" db:"id"`
	Group string `json:"group" db:"group_name" binding:"required"`
	Song  string `json:"song" db:"song" binding:"required"`
}

type Detail struct {
	ReleaseDate string `json:"releaseDate" db:"release_date" binding:"required"`
	Text        string `json:"text" db:"text" binding:"required"`
	Link        string `json:"link" db:"link" binding:"required"`
}
