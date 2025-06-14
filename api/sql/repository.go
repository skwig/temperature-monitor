package sql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type SensorReading struct {
	Id          int64   `db:"id"`
	Session     string  `db:"session"`
	SensorTime  string  `db:"sensorTime"`
	Temperature float32 `db:"temperature"`
	Humidity    float32 `db:"humidity"`
}

type Repository interface {
	Setup() error
	Save(reading *SensorReading) error
}

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

func (repo *SqliteRepository) Setup() error {
	db, err := sqlx.Connect("sqlite3", repo.path)
	if err != nil {
		return err
	}

	db.MustExec(`
		CREATE TABLE IF NOT EXISTS sensor_readings (
			id           INTEGER PRIMARY KEY,
			session      TEXT,
			sensor_time  TEXT,
			temperature  REAL,
			humidity     REAL
		);
	`)

	repo.db = db

	return nil
}

func (repo *SqliteRepository) Save(reading *SensorReading) error {

	tx := repo.db.MustBegin()

	result := tx.MustExec("INSERT INTO sensor_readings (session, sensor_time, temperature, humidity) VALUES ($1, $2, $3, $4)",
		reading.Session, reading.SensorTime, reading.Temperature, reading.Humidity)
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
