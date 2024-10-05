package repository

import (
	"app"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// group_name, song, release_date, text, link	var id int
func (r *AuthPostgres) DeleteSong(song app.Song) error {
	var id int
	query := fmt.Sprintf("DELETE FROM %s WHERE group_name=$1 AND song=$2 RETURNING id", songTable)
	user_row := r.db.QueryRow(query, song.Group, song.Song)
	if err := user_row.Scan(&id); err != nil {
		return err
	}
	return nil
}

// group_name, song, release_date, text, link	var id int
func (r *AuthPostgres) UpdateSong(song app.Song) (int, error) {
	var id int
	query := fmt.Sprintf("UPDATE %s "+
		"SET group_name = COALESCE(NULLIF($1, ''), group_name), "+
		"song = COALESCE(NULLIF($2, ''), song), "+
		"release_date = COALESCE(NULLIF($3, ''), release_date), "+
		"text = COALESCE(NULLIF($4, ''), text), "+
		"link = COALESCE(NULLIF($5, ''), link) "+
		"WHERE id = $6 "+
		"RETURNING id",
		songTable)
	user_row := r.db.QueryRow(query, song.Group, song.Song, song.ReleaseDate, song.Text, song.Link, song.Id)
	if err := user_row.Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

func (r *AuthPostgres) PostNewSong(song app.Song) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (group_name, song, release_date, text, link) values ($1, $2, $3, $4, $5) RETURNING id", songTable)
	user_row := r.db.QueryRow(query, song.Group, song.Song, song.ReleaseDate, song.Text, song.Link)
	if err := user_row.Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}
