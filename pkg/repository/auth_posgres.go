package repository

import (
	"fmt"
	"online_music/base"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// group_name, song, release_date, text, link	var id int
func (r *AuthPostgres) GetFilterData(song base.Song) ([]base.Song, error) {
	var info []base.Song
	query := fmt.Sprintf("SELECT * FROM %s "+
		"WHERE group_name = COALESCE(NULLIF($1, ''), group_name) AND "+
		"song = COALESCE(NULLIF($2, ''), song) AND "+
		"release_date = COALESCE(NULLIF($3, ''), release_date) AND "+
		"text = COALESCE(NULLIF($4, ''), text) AND "+
		"link = COALESCE(NULLIF($5, ''), link)",
		songTable)
	if err := r.db.Select(&info, query, song.Group, song.Song, song.ReleaseDate, song.Text, song.Link); err != nil {
		return info, err
	}
	return info, nil
}

func (r *AuthPostgres) GetTextSong(id int) (string, error) {
	var text string
	query := fmt.Sprintf("SELECT text FROM %s WHERE id=$1", songTable)
	if err := r.db.Get(&text, query, id); err != nil {
		return "", err
	}
	return text, nil
}

func (r *AuthPostgres) DeleteSongByID(id int) error {
	var res int
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 RETURNING id", songTable)
	user_row := r.db.QueryRow(query, id)
	if err := user_row.Scan(&res); err != nil {
		return err
	}
	return nil
}

func (r *AuthPostgres) UpdateSongByID(song base.Song) error {
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
		return err
	}
	return nil
}

func (r *AuthPostgres) PostNewSong(song base.Song) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (group_name, song, release_date, text, link) values ($1, $2, $3, $4, $5) RETURNING id", songTable)
	user_row := r.db.QueryRow(query, song.Group, song.Song, song.ReleaseDate, song.Text, song.Link)
	if err := user_row.Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}
