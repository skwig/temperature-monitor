package sql

import (
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type SensorReading struct {
	Id             int64     `db:"id"`
	Session        uuid.UUID `db:"session"`
	SensorTimeUnix int64     `db:"sensor_time"`
	ServerTimeUnix int64     `db:"server_time"`
	Temperature    float32   `db:"temperature"`
	Humidity       float32   `db:"humidity"`
}

type Repository interface {
	Setup() error
	Save(reading *SensorReading) error
	GetAll() ([]SensorReading, error)
}
