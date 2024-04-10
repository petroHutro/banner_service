package storage

import "database/sql"

// TO DO
type UserStorage interface {
}

type Storage interface {
	UserStorage
}

type storageData struct {
	db *sql.DB
}
