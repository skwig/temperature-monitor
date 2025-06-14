package sql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var _ Repository = (*SqliteRepository)(nil)

type SqliteRepository struct {
	path string
	db   *sqlx.DB
}

func NewDefaultSqliteRepository() (Repository, error) {
	// TODO: Some sort of config?
	return NewSqliteRepository("./db/sqlite.db")
}

func NewSqliteRepository(path string) (Repository, error) {
	// os.Remove(path)
	repo := SqliteRepository{path, nil}
	err := repo.Setup()
	if err != nil {
		return nil, err
	}

	return &repo, nil
}

// Setup implements Repository.
func (repo *SqliteRepository) Setup() error {
	db, err := sqlx.Connect("sqlite3", repo.path)
	if err != nil {
		return err
	}

	db.MustExec(`
		CREATE TABLE IF NOT EXISTS sensor_readings (
			id           INTEGER PRIMARY KEY,
			session      TEXT,
			sensor_time  INTEGER,
			server_time  INTEGER,
			temperature  REAL,
			humidity     REAL
		);
	`)

	repo.db = db

	return nil
}

// Save implements Repository.
func (repo *SqliteRepository) Save(reading *SensorReading) error {

	tx := repo.db.MustBegin()

	result, err := tx.NamedExec("INSERT INTO sensor_readings (session, server_time, sensor_time, temperature, humidity) VALUES (:session, :server_time, :sensor_time, :temperature, :humidity)", reading)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	reading.Id = id

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// GetAll implements Repository.
func (repo *SqliteRepository) GetAll() ([]SensorReading, error) {
	rows := []SensorReading{}
	err := repo.db.Select(&rows, "SELECT session, server_time, sensor_time, temperature, humidity FROM sensor_readings")
	if err != nil {
		return nil, err
	}

	return rows, nil
}
