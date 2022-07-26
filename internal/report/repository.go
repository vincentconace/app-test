package report

import (
	"database/sql"
	"time"
)

type Repository interface {
	SaveReportPosts(userName string, idposts int, text string, date time.Time) (int, error)
	SaveReportComments(userName string, idcomments int, text string, date time.Time) (int, error)
	Exists(userName string) bool
	ExistsPosts(idPost int) bool
	ExistsComment(idPost int) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) SaveReportPosts(userName string, idposts int, text string, date time.Time) (int, error) {
	query := `INSERT INTO reports_posts (id_posts, user_name, text, create_date) VALUES(?, ?, ?, ?);`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(idposts, userName, text, date)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) SaveReportComments(userName string, idcomments int, text string, date time.Time) (int, error) {
	query := `INSERT INTO reports (user_name, id_posts, text, create_date) VALUES(?, ?, ?, ?)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(userName, idcomments, text, date)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Exists(userName string) bool {
	var user string
	query := "SELECT user_name FROM users WHERE user_name=?;"
	row := r.db.QueryRow(query, userName)
	err := row.Scan(&user)
	return err == nil
}

func (r *repository) ExistsPosts(idPost int) bool {
	var id int
	query := "SELECT id FROM posts WHERE id=?;"
	row := r.db.QueryRow(query, idPost)
	err := row.Scan(&id)
	return err == nil
}

func (r *repository) ExistsComment(idPost int) bool {
	var id int
	query := "SELECT id FROM comments WHERE id=?;"
	row := r.db.QueryRow(query, idPost)
	err := row.Scan(&id)
	return err == nil
}
