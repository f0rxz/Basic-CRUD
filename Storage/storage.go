package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db     *sql.DB
	Config *Config
}

func NewStorage() *Storage {
	s := &Storage{
		Config: NewConfig(),
	}

	err := s.Open()
	if err != nil {
		panic(err)
		return nil
	}

	return s
}
func (s *Storage) Open() error {

	url := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		s.Config.Login,
		s.Config.Password,
		s.Config.Adress,
		s.Config.DbName,
	)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	s.Db = db
	return nil
}

func (s *Storage) Close() error {
	err := s.Db.Close()
	if err != nil {

		return err
	}
	return nil
}
