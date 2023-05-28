package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"http-calendar/internal/models"
	"time"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

func (pr *PostgresRepository) Close() {
	pr.db.Close()
}

func (pr *PostgresRepository) Create(event models.Event) (int64, error) {
	query := `insert into events (title, date) values ($1, $2)`
	res, err := pr.db.Exec(query, event.Title, event.Date)
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (pr *PostgresRepository) Update(id int64, title string, date time.Time) error {
	query := `update events set title = $1, date = $2 where id_event = $3`

	_, err := pr.db.Exec(query, title, date, id)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PostgresRepository) Delete(id int) error {
	query := `delete from events where id_event = $1`
	_, err := pr.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PostgresRepository) GetForDay(date time.Time) []models.Event {
	return nil
}

func (pr *PostgresRepository) GetForWeek(date time.Time) []models.Event {
	return nil
}

func (pr *PostgresRepository) GetForMonth(date time.Time) []models.Event {
	return nil
}
